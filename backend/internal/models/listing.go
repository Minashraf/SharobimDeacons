package db

import (
	"context"
	"fmt"
)

var allowedSortFields = map[string]bool{
	"first_name":     true,
	"date_of_birth":  true,
	"country":        true,
	"deacon_rank_id": true,
}

var allowedFilteredFields = map[string]bool{
	"country":        true,
	"deacon_rank_id": true,
}

var allowedSortDirections = map[string]bool{
	"asc":  true,
	"desc": true,
}

func (q *Queries) ListDeacons(ctx context.Context, sorting map[string]string, deaconPage GetHistoryServiceByDeaconIdParams) ([]GetDeaconByIdRow, error) {
	sortField := sorting["Field"]
	sortDirection := sorting["Direction"]
	filterField := sorting["FilterField"]
	filterValue := sorting["FilterValue"]
	if !allowedSortFields[sortField] || sortField == "first_name" {
		sortField = fmt.Sprintf("first_name %s, last_name", sortDirection)
	}

	if !allowedSortDirections[sortDirection] {
		sortDirection = "asc"
	}

	if (!allowedFilteredFields[filterField]) || (allowedFilteredFields[filterField] && filterValue == "") {
		filterField = "1"
		filterValue = "1"
	}

	query := fmt.Sprintf(`
        SELECT d.id, d.first_name, d.last_name, d.phone_number, d.date_of_birth, d.country, r.rank_name
        FROM deacons.deacons.deacons d JOIN deacons.deacons.deacon_ranks r ON d.deacon_rank_id = r.id
        WHERE %s = '%s'
        ORDER BY %s %s
        LIMIT $1 OFFSET $2
    `, filterField, filterValue, sortField, sortDirection)

	rows, err := q.db.QueryContext(ctx, query, deaconPage.Limit, deaconPage.Offset)
	if err != nil {
		return nil, err
	}

	deacons := make([]GetDeaconByIdRow, 0)
	for rows.Next() {
		var d GetDeaconByIdRow
		if err := rows.Scan(&d.ID, &d.FirstName, &d.LastName, &d.PhoneNumber, &d.DateOfBirth, &d.Country, &d.RankName); err != nil {
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
