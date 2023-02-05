package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type usersEmail [100_000]string

func getUsers(r io.Reader) (result usersEmail, err error) {
	scanner := bufio.NewScanner(r)

	user := &User{}
	i := 0
	for scanner.Scan() {
		if err = user.UnmarshalJSON(scanner.Bytes()); err != nil {
			return usersEmail{}, err
		}
		result[i] = user.Email
		i++
		*user = User{}
	}
	return result, nil
}

func countDomains(u usersEmail, domain string) (DomainStat, error) {
	result := make(DomainStat)

	searchDomain := "." + domain
	for _, userEmail := range u {
		if strings.Contains(userEmail, searchDomain) {
			result[strings.ToLower(strings.SplitN(userEmail, "@", 2)[1])]++
		}
	}
	return result, nil
}
