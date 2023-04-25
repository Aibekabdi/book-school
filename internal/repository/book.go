package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type BookRepository struct {
	db *sql.DB
}

func newBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) GetAllForTest(ctx context.Context) ([]models.Book, error) {
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id NOT IN (SELECT book_id FROM %s)", bookTable, testTable)

	prep, err := b.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	var books []models.Book

	rows, err := prep.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.Id, &book.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) CheckCompleteBook(ctx context.Context, bookId, studentId uint) error {
	query := fmt.Sprintf("SELECT id FROM %s WHERE book_id = $1 AND student_id = $2", completeBooksTable)

	prep, err := b.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("book repository: check complete book: %w", err)
	}
	defer prep.Close()

	var id uint

	err = prep.QueryRowContext(ctx, bookId, studentId).Scan(&id)
	if err != nil {
		return fmt.Errorf("book repository: check complete book: %w", err)
	}

	return nil
}

func (b *BookRepository) Complete(ctx context.Context, bookId, studentId, points uint) error {
	query := fmt.Sprintf("INSERT INTO %s(book_id, student_id, points) VALUES($1, $2, $3);", completeBooksTable)

	prep, err := b.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("book repository: complete: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, bookId, studentId, points); err != nil {
		return fmt.Errorf("book repository: complete: %w", err)
	}

	return nil
}

func (b *BookRepository) CreateBook(input *models.Book, hashed string) (int, error) {
	query := "INSERT INTO books(name, category, class, hashed_id, language) VALUES($1, $2, $3, $4, $5) RETURNING id;"

	prep, err := b.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	var id int
	if err := prep.QueryRow(input.Name, input.Category, input.Class, hashed, input.Language).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BookRepository) DeleteBook(hashed string) error {
	query := "DELETE FROM books WHERE hashed_id=$1;"

	prep, err := b.db.Prepare(query)
	if err != nil {
		return err
	}

	if _, err := prep.Exec(hashed); err != nil {
		return err
	}
	return nil
}

func (b *BookRepository) Get(ctx context.Context) ([]models.Book, error) {
	args := []interface{}{}

	offset := ""

	classGrade := ctx.Value(models.ClassGrade)
	if classGrade != nil {
		offset = " WHERE class = $1"
		args = append(args, classGrade.(string))
	}

	bookId := ctx.Value(models.BookId)
	if bookId != nil {
		offset = " WHERE id = $1"
		args = append(args, bookId.(uint))
	}

	query := fmt.Sprintf("SELECT * FROM %s%s;", bookTable, offset)

	prep, err := b.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("book repository: get: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("book repository: get: %w", err)
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.Id, &book.Name, &book.Category, &book.Class, &book.Hashed_ID, &book.Language); err != nil {
			return nil, fmt.Errorf("book repository: get: %w", err)
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepository) GetBookHashedId(id int) (string, error) {
	query := fmt.Sprintf("SELECT hashed_id FROM %s WHERE id = $1;", bookTable)

	prep, err := b.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer prep.Close()

	var input string
	err = prep.QueryRow(id).Scan(&input)
	if err != nil {
		return "", fmt.Errorf("GetBookHashedId repository: get hash: %w", err)
	}

	return input, nil
}

func (b *BookRepository) GetAll(ctx context.Context, category []string) ([]models.Book, error) {
	args := []interface{}{}

	offset := ""

	classGrade := ctx.Value(models.ClassGrade)
	check := false
	if classGrade != nil {
		check = true
		offset = " WHERE class = $1"
		args = append(args, classGrade.(string))
	}

	bookId := ctx.Value(models.BookId)
	if bookId != nil {
		check = true
		offset = " WHERE id = $1"
		args = append(args, bookId.(uint))
	}
	if check {
		offset += " AND (category = $2 OR category = $3)"
	} else {
		offset += " WHERE category = $1 OR category = $2"
	}
	args = append(args, category[0])
	args = append(args, category[1])

	query := fmt.Sprintf("SELECT * FROM %s%s;", bookTable, offset)

	prep, err := b.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("book repository: get: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("book repository: get: %w", err)
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.Id, &book.Name, &book.Category, &book.Class, &book.Hashed_ID, &book.Language); err != nil {
			return nil, fmt.Errorf("book repository: get: %w", err)
		}
		books = append(books, book)
	}

	return books, nil
}
