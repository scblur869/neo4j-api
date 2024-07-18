package services

import (
	"fmt"
	"net/http"

	s "github.com/scblur869/neo4j-api/structs"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) DeleteNode(c *gin.Context) {

	var node s.Node

	// Bind the JSON data to the Node struct
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusUnprocessableEntity, node.Properties)
		return
	}

	// Open a new session
	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := CypherRemove(node)
		fmt.Println(query)
		_, err := session.Run(n.Ctx,
			query, node.Properties)
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(http.StatusAccepted, node)
	}

}

func CypherRemove(node s.Node) string {
	var meta = ""

	var properties = node.Properties
	var l = len(properties)
	var c = 0
	for key, value := range properties {
		fmt.Println(fmt.Sprint(value))
		c++
		if c == l {
			meta += key + ":" + "$" + key
		} else {
			meta += key + ":" + "$" + key + ", "
		}

	}
	var cr = "MATCH (n:" + node.Label + " {" + meta + "}) DETACH DELETE (n)"
	return cr
}
