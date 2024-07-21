import json
import requests

node = open('nodes.json')
rel = open('relations.json')
node_data = json.load(node)
relation_data = json.load(rel)
mr = ["company","equipment","workorders","alerts","data-sources", "alertgroup","personnel"]

headers = {
  'Content-Type': 'application/json'
}
for j in mr:
    for i in node_data[j]:
       r = requests.post('http://localhost:4000/api/v1/addNode',headers=headers, data=json.dumps(i))
       r.json()
       print(r.text)
node.close()


for z in relation_data["relations"]:
      
       r = requests.post('http://localhost:4000/api/v1/addRelation',headers=headers, data=json.dumps(z))
       r.json()
       print(r.text)
rel.close()