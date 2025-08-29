-- Majors / Programs
MAJOR (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    aim TEXT,
    opportunities TEXT[]
);

-- Types of programs (Diploma, Advanced Diploma, Bachelor)
PROGRAM (
    id SERIAL PRIMARY KEY,
    major_id INT REFERENCES MAJOR(id) ON DELETE CASCADE,
    type TEXT CHECK (type IN ('diploma','advanced_diploma','bachelor')),
    duration_years INT NOT NULL
);

-- Each year in a program
PROGRAM_YEAR (
    id SERIAL PRIMARY KEY,
    program_id INT REFERENCES PROGRAM(id) ON DELETE CASCADE,
    year_number INT NOT NULL
);

-- Each year has semesters
SEMESTER (
    id SERIAL PRIMARY KEY,
    year_id INT REFERENCES PROGRAM_YEAR(id) ON DELETE CASCADE,
    semester_number INT NOT NULL
);

-- Courses
COURSE (
    id SERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT
);

-- Junction: courses in a semester
SEMESTER_COURSE (
    semester_id INT REFERENCES SEMESTER(id) ON DELETE CASCADE,
    course_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (semester_id, course_id)
);

-- Junction: prerequisites (self-referencing many-to-many)
COURSE_PREREQUISITE (
    course_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    prereq_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, prereq_id)
);

-- Junction: electives
COURSE_ELECTIVE (
    course_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    elective_id INT REFERENCES COURSE(id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, elective_id)
);
