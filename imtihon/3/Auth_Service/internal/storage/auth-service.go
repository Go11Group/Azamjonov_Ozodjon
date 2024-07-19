package storage

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	pb "github.com/imtihon/3/Auth_Service/generated/genprotos/auth_pb"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	AuthSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
		logger       *slog.Logger
	}
)

func New(logger *slog.Logger) (*AuthSt, error) {

	db, err := Conn()
	if err != nil {
		return nil, err
	}

	return &AuthSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:       logger,
	}, nil
}

func (s *AuthSt) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user_id := uuid.New().String()
	created_at := time.Now()
	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	query, args, err := s.queryBuilder.Insert("users").
		Columns(
			"user_id",
			"username",
			"email",
			"password",
			"full_name",
			"user_type",
			"created_at").
		Values(
			user_id,
			in.Username,
			in.Email,
			hashedPassword,
			in.FullName,
			in.UserType,
			created_at).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId:    user_id,
		Username:  in.Username,
		Email:     in.Email,
		FullName:  in.FullName,
		UserType:  in.UserType,
		CreatedAt: created_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 2
func (s *AuthSt) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// query, args, err := s.queryBuilder.Select("username", "email", "user_type").
	return nil, nil
}

// 3
func (s *AuthSt) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"user_id",
		"username",
		"email",
		"full_name",
		"user_type",
		"address",
		"phone_number",
		"created_at",
		"updated_at").
		From("users").
		Where(sq.Eq{"user_id": in.UserId}).
		ToSql()
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response pb.GetProfileResponse
	var tempAddress, tempPhoneNumber sql.NullString

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&response.UserId,
		&response.Username,
		&response.Email,
		&response.FullName,
		&response.UserType,
		&tempAddress,
		&tempPhoneNumber,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	if tempAddress.Valid {
		response.Address = tempAddress.String
	}
	if tempPhoneNumber.Valid {
		response.PhoneNumber = tempPhoneNumber.String
	}

	return &response, nil
}

// 4
func (s *AuthSt) UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	updated_at := time.Now()

	query, args, err := s.queryBuilder.Update("users").
		Set("full_name", in.FullName).
		Set("address", in.Address).
		Set("phone_number", in.PhoneNumber).
		Set("updated_at", updated_at).
		Where(sq.Eq{"user_id": in.UserId}).
		Suffix("RETURNING username, email, user_type").
		ToSql()

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	row := s.db.QueryRowContext(ctx, query, args...)
	var username, email, user_type string
	err = row.Scan(
		&username,
		&email,
		&user_type,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		UserId:      in.UserId,
		Username:    username,
		Email:       email,
		FullName:    in.FullName,
		UserType:    user_type,
		Address:     in.Address,
		PhoneNumber: in.PhoneNumber,
		UpdatedAt:   updated_at.Format("2006-01-02 15:04:05"),
	}, nil
}

// 5
func (s *AuthSt) ResetPassword(ctx context.Context, in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		s.logger.Error("Failed to hash password", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	query, args, err := s.queryBuilder.Update("users").
		Set("password", hashedPassword).
		Where(sq.Eq{"email": in.Email}).
		ToSql()
	if err != nil {
		s.logger.Error("Failed to build SQL query", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	result, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Failed to execute SQL query", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Failed to get affected rows", "error", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	if rowsAffected == 0 {
		s.logger.Warn("No user found with the provided email")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.ResetPasswordResponse{Message: "Password successfully updated"}, nil
}

// 6
func (s *AuthSt) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return nil, nil
}

// 7
func (s *AuthSt) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}
