package repository

import (
	"context"
	"effective/internal/domain"
	"fmt"
	"strconv"
	"strings"
)

func (r *repositoryPostgres) AddCar(ctx context.Context, car domain.Car) error {
	query := `
		INSERT INTO cars (reg_num,mark,model,year,owner_name,owner_surname,owner_patronymic)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.db.ExecContext(ctx, query, car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryPostgres) GetCarById(ctx context.Context, id int) (domain.Car, error) {
	var car domain.Car
	query := `
		SELECT * FROM cars 
		WHERE id = $1  
		LIMIT 1
	`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic); err != nil {
		return car, err
	}
	return car, nil
}

func (r *repositoryPostgres) GetCarsList(ctx context.Context, params domain.CarsListParams) ([]domain.Car, error) {
	cars := make([]domain.Car, 0)
	var (
		sql     string
		args    []interface{}
		where   []string
		orderBy []string
	)

	if params.RegNumber != "" {
		where = append(where, "reg_num = $")
		args = append(args, params.RegNumber)
	}
	if params.Mark != "" {
		where = append(where, "mark = $")
		args = append(args, params.Mark)
	}
	if params.Model != "" {
		where = append(where, "model = $")
		args = append(args, params.Model)
	}
	if params.Year > 0 {
		where = append(where, "year = $")
		args = append(args, params.Year)
	}
	if params.Name != "" {
		where = append(where, "owner_name = $")
		args = append(args, params.Name)
	}
	if params.Surname != "" {
		where = append(where, "owner_surname = $")
		args = append(args, params.Surname)
	}
	if params.Patronymic != "" {
		where = append(where, "owner_patronymic = $")
		args = append(args, params.Patronymic)
	}

	if len(where) > 0 {
		sql = " WHERE " + strings.Join(where, " AND ")
	} else {
		sql = ""
	}

	if params.Limit > 0 {
		sql += " LIMIT $" + strconv.Itoa(len(args)+1)
		args = append(args, params.Limit)
	}
	if params.Offset > 0 {
		sql += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, params.Offset)
	}

	for i := range args {
		sql = strings.Replace(sql, "$", fmt.Sprintf("$%d", i+1), -1)
	}

	sql = "SELECT * FROM cars" + sql + (strings.Join(orderBy, ", "))

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car domain.Car
		err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year,
			&car.Owner.Name, &car.Owner.Surname, &car.Owner.Patronymic)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *repositoryPostgres) UpdateCarInfo(ctx context.Context, carID int, params domain.CarDataUpdatingRequest) error {

	query := `
	UPDATE cars SET
	reg_num = COALESCE(NULLIF($1, ''), reg_num),
	mark = COALESCE(NULLIF($2, ''), mark),
	model = COALESCE(NULLIF($3, ''), model),
	year = COALESCE(NULLIF($4, 0), year),
	owner_name = COALESCE(NULLIF($5, ''), owner_name),
	owner_surname = COALESCE(NULLIF($6, ''), owner_surname),
	owner_patronymic = COALESCE(NULLIF($7, ''), owner_patronymic)
	WHERE id = $8
	`
	_, err := r.db.ExecContext(ctx, query,
		params.RegNumber, params.Mark, params.Model, params.Year, params.Owner.Name, params.Owner.Surname, params.Owner.Patronymic, carID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryPostgres) DeleteCarById(ctx context.Context, id int) error {
	query := `
		DELETE FROM cars WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
