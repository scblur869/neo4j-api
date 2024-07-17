curl --location 'localhost:4000/api/v1/deleteNode' \
--header 'Content-Type: application/json' \
--data '{
    "label": "DataSource",
    "properties": {
        "name": "tsdc-2"
    }
}'


curl --location 'localhost:4000/api/v1/deleteRelation' \
--header 'Content-Type: application/json' \
--data '{
    "fromAlias": "jprobot02",
    "fromLabel": "Equipment",
    "fromProperties": {
        "name": "sjprobot02-bs"
    },
    "relation": "SENDS_DATA_TO",
    "relationProperties": {
        "action": "data stream"
    },
    "toAlias": "tsdc02",
    "toLabel": "DataSource",
    "toProperties": {
        "name": "tsdc-2"
    }
}'