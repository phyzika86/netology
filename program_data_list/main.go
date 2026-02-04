package main

import "fmt"

// Структура для хранения данных о сотруднике
type Employee struct {
	FirstName string
	LastName  string
	Age       int8
	Position  string
	Salary    int64
}

// Интерфейс для вывода информации
type Displayable interface {
	Display()
}

// Реализация метода Display для Employee
func (e Employee) Display() {
	fmt.Println("Имя", e.FirstName)
	fmt.Println("Фамилия", e.LastName)
	fmt.Println("Возраст", e.Age)
	fmt.Println("Должность", e.Position)
	fmt.Println("Зарплата", e.Salary)
	fmt.Println("----------------")
}

// Реализация метода AddEmployee для Employee для добавления в список сотрудников
func AddEmployee(employees []Employee, employee Employee) []Employee {
	employees = append(employees, employee)

	return employees
}

func CreateEmployee(firstName string, lastName string, age int8, position string, salary int64) Employee {
	e := Employee{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Position:  position,
		Salary:    salary,
	}
	return e
}

// Функция предикат для фильтрации
func predicate(e Employee, minAge int8, minSalary int64) bool {
	return e.Age >= minAge && e.Salary >= minSalary
}

// Функция фильтрации сотрудников по возрасту и зарплате
func FilterEmployees(employees []Employee, minAge int8, minSalary int64) []Employee {
	var filtered []Employee
	for _, emp := range employees {
		if predicate(emp, minAge, minSalary) {
			filtered = append(filtered, emp)
		}
	}
	return filtered
}

func main() {
	// Инициализация списка сотрудников
	empOne := CreateEmployee("Иван", "Петров", 20, "стажер", 15_000)
	empTwo := CreateEmployee("Николай", "Жуков", 40, "Начальник", 100_000)
	empThird := CreateEmployee("Василий", "Петров", 35, "Заместитель", 80_000)
	empFour := CreateEmployee("Владимир", "Журавлев", 50, "Продавец", 70_000)

	employees := []Employee{
		empOne, empTwo, empThird,
	}

	fmt.Println("Список сотрудников изначальный", employees)

	employees = AddEmployee(employees, empFour)

	fmt.Println("Список сотрудников обновленный", employees)

	// Параметры фильтрации
	var minAge int8 = 40
	var minSalary int64 = 60

	// Фильтрация и вывод
	filteredEmployees := FilterEmployees(employees, minAge, minSalary)
	for index, emp := range filteredEmployees {
		fmt.Println("Сотрудник номер ", index+1)
		emp.Display()
	}
}
