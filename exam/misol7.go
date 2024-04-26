package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// type Student struct {
// 	ID      string   `json:"id"`
// 	Name    string   `json:"name"`
// 	Courses []Course `json:"courses"`
// 	Address Address  `json:"address"`
// }
// type Course struct {
// 	CourseID  string `json:"courseId"`
// 	Grade     string `json:"grade"`
// 	Professor string `json:"professor"`
// }
// type Address struct {
// 	Street  string `json:"street"`
// 	City    string `json:"city"`
// 	Country string `json:"country"`
// }

// func main() {
// 	// Bularni write qilganda ochib qo'ying write qilib olgandan keyin berkitib qo'ying :)

// 	azizbek, anwar, kamal := makeStudents()
// 	students := []Student{azizbek, anwar, kamal}
// 	writeStudentsToFile(students, "students.json")

// 	// -----------------------------------------------------

// 	// Bularni read qilganda ochib qo'ying :)
// 	// students, err := readStudentsFromFile("students.json")
// 	// if err != nil{ panic(err) }

// 	// printStudentDetails(students)
// 	// fmt.Println()

// 	// top, _ := findTopScoringStudent(students)
// 	// fmt.Printf("Eng ko'p bali bor student - ID:%s Ismi: %s, Bali: %d\n", top.ID, top.Name, top.getAllGrade())

// 	// course101Students, err := groupStudentsByCategory(students, "101")
// 	// if err != nil{ panic(course101Students) }

// 	// fmt.Println()
// 	// fmt.Println("101 kursda o'qiydigan studentlar: ")
// 	// printStudentDetails(course101Students["101"])
// }


// func makeStudents() (Student, Student, Student){
// 	student1 := Student{
// 		ID:   "1",
// 		Name: "Azizbek",
// 		Courses: []Course{
// 			Course{
// 				CourseID:  "101",
// 				Grade:     "11",
// 				Professor: "Dr.Black",
// 			},
// 			Course{
// 				CourseID:  "105",
// 				Grade:     "1",
// 				Professor: "Pr.Smith",
// 			},
// 		},
// 		Address: Address{
// 			Street:  "Binokor",
// 			City:    "Andijan",
// 			Country: "Uzbekistan",
// 		},
// 	}
// 	student2 := Student{
// 		ID:   "2",
// 		Name: "Anwar",
// 		Courses: []Course{
// 			Course{
// 				CourseID:  "101",
// 				Grade:     "5",
// 				Professor: "Dr.Adam",
// 			},
// 			Course{
// 				CourseID:  "105",
// 				Grade:     "8",
// 				Professor: "Pr.John",
// 			},
// 		},
// 		Address: Address{
// 			Street:  "Uzbekistan",
// 			City:    "Andijan",
// 			Country: "Uzbekistan",
// 		},
// 	}
// 	student3 := Student{
// 		ID: "3",
// 		Name: "Kamal",
// 		Courses: []Course{
// 			Course{
// 				CourseID: "101",
// 				Grade: "6",
// 				Professor: "Dr.Clean",
// 			},
// 			Course{
// 				CourseID: "105",
// 				Grade: "9",
// 				Professor: "Pr.Peter",
// 			},
// 		},
// 		Address: Address{
// 			Street: "Oqtepa",
// 			City: "Tashkent",
// 			Country: "Uzbekistan",
// 		},
// 	}

// 	return student1, student2, student3
// }

// func writeStudentsToFile(students []Student, filename string) error{
// 	file, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)
// 	if err != nil{
// 		return err
// 	}
// 	err = json.NewEncoder(file).Encode(students)
// 	if err != nil{
// 		return err
// 	}

// 	fmt.Println("JSON yozildi!")
// 	return nil
// }

// func readStudentsFromFile(filename string) ([]Student, error){
// 	students := []Student{}

// 	file, err := os.Open(filename)
// 	if err != nil{
// 		return students, err
// 	}
// 	err = json.NewDecoder(file).Decode(&students)
// 	if err != nil {
// 		return students, err
// 	}
// 	return students, nil
// }

// func printStudentDetails(students []Student){
// 	fmt.Println("ID    Ism           Courselar soni   Address")
// 	for _, v := range students{
// 		fmt.Printf("%s     %s            %d             %s, %s, %s\n", v.ID, v.Name, len(v.Courses), v.Address.Country, v.Address.City, v.Address.Street)
// 	}
// }

// func findTopScoringStudent(students []Student) (Student, error){
// 	maxOne := students[0]
// 	for _, student := range students{
// 		if student.getAllGrade() > maxOne.getAllGrade(){
// 			maxOne = student
// 		}
// 	}
// 	return maxOne, nil
// }

// func (student *Student) getAllGrade() int{
// 	allGrade := 0
// 	for _, course := range student.Courses{
// 		grade, _ := strconv.Atoi(course.Grade)
// 		allGrade += grade
// 	}
// 	return allGrade
// }

// func groupStudentsByCategory(students []Student, category string)(map[string][]Student, error){
// 	sortedStudents := []Student{}
// 	categoryID := category
	
// 	for _, student := range students{
// 		for _, course := range student.Courses{
// 			if course.CourseID == category{
// 				sortedStudents = append(sortedStudents, student)
// 			}
// 		}
// 	}

// 	result := map[string][]Student{categoryID: sortedStudents}
// 	return result, nil
// }

// Done