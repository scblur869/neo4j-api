package services

import (
	"encoding/json"
	"fmt"
	"io"
	s "local/neo4j-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *NeoHandler) CreatRelationship(c *gin.Context) {

	// Get JSON data
	relation := s.Relationship{}

	var jsonData []byte
	if c.Request.Body != nil {
		jsonData, _ = io.ReadAll(c.Request.Body)
	}
	// map to hold json / request body
	m := make(map[string]any)
	err_a := json.Unmarshal(jsonData, &m)
	if err_a != nil {
		c.JSON(http.StatusUnprocessableEntity, m)
		return
	}
	// assign unmarshalled values from map[string]any to struct
	relation.FromAlias = m["fromAlias"].(string)
	relation.FromLabel = m["fromLabel"].(string)
	relation.FromProperties = m["fromProperties"].(map[string]any)
	relation.ToAlias = m["toAlias"].(string)
	relation.ToLabel = m["toLabel"].(string)
	relation.ToProperties = m["toProperties"].(map[string]any)
	relation.Relation = m["relation"].(string)
	relation.RelationProperties = m["relationProperties"].(map[string]any)
	query := CypherCreateRelationship(relation)
	fmt.Println(query)

	// Open a new session
	if n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}) != nil {
		query := CypherCreateRelationship(relation)
		fmt.Println(query)
		session := n.Driver.NewSession(n.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer UnsafeClose(n.Ctx, session)
		result, err := session.ExecuteWrite(n.Ctx,
			func(tx neo4j.ManagedTransaction) (any, error) {
				records, err := tx.Run(n.Ctx, query, relation.RelationProperties)
				// since i had to have atleast one parameter, i chose a relatiohip property
				if err != nil {
					c.JSON(http.StatusUnprocessableEntity, err)
				}
				return records, err
			})
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
		}
		c.JSON(http.StatusAccepted, result)
	}
}

func CypherCreateRelationship(rel s.Relationship) string {

	var propertiesTo = propertyChain(rel.ToProperties)
	var propertiesFrom = propertyChain(rel.FromProperties)
	var properties = propertyChain(rel.RelationProperties)

	var cr = "MATCH (" + rel.FromAlias + ":" + rel.FromLabel + " {" + propertiesFrom + "}), (" + rel.ToAlias + ":" + rel.ToLabel + " {" + propertiesTo + "}) CREATE (" + rel.FromAlias + ")-[:" + rel.Relation + " {" + properties + "}]->(" + rel.ToAlias + ")"
	return cr
}

func propertyChain(m map[string]any) string {
	var meta = ""

	var l = len(m)
	var c = 0
	for key, value := range m {
		fmt.Println(fmt.Sprint(value))
		c++
		if c == l {
			meta += key + ":\"" + value.(string) + "\""
		} else {
			meta += key + ":\"" + value.(string) + "\", "
		}

	}
	return meta
}

/*
func kvFormatter(m map[string]any) string {
	var meta = ""

	var l = len(m)
	var c = 0
	for key, value := range m {
		fmt.Println(fmt.Sprint(value))
		c++
		if c == l {
			meta += key + ":" + "$" + key
		} else {
			meta += key + ":" + "$" + key + ", "
		}

	}
	return meta
}
*/
