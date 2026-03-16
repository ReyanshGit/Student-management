package models

import "gorm.io/gorm"

// Student — ek student ki details
type Student struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Phone   string `json:"phone"`
	Course  string `json:"course"`
	Age     int    `json:"age"`
}

// StudentInput — student add karte waqt jo data aata hai
type StudentInput struct {
	Name   string `json:"name"   binding:"required"`
	Email  string `json:"email"  binding:"required"`
	Phone  string `json:"phone"  binding:"required"`
	Course string `json:"course" binding:"required"`
	Age    int    `json:"age"    binding:"required"`
}

/**
```

---

**Samjho kya likha:**
```
Student struct  → database table banega
gorm.Model      → ID, CreatedAt automatic
binding:required → field khali nahi hona chahiye
json:"name"     → JSON mein "name" key hogi
gorm:"unique"   → ek email ek baar

*/