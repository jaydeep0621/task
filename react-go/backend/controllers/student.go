package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/database"
	"server/models"
	"strconv"
)

func getStudentIDFromHeader(r *http.Request) (int, error) {
	StudentIDHeader := r.Header.Get("student-id")
	if StudentIDHeader == "" {
		return 0, http.ErrMissingFile
	}
	studentID, err := strconv.Atoi(StudentIDHeader)
	if err != nil {
		return 0, err
	}
	return studentID, nil
}

// GetStudents godoc
// @Summary Get all students
// @Description Get all students
// @Tags students
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {string} string "Internal Server Error"
// @Router /students [get]
func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	rows, err := database.DB.Query("SELECT id, first_name, last_name, dob, gender, address_line1, address_line2, city, contact_number, contact_number_guardian, supervisor_id, remarks, home_long, home_lat, employer_id, check_in_time, check_out_time FROM student")
	if err != nil {
		log.Printf("Error fetching students: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.FirstName, &s.LastName, &s.DOB, &s.Gender, &s.AddressLine1, &s.AddressLine2, &s.City, &s.ContactNumber, &s.ContactNumberGuardian, &s.SupervisorID, &s.Remarks, &s.HomeLong, &s.HomeLat, &s.EmployerID, &s.CheckInTime, &s.CheckOutTime); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}
	log.Printf("Fetched %d students", len(students))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// GetStudent godoc
// @Summary Get a student by ID
// @Description Get a student by ID
// @Tags students
// @Produce json
// @Param student-id header string true "Student ID"
// @Success 200 {object} models.Student
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /get-student [get]
func GetStudent(w http.ResponseWriter, r *http.Request) {
	studentID, err := getStudentIDFromHeader(r)
	if err != nil {
		log.Printf("Error extracting student-id: %v", err)
		http.Error(w, "Invalid or missing student-id header", http.StatusBadRequest)
		return
	}
	var s models.Student
	err = database.DB.QueryRow("SELECT id, first_name, last_name, dob, gender, address_line1, address_line2, city, contact_number, contact_number_guardian, supervisor_id, remarks, home_long, home_lat, employer_id, check_in_time, check_out_time FROM student WHERE id = $1", studentID).Scan(&s.ID, &s.FirstName, &s.LastName, &s.DOB, &s.Gender, &s.AddressLine1, &s.AddressLine2, &s.City, &s.ContactNumber, &s.ContactNumberGuardian, &s.SupervisorID, &s.Remarks, &s.HomeLong, &s.HomeLat, &s.EmployerID, &s.CheckInTime, &s.CheckOutTime)
	if err != nil {
		log.Printf("Error fetching student with ID %d: %v", studentID, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Printf("Fetched student with ID %d: %+v", studentID, s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student"
// @Success 201 {object} models.Student
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /students [post]
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := `INSERT INTO student (first_name, last_name, dob, gender, address_line1, address_line2, city, contact_number, contact_number_guardian, supervisor_id, remarks, home_long, home_lat, employer_id, check_in_time, check_out_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id`
	err := database.DB.QueryRow(query, s.FirstName, s.LastName, s.DOB, s.Gender, s.AddressLine1, s.AddressLine2, s.City, s.ContactNumber, s.ContactNumberGuardian, s.SupervisorID, s.Remarks, s.HomeLong, s.HomeLat, s.EmployerID, s.CheckInTime, s.CheckOutTime).Scan(&s.ID)
	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"data": s})
}

// UpdateStudent godoc
// @Summary Update a student by ID
// @Description Update a student by ID
// @Tags students
// @Accept json
// @Produce json
// @Param id path string true "Student ID"
// @Param student body models.Student true "Student"
// @Success 200 {object} models.Student
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /students/{id} [put]
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	
	// Write your code here

}

// DeleteStudent godoc
// @Summary Delete a student by ID
// @Description Delete a student by ID
// @Tags students
// @Param id path string true "Student ID"
// @Success 204 {string} string "No Content"
// @Failure 500 {string} string "Internal Server Error"
// @Router /students/{id} [delete]
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	
	// Write your code here

}
