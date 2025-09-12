package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/scraper/database"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/scraper/internals/models"
)

func NewMajor(major models.Major) (int, error) {
	pool := database.GetPool()

	row := pool.QueryRow(
		context.Background(),
		`
			INSERT INTO major(title, aim, opportunities) 
			VALUES ($1, $2, $3) 
			RETURNING id;
		`,
		major.Title, major.Aim, major.Opportunities,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, errors.New(fmt.Sprintf("Error while adding major: %s: %s", major.Title, err.Error()))
	}

	return id, nil
}

func NewProgram(program models.Program) (int, error) {
	pool := database.GetPool()

	row := pool.QueryRow(
		context.Background(),
		`
			INSERT INTO program (major_id, type, duration)
			VALUES($1, $2, $3)
			RETURNING id;
		`,
		program.MajorId,
		program.Type,
		program.Duration,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, errors.New(fmt.Sprintf("Error while adding program with major id: %d: %s", program.MajorId, err.Error()))
	}

	return id, nil
}

func NewProgramYear(programYear models.ProgramYear) (int, error) {
	pool := database.GetPool()

	row := pool.QueryRow(
		context.Background(),
		`
			INSERT INTO program_year (program_id, year_number)
			VALUES($1, $2)
			RETURNING id;
		`,
		programYear.ProgramId,
		programYear.YearNumber,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, errors.New(fmt.Sprintf("Error while adding program year with program id: %d: %s", programYear.ProgramId, err.Error()))
	}

	return id, nil
}

func NewSemester(semester models.Semester) (int, error) {
	pool := database.GetPool()

	row := pool.QueryRow(
		context.Background(),
		`
			INSERT INTO semester (year_id, semester_number)
			VALUES($1, $2)
			RETURNING id;
		`,
		semester.YearId,
		semester.SemesterNumber,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, errors.New(fmt.Sprintf("Error while adding semester with year id: %d: %s", semester.YearId, err.Error()))
	}

	return id, nil
}
