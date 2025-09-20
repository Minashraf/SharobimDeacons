package db

import (
	"context"
	"fmt"
	"strings"
)

var allowedSortFields = map[string]bool{
	"first_name":    true,
	"date_of_birth": true,
	"country":       true,
	"rank_name":     true,
}

var allowedSortDirections = map[string]bool{
	"asc":  true,
	"desc": true,
}

func (q *Queries) ListDeacons(ctx context.Context, sorting map[string]string, deaconPage GetHistoryServiceByDeaconIdParams) ([]GetDeaconByIdRow, error) {
	sortField := strings.ToLower(sorting["Field"])
	sortDirection := strings.ToLower(sorting["Direction"])

	if !allowedSortFields[sortField] {
		sortField = "first_name"
	}

	if !allowedSortDirections[sortDirection] {
		sortDirection = "asc"
	}

	query := fmt.Sprintf(`
        SELECT d.first_name, d.last_name, d.phone_number, d.date_of_birth, d.country, r.rank_name
        FROM deacons.deacons.deacons d JOIN deacons.deacons.deacon_ranks r ON d.deacon_rank_id = r.id
        ORDER BY %s %s
        LIMIT $1 OFFSET $2
    `, sortField, sortDirection)

	rows, err := q.db.QueryContext(ctx, query, deaconPage.Limit, deaconPage.Offset)
	if err != nil {
		return nil, err
	}

	var deacons []GetDeaconByIdRow
	for rows.Next() {
		var d GetDeaconByIdRow
		if err := rows.Scan(&d.FirstName, &d.LastName, &d.PhoneNumber, &d.DateOfBirth, &d.Country, &d.RankName); err != nil {
			return nil, err
		}
		deacons = append(deacons, d)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return deacons, nil
}
