package payload

type Deacon struct {
	FirstName  string        `json:"first_name" binding:"required"`
	LastName   string        `json:"last_name" binding:"required"`
	Address    string        `json:"address"`
	Email      string        `json:"email" binding:"omitempty,email"`
	Phone      string        `json:"phone" binding:"omitempty,e164"`
	DOB        string        `json:"dob" `
	Country    string        `json:"country" binding:"required"`
	DeaconRank int32         `json:"deacon_rank" binding:"required"`
	Skills     []DeaconSkill `json:"skills"`
}

type DeaconSkill struct {
	SkillID int32 `json:"skill_id" binding:"required"`
	Score   int32 `json:"score" binding:"required"`
}
