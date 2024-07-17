package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

// NeoHandler allows for passing the neo driver to receiver functions
type NeoHandler struct {
	Ctx    context.Context
	Driver neo4j.DriverWithContext
	Config *Neo4jConfiguration
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func ParseConfiguration() *Neo4jConfiguration {
	database := lookupEnvOrGetDefault("NEO4J_DATABASE", "neo4j")
	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4J_VERSION", "4"), "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url:      lookupEnvOrGetDefault("NEO4J_URI", "neo4j+s://localhost:7687"),
		Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "1cartman"),
		Database: database,
	}
}

func UnsafeClose(ctx context.Context, closeable interface{ Close(context.Context) error }) {
	if err := closeable.Close(ctx); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
