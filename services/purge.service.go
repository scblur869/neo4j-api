package services

import (
	s "local/cypher-api/structs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) PurgeDbData(c *gin.Context) {

	var dbAction s.DbActions

	// Bind the JSON data to the Node struct
	if err := c.ShouldBindJSON(&dbAction); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dbAction.Action)
		return
	}

	// Open a new session
	if dbAction.Action == "DELETE" && dbAction.Key == os.Getenv("AESKEY") {
		if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

			session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
			defer UnsafeClose(n.Ctx, session)

			result, err := session.Run(n.Ctx,
				`MATCH (n) DETACH DELETE n`, nil)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err)
			}

			c.JSON(http.StatusAccepted, result)
		}
	} else {
		c.JSON(http.StatusForbidden, "key auth failed")
	}

}
