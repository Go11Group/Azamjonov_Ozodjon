package psql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Id, First_name, Last_name, Gender, Nation, Field, Parent_name, City string
	Age                                                                 int
}

type Filter struct {
	Age           int
	Gender        string
	Nation        string
	Field         string
	Limit, Offset int
}

type UserRepo struct {
	Db *sql.DB
}

func CreateUser(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (user *UserRepo) GetAll(f Filter) ([]User, error) {

	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)

	query := `select * from users`

	filter := ``

	if len(f.Gender) > 0 {
		params["gender"] = f.Gender
		filter += ` and gender = :gender `
	}

	if len(f.Nation) > 0 {
		params["nation"] = f.Gender
		filter += ` and nation = :nation `
	}

	if len(f.Field) > 0 {
		params["field"] = f.Gender
		filter += ` and field = :field `
	}

	if f.Age > 0 {
		params["age"] = f.Gender
		filter += ` and age = :age `
	}

	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = (f.Offset - 1) * f.Limit
		limit += ` OFFSET :offset`
	}

	if len(filter) > 0 {
		query = query + ` where true` + filter + limit
	}

	query, arr = replaceQueryParams(query, params)

	rows, err := user.Db.Query(query, arr...)

	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		user := User{}

		err = rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Gender, &user.Nation, &user.Field, &user.Parent_name, &user.City, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println(query)
	return users, err

}

func replaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
