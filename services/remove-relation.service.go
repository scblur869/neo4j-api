package services

import (
	"fmt"
	s "local/cypher-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) DeleteRelation(c *gin.Context) {
	// get the relation id from the request

	var rel s.Relationship

	// Bind the JSON data to the Node struct
	if err := c.ShouldBindJSON(&rel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, rel.RelationProperties)
		return
	}
	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := CypherRemoveOutGoingRelation(rel)
		fmt.Println(query)
		result, err := session.Run(n.Ctx,
			query, rel.FromProperties)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, result)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Outgoing relation deleted successfully"})
}

func CypherRemoveOutGoingRelation(rel s.Relationship) string {
	var meta = ""

	var properties = rel.FromProperties
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
	var cr = "MATCH (b:" + rel.FromLabel + " {" + meta + "})-[r:" + rel.Relation + "]->() DELETE r"
	return cr
}
