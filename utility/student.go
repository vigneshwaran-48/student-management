package utility

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DATABASE = "student.txt"

type Student struct {
	Id          int32
	Name, Email string
	Age         int8
}

func AddStudent(student Student) {
	id := getStudentId()
	student.Id = id

	file, err := os.OpenFile(DATABASE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	record := convertStudentDTO(student)

	_, err = file.WriteString(record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Added student!")
}

func GetAllStudent() []Student {
	file, err := os.OpenFile(DATABASE, os.O_RDONLY|os.O_CREATE, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var students []Student

	for scanner.Scan() {
		recordText := scanner.Text()
		if len(recordText) <= 0 {
			continue
		}
		record := strings.Split(recordText, ",")
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		age, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
		}
		student := Student{
			Id:    int32(id),
			Name:  record[1],
			Email: record[2],
			Age:   int8(age),
		}
		students = append(students, student)
	}
	return students
}

func DeleteStudent(email string) {
	students := GetAllStudent()
	if len(students) == 0 {
		return
	}
	var studentsContent strings.Builder

	for _, student := range students {
		if student.Email == email {
			continue
		}
		record := convertStudentDTO(student)
		studentsContent.WriteString(record)
	}

	file, err := os.Create(DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(studentsContent.String())
}

func UpdateStudent(student Student) {
	students := GetAllStudent()
	if len(students) == 0 {
		return
	}
	studentExists := false
	var studentsContent strings.Builder
	for _, s := range students {
		if s.Id == student.Id || s.Email == student.Email {
			studentExists = true
			studentsContent.WriteString(convertStudentDTO(student))
			continue
		}
		studentsContent.WriteString(convertStudentDTO(s))
	}
	if !studentExists {
		log.Fatal("Student not exists to update!")
	}

	file, err := os.Create(DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(studentsContent.String())
}

func getStudentId() int32 {
	students := GetAllStudent()
	if len(students) == 0 {
		return 1
	}
	return students[len(students)-1].Id + 1
}

func convertStudentDTO(student Student) string {
	var record strings.Builder
	record.WriteString(fmt.Sprintf("%d", student.Id))
	record.WriteString(",")
	record.WriteString(student.Name)
	record.WriteString(",")
	record.WriteString(student.Email)
	record.WriteString(",")
	record.WriteString(fmt.Sprintf("%d", student.Age))
	record.WriteString("\n")
	return record.String()
}
