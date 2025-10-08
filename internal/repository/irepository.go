package repository

import "fmt"

type IRepository interface {
	Category() ICategory
}

// refactory review this code may not be here
type ConnectionConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func (c ConnectionConfig) CastToString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		c.Host, c.Port, c.User, c.DBName, c.Password)
}
