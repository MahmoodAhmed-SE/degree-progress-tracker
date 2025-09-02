CREATE DATABASE degree_progress_db;

\connect degree_progress_db;

-- Majors / Programs
CREATE TABLE IF NOT EXISTS MAJOR (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    aim TEXT,
    opportunities TEXT[]
);

-- Types of programs (Diploma, Advanced Diploma, Bachelor)
CREATE TABLE IF NOT EXISTS PROGRAM (
    id SERIAL PRIMARY KEY,
    major_id INT REFERENCES MAJOR(id) ON DELETE CASCADE,
    type TEXT CHECK (type IN ('diploma','advanced_diploma','bachelor')),
    duration_years INT NOT NULL
);

-- Each year in a program
CREATE TABLE IF NOT EXISTS PROGRAM_YEAR (
    id SERIAL PRIMARY KEY,
    program_id INT REFERENCES PROGRAM(id) ON DELETE CASCADE,
    year_number INT NOT NULL
);

-- Each year has semesters
CREATE TABLE IF NOT EXISTS SEMESTER (
    id SERIAL PRIMARY KEY,
    year_id INT REFERENCES PROGRAM_YEAR(id) ON DELETE CASCADE,
    semester_number INT NOT NULL
);

-- Courses
CREATE TABLE IF NOT EXISTS COURSE (
    id SERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT
);

-- Junction: courses in a semester
CREATE TABLE IF NOT EXISTS SEMESTER_COURSE (
    semester_id INT REFERENCES SEMESTER(id) ON DELETE CASCADE,
    course_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (semester_id, course_id)
);

-- Junction: prerequisites (self-referencing many-to-many)

/*
    this is so that when we have course 1 and prerequistes are course 2 and course 3 we make
    2 groups with one option each.
    BUT, if we have course 1 and prerequistes are course 2 OR course 3 we make
    1 group with 2 options (2 and 3)
*/ 
CREATE TABLE COURSE_PREREQ_GROUP (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES COURSE(id) ON DELETE CASCADE
);

CREATE TABLE COURSE_PREREQ_OPTION (
    group_id INT REFERENCES COURSE_PREREQ_GROUP(id) ON DELETE CASCADE,
    prereq_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, prereq_id)
);


-- Junction: electives
CREATE TABLE IF NOT EXISTS SEMESTER_ELECTIVE (
    id SERIAL PRIMARY KEY,
    semester_id INT REFERENCES SEMESTER(id) ON DELETE CASCADE,
    title TEXT
);


CREATE TABLE IF NOT EXISTS SEMESTER_ELECTIVE_COURSE (
    semester_elective_id INT REFERENCES SEMESTER_ELECTIVE(id) ON DELETE CASCADE,
    elective_course INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (semester_elective_id, elective_course)
);
