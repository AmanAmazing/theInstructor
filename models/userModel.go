package models

import (
	"time"

	"gorm.io/gorm"
)
type User struct {
    gorm.Model
    FirstName     string `gorm:"not null" json:"firstName"`
    LastName      string `gorm:"not null" json:"lastName"`
    Email         string `gorm:"unique;not null" json:"email"`
    PasswordHash  string `gorm:"not null" json:"password"`
    Role          string `gorm:"check:role IN ('instructor', 'student');not null" json:"role"`
    CreatedLessons []Lesson `gorm:"foreignkey:InstructorID" json:"createdLessons"` // Lessons created by instructors
    BookedLessons  []Booking `gorm:"foreignkey:StudentID" json:"bookedLessons"`   // Lessons booked by students
}

type Lesson struct {
    gorm.Model
    InstructorID  uint `gorm:"not null" json:"instructorID"`
    Title         string `json:"title"`
    DateAndTime   time.Time `json:"dateTime"`
    DurationMinutes int `json:"duration"`
    Location      string `json:"location"`
    Price         string `json:"price"`
}

type Booking struct {
    gorm.Model
    StudentID     uint `json:"studentID"`
    LessonID      uint `json:"lessonID"`
    PaymentStatus string `gorm:"check:payment_status IN ('pending', 'completed');not null" json:"paymentStatus"`
}
