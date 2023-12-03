package db

import (
	"github.com/lib/pq"
	"github.com/mounis-bhat/go-bank/types"
)

func (s *PostgresStorage) GetAccountByUsername(username string) (*types.Account, error) {
	query := `
	SELECT *
	FROM accounts
	WHERE username = $1
	`
	row := s.db.QueryRow(query, username)

	account := &types.Account{}

	if err := row.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Balance, &account.CreatedAt, &account.Username, &account.Password, pq.Array(&account.Roles)); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *PostgresStorage) CreateAccount(account *types.Account) (int, error) {
	query := `
	INSERT INTO accounts (
			first_name,
			last_name,
			balance,
			created_at,
			username,
			password,
			roles
	)
	VALUES
			($1, $2, $3, $4, $5, $6, $7)
	RETURNING account_id
	`

	var accountID int
	err := s.db.QueryRow(query, account.FirstName, account.LastName, account.Balance, account.CreatedAt, account.Username, account.Password, pq.Array(account.Roles)).Scan(&accountID)
	if err != nil {
		return 0, err
	}

	account.ID = accountID
	return accountID, nil
}
