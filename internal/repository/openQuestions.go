package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"book-school/internal/models"
)

type CreativeTaskRepository struct {
	db *sql.DB
}

func newCreativeTaskRepository(db *sql.DB) *CreativeTaskRepository {
	return &CreativeTaskRepository{
		db: db,
	}
}

func (g *CreativeTaskRepository) GetCreativeTask(user models.User, category string, isCreative bool) ([]models.CreativeTask, error) {
	table := ""
	if isCreative {
		table = "creative_questions"
	} else {
		table = "open_questions"
	}
	args := []interface{}{}
	query := fmt.Sprintf(`SELECT 
			q.id,
			q.category,
			q.question,
			q.audio_link
			FROM %s`, table+" q;")
	if category != "" {
		query = strings.ReplaceAll(query, ";", " WHERE category = $1;")
		args = append(args, category)
	}
	if user.Role == models.StudentRole {
		if category != "" {
			query = strings.ReplaceAll(query, ";", " AND ;")
		} else {
			query = strings.ReplaceAll(query, ";", " WHERE ;")
		}
		query = strings.ReplaceAll(query, ";", fmt.Sprintf("q.id NOT IN (select a.open_questions_id from open_answers a where a.student_id = %v);", user.Id))
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("CreativeTask repository: get all: %w", err)
	}
	defer prep.Close()
	rows, err := prep.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("CreativeTask repository: get all: %w", err)
	}
	defer rows.Close()
	CreativeTask := []models.CreativeTask{}
	for rows.Next() {
		currentQuestion := models.CreativeTask{}
		if err := rows.Scan(&currentQuestion.Id, &currentQuestion.Category, &currentQuestion.Question, &currentQuestion.Audio); err != nil {
			return nil, fmt.Errorf("CreativeTask repository: get all: %w", err)
		}
		CreativeTask = append(CreativeTask, currentQuestion)
	}
	return CreativeTask, nil
}

