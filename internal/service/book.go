package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/pkg/tts"
	"book-school/pkg/utils"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type BookService struct {
	bookRepo    repository.Book
	studentRepo repository.Student
	classRepo   repository.Class
	imageUrl    string
}

func newBookService(bookRepo repository.Book, studentRepo repository.Student, classRepo repository.Class, imageUrl string) *BookService {
	return &BookService{
		bookRepo:    bookRepo,
		studentRepo: studentRepo,
		classRepo:   classRepo,
		imageUrl:    imageUrl,
	}
}

func (b *BookService) GetAllForTest(ctx context.Context) ([]models.Book, error) {
	books, err := b.bookRepo.GetAllForTest(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookService) TotalPages(ctx context.Context, user models.User) (int, error) {
	if user.Role == models.StudentRole {
		student, err := b.studentRepo.GetById(ctx, user.Id)
		if err != nil {
			return 0, err
		}

		class, err := b.classRepo.GetAll(context.WithValue(ctx, models.ClassId, student.ClassId))
		if err != nil {
			return 0, err
		}

		ctx = context.WithValue(ctx, models.ClassGrade, class[0].Grade)
	}

	books, err := b.bookRepo.Get(ctx)
	if err != nil {
		return 0, fmt.Errorf("book service: total pages: %w", err)
	}

	return len(books) / 10, nil
}

func (b *BookService) Complete(ctx context.Context, bookId, studentId uint) (uint, error) {
	err := b.bookRepo.CheckCompleteBook(ctx, bookId, studentId)
	if err == nil {
		return 0, fmt.Errorf("book service: complete: you passed book once")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("book service: complete: %w", err)
	}

	if err := b.studentRepo.GivePoints(ctx, studentId, models.BookPoints); err != nil {
		return 0, fmt.Errorf("book service: complete: %w", err)
	}

	if err := b.bookRepo.Complete(ctx, bookId, studentId, models.BookPoints); err != nil {
		if err := b.studentRepo.TakePoints(ctx, studentId, models.BookPoints); err != nil {
			return 0, fmt.Errorf("book service: complete: %w", err)
		}
		return 0, fmt.Errorf("book service: complete: %w", err)
	}

	return models.BookPoints, nil
}

func (b *BookService) GetAll(ctx context.Context, user models.User) ([]models.BooksStruct, error) {
	var categories [][]string
	if user.Role == models.StudentRole {
		student, err := b.studentRepo.GetById(ctx, user.Id)
		if err != nil {
			return nil, err
		}

		class, err := b.classRepo.GetAll(context.WithValue(ctx, models.ClassId, student.ClassId))
		if err != nil {
			return nil, err
		}
		log.Println(class[0].Grade[1:])
		categories = categoryMassive(class[0].Grade[2:])
		ctx = context.WithValue(ctx, models.ClassGrade, class[0].Grade)
	}

	if user.Role == models.AdminRole {
		categories = categoryMassive("all")
	}

	if user.Role == models.TeacherRole {
		classes, err := b.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, user.Id))
		if err != nil {
			return nil, err
		}
		underage := false
		upperage := false
		for _, class := range classes {
			if class.Grade[2:] == "класс" {
				upperage = true
			}
			if class.Grade[2:] == "год" {
				underage = true
			}
			if underage && upperage {
				break
			}
		}
		if underage && upperage {
			categories = categoryMassive("all")
		} else if underage {
			categories = categoryMassive("год")
		} else {
			categories = categoryMassive("класс")
		}
	}

	var booksStruct []models.BooksStruct

	for _, category := range categories {
		var books []models.Book
		books, err := b.bookRepo.GetAll(ctx, category)
		if err != nil {
			return nil, err
		}
		for index, book := range books {
			books[index].Preview = b.imageUrl + "/static/books/" + book.Hashed_ID + "/preview.png"
		}
		booksStruct = append(booksStruct, models.BooksStruct{
			Categories: category[0],
			Books:      books,
		})
	}
	log.Println(categories)

	return booksStruct, nil
}

