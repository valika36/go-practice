package Models

import (
	"github.com/valikak/go-practice/student/Config"
)

func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}

func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudent(student *Student, id string) (err error) {
	if err := Config.DB.Save(student).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(student *Student, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).Delete(student).Error; err != nil {
		return err
	}
	return nil
}