func (g *CreativeTaskRepository) CreateCreativeTask(question models.CreativeTask, isCreative bool) (int, error) {
	if question.Question == "" {
		return -1, errors.New("nil text")
	}
	var id int
	query := ""
	if isCreative {
		query = "INSERT INTO creative_questions(question, category) VALUES($1, $2) RETURNING id;"
	} else {
		query = "INSERT INTO open_questions(question, category) VALUES($1, $2) RETURNING id;"

	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer prep.Close()
	err = prep.QueryRow(question.Question, question.Category).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (g *CreativeTaskRepository) DeleteCreativeTask(id int, isCreative bool) error {
	query := ""
	if isCreative {
		query = "DELETE FROM creative_questions WHERE id=$1;"
	} else {
		query = "DELETE FROM open_questions WHERE id=$1;"
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err := prep.Exec(id); err != nil {
		return err
	}
	return nil
}

func (g *CreativeTaskRepository) UpdateCreativeTask(question interface{}, isCreative bool) error {
	table := ""
	if isCreative {
		table = "creative_questions"
	} else {
		table = "open_questions"
	}
	input, ok := question.(map[string]interface{})
	if !ok {
		return errors.New("cannot decode input in update creative task")
	}
	if input["id"] == nil {
		return errors.New("not valid id")
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input["category"] != nil {
		setValues = append(setValues, fmt.Sprintf("category = $%d", argId))
		args = append(args, input["category"])
		argId++
	}

	if input["question"] != nil {
		if len(input["question"].(string)) > 200 {
			return errors.New("question's letter is over limitted")
		}
		setValues = append(setValues, fmt.Sprintf("question = $%d", argId))
		args = append(args, input["question"])
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d;", table, setQuery, argId)
	args = append(args, input["id"])
	prep, err := g.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("UpdateCreativeTask repository: update: %w", err)
	}

	defer prep.Close()

	if _, err = prep.Exec(args...); err != nil {
		return fmt.Errorf("UpdateCreativeTask repository: update: %w", err)
	}
	return nil
}

func (g *CreativeTaskRepository) CreatePassCreativeTask(answer models.PassCreativeTask, isCreative bool) error {
	if _, err := g.GetCurrentStudentPass(answer.BookId, answer.StudentId, answer.QuestionId, isCreative); err == nil {
		return fmt.Errorf("CreatePassCreativeTask repository: insert answer: %w", errors.New("students answer is already exists"))
	}
	query := ""
	if !isCreative {
		query = "INSERT INTO open_answers(answer, open_questions_id, student_id, book_id) VALUES($1, $2, $3, $4);"
	} else {
		query = "INSERT INTO creative_answers(answer, creative_questions_id, student_id, book_id) VALUES($1, $2, $3, $4);"
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("CreatePassCreativeTask repository: insert answer: %w", err)
	}
	defer prep.Close()
	if err != nil {
		return fmt.Errorf("CreatePassCreativeTask repository: insert answer: %w", err)
	}
	if _, err := prep.Exec(answer.Answer, answer.QuestionId, answer.StudentId, answer.BookId); err != nil {
		return fmt.Errorf("CreatePassCreativeTask repository: insert answer: %w", err)
	}
	return nil
}

func (g *CreativeTaskRepository) GetCurrentStudentPass(bookId int, studentId uint, questionId int, isCreative bool) (models.PassCreativeTask, error) {
	query := ""
	if !isCreative {
		query = `SELECT 
	o.id,
	o.answer,
	o.open_questions_id,
	o.student_id,
	o.book_id,
	q.question
	FROM open_answers o 
	inner join open_questions q on q.id = o.open_questions_id
	WHERE open_questions_id = $1 AND student_id = $2 AND book_id = $3;`
	} else {
		query = `SELECT 
		o.id,
		o.answer,
		o.creative_questions_id,
		o.student_id,
		o.book_id,
		q.question
		FROM creative_answers o 
		inner join creative_questions q on q.id = o.creative_questions_id
		WHERE creative_questions_id = $1 AND student_id = $2 AND book_id = $3;`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return models.PassCreativeTask{}, err
	}
	defer prep.Close()
	var input models.PassCreativeTask
	err = prep.QueryRow(questionId, studentId, bookId).Scan(&input.Id, &input.Answer, &input.QuestionId, &input.StudentId, &input.BookId, &input.Question)
	if err != nil {
		return models.PassCreativeTask{}, err
	}
	return input, nil
}

func (g *CreativeTaskRepository) GetPassedStudents(classId int, bookId int, isCreative bool) ([]models.Student, error) {
	query := `
	select 
	DISTINCT ON (s.id)
		s.id, 
		s.username,
		s.first_name, 
		s.second_name,
		b.category
	from students s
	JOIN books b ON b.id = $2 
	where 
	s.class_id = $1
	group by s.id, s.first_name, s.second_name, s.username, b.category;
	`
	prep, err := g.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("GetPassedStudents repository: get all: %w", err)
	}
	defer prep.Close()
	rows, err := prep.Query(classId, bookId)
	if err != nil {
		return nil, fmt.Errorf("GetPassedStudents repository: get all: %w", err)
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		current := models.Student{}
		current.ClassId = uint(classId)
		var category string
		if err := rows.Scan(&current.Id, &current.Username, &current.FirstName, &current.SecondName, &category); err != nil {
			return nil, fmt.Errorf("GetPassedStudents repository: get all: %w", err)
		}
		allPasses, _ := g.CountAllPasses(bookId, current.Id, isCreative)
		current.CountPass = uint(allPasses)
		students = append(students, current)
	}
	return students, nil
}

func (g *CreativeTaskRepository) GetStudentAllPasses(bookId int, studentId uint, isCreative bool) ([]models.PassCreativeTask, error) {
	query := ""
	if !isCreative {
		query = `SELECT 
	o.id,
	o.answer,
	o.open_questions_id,
	o.student_id,
	o.book_id,
	q.question
	FROM open_answers o 
	inner join open_questions q on q.id = o.open_questions_id
	WHERE o.student_id = $1 AND o.book_id = $2
	AND o.id NOT IN (SELECT c.answer_id from open_comments c)
	;`
	} else {
		query = `SELECT 
	o.id,
	o.answer,
	o.creative_questions_id,
	o.student_id,
	o.book_id,
	q.question
	FROM creative_answers o 
	inner join creative_questions q on q.id = o.creative_questions_id
	WHERE o.student_id = $1 AND o.book_id = $2
	AND o.id NOT IN (SELECT c.answer_id from creative_comments c)
	;`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()
	var input []models.PassCreativeTask
	rows, err := prep.Query(studentId, bookId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var current models.PassCreativeTask
		if err := rows.Scan(&current.Id, &current.Answer, &current.QuestionId, &current.StudentId, &current.BookId, &current.Question); err != nil {
			return nil, err
		}
		input = append(input, current)
	}
	return input, nil
}

func (g *CreativeTaskRepository) PostCommentStudent(comment models.CheckCreativePass, isCreative bool) (uint, error) {
	query := ""
	if !isCreative {
		query = `INSERT INTO open_comments(answer_id, teacher_id, student_id, comment, points) VALUES($1, $2, $3, $4, $5) RETURNING id;`

	} else {
		query = `INSERT INTO creative_comments(answer_id, teacher_id, student_id, comment, points) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("PostCommentStudent repository: insert comment: %w", err)
	}
	defer prep.Close()
	if err != nil {
		return 0, fmt.Errorf("PostCommentStudent repository: insert comment: %w", err)
	}
	var id uint
	if err := prep.QueryRow(comment.AnswerId, comment.TeacherId, comment.StudentId, comment.Comment, comment.Point).Scan(&id); err != nil {
		return 0, fmt.Errorf("PostCommentStudent repository: insert comment: %w", err)
	}
	if err := g.CreateNotifications(id, false, isCreative); err != nil {
		return 0, fmt.Errorf("PostCommentStudent repository: insert comment: %w", err)
	}
	return id, nil
}

func (g *CreativeTaskRepository) CreateNotifications(id uint, check bool, isCreative bool) error {
	query := ""
	if !isCreative {
		query = `INSERT INTO open_notifications (comments_id, is_read) VALUES($1, $2);`

	} else {
		query = `INSERT INTO creative_notifications (comments_id, is_read) VALUES($1, $2);`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("CreateNotifications repository: insert open_notifications: %w", err)
	}
	defer prep.Close()
	if err != nil {
		return fmt.Errorf("CreateNotifications repository: insert open_notifications: %w", err)
	}
	if _, err := prep.Exec(id, check); err != nil {
		return fmt.Errorf("CreateNotifications repository: insert open_notifications: %w", err)
	}
	return nil
}

func (g *CreativeTaskRepository) UpdateNotifications(id uint, check bool, isCreative bool) error {
	query := ""
	if !isCreative {
		query = `UPDATE open_notifications SET is_read = $1 WHERE comments_id = $2`

	} else {
		query = `UPDATE creative_notifications SET is_read = $1 WHERE comments_id = $2`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("UpdateNotifications repository: update open_notifications: %w", err)
	}
	defer prep.Close()
	if err != nil {
		return fmt.Errorf("UpdateNotifications repository: update open_notifications: %w", err)
	}
	if _, err := prep.Exec(check, id); err != nil {
		return err
	}
	return nil
}

func (g *CreativeTaskRepository) CountAllPasses(bookId int, studentId uint, isCreative bool) (int, error) {
	query := ""
	if !isCreative {
		query = `SELECT 
		COUNT(o.id)
		FROM open_answers o 
		WHERE o.student_id = $1 AND o.book_id = $2;`

	} else {
		query = `SELECT 
		COUNT(o.id)
		FROM creative_answers o 
		WHERE o.student_id = $1 AND o.book_id = $2;`
	}

	prep, err := g.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer prep.Close()
	var count int
	if err := prep.QueryRow(studentId, bookId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (g *CreativeTaskRepository) GetComments(studentId uint, imageUrl string, isCreative bool) ([]models.CreativeNotifications, error) {
	var comments []models.CreativeNotifications
	query := ""
	if !isCreative {
		query = `
		SELECT 
		c.id,
		c.comment,
		c.audio_link,
		a.answer,
		q.question,
		b.name
		FROM open_comments c
		JOIN open_answers a on a.id = c.answer_id
		JOIN open_questions q on q.id = a.open_questions_id
		JOIN books b on b.id = a.book_id
		WHERE c.student_id = $1
		AND c.id IN (SELECT n.comments_id from open_notifications n WHERE NOT n.is_read)
		;
	`

	} else {
		query = `
		SELECT 
		c.id,
		c.comment,
		c.audio_link,
		a.answer,
		q.question,
		b.name
		FROM creative_comments c
		JOIN creative_answers a on a.id = c.answer_id
		JOIN creative_questions q on q.id = a.creative_questions_id
		JOIN books b on b.id = a.book_id
		WHERE c.student_id = $1
		AND c.id IN (SELECT n.comments_id from creative_notifications n WHERE NOT n.is_read)
		;
	`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer prep.Close()
	rows, err := prep.Query(studentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			current   models.CreativeNotifications
			commentId uint
		)
		if err := rows.Scan(&commentId, &current.Comment, &current.Audio, &current.Answer, &current.Question, &current.BookName); err != nil {
			return nil, err
		}
		if err := g.UpdateNotifications(commentId, true, isCreative); err != nil {
			return nil, err
		}
		comments = append(comments, current)
	}

	return comments, nil
}

func (g *CreativeTaskRepository) GetAnswerById(answerId uint, isCreative bool) (models.PassCreativeTask, error) {
	log.Println(answerId)
	var answer models.PassCreativeTask
	query := ""
	if !isCreative {
		query = `SELECT * FROM open_answers where id = $1;`

	} else {
		query = `SELECT * FROM creative_answers where id = $1;`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return answer, err
	}
	defer prep.Close()
	if err := prep.QueryRow(answerId).Scan(&answer.Id, &answer.Answer, &answer.QuestionId, &answer.StudentId, &answer.BookId); err != nil {
		return answer, err
	}
	return answer, nil
}

func (g *CreativeTaskRepository) AddAudioLinkComment(audio string, commentId uint, isCreative bool) error {
	query := ""
	if !isCreative {
		query = `UPDATE open_comments SET audio_link = $1 WHERE id = $2`
	} else {
		query = `UPDATE creative_comments SET audio_link = $1 WHERE id = $2`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err := prep.Exec(audio, commentId); err != nil {
		return err
	}
	return nil
}

func (g *CreativeTaskRepository) AddAudioLinkQuestion(audio string, questionId uint, isCreative bool) error {
	query := ""
	if !isCreative {
		query = `UPDATE open_questions SET audio_link = $1 WHERE id = $2`
	} else {
		query = `UPDATE creative_questions SET audio_link = $1 WHERE id = $2`
	}
	prep, err := g.db.Prepare(query)
	if err != nil {
		return err
	}
	defer prep.Close()
	if _, err := prep.Exec(audio, questionId); err != nil {
		return err
	}
	return nil
}
