package postgres

import (
	pb "auth_service/generated/genproto"
	"auth_service/token"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type UserAuthRepo struct {
	db *sql.DB
}

func NewUserAuthRepo(db *sql.DB) *UserAuthRepo { return &UserAuthRepo{db: db} }

func (r *UserAuthRepo) RegisterUser(req *pb.RegisterUserReq) (*pb.Register, error) {
	response := &pb.Register{}
	query := `
    INSERT INTO users (
        username,
        email, 
        password_hash, 
        full_name
    ) VALUES ($1, $2, $3, $4)
    RETURNING id, username, email, full_name, created_at`
	err := r.db.QueryRow(
		query,
		req.Username,
		req.Email,
		req.Password,
		req.FullName,
	).Scan(&response.Id, &response.Username, &response.Email, &response.FullName, &response.CreatedAt)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return response, nil
}

func (r *UserAuthRepo) LoginUser(req *pb.LoginReq) (*pb.Token, error) {
	var userId, passwordHash, emailDB string
	query := "SELECT id, password_hash, email FROM users WHERE email = $1"

	err := r.db.QueryRow(query, req.Email).Scan(&userId, &passwordHash, &emailDB)
	if err != nil {
		return nil, err
	}

	// Generate JWT tokens
	accessToken, err := token.GenerateJWTToken(userId, req.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.GenerateRefreshToken(userId, req.Email)
	if err != nil {
		return nil, err
	}

	// Set token expiration time
	expiresIn := int64(time.Hour.Seconds()) // assuming the token expires in 1 hour

	response := &pb.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}

	return response, nil
}

func (r *UserAuthRepo) GetUser(req *pb.ById) (*pb.Register, error) {
	query := `SELECT id, username, email, full_name, eco_points, created_at, updated_at
              FROM users WHERE id = $1`
	row := r.db.QueryRow(query, req.Id)

	res := &pb.Register{}
	err := row.Scan(&res.Id, &res.Username, &res.Email, &res.FullName, &res.EcoPoints, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (r *UserAuthRepo) UpdateUser(req *pb.UpdateUserReq) (*pb.Register, error) {
	if req.Id == "" {
		return nil, errors.New("user ID is required")
	}

	// Start building the update query
	baseQuery := `UPDATE users SET `
	var params []interface{}
	paramCount := 1

	if req.FullName != "" {
		baseQuery += fmt.Sprintf("full_name = $%d, ", paramCount)
		params = append(params, req.FullName)
		paramCount++
	}
	if req.Bio != "" {
		baseQuery += fmt.Sprintf("bio = $%d, ", paramCount)
		params = append(params, req.Bio)
		paramCount++
	}

	baseQuery += fmt.Sprintf("updated_at = $%d WHERE id = $%d AND deleted_at IS NULL RETURNING id, username, email, full_name, bio, eco_points, created_at, updated_at", paramCount, paramCount+1)
	params = append(params, time.Now(), req.Id)

	// Execute the update query
	row := r.db.QueryRow(baseQuery, params...)
	res := &pb.Register{}
	err := row.Scan(&res.Id, &res.Username, &res.Email, &res.FullName, &res.Bio, &res.EcoPoints, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

func (r *UserAuthRepo) GetAllUsers(req *pb.PageLimit) (*pb.GetAllUsersRes, error) {
	// Calculating offset for pagination
	offset := (req.Page - 1) * req.Limit

	// Query to get users with pagination
	query := `SELECT id, username, full_name, eco_points 
              FROM users 
              LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []*pb.Register

	// Iterating over the rows and scanning the results into the users slice
	for rows.Next() {
		user := &pb.Register{}
		err := rows.Scan(&user.Id, &user.Username, &user.FullName, &user.EcoPoints)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	// Checking for errors after the loop
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	// Query to get the total count of users
	var total int64
	err = r.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Creating the response
	res := &pb.GetAllUsersRes{
		Users: users,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}

	return res, nil
}

func (r *UserAuthRepo) DeleteUser(req *pb.ById) (*pb.Success, error) {
	query := `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 and deleted_at = 0`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Success{Message: "NICE DELETED"}, nil
}

func (r *UserAuthRepo) ResetPassword(req *pb.ByEmail) (*pb.Success, error) {
	query := `SELECT *
              FROM users WHERE email = $1`
	row := r.db.QueryRow(query, req.Email)

	res := &pb.Register{}
	err := row.Scan(&res.Id, &res.Username, &res.Email, &res.FullName, &res.EcoPoints, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return &pb.Success{
			Message: "no user with email found",
		}, err
	}
	return &pb.Success{
		Message: "email found successfully",
	}, nil
}

func (r *UserAuthRepo) RefreshToken(req *pb.RefreshTokenReq) (*pb.Token, error) {
	// Verify the refresh token
	claims, err := token.ExtractClaims(req.RefreshToken)
	if err != nil {
		log.Println("Invalid refresh token:", err)
		return nil, errors.New("invalid refresh token")
	}

	// Ensure claims have required fields
	userId, userIdOk := claims["user_id"].(string)
	email, emailOk := claims["email"].(string)
	if !userIdOk || !emailOk {
		return nil, errors.New("invalid token claims")
	}

	// Generate a new access token
	newAccessToken, err := token.GenerateJWTToken(userId, email)
	if err != nil {
		log.Println("Error generating new access token:", err)
		return nil, errors.New("could not generate access token")
	}

	// Optionally, generate a new refresh token if the old one is close to expiry
	var newRefreshToken string
	expiresAt, expiresAtOk := claims["exp"].(int64)
	if !expiresAtOk {
		return nil, errors.New("invalid token expiry claim")
	}
	if expiresAt-time.Now().Unix() < 3600 { // less than 1 hour to expire
		newRefreshToken, err = token.GenerateRefreshToken(userId, email)
		if err != nil {
			log.Println("Error generating new refresh token:", err)
			return nil, errors.New("could not generate refresh token")
		}
	} else {
		newRefreshToken = req.RefreshToken
	}

	// Prepare the response
	res := &pb.Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    3600, // Assuming access token expires in 1 hour
	}

	return res, nil
}

func (r *UserAuthRepo) InvalidateToken(token string) error {
	_, err := r.db.Exec("INSERT INTO token_blacklist (token) VALUES ($1)", token)
	if err != nil {
		return err
	}
	return nil
}

//func (r *UserAuthRepo) IsTokenBlacklisted(token string) (bool, error) {
//	var count int
//	err := r.db.QueryRow("SELECT COUNT(*) FROM token_blacklist WHERE token=$1", token).Scan(&count)
//	if err != nil {
//		return false, err
//	}
//	return count > 0, nil
//}
