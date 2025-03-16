package hspscraper

import (
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"golang.org/x/net/html"
	"strings"
	"time"
)

var xPathCourseRowTemplate = "//table[@class='bs_kurse']/tbody//td[text() = '%s']/parent::tr"

var xPathCourseRows = xpath.MustCompile("//table[@class='bs_kurse']/tbody/tr")
var xPathCourseNumber = xpath.MustCompile("//td[@class='bs_sknr']/text()")
var xPathCourseDetails = xpath.MustCompile("//td[@class='bs_sdet']/text()")
var xPathCourseDay = xpath.MustCompile("//td[@class='bs_stag']/text()")
var xPathCourseTime = xpath.MustCompile("//td[@class='bs_szeit']/text()")
var xPathCourseLocation = xpath.MustCompile("//td[@class='bs_sort']/a/text()")
var xPathCourseDate = xpath.MustCompile("//td[2]//text()")
var xPathCourseDateTime = xpath.MustCompile("//td[3]//text()")
var xPathCourseManagement = xpath.MustCompile("//td[@class='bs_skl']/text()")
var xPathCoursePrice = xpath.MustCompile("//td[@class='bs_spreis']///text()")
var xPathCourseState = xpath.MustCompile("//td[@class='bs_sbuch']/input")
var xPathCourseBookingID = xpath.MustCompile("//td[@class='bs_sbuch']/input")

var courseIDPrefix = "bs_tr"
var hspCourseDatesTemplate = "https://buchung.hochschulsport-hamburg.de/cgi/webpage.cgi?kursinfo=%s"

type CourseState string

const (
	CourseStateOpen        = "Vormerkliste"
	CourseStateWaitingList = "Warteliste"
)

type Course struct {
	Number     string      `json:"number"`
	Details    string      `json:"details"`
	Day        string      `json:"day"`
	Time       string      `json:"time"`
	Location   string      `json:"location"`
	Id         string      `json:"id"`
	Management string      `json:"management"`
	Price      string      `json:"price"`
	State      CourseState `json:"state"`
	BookingID  string      `json:"bookingId"`
}

type CourseDate struct {
	Date     time.Time      `json:"date"`
	Duration *time.Duration `json:"duration"`
	Updated  time.Time      `json:"updated"`
}

func FindCourse(sport string, courseNumber string) (*Course, error) {
	doc, err := htmlquery.LoadURL(getHspSportUrl(sport))
	if err != nil {
		return nil, err
	}

	xPathCourseRow, err := xpath.Compile(fmt.Sprintf(xPathCourseRowTemplate, courseNumber))
	if err != nil {
		return nil, err
	}

	if tr := htmlquery.QuerySelector(doc, xPathCourseRow); tr != nil {
		return parseCourseRow(tr)
	}
	return nil, errors.New("course not found")
}

func GetAllCourses(sport *Sport) ([]*Course, error) {
	return getAllCourses(sport)
}

func GetCourseDates(couse *Course) ([]CourseDate, error) {
	return parseDates(fmt.Sprintf(hspCourseDatesTemplate, couse.Id))
}

func getAllCourses(sport *Sport) ([]*Course, error) {
	doc, err := htmlquery.LoadURL(sport.Href)
	if err != nil {
		return nil, err
	}
	trs := htmlquery.QuerySelectorAll(doc, xPathCourseRows)

	var courses = make([]*Course, 0)
	for _, tr := range trs {
		course, err := parseCourseRow(tr)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func parseCourseRow(tr *html.Node) (*Course, error) {
	id := ""
	for _, attr := range tr.Attr {
		if attr.Key == "id" {
			id = strings.TrimPrefix(strings.ToUpper(attr.Val), strings.ToUpper(courseIDPrefix))
		}
	}

	number := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseNumber); n != nil {
		number = n.Data
	}
	details := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseDetails); n != nil {
		details = n.Data
	}
	day := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseDay); n != nil {
		day = n.Data
	}
	t := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseTime); n != nil {
		t = n.Data
	}
	location := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseLocation); n != nil {
		location = n.Data
	}

	management := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseManagement); n != nil {
		management = n.Data
	}
	price := ""
	if n := htmlquery.QuerySelector(tr, xPathCoursePrice); n != nil {
		price = n.Data
	}
	state := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseState); n != nil {
		state = htmlquery.SelectAttr(n, "value")
	}
	bookingID := ""
	if n := htmlquery.QuerySelector(tr, xPathCourseBookingID); n != nil {
		bookingID = htmlquery.SelectAttr(n, "name")
	}

	course := Course{
		Id:         id,
		BookingID:  bookingID,
		Number:     number,
		Details:    details,
		Day:        day,
		Time:       t,
		Location:   location,
		Management: management,
		Price:      price,
		State:      CourseState(state),
	}
	return &course, nil
}

func parseDates(href string) ([]CourseDate, error) {
	datesDoc, err := htmlquery.LoadURL(href)
	if err != nil {
		return nil, err
	}
	trs := htmlquery.QuerySelectorAll(datesDoc, xPathCourseRows)
	courseDates := make([]CourseDate, 0)
	if trs == nil {
		return nil, errors.New(fmt.Sprintf("given Link cannot be parsed: %s", href))
	}
	for _, tr := range trs {
		var date time.Time
		if n := htmlquery.QuerySelector(tr, xPathCourseDate); n != nil {
			date, err = time.Parse("02.01.2006", n.Data)
			if err != nil {
				return nil, err
			}
		}

		var duration *time.Duration = nil
		if n := htmlquery.QuerySelector(tr, xPathCourseDateTime); n != nil {
			f, t, ok := strings.Cut(n.Data, "-")
			if !ok {
				goto creation
			}
			from, err := time.Parse("15.04", f)
			if err != nil {
				goto creation
			}
			date = time.Date(date.Year(), date.Month(), date.Day(), from.Hour(), from.Minute(), 0, 0, time.UTC)
			to, err := time.Parse("15.04", t)
			if err != nil {
				goto creation
			}
			d := to.Sub(from)
			duration = &d
		}
	creation:
		courseDate := CourseDate{
			Date:     date,
			Duration: duration,
			Updated:  time.Now(),
		}
		courseDates = append(courseDates, courseDate)
	}
	return courseDates, nil
}
