package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
}

// Create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get a user
func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get users
func GetUsers(db *gorm.DB, User *User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update a user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// Delete a user
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
