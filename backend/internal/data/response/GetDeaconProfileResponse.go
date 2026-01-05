package response

import (
	db "Backend/internal/models"
	"database/sql"
)

type GetDeaconById struct {
	ID          int64
	FirstName   string
	LastName    string
	DateOfBirth sql.NullTime
	Address     sql.NullString
	Email       sql.NullString
	PhoneNumber sql.NullString
	Country     string
	RankName    string
	Skills      []db.GetDeaconSkillByIdRow
}
