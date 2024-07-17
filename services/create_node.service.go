package services

import (
	"encoding/json"
	"fmt"
	"io"
	s "local/cypher-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) CreateNode(c *gin.Context) {

	// Get JSON data
	node := s.Node{}

	var jsonData []byte
	if c.Request.Body != nil {
		jsonData, _ = io.ReadAll(c.Request.Body)
	}
	m := make(map[string]any)
	err_a := json.Unmarshal(jsonData, &m)
	if err_a != nil {
		c.JSON(http.StatusUnprocessableEntity, m)
		return
	}

	node.Alias = m["alias"].(string)
	node.Label = m["label"].(string)
	node.Properties = m["properties"].(map[string]any)

	// Open a new session

	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := CypherCreate(node)
		fmt.Println(query)
		_, err := session.Run(n.Ctx,
			query, node.Properties)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusAccepted, node)
	}
}

func CypherCreate(node s.Node) string {
	var meta = ""

	var props = node.Properties
	var l = len(props)
	var c = 0
	for key, value := range props {
		fmt.Println(fmt.Sprint(value))
		c++
		if c == l {
			meta += key + ":" + "$" + key
		} else {
			meta += key + ":" + "$" + key + ", "
		}

	}
	var cr = "CREATE (" + node.Alias + ":" + node.Label + " {" + meta + "})"
	return cr
}
