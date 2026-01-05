package payload

type Attendance struct {
	ESLId int32  `json:"esl_id" binding:"required"`
	Date  string `json:"date" binding:"required"`
}
