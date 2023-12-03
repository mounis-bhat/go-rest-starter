package server

import (
	"encoding/json"
	"net/http"

	"github.com/mounis-bhat/go-bank/pkg/lib"
	"github.com/mounis-bhat/go-bank/types"
)

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {
	validatedAccount, err := lib.GetAccountAndValidate(r)
	if err != nil {
		lib.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	account, err := s.dbConnection.GetAccountByUsername(validatedAccount.Username)
	if err != nil {

		lib.WriteJSON(w, http.StatusNotFound, "Account not found")
		return
	}

	lib.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var request types.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, "Invalid body")
	}

	if request.FirstName == "" || request.LastName == "" || request.Username == "" || request.Password == "" || request.Roles == nil {
		lib.WriteJSON(w, http.StatusBadRequest, "Invalid body")
		return
	}

	isValid := lib.IsValidPassword(request.Password)

	if !isValid {
		lib.WriteJSON(w, http.StatusBadRequest, "The password must be at least 8 characters long, contain at least one uppercase/lowercase letter, at least one number and at least one special character")
		return
	}

	hashedAndSaltedPassword, err := lib.HashAndSaltPassword(request.Password)
	if err != nil {
		lib.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error hash")
		return
	}

	account := types.NewAccount(request.FirstName, request.LastName, request.Username, hashedAndSaltedPassword, request.Roles)

	accountID, err := s.dbConnection.CreateAccount(account)
	if err != nil {
		lib.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error create")
		return
	}

	account.ID = accountID

	token, err := lib.GenerateJWTToken(account)
	if err != nil {
		lib.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error generate")
		return
	}

	lib.WriteJSON(w, http.StatusOK, token)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var request types.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		lib.WriteJSON(w, http.StatusBadRequest, "Invalid body")
	}

	if request.Username == "" || request.Password == "" {
		lib.WriteJSON(w, http.StatusBadRequest, "Access denied")
		return
	}

	account, err := s.dbConnection.GetAccountByUsername(request.Username)
	if err != nil {
		lib.WriteJSON(w, http.StatusNotFound, "Access denied")
		return
	}

	isValid := lib.ComparePasswords(account.Password, request.Password)

	if !isValid {
		lib.WriteJSON(w, http.StatusUnauthorized, "Access denied")
		return
	}

	token, err := lib.GenerateJWTToken(account)
	if err != nil {
		lib.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	lib.WriteJSON(w, http.StatusOK, token)
}
