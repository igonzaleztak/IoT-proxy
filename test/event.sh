curl 10.10.46.20:5053/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -X POST -d @- <<EOF
{
  "id":"urn:ngsi-ld:TrafficFlowObserved:santander:traffic:flow:1001",
  "type":"TrafficFlowObserved",
  "dateModified":{
      "type":"ISO8601",
      "value":"2020-09-09T11:58:00.00Z",
      "metadata":{
        
      }
  },
  "dateObserved":{
      "type":"ISO8601",
      "value":"2020-09-09T11:58:00.00Z",
      "metadata":{
        
      }
  },
  "intensity":{
      "type":"Number",
      "value":310,
      "metadata":{
        
      }
  },
  "laneId":{
      "type":"Number",
      "value":0,
      "metadata":{
        
      }
  },
  "location":{
      "type":"geo:json",
      "value":{
        "type":"Point",
        "coordinates":[
            -3.8295937,
            43.4535859
        ]
      },
      "metadata":{
        
      }
  },
  "occupancy":{
      "type":"Number",
      "value":0.08,
      "metadata":{
        
      }
  },
  "roadLoad":{
      "type":"Number",
      "value":17,
      "metadata":{
        
      }
  }
}
EOF