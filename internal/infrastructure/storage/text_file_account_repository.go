package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fraylopez/bank-go/internal/domain/account"
	"github.com/fraylopez/bank-go/internal/domain/money"
)

type TextFileAccountRepository struct {
	filePath string
}

func NewTextFileAccountRepository(filename string) *TextFileAccountRepository {
	filePath, err := getFilePath(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to get file directory: %v", err))
	}

	return &TextFileAccountRepository{filePath: filePath}
}

func (r *TextFileAccountRepository) OpenAccount(account *account.Account) error {
	accountsFile, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer accountsFile.Close()

	entry := fmt.Sprintf("%s,%s,%s,%f\n", account.Id, account.Holder, account.Currency, account.Balance.Amount)
	_, err = accountsFile.WriteString(entry)
	if err != nil {
		return err
	}
	return nil
}

func (r *TextFileAccountRepository) GetAccountById(id string) (*account.Account, error) {
	accountsFile, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer accountsFile.Close()

	scanner := bufio.NewScanner(accountsFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}
		if fields[0] == id {
			holder := fields[1]
			currency := fields[2]
			balance, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return nil, err
			}
			return &account.Account{
				Id:       id,
				Holder:   holder,
				Currency: money.Currency(currency),
				Balance:  money.Money{Amount: balance, Currency: money.Currency(currency)},
			}, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("account with id %s not found", id)
}

func (r *TextFileAccountRepository) UpdateAccount(account *account.Account) error {
	accountsFile, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer accountsFile.Close()
	var lines []string
	scanner := bufio.NewScanner(accountsFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}
		if fields[0] == account.Id {
			entry := fmt.Sprintf("%s,%s,%s,%f\n", account.Id, account.Holder, account.Currency, account.Balance.Amount)
			lines = append(lines, entry)
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if len(lines) == 0 {
		return fmt.Errorf("account with id %s not found", account.Id)
	}
	err = os.WriteFile(r.filePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return err
	}
	return nil
}

func getFilePath(filename string) (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// find go.mod file directory
	for {
		if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
			break
		}
		root = filepath.Dir(root)
	}
	// create a db directory if it doesn't exist
	dbDir := filepath.Join(root, "db")
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.Mkdir(dbDir, os.ModePerm); err != nil {
			return "", err
		}
	}
	// create a db/accounts directory if it doesn't exist

	// create a db/accounts/filename directory if it doesn't exist
	filePath := filepath.Join(dbDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
			return "", err
		}
	}
	return filePath, nil
}
