package models

type User struct {
        // gorm.Model
        ID uint          `json:"id" binding:"required" gorm:"column:id"`
        Name string     `json:"name" binding:"required" gorm:"cloumn:name"`
        Email string    `json:"email" binding:"required" gorm:"column:email"`
        // Created_at time.Time
}
