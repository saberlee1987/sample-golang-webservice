package dto

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

func (c Customer) String() string {
	json, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("{\"id\":%d,\"firstName\":\"%s\",\"lastName\":\"%s\",\"email\":\"%s\"}",
			c.Id, c.FirstName, c.LastName, c.Email)
	}
	return string(json)
}
