package service

import (
	"time"
)

type m map[string]interface{}

// slice of map
type arrayMap []m

var students = arrayMap{}

// GetStudentInfos model returned when request single student infos
type GetStudentInfos struct {
	Student *Student `json:"student"`
	Exams   []*Exam  `json:"exams"`
}

// StudentsListOut model
type StudentsListOut struct {
	Students []*Student `json:"students"`
	Count    uint       `json:"count"`
}

// ExamListsOut model
type ExamListsOut struct {
	ExamLists []*ExamList `json:"examLists"`
	Count     uint        `json:"count"`
}

// MultiLanguageField model
type MultiLanguageField struct {
	FR string `json:"fr"`
	AR string `json:"ar"`
}

//Student model
type Student struct {
	ID             uint       `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt      time.Time  `sql:"DEFAULT:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `sql:"index"`
	FileNumber     string     `gorm:"unique" json:"file_number,omitempty"`
	FirstName      string     `gorm:"type:varchar(512)" json:"first_name,omitempty"`
	LastName       string     `gorm:"type:varchar(512);index:last_name" json:"last_name,omitempty"`
	MaidenName     string     `gorm:"type:varchar(512)" json:"maiden_name,omitempty"`
	PhoneNumber    string     `json:"phone_number,omitempty"`
	Job            string     `json:"job,omitempty"`
	BirthDay       time.Time  `json:"birthday,omitempty"`
	Gender         string     `json:"gender,omitempty"`
	Country        string     `json:"country,omitempty"`
	City           string     `json:"city,omitempty"`
	Department     string     `json:"department,omitempty"`
	AddressStreet  string     `gorm:"type:varchar(512)" json:"address_street,omitempty"`
	RegistredDate  time.Time  `json:"registred_date,omitempty"`
	Image          string     `gorm:"default:'imageURL'" json:"image,omitempty"`
	WinLicenceDate *time.Time `sql:"index" json:"win_licence_date,omitempty"`
	ExamLevel      uint       `sql:"default:1" json:"exam_level"`
	LastExamDate   *time.Time `json:"last_exam_date"`
	LastExamStatus *bool      `json:"last_exam_status"`
	Archived       bool       `json:"archived"`
}

// TableName :Database table name
func (s Student) TableName() string {
	return "students"
}

// Exam an exam passed by a student
type Exam struct {
	ID         uint       `gorm:"unique_index" json:"id,omitempty"`
	CreatedAt  time.Time  `sql:"DEFAULT:current_timestamp" json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `sql:"index"`
	ExamLevel  uint8      `sql:"default:1" json:"exam_level"`
	DateExam   time.Time  `json:"date_exam,omitempty"`
	Status     bool       `gorm:"default:false" json:"status"`
	StudentID  uint       `json:"student_id,omitempty"`
	Student    *Student   `json:"student,omitempty"`
	ExamListID uint       `json:"exam_list_id,omitempty"`
}

// TableName :Database table name
func (e Exam) TableName() string {
	return "exams"
}

//ExamList a list of students who are added to pass an exam
type ExamList struct {
	ID            uint       `gorm:"unique_index" json:"id,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `sql:"index"`
	DateExam      time.Time  `json:"date_exam,omitempty"`
	Examiner      string     `json:"examiner,omitempty"`
	Archived      bool       `json:"archived"`
	StudentsExams []*Exam    `json:"students_exams,omitempty"`
}

// TableName :Database table name
func (s ExamList) TableName() string {
	return "exam_lists"
}
