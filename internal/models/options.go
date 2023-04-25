package models

type ctx string

var (
	StudentId      ctx = "studentId"
	ClassId        ctx = "classId"
	MinClass       ctx = "minClass"
	MaxClass       ctx = "maxClass"
	TeacherId      ctx = "teacherId"
	SchoolId       ctx = "schoolId"
	BookId         ctx = "bookId"
	TestId         ctx = "testId"
	CompleteTestId ctx = "completeTestId"
	Page           ctx = "pages"
)

const (
	BookPoints      = 40
	AudioPoints     = 40
	PointsForAnswer = 10
)
