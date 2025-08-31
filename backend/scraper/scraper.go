package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type course struct {
	Code               string
	Title              string
	Description        string
	PrerequisiteCourse []string
	isElective         bool
}

func ScrapeMajors() {
	var c *colly.Collector = colly.NewCollector()

	sites := []string{
		"https://www.utas.edu.om/ccis/Programs/Software-Engineering/Bachelor-in-Software-Engineering",
		"https://www.utas.edu.om/ccis/Programs/Information-System/Bachelor-in-Information-Systems",
		"https://www.utas.edu.om/ccis/Programs/Cyber-and-Information-Security/Bachelor-in-Cyber-and-Information-Security",
		"https://www.utas.edu.om/ccis/Programs/Network-Computing/Bachelor-in-Network-computing",
		"https://www.utas.edu.om/ccis/Programs/Data-Science-and-Artificial-Intelligence/Bachelor-Data-Science-and-Artificial-Intelligence",
	}

	c.OnHTML("body", func(h *colly.HTMLElement) {

		majorTitle := h.DOM.Find(".program-header h1").Text()
		majorAim := h.DOM.Find("#aims p").Eq(1).Text()
		var majorOpportunities []string
		h.DOM.Find("#career ul").Children().Each(func(i int, s *goquery.Selection) {
			if career := s.Has("i"); career != nil {
				majorOpportunities = append(majorOpportunities, strings.TrimSpace(career.Text()))
			}
		})
		log.Println(majorTitle)
		log.Println(majorAim)
		log.Println(majorOpportunities)

		// **** supporting bachelor only for now. ****
		// programType := "bachelor"
		programDuration := 4 // years

		for i := 1; i <= programDuration; i++ {
			year := fmt.Sprintf("year%d", i)
			h.DOM.Find(fmt.Sprintf("#%s div", year)).
				First().
				Children().
				Each(func(i int, s *goquery.Selection) {
					semester := i + 1
					var courses []course

					elective := false
					s.Children().Eq(1).Children().First().Children().Each(func(k int, s *goquery.Selection) {
						if (k+1)%2 != 0 {
							// course title
							courseHeader := s.Find("h4 a").Text()

							re := regexp.MustCompile(`(?i)elective`)
							if isMatched := re.Match([]byte(courseHeader)); isMatched {
								log.Println(courseHeader)
								elective = true
								return
							}

							courseCode := strings.TrimSpace(strings.Split(courseHeader, " - ")[0])
							courseTitle := strings.TrimSpace(strings.Join(strings.Split(courseHeader, " - ")[1:], " - "))

							courses = append(courses, course{Code: courseCode, Title: courseTitle})
						} else {
							if elective {
								ulEle := s.Find("div ul li")
								if ulEle.Length() > 0 {
									ulEle.Each(func(i int, s *goquery.Selection) {
										courseCode := strings.TrimSpace(strings.Split(s.Text(), "-")[0])
										courseTitle := strings.TrimSpace(strings.Split(s.Text(), "-")[1])

										electiveCourse := course{
											Code:       courseCode,
											Title:      courseTitle,
											isElective: true,
										}

										courses = append(courses, electiveCourse)
									})

								} else {
									courseDescription := s.Find("div p").First().Text()

									// go through all elective courses in this semester..
									for _, electiveItemLine := range strings.Split(courseDescription, "\n") {
										if len(strings.TrimSpace(electiveItemLine)) < 1 {
											continue
										}
										electiveItemData := strings.Split(electiveItemLine, "-")
										courseCode := strings.TrimSpace(electiveItemData[0])
										courseTitle := strings.TrimSpace(electiveItemData[1])
										var prerequisites []string
										var prereqs string
										fmt.Sscanf(electiveItemData[2], "Prerequisite:%s", &prereqs)

										prerequisites = append(prerequisites, strings.Split(prereqs, "+")...)

										for j, prerequisite := range prerequisites {
											prerequisites[j] = strings.TrimSpace(prerequisite)
										}

										electiveCourse := course{
											Code:               courseCode,
											Title:              courseTitle,
											PrerequisiteCourse: prerequisites,
											isElective:         true,
										}

										log.Println(electiveCourse)

										courses = append(courses, electiveCourse)
									}
								}

								elective = false
							} else {
								// course description
								courseDescription := s.Find("div p").First().Text()

								var prerequisites []string
								if strings.HasPrefix(courseDescription, "Prerequisite:") {

									splittedDescription := strings.Split(courseDescription, "\n")

									// remove the prerequisite prefix:
									fmt.Sscanf(splittedDescription[0], "Prerequisite:%s", &splittedDescription[0])

									prerequisites = append(prerequisites, strings.Split(splittedDescription[0], "+")...)

									for j, prerequisite := range prerequisites {
										prerequisites[j] = strings.TrimSpace(prerequisite)
									}

									courses[len(courses)-1].PrerequisiteCourse = prerequisites

									if len(splittedDescription) > 1 {
										courses[len(courses)-1].Description = strings.Join(splittedDescription[1:], "\n")
									}
								} else {
									courses[len(courses)-1].Description = courseDescription
								}
							}

						}
					})

					log.Println(year)
					log.Println(semester)
					log.Println(len(courses))
					log.Println("")
				})
		}

	})

	for _, site := range sites {
		c.Visit(site)

		<-time.After(time.Second * 2)
	}

}
