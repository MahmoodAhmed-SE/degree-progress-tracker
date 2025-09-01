package models

import (
	"fmt"
	"log"
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
}

func Scraper() {

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

					s.Children().Eq(1).Children().First().Children().Each(func(k int, s *goquery.Selection) {
						if (k+1)%2 != 0 {
							// course title
							courseHeader := s.Find("h4 a").Text()
							courseCode := strings.TrimSpace(strings.Split(courseHeader, " - ")[0])
							courseTitle := strings.TrimSpace(strings.Join(strings.Split(courseHeader, " - ")[1:], " - "))

							courses = append(courses, course{Code: courseCode, Title: courseTitle})
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