func categoryMassive(kind string) [][]string {
	category := [][]string{}
	if kind == "класс" || kind == "all" {
		category = append(category, []string{"Мир за окном", "Таң қаларлық дүние"})
		category = append(category, []string{"Путешествие в сказку", "Халық тағылымы"})
		category = append(category, []string{"Что хорошо, а что плохо", "Әдеп әлемі"})
		category = append(category, []string{"Удивительное рядом", "Ғажайып қасымызда"})
		category = append(category, []string{"Поэзия", "Поэзия"})
	}
	if kind == "год" || kind == "all" {
		category = append(category, []string{"Мир сказок", "Ертегілер әлемі"})
		category = append(category, []string{"Всё о природе", "Табиғатпен танысу"})
		category = append(category, []string{"Стихи и потешки", "Өлеңдер мен тақпақтар"})
		category = append(category, []string{"Азбука поведения", "Өнегелі ертегілер"})
		category = append(category, []string{"Книги развивашки", "Дамыту кітаптар"})
	}
	return category
}

func (b *BookService) GetJsonBook(id int) (models.Book, error) {
	var book models.Book
	hashed_id, err := b.bookRepo.GetBookHashedId(id)
	if err != nil {
		return models.Book{}, err
	}
	jsonFile, err := os.Open("./static/books/" + hashed_id + "/" + hashed_id + ".json")
	if err != nil {
		return models.Book{}, fmt.Errorf("book service: get json book: %w", err)
	}
	defer jsonFile.Close()

	if err = json.NewDecoder(jsonFile).Decode(&book); err != nil {
		return models.Book{}, fmt.Errorf("book service: get json book: %w", err)
	}

	return book, nil
}

func (b *BookService) DeleteBook(id int) error {
	hashed, err := b.bookRepo.GetBookHashedId(id)
	if err != nil {
		return err
	}
	err = os.RemoveAll("./static/books/" + hashed)
	if err != nil {
		return fmt.Errorf("book service: delete book: %w", err)
	}
	return b.bookRepo.DeleteBook(hashed)
}

