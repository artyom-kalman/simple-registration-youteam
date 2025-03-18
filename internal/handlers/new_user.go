package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/artyom-kalman/simple-registration-youteam/internal/models"
	"github.com/artyom-kalman/simple-registration-youteam/internal/models/dto"
	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received new user request")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	newUser, err := getRegistrationDataFromRequest(r)
	if err != nil {
		logger.Error("Failed to get registration data: %v", err)
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = saveUser(r.Context(), newUser)
	if err != nil {
		logger.Error("Failed to save user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:8000", http.StatusSeeOther)
	logger.Info("Successfully added new user")
}

func getRegistrationDataFromRequest(r *http.Request) (*dto.NewUserDto, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, fmt.Errorf("invalid request body")
	}

	newUser := dto.NewUserDto{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if newUser.Email == "" || newUser.Password == "" {
		return nil, errors.New("email and password are required")
	}

	return &newUser, nil
}

func saveUser(ctx context.Context, user *dto.NewUserDto) (*models.User, error) {
	logger.Info("Saving new user")

	query := fmt.Sprintf("INSERT INTO users (email, password) VALUES ('%s', '%s') RETURNING id, email, created_at",
		user.Email, user.Password,
	)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	defer rows.Close()

	isResult := rows.Next()
	if !isResult {
		return nil, fmt.Errorf("failed to create user")
	}

	var newUser models.User
	if err := rows.Scan(&newUser.ID, &newUser.Email, &newUser.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan user data: %w", err)
	}

	return &newUser, nil
}
