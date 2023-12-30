package controller

import (
	"fmt"
	"log"
	"shop/models"
)

func (c Controller) CreateUser() {
	user := getUsersInfo()

	//if !checkPhoneNumber(user.Phone) {
	//	fmt.Println("error is while input phone, format phone not correct")
	//	return
	//}

	if _, err := c.Store.UserRepo.Insert(user); err != nil {
		log.Fatalln("error is while inserting", err.Error())
		return
	}
	fmt.Println("user added")
}

func (c Controller) GetUserByID() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	user, err := c.Store.UserRepo.GetByID(id)
	if err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return
	}
	fmt.Println("user: ", user)
}

func (c Controller) GetUserList() {
	users, err := c.Store.UserRepo.GetList()
	if err != nil {
		fmt.Println("error is while getting list", err.Error())
		return
	}
	fmt.Println(users)
}

func (c Controller) UpdateUser() {
	us := getUsersInfo()
	if err := c.Store.UserRepo.Update(us); err != nil {
		fmt.Println("error is while updating data", err.Error())
		return
	}
	fmt.Println("user updated")
}

func (c Controller) DeleteUser() {
	id := ""
	fmt.Print("input id: ")
	fmt.Scan(&id)

	if err := c.Store.UserRepo.Delete(id); err != nil {
		fmt.Println("error is while deleting user", err.Error())
		return
	}
	fmt.Println("user deleted!")
}

func getUsersInfo() models.User {
	var (
		id, fname, lname, email, phone string
		cmd                            int
	)
	fmt.Print(`
					1 - create
					2 - update
`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("input id: ")
		fmt.Scan(&id)

		fmt.Print("input first name: ")
		fmt.Scan(&fname)

		fmt.Print("input last name: ")
		fmt.Scan(&lname)

		fmt.Print("input email: ")
		fmt.Scan(&email)

		fmt.Print("input phone: ")
		fmt.Scan(&phone)
	} else if cmd == 1 {
		fmt.Print("input first name: ")
		fmt.Scan(&fname)

		fmt.Print("input last name: ")
		fmt.Scan(&lname)

		fmt.Print("input email: ")
		fmt.Scan(&email)

		fmt.Print("input phone: ")
		fmt.Scan(&phone)
	} else {
		fmt.Println("not found")
	}

	if id != "" {
		return models.User{
			ID:        id,
			FirstName: fname,
			LastName:  lname,
			Email:     email,
			Phone:     phone,
		}
	}

	return models.User{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Phone:     phone,
	}
}

//func checkPhoneNumber(phone string) bool {
//	for _, r := range phone {
//		if r > '0' || r < '9' || r != '+' {
//			return false
//		}
//	}
//	return true
//}
