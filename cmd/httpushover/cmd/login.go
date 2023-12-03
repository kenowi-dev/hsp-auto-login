package cmd

import (
	"github.com/kenowi-dev/hspscraper"
	"github.com/spf13/cobra"
	"log"
	"time"
)

type LoginConfig struct {
	sport    string
	courseId string
	email    string
	password string
	retry    int
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to HSP-Course",
	Long:  "Automatically login to a given Hochschulsport Course",
	Args:  cobra.NoArgs,
	Run:   runLogin,
}
var loginConfig LoginConfig

func init() {

	flags := loginCmd.Flags()

	flags.StringVarP(&loginConfig.sport, "sport", "s", "", "The name of the sport. Can also be the url to the sport.")
	flags.StringVarP(&loginConfig.courseId, "courseId", "c", "", "The course number")
	flags.StringVarP(&loginConfig.email, "email", "e", "", "The email for the login")
	flags.StringVarP(&loginConfig.password, "password", "p", "", "The password for the login")
	flags.IntVarP(&loginConfig.retry, "retry", "r", -1, "Interval between looking if the course is open. Will not retry if omitted.")

	_ = loginCmd.MarkFlagRequired("sport")
	_ = loginCmd.MarkFlagRequired("courseId")
	_ = loginCmd.MarkFlagRequired("email")
	_ = loginCmd.MarkFlagRequired("password")

}

func runLogin(_ *cobra.Command, _ []string) {

	course, err := hspscraper.FindCourse(loginConfig.sport, loginConfig.courseId)
	if err != nil {
		return
	}

	err = hspscraper.Register(course, loginConfig.email, loginConfig.password)
	if err == nil {
		log.Println("Registered")
		return
	}

	log.Println(err.Error())
	if loginConfig.retry > 0 {
		err = retryLogin(course, loginConfig.email, loginConfig.password)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func retryLogin(course *hspscraper.Course, email string, password string) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		open, err := hspscraper.IsCourseOpen(course)
		if err != nil {
			return err
		}
		if open {
			return hspscraper.Register(course, email, password)
		}
	}
}
