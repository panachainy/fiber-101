package build

import "log"

var (
	Time    string
	User    string
	Version = "development"
)

func PrintBuildDetail() {
	log.Println("Version:\t", Version)
	log.Println("build.Time:\t", Time)
	log.Println("build.User:\t", User)
}
