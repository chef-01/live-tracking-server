package graph

import "time"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var resolvers = map[string]interface{}{
	"Mutation": map[string]interface{}{
		"createUser": func(name, email string) map[string]interface{} {
			user := map[string]interface{}{
				"id":    time.Now().Unix(),
				"name":  name,
				"email": email,
			}
			// You would typically save the user to a database here
			return user
		},
	},
}
