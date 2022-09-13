package dto

import "time"

type TransactionRequest struct {
	StartDate time.Time            		`json:"startDate"`
	DueDate	  time.Time     	   		`json:"dueDate"`
	UserID	 int				 	    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Attache	  string					`json:"attache" gorm:"type:varchar(50)" validate:"required"`
	Status    string               		`json:"status" gorm:"type:varchar(50)" validate:"required"`
}

type UpdateTransactionRequest struct {
	StartDate time.Time            		`json:"startDate" `
	DueDate	  time.Time     	   		`json:"dueDate" `
	UserID	 int				 	    `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Attache	  string					`json:"attache" gorm:"type:varchar(50)" validate:"required"`
	Status    string               		`json:"status" gorm:"type:varchar(50)" validate:"required"`
}