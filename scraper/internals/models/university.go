package models

type Major struct {
	Id            int      `db:"id"`
	Title         string   `db:"title"`
	Aim           string   `db:"aim"`
	Opportunities []string `db:"opportunities"` // NOTE: in DB it's TEXT[], here it's string
}

type Program struct {
	Id       int    `db:"id"`
	MajorId  int    `db:"major_id"`
	Type     string `db:"type"`
	Duration int    `db:"duration_years"`
}

type ProgramYear struct {
	Id         int `db:"id"`
	ProgramId  int `db:"program_id"`
	YearNumber int `db:"year_number"`
}

type Semester struct {
	Id             int `db:"id"`
	YearId         int `db:"year_id"`
	SemesterNumber int `db:"semester_number"`
}

type Course struct {
	Id          int    `db:"id"`
	Code        string `db:"code"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type SemesterCourse struct {
	SemesterId int `db:"semester_id"`
	CourseId   int `db:"course_id"`
}

type CoursePrereqGroup struct {
	Id       int `db:"id"`
	CourseId int `db:"course_id"`
}

type CoursePrereqOption struct {
	GroupId  int `db:"group_id"`
	PrereqId int `db:"prereq_id"`
}

type SemesterElective struct {
	Id         int    `db:"id"`
	SemesterId int    `db:"semester_id"`
	Title      string `db:"title"`
}

type SemesterElectiveCourse struct {
	SemesterElectiveId int `db:"semester_elective_id"`
	ElectiveCourse     int `db:"elective_course"`
}
