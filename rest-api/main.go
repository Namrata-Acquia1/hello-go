package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Employee struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Active      bool   `json:"active"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	Designation string `json:"designation"`
}

var employees = []Employee{
	{
		ID:          "464eee0b-802f-4193-8405-c5f39c1ecd72",
		Name:        "John Doe",
		Email:       "john.doe@dummy.com",
		Active:      true,
		CreatedAt:   1720430892,
		UpdatedAt:   1720430892,
		Designation: "Software developer",
	},
	{
		ID:          "c5f12d01-8113-4275-9a2a-77a1c40690f4",
		Name:        "James Bond",
		Email:       "bond007@dummy.com",
		Active:      true,
		CreatedAt:   1720430892,
		UpdatedAt:   1720430892,
		Designation: "Software developer",
	},
}

// GET all employee function
func getAllEmployees(context *gin.Context) {
	context.JSON(http.StatusOK, employees)
}

// GET employee by id function
func getEmployeeByID(context *gin.Context) {
	id := context.Param("id")
	for _, emp := range employees {
		if emp.ID == id {
			context.JSON(http.StatusOK, emp)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// POST employee function
func postEmployees(context *gin.Context) {
	var newEmployee Employee

	if err := context.BindJSON(&newEmployee); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEmployee.ID = uuid.New().String()

	for _, emp := range employees {
		if emp.ID == newEmployee.ID {
			context.JSON(http.StatusConflict, gin.H{"error": "Employee with same ID already exists"})
			return
		}
	}

	currentTime := getCurrentTimestamp()
	newEmployee.CreatedAt = currentTime
	newEmployee.UpdatedAt = currentTime

	employees = append(employees, newEmployee)
	context.JSON(http.StatusCreated, employees)
}

// PATCH employee data function
func patchEmployee(context *gin.Context) {
	id := context.Param("id")
	var patchedEmployee Employee

	if err := context.BindJSON(&patchedEmployee); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			if emp.Active == false {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Employee is not active and cannot be updated"})
				return
			}

			if patchedEmployee.Name != "" {
				employees[i].Name = patchedEmployee.Name
			}
			if patchedEmployee.Email != "" {
				employees[i].Email = patchedEmployee.Email
			}
			if patchedEmployee.Designation != "" {
				employees[i].Designation = patchedEmployee.Designation
			}

			employees[i].UpdatedAt = getCurrentTimestamp()

			context.JSON(http.StatusOK, employees[i])
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// PUT employee data function
func putEmployee(context *gin.Context) {
	id := context.Param("id")
	var updatedEmployee Employee

	if err := context.BindJSON(&updatedEmployee); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			if emp.Active == false {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Employee is not active and cannot be updated"})
				return
			}

			updatedEmployee.ID = id // ID preserved
			employees[i] = updatedEmployee
			employees[i].UpdatedAt = getCurrentTimestamp()

			context.JSON(http.StatusOK, employees[i])
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

// DELETE employee data function
func deleteEmployee(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			if emp.Active == true {
				context.JSON(http.StatusBadRequest, gin.H{"error": "Active employee cannot be deleted"})
				return
			}

			employees = append(employees[:i], employees[i+1:]...)

			context.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func main() {
	router := gin.Default()
	//GET all employees
	router.GET("/api/employees", getAllEmployees)
	//GET employee using id
	router.GET("/api/employees/:id", getEmployeeByID)
	//POST employee
	router.POST("/api/employees", postEmployees)
	//PATCH employee
	router.PATCH("/api/employees/:id", patchEmployee)
	//PUT employee
	router.PUT("/api/employees/:id", putEmployee)
	//DELETE employee
	router.DELETE("/api/employees", deleteEmployee)
	router.Run("localhost:9090")
}
