package hsp

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/db"
	"github.com/kenowi-dev/hspscraper"
	"github.com/kenowi-dev/observer"
	"log"
	"sync"
	"time"
)

type HSP interface {
	AutoLogin(data *Data) error
	GetRegistrationListeners() []*Data
}

type hsp struct {
	mutex     sync.Mutex
	observers map[string]observer.RunningObserver
	db        db.DB[Data]
}

type Data struct {
	Sport        string
	CourseNumber string
	Email        string
	Password     string
}

func New(db db.DB[Data]) HSP {
	return &hsp{
		mutex:     sync.Mutex{},
		db:        db,
		observers: make(map[string]observer.RunningObserver),
	}
}

func (h *hsp) AutoLogin(data *Data) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	err := h.db.Save(data)

	course, err := hspscraper.FindCourse(data.Sport, data.CourseNumber)
	if err != nil {
		return err
	}

	obs := h.getCourseObserver(course)
	obs.AddCallback(func(i int) {
		err := hspscraper.Register(course, data.Email, data.Password)
		if err != nil {
			log.Println(err.Error())
		}
		obs.Unregister(i)
	})

	return nil
}

func (h *hsp) GetRegistrationListeners() []*Data {
	var cp []*Data
	h.mutex.Lock()
	defer h.mutex.Unlock()
	copy(cp, h.db.GetAll())

	return cp
}

func (h *hsp) getCourseObserver(course *hspscraper.Course) observer.RunningObserver {
	key := course.Sport + course.CourseNumber
	if v, ok := h.observers[key]; ok {
		return v
	} else {
		obs := h.newCourseObserver(course)
		h.observers[key] = obs
		return obs
	}
}

func (h *hsp) register(data *Data) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		course, err := hspscraper.FindCourse(data.Sport, data.CourseNumber)
		if err != nil {
			return err
		}
		if course.IsOpen {
			return hspscraper.Register(course, data.Email, data.Password)
		}
	}
}

func (h *hsp) newCourseObserver(course *hspscraper.Course) observer.RunningObserver {
	return observer.NewIntervalObs(func() bool {
		course, err := hspscraper.FindCourse(course.Sport, course.CourseNumber)
		if err != nil {
			return false
		}
		return course.IsOpen
	}, 5*time.Second).PeriodicAsync()
}
