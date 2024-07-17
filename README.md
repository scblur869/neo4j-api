
# NEO4j-API
## Purpose
Purpose of this API is to accelerate the onboarding of graphdb data in cypher query language format to Neo4j


### Use
Build this service as a docker container (docker build)
OR
Clone this repo and compile and run it locally via the console (go build)

### Endpoints
```console
  [POST] /api/v1/purge   --> purges neo4j data, requires key
  [POST] /api/v1/addNode --> adds node
  [POST] /api/v1/addRelation --> adds relation for nodes
  [POST] /api/v1/deleteNode  --> deletes node by name property
  [POST] /api/v1/deleteRelation --> deletes relation by name property
  [GET]  /api/v1/health  --> health check, good when its containerized
  ```

  #### Node and Relation Properties
  One of the challenges on the current Neo4j docs were the ability to add properties unknown before hand. I added some extra utility functions to handle this and leveraged the power of Go's maps..
  ```go
  node.Properties = m["properties"].(map[string]any)
  ```
### Testing and Seeding
 run the seed.sh file in a bash shell. Basically these are a set of curl statements to purge and load the database with some data. 
 ```bash
  sh ./seed.sh
  ```
  **/api/v1/purge** [POST] requires the 32 bit key included in the payload. you can get this either from the console (its printed when the api starts) or from the key.txt file which is auto created every time the api / service starts. **This key changes everytime the service / api starts**.
  **Note**, if you containerize this to run as a micro service, then the key.txt will be in the container. You can access this via a 'docker exec' command
```json 
{
    "action": "DELETE",
    "key":"679d985808e8dd0178971e838752115b218f2d053cd6eabc0ecb2cd1a2a781bc"
}
```

run the test.sh file to test deletions
```bash
sh ./test.sh
```
#### extra
mr_data.json contains more example data in native json format. This came from my initial thoughts about the data.

#### Payload structs
this maybe helpful for inisight into the JSON payload structure
```go
type Node struct {
	Alias      string         `json:"alias"`
	Label      string         `json:"label"`
	ElementId  string         `json:"elementId,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
}

type Relationship struct {
	Relation           string         `json:"relation"`
	RelationProperties map[string]any `json:"relationProperties"`
	FromAlias          string         `json:"fromAlias"`
	FromLabel          string         `json:"fromLabel"`
	FromProperties     map[string]any `json:"fromProperties"`
	ToAlias            string         `json:"toAlias"`
	ToLabel            string         `json:"toLabel"`
	ToProperties       map[string]any `json:"toProperties"`
}

type DbActions struct {
	Action string `json:"action"`
	Node   string `json:"node"`
	Key    string `json:"key"`
}
```
#### The Neo Handler
I didn't like the example code from Neo4j, so i decided to use Receiver Functions and inject a database handler for any function call that needed an neo4j.session
```go
type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

// NeoHandler allows for passing the neo driver from Main() to receiver functions for calling Neo4j sessions..
type NeoHandler struct {
	Ctx    context.Context
	Driver neo4j.DriverWithContext
	Config *Neo4jConfiguration
}
``` 


### Running Neo4j as a container
```bash
docker run -itd -p 0.0.0.0:7474:7474 -p 0.0.0.0:7687:7687 -v $PWD/neo4j-data:/data -v $PWD/neo4j-plugins:/plugins --name neo4j-apoc -e NEO4J_apoc_export_file_enabled=true -e NEO4J_apoc_import_file_enabled=true -e NEO4J_apoc_import_file_use__neo4j__config=true -e NEO4JLABS_PLUGINS=\[\"apoc\"\] neo4j:latest
```
- You need to run this with APOC plugins enabled...

Enjoy!