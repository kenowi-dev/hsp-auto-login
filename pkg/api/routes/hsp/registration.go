package hsp

import (
	"github.com/kenowi-dev/hspscraper"
	"log"
	"sync"
	"time"
)

type Registration struct {
	Sport        string
	CourseNumber string
	Date         time.Time
	Email        string
	Password     string
}

type registrationKey struct {
	sport        string
	courseNumber string
}

type RegistrationHandler interface {
	AddRegistration(registration *Registration) error
	GetRegistrationsByUser(email, password string) []*Registration
	GetAllRegistrations() []*Registration
	Close()
}

type registrationHandler struct {
	db            DB[*Registration]
	regLock       sync.Mutex
	registrations map[registrationKey][]*Registration
	close         chan<- bool
}

func newRegistrationHandler(db DB[*Registration]) RegistrationHandler {
	regHandler := registrationHandler{
		db:            db,
		regLock:       sync.Mutex{},
		registrations: make(map[registrationKey][]*Registration),
	}

	registrations := db.GetAll()
	for _, registration := range registrations {
		regHandler.addRegistration(registration)
	}
	regHandler.close = regHandler.startObserver()
	return &regHandler
}

func (rh *registrationHandler) AddRegistration(registration *Registration) error {

	rh.addRegistration(registration)
	return rh.db.Save(registration)
}

func (rh *registrationHandler) Close() {
	rh.close <- true
}

func (rh *registrationHandler) GetRegistrationsByUser(email, password string) []*Registration {
	rh.regLock.Lock()
	defer rh.regLock.Unlock()
	registrations := make([]*Registration, 0)
	for _, rs := range rh.registrations {
		for _, r := range rs {
			if r.Email == email && r.Password == password {
				registrations = append(registrations, r)
			}
		}
	}
	return registrations
}

func (rh *registrationHandler) GetAllRegistrations() []*Registration {
	rh.regLock.Lock()
	defer rh.regLock.Unlock()
	registrations := make([]*Registration, 0)
	for _, v := range rh.registrations {
		registrations = append(registrations, v...)
	}
	return registrations
}

func (rh *registrationHandler) addRegistration(registration *Registration) {
	rh.regLock.Lock()
	regKey := registrationKey{
		sport:        registration.Sport,
		courseNumber: registration.CourseNumber,
	}
	var regList []*Registration
	if v, ok := rh.registrations[regKey]; ok {
		regList = v
	} else {
		regList = make([]*Registration, 0)
		rh.registrations[regKey] = regList
	}
	rh.registrations[regKey] = append(regList, registration)
	rh.regLock.Unlock()
}

func (rh *registrationHandler) startObserver() chan<- bool {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool, 1)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				rh.observe()
			}
		}
	}()
	return quit
}

func (rh *registrationHandler) observe() {
	rh.regLock.Lock()
	// Copy Registrations to not lock for too long
	registrationsCopy := make(map[registrationKey][]*Registration, len(rh.registrations))
	for k, v := range rh.registrations {
		registrationsCopy[k] = v
	}
	rh.regLock.Unlock()

	for key, registrationList := range registrationsCopy {
		course, err := hspscraper.FindCourse(key.sport, key.courseNumber)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if course.State == hspscraper.CourseStateOpen {
			go rh.registerAll(registrationList)
		}
	}
}

func (rh *registrationHandler) registerAll(regList []*Registration) {
	for _, reg := range regList {
		go func(r *Registration) {
			err := hspscraper.Register(r.Sport, r.CourseNumber, r.Email, r.Password, r.Date)
			if err != nil {
				log.Println(err.Error())
			}
		}(reg)
	}
}
