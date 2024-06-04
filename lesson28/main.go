package main

import (
	"fmt"
	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	// db, err := postgres.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// st := postgres.NewStudentRepo(db)

	// users, err := st.GetAllStudents()
	// if err != nil {
	// 	panic(err)
	// }

	// user, _ := st.GetByID(users[2].ID)

	// fmt.Println(users)

	// fmt.Println(user)

	// cr := postgres.NewCourseRepo(db)
	// _ = cr.Create(&model.Course{})

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	cr := postgres.NewCourseRepo(db)

	defer db.Close()
	
	for{
		fmt.Println("What do you want to do?")
		fmt.Print(`[1]-> add course
[2]-> show all cocurse
[3]-> get by id
[4]-> update course
[5]-> delete course
[6]-> add student
[7]-> show all students
[8]-> get by id student
[9]-> update student
[10]-> delete student
[0]-> exit : `)
		var a int
		fmt.Scan(&a)
		switch a{
		case 1:
			fmt.Println("Add course")

			course :=  model.Course{}

			fmt.Print("Enter name for new course : ")
			fmt.Scan(&course.Name)

			fmt.Print("Enter field for new course : ")
			fmt.Scan(&course.Field)

			err = cr.Create(course)
			if err != nil{
				panic(err)
			}
			fmt.Println()
			fmt.Println("COURSE ADDED")
			fmt.Println()
		case 2:
			fmt.Println("Show all course")
			fmt.Println()

			courses, err := cr.ShowAllCourse()
			if err != nil{
				panic(err)
			} 
			for _, i := range courses{
				fmt.Printf("%+v\n", i)
			}
			fmt.Println()
		case 3:
			fmt.Println("Get by Id")

			var id string
			fmt.Print("Enter Id : ")
			fmt.Scan(&id)

			course, err := cr.GetById(id)
			if err != nil{
				panic(err)
			}
			fmt.Println()
			fmt.Printf("%+v\n", course)
			fmt.Println()
		case 4:
			fmt.Println("Update course")

			course :=  model.Course{}

			fmt.Print("Enter id course which you want update : ")
			fmt.Scan(&course.Id)

			fmt.Print("Enter new name  course : ")
			fmt.Scan(&course.Name)

			fmt.Print("Enter new field course : ")
			fmt.Scan(&course.Field)

			err = cr.Update(course)
			if err != nil{
				panic(err)
			}
			fmt.Println()
		case 5:
			fmt.Println("Delete course")

			var id string
			fmt.Print("Enter Id : ")
			fmt.Scan(&id)

			err = cr.Delete(id)
			if err != nil {
				panic(err)
			}
			fmt.Println()
			fmt.Println("COURSE DELETED")
			fmt.Println()		
		case 6:
			fmt.Println("Add student")
			fmt.Println()

			student := model.Student{}
			fmt.Print("Enter name : ")
			fmt.Scan(&student.Name)

			fmt.Print("Enter age : ")
			fmt.Scan(&student.Age)

			fmt.Print("Enter gender m or f : ")
			fmt.Scan(&student.Gender)

			fmt.Print("Enter course id : ")
			fmt.Scan(&student.CourseId)

			err = st.Create(student)
			if err != nil{
				panic(err)
			}
			fmt.Println()
			fmt.Println("STUDENT ADDED")
			fmt.Println()
		case 7:
			fmt.Println("Show all students")
			fmt.Println()

			students, err := st.GetAllStudents()
			if err != nil{
				panic(err)
			} 
			for _, i := range students{
				fmt.Printf("%+v\n", i)
			}
			fmt.Println()
		case 8:
			fmt.Println("Get by Id")

			var id string
			fmt.Print("Enter Id : ")
			fmt.Scan(&id)

			student, err := st.GetByID(id)
			if err != nil{
				panic(err)
			}
			fmt.Println()
			fmt.Printf("%+v\n", student)
			fmt.Println()
		case 9:
			fmt.Println("Update student")

			student :=  model.Student{}

			fmt.Print("Enter student id  which you want update : ")
			fmt.Scan(&student.ID)

			fmt.Print("Enter new name  student : ")
			fmt.Scan(&student.Name)

			fmt.Print("Enter new age student : ")
			fmt.Scan(&student.Age)

			fmt.Print("Enter new gender [m] or [f] student : ")
			fmt.Scan(&student.Gender)

			fmt.Print("Enter new course id student : ")
			fmt.Scan(&student.CourseId)

			err = st.Update(student)
			if err != nil{
				panic(err)
			}
			fmt.Println()
			fmt.Printf("STUDENT WITH ID %s UPDATED\n", student.ID)
			fmt.Println()
		case 10:
			fmt.Println()
			fmt.Println("Delete student")

			var id string
			fmt.Print("Enter Id : ")
			fmt.Scan(&id)

			err = st.Delete(id)
			if err != nil {
				panic(err)
			}
			fmt.Println()
			fmt.Println("STUDENT DELETED")
			fmt.Println()
		default:
			return 
		}
	}
}
