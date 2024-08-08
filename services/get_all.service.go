package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) GetAllNodesAndRelationships(c *gin.Context) {

	// Open a new session

	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := "MATCH (n) MATCH (n)-[r]-() RETURN n,r"
		result, _ := neo4j.ExecuteQuery(n.Ctx, n.Driver,
			query, nil, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))

		c.JSON(http.StatusAccepted, result.Records)
	}
}

func (n *NeoHandler) GetAllNodes(c *gin.Context) {

	// Open a new session

	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := "MATCH (n) RETURN n"
		result, _ := neo4j.ExecuteQuery(n.Ctx, n.Driver,
			query, nil, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))

		c.JSON(http.StatusAccepted, result.Records)
	}
}

func (n *NeoHandler) GetAllRelationships(c *gin.Context) {

	// Open a new session

	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {

		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		query := "MATCH (r) RETURN r"
		result, _ := neo4j.ExecuteQuery(n.Ctx, n.Driver,
			query, nil, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))

		c.JSON(http.StatusAccepted, result.Records)
	}
}
