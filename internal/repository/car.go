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
		where = append(where, "reg_num = $1")
		args = append(args, params.RegNumber)
	}
	if params.Mark != "" {
		where = append(where, "mark = $2")
		args = append(args, params.Mark)
	}
	if params.Model != "" {
		where = append(where, "model = $3")
		args = append(args, params.Model)
	}
	if params.Year > 0 {
		where = append(where, "year = $4")
		args = append(args, params.Year)
	}
	if params.Name != "" {
		where = append(where, "owner_name = $5")
		args = append(args, params.Name)
	}
	if params.Surname != "" {
		where = append(where, "owner_surname = $6")
		args = append(args, params.Surname)
	}
	if params.Patronymic != "" {
		where = append(where, "owner_patronymic = $7")
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
	sql := `UPDATE cars SET`

	setValues := []string{}
	args := []interface{}{}
	fmt.Println(params)
	if params.RegNumber != "" {
		setValues = append(setValues, " reg_num = $1")
		args = append(args, params.RegNumber)
	}
	if params.Mark != "" {
		setValues = append(setValues, " mark = $2")
		args = append(args, params.Mark)
	}
	if params.Model != "" {
		setValues = append(setValues, " model = $3")
		args = append(args, params.Model)
	}
	if params.Year > 0 {
		setValues = append(setValues, " year = $4")
		args = append(args, params.Year)
	}
	if params.Owner.Name != "" {
		setValues = append(setValues, " owner_name = $5")
		args = append(args, params.Owner.Name)
	}
	if params.Owner.Surname != "" {
		setValues = append(setValues, " owner_surname = $6")
		args = append(args, params.Owner.Surname)
	}
	if params.Owner.Patronymic != "" {
		setValues = append(setValues, " owner_patronymic = $7")
		args = append(args, params.Owner.Patronymic)
	}

	args = append(args, carID) // Add carID as the last argument for the WHERE clause
	sql += strings.Join(setValues, ",\n") + ` WHERE id = $` + strconv.Itoa(len(args))
	fmt.Println(sql)
	result, err := r.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated for car with ID %d", carID)
	}

	return nil
}
