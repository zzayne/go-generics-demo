package main

import "fmt"

type user interface {
	userBase | userDetail
}

type userBase struct {
	ID   string
	Name string
}
type userDetail struct {
	UserID   string
	Name     string
	Age      int
	Location string
}

// different with interface{}, we can use type constraint user to limit param type
// reference:https://groups.google.com/g/golang-nuts/c/iAD0NBz3DYw/m/VcXSK55XAwAJ?pli=1

func getUserID[T user](u T) string {
	switch v := any(u).(type) {
	case userBase:
		return v.ID
	case userDetail:
		return v.UserID
	default:
		return ""
	}
}

func main() {
	fmt.Println(getUserID(userBase{
		ID:   "111",
		Name: "Jack",
	}))
	fmt.Println(getUserID(userDetail{
		UserID:   "222",
		Name:     "Tom",
		Age:      20,
		Location: "us",
	}))
}
