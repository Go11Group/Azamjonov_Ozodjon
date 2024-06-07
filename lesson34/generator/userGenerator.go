package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"users_and_produncts/model"
)

var firstNames = []string{
	"John", "Jane", "Michael", "Michelle", "Chris", "Christina",
	"James", "Jennifer", "Robert", "Linda", "William", "Elizabeth",
	"David", "Susan", "Richard", "Karen", "Joseph", "Nancy",
	"Thomas", "Lisa",
}

var lastNames = []string{
	"Smith", "Johnson", "Brown", "Williams", "Jones", "Garcia",
	"Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez",
	"Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore",
	"Jackson", "Martin",
}

var domains = []string{
	"example.com", "test.com", "sample.org", "demo.net", "mail.com", "service.com",
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func randomUsername(firstName, lastName string) string {
	return strings.ToLower(firstName[:1] + lastName + fmt.Sprint(rand.Intn(100)))
}

func randomEmail(username string) string {
	domain := domains[rand.Intn(len(domains))]
	return fmt.Sprintf("%s@%s", username, domain)
}

func GenerateUsers(count int) []model.User {
	rand.Seed(time.Now().UnixNano())

	users := make([]model.User, count)
	for i := 0; i < count; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		username := randomUsername(firstName, lastName)
		email := randomEmail(username)
		password := randomString(10)

		users[i] = model.User{
			Username: username,
			Email:    email,
			Password: password,
		}
	}
	return users
}
