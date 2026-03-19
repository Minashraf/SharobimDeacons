package payload

type BulkAttendance struct {
	ESLId           int32   `json:"esl_id" binding:"required"`
	Date            string  `json:"date" binding:"required"`
	DeaconId        []int64 `json:"deacon_id" binding:"required"`
	OverrideWarning bool    `json:"override"`
}