func (b *BookService) CreateBook(input *models.Book, preview multipart.File) error {
	cats := []string{
		"Мир за окном", "Путешествие в сказку", "Что хорошо, а что плохо", "Удивительное рядом", "Поэзия", "Мир сказок", "Всё о природе", "Стихи и потешки", "Азбука поведения",
		"Книги развивашки", "Таң қаларлық дүние", "Халық тағылымы", "Әдеп әлемі", "Ғажайып қасымызда", "Табиғатпен танысу", "Ертегілер әлемі", "Өлеңдер мен тақпақтар", "Өнегелі ертегілер", "Дамыту кітаптар",
	}

	isValidCat := false

	for _, cat := range cats {
		if cat == input.Category {
			isValidCat = true
		}
	}

	if !isValidCat {
		return errors.New("not correct category")
	}

	if err := validationBook(input, preview); err != nil {
		return err
	}

	pReg, err := regexp.Compile("(<p>)(.*?)(</p>)")
	if err != nil {
		return err
	}

	imgReg, err := regexp.Compile(`(<img src="data:image\/.*">)`)
	if err != nil {
		return err
	}

	id := hashingId(input.Category + input.Class + input.Name)
	input.Hashed_ID = id
	if err := utils.CreateDirectory("./static/books/"); err != nil {
		return err
	}
	if utils.IsNotExistsDirectory("./static/books/" + id) {
		if err := utils.CreateDirectory("./static/books/" + id); err != nil {
			return err
		}
	} else {
		return errors.New("the directory is already exists")
	}

	isKaz := true
	if input.Language == "RU" {
		isKaz = false
	}

	errs, _ := errgroup.WithContext(context.Background())
	var mu sync.Mutex

	for i, page := range input.Pages {
		func(i int, page string) {
			errs.Go(func() error {
				mu.Lock()
				defer mu.Unlock()
				temp := pReg.ReplaceAllString(page, "$2\n")
				temp = strings.ReplaceAll(temp, "<br>\n", "")
				temp = imgReg.ReplaceAllString(temp, "")
				temp = strings.ReplaceAll(temp, "\n\n", "\n")
				temp = temp[:len(temp)-1]

				path, err := tts.TextToSpeech(temp, "./static/books/"+id+"/", i+1, isKaz)
				if err != nil {
					fmt.Println(err)
					err = os.RemoveAll("./static/books/" + id)
					if err != nil {
						return err
					}
					return err
				}

				input.Audio = append(input.Audio, b.imageUrl+path[1:])
				return nil
			})
		}(i, page)
	}

	err = errs.Wait()
	if err != nil {
		log.Println(err)
		err = os.RemoveAll("./static/books/" + id)
		if err != nil {
			return err
		}
		return err
	}

	sort.Slice(input.Audio, func(i, j int) bool {
		return input.Audio[i] < input.Audio[j]
	})

	if err := savePreview("./static/books/"+id+"/preview.png", preview); err != nil {
		err = os.RemoveAll("./static/books/" + id)
		if err != nil {
			return err
		}
		return err
	}

	pages, err := contentParser(input, "./static/books/"+id+"/", b.imageUrl)
	if err != nil {
		err = os.RemoveAll("./static/books/" + id)
		if err != nil {
			return err
		}
		return err
	}

	input.Id, err = b.bookRepo.CreateBook(input, id)
	if err != nil {
		err = os.RemoveAll("./static/books/" + id)
		if err != nil {
			return err
		}
		return err
	}

	input.Pages = pages
	input.Preview = b.imageUrl + "/static/books/" + id + "/preview.png"
	err = createJsonBook("./static/books/"+id+"/", input, id)
	if err != nil {
		err = os.RemoveAll("./static/books/" + id)
		if err != nil {
			return err
		}
		return err
	}

	return nil
}

func createJsonBook(filePath string, input *models.Book, id string) error {
	input.Hashed_ID = id
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(input)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath+id+".json", bytes.ReplaceAll(bytes.TrimRight(buffer.Bytes(), "\n"), []byte(`\"`), []byte(``)), 0o644)
	if err != nil {
		return err
	}
	return nil
}

func validationBook(input *models.Book, preview multipart.File) error {
	if len(input.Name) == 0 {
		return errors.New("name is empty")
	}
	if len(input.Name) > 500 {
		return errors.New("name's letter is overlimitted than 50")
	}

	if input.Class != "2 год" && input.Class != "3 год" && input.Class != "4 год" && input.Class != "5 год" &&
		input.Class != "1 класс" && input.Class != "2 класс" && input.Class != "3 класс" && input.Class != "4 класс" && input.Class != "2 жас" &&
		input.Class != "3 жас" && input.Class != "4 жас" && input.Class != "5 жас" && input.Class != "1 сынып" && input.Class != "2 сынып" && input.Class != "3 сынып" && input.Class != "4 сынып" {
		return errors.New("class is empty")
	}
	switch input.Class {
	case "2 жас":
		input.Class = "2 год"
	case "3 жас":
		input.Class = "3 год"
	case "4 жас":
		input.Class = "4 год"
	case "5 жас":
		input.Class = "5 год"
	case "1 сынып":
		input.Class = "1 класс"
	case "2 сынып":
		input.Class = "2 класс"
	case "3 сынып":
		input.Class = "3 класс"
	case "4 сынып":
		input.Class = "4 класс"
	}
	if len(input.Pages) == 0 {
		return errors.New("pages are empty")
	}
	if preview == nil {
		return errors.New("preview is empty")
	}
	return nil
}
