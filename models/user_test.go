package models

import (
	"fmt"
	"testing"
	"time"
)

func testingUserService() (*UserService, error)  {

	const (
		host = "localhost"
		port = 5434
		user = "admin"
		password = "123456"
		dbname = "lenslocked_test"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable", host, port, password, user, dbname)
	service, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}
	service.db.Migrator().DropTable(&User{})
	service.db.AutoMigrate(&User{})
	return service, nil
}

func seedUser(userService *UserService) (*User, error) {
	user := User{
		Name:  "Thizao",
		Email: "thiago@studydrive.net",
	}

	if err := userService.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func TestCreate(t *testing.T) {
	userService, err := testingUserService()
	if err != nil{
	    t.Fatal(err)
	}
	user := User{
		Name:  "Thizao",
		Email: "thiago@studydrive.net",
	}

	if err := userService.Create(&user); err != nil {
		t.Fatal(err)
	}

	if user.ID == 0{
		t.Errorf("Expected ID > 0. Received %d", user.ID)
	}

	if time.Since(user.CreatedAt) > time.Duration(5 * time.Second){
		t.Errorf("Expected CreatedAt to be recent. Received %s",user.CreatedAt)
	}
}

func TestUpdate(t *testing.T) {
	userService, err := testingUserService()
	if err != nil{
		t.Fatal(err)
	}
	user, err := seedUser(userService)
	if err != nil{
	    panic(err)
	}

	user.Name = "Baclava"
	user.Email = "fake@fake.com"

	err = userService.Update(user)

	if err != nil{
	    t.Fatal("Failed to update the User", err)
	}

	user, err = userService.ByID(user.ID)

	if err != nil{
		t.Fatal("Failed to find the newly updated User", err)
	}

	if user.ID == 0{
		t.Errorf("Expected ID > 0. Received %d", user.ID)
	}

	if time.Since(user.UpdatedAt) > time.Duration(5 * time.Second){
		t.Errorf("Expected UpdatedAt to be recent. Received %s",user.UpdatedAt)
	}

	if !user.UpdatedAt.After(user.CreatedAt) {
		t.Errorf("Expected UpdatedAt to be more recent than CreatedAt.")
	}
}