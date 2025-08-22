# Schema Design
```sql
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

```

## If I want to get all courses in Semester 1 of Software Engineering Diploma:

```sql
SELECT c.code, c.title
FROM COURSE c
JOIN SEMESTER_COURSE sc ON sc.course_id = c.id
JOIN SEMESTER s ON s.id = sc.semester_id
JOIN PROGRAM_YEAR y ON y.id = s.year_id
JOIN PROGRAM p ON p.id = y.program_id
JOIN MAJOR m ON m.id = p.major_id
WHERE m.title = 'Software Engineering'
  AND p.type = 'diploma'
  AND y.year_number = 1
  AND s.semester_number = 1;
```