BEGIN;
-- Majors / Programs
DROP TABLE IF EXISTS majors;

-- Types of programs (Diploma, Advanced Diploma, Bachelor)
DROP TABLE IF EXISTS programs;

-- Each year in a program
DROP TABLE IF EXISTS program_year;

-- Each year has semesters
DROP TABLE IF EXISTS semesters;

-- Courses
DROP TABLE IF EXISTS courses;

-- Junction: courses in a semester
DROP TABLE IF EXISTS semester_courses;

-- Junction: prerequisites (self-referencing many-to-many)
DROP TABLE IF EXISTS course_prerequisite;

-- Junction: electives
DROP TABLE IF EXISTS course_elective;

COMMIT;
