# degree-progress-tracker
Full-stack web application that gamifies the pursuit of degree progress by mapping completed courses to curriculum milestones.


## Requirements & Planning:

### Vision

Create an engaging web-based solution that enables UTAS students to track their academic progress, while leveraging gamification to promote continuous learning, motivation, and improvement.

### User Stories (Initial Backlog Items):

#### Progress Tracking & Planning

- As a student, I want to see my GPA/grades for completed courses, so that I understand my academic standing. 💫

- As a student, I want to see a visual progress bar or dashboard, so that I can quickly understand how far I am toward completing my degree. 💫

- As a student, I want to plan my future semesters (e.g., select courses tentatively), so that I can stay on track for graduation.


#### Gamification & Motivation

- As a student, I want to earn badges/achievements when I finish milestones (e.g., first semester, 50% completion), so that I feel rewarded.

- As a student, I want to compare my progress with friends/classmates (optional leaderboard), so that I feel motivated through healthy competition.


#### Notifications & Guidance

- As a student, I want to be notified when I am eligible to take a new course, so that I don’t miss opportunities.

- As a student, I want to be warned when I am falling behind in course completion, so that I can adjust my study plan.

- As a student, I want recommendations for elective courses based on my progress, so that I can make informed choices.


#### Administrative / System Support

- As a system, I want to be able to lightly scrap specialization program diploma -> adv. diploma -> bachelor (one request each sepecialization year), so that I can take it as a reference for student degree tracking. 💫

- As an administrator, I want to update degree requirements easily, so that the system always reflects the latest curriculum.

- As a system, I want to verify prerequisite completion automatically, so that students cannot enroll incorrectly.



### Deliverables:

- Product Vision Statement

- Initial Product Backlog (user stories)

- Prioritization criteria (to be refined in Sprint Planning)



## Backlog Items refinement (prioritization)

- Scrap specialization program diploma -> adv. diploma -> bachelor (one request each sepecialization year) using colly golang external package.
- Use unidoc/unipdf golang package to read PDF files and calculate GPA/grades for completed courses.
- Tree-like network of courses that are completed, remaining, unable to take yet because of prerequisite not being taken + info description of each course (if available).


