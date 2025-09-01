BEGIN;
-- Majors / Programs
CREATE TABLE IF NOT EXISTS majors (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    aim TEXT,
    opportunities JSONB DEFAULT '[]'::JSONB
);

-- Types of programs (Diploma, Advanced Diploma, Bachelor)
CREATE TABLE IF NOT EXISTS programs (
    id SERIAL PRIMARY KEY,
    major_id INT REFERENCES majors (id) ON DELETE CASCADE,
    program_type TEXT CHECK (program_type IN ('diploma', 'advancd_diploma', 'bachelor')),
    duration_years INT NOT NULL
);

-- Each year in a program
CREATE TABLE IF NOT EXISTS program_year (
    id SERIAL PRIMARY KEY,
    program_id INT REFERENCES programs (id) ON DELETE CASCADE,
    year_number INT NOT NULL
);

-- Each year has semesters
CREATE TABLE IF NOT EXISTS semesters (
    id SERIAL PRIMARY KEY,
    year_id INT REFERENCES program_year (id) ON DELETE CASCADE,
    semester_number INT NOT NULL
);

-- Courses
CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT
);

-- Junction: courses in a semester
CREATE TABLE IF NOT EXISTS semester_courses (
    semester_id INT REFERENCES semesters (id) ON DELETE CASCADE,
    course_id INT REFERENCES courses (id) ON DELETE CASCADE,
    PRIMARY KEY (semester_id, course_id)
);

-- Junction: prerequisites (self-referencing many-to-many)
CREATE TABLE IF NOT EXISTS course_prerequisite (
    course_id INT REFERENCES courses (id) ON DELETE CASCADE,
    prereq_id INT REFERENCES courses (id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, prereq_id)
);

-- Junction: electives
CREATE TABLE IF NOT EXISTS course_elective (
    course_id INT REFERENCES courses (id) ON DELETE CASCADE,
    elective_id INT REFERENCES courses (id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, elective_id)
);

CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY,
  domain VARCHAR(50) NOT NULL,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS roles_group (
  id SERIAL PRIMARY KEY,
  role_id INT  NOT NULL,
  group_id INT NOT NULL
);

COMMIT;
