curl --location 'localhost:4000/api/db/purge' \
--header 'Content-Type: application/json' \
--data '{
    "action": "DELETE",
    "key":"679d985808e8dd0178971e838752115b218f2d053cd6eabc0ecb2cd1a2a781bc"
}'

curl --location 'localhost:4000/api/v1/addNode' \
--header 'Content-Type: application/json' \
--data '{
    "alias": "tsdc02",
    "label": "DataSource",
    "properties": {
        "name": "tsdc-2",
        "details": "time stream data source v2.1",
        "owner": "shampton",
        "location": "body shop",
        "collects_from": "jprobot",
        "release": "2.1"
    }
}'

curl --location 'localhost:4000/api/v1/addNode' \
--header 'Content-Type: application/json' \
--data ' {
            "alias": "jprobot02",
            "label": "Equipment",
            "properties": {
                "name": "jprobot02-bs",
                "type": "robotic welder",
                "serial": "134f2ccjp56",
                "role": "body welder",
                "owner": "jsmith",
                "uptime": "37.6",
                "area": "bodyshop",
                "status": "active",
                "lastMaintenance": "2023-08-03T00:00:00Z",
                "lastWorkOrder_id": "wo15463",
                "currentWorkOrder_id": "wo773454",
                "workOrderStatus": "ordered"
            }
        }
  
'
curl --location 'localhost:4000/api/v1/addRelation' \
--header 'Content-Type: application/json' \
--data '{
    "fromAlias": "jprobot02",
    "fromLabel": "Equipment",
    "fromProperties": {
        "name": "jprobot02-bs"
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