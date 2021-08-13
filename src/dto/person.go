package dto

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (c Customer) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("{\"id\":%d,\"firstName\":\"%s\",\"lastName\":\"%s\",\"email\":\"%s\"}",
			c.Id, c.FirstName, c.LastName, c.Email)
	}
	return string(marshal)
}
