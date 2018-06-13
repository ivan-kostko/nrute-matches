package application

import "time"

type Movement struct {
	Id       string
	Type     string
	Option   string
	Date     time.Time
	Branch   Branch
	Workflow Workflow
	User     User
	Vehicle  Vehicle
}

type Branch struct {
	Id string
}

type Workflow struct {
	Id     string
	Type   string
	Factor string
}

type User struct {
	Id         string
	Contractor *string
}

type Vehicle struct {
	Id   string
	Type string
}
