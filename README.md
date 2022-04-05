
# go-infinity
Golang client for [infinity graph db](https://infinitegraph.com/). 

Implements infinity graph v2021.3 REST API [specification](https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2FwelcomeInfiniteGraphPlatform.html%23), v1, as per March 2022.

See [example code](/example) and [tests](/test) for usage examples.


## Api implemented (March / 2022)

- [x] Server endpoint
- [x] FDB Endpoint 
- [x] Schema Endpoint
- [x] Object Endpoint
- [x] Query Endpoint
- [x] Index Endpoint
- [x] Transaction Endpoint
- [x] Tool Endpoint (runs administrative tools on the infinity server and returns the output)

Known issues:
* Transaction through the endpoint may not work. Test fails with server side error. 
* Class schema update doesn't work. Test fails for unknown reasons.
* Get Index & GetIndexClass *not* implemented because of server side error 
* Delete operations do not return anything so it's impossible to know at runtime if an delete actually succeeded.  

Workarounds:
* Write transactions, schema & all index operations as Do Query and send them through the query endpoint
* Delete only through Do query operations to circumvent the no-return bug. 

## Docker setup

1) Clone the repo

```bash 
    git clone https://github.com/marvin-hansen/go-infinity.git
```

2) Check if all requirements are installed. Requires make. 

```bash 
    make check  
```

3) Build the container:  

```bash 
   make graph-build-image
```

4) Run a new container

```bash 
   make graph-create-container 
```

## Managing the container 

Start the container

```bash 
   make graph-start
```


Stop the container

```bash 
   make graph-stop  
```

## DB Setup
The client assumes a DB to be present and the API server online. If not, please create one from the command line.

1) exec into the container:
```bash 
    make graph-connect-container
```

2) Create DB: 

```bash 
    objy CreateFd -fdNa DBNAME -fdDirPath /app/data
```

3) Start REST API server

```bash 
    // ensure the DB name is  in the restApiConfig.xml file
    objy StartRESTServer -fdAlias kb -configFile restApiConfig.xml
```

4)  Verify that the API server is running:

```bash 
    objy CheckRESTServer -port 8185
```

When DB & APi server are up & running, exit container. 


## Getting started

```go 

import (
	"go-infinity/rest"
	"go-infinity/rest/object"
	"go-infinity/rest/schema"
	"log"
)

func main() {
	// Ensure DB & APi server are up & running.

	// creates a client with default config i.e. localhost.
	c := rest.NewClient(nil)

	// creates a client with config i.e. for remote host or different port
	config := rest.NewClientConfig("localhost", "8185", 3)
	c = rest.NewClient(config)

	testSchema := []byte(TestSchema) // schema must be JSON byte array
	res1, err1 := c.AddSchema(schema.NewRequestForAddSchema(testSchema))
	if err1 != nil {
		log.Fatalf("can't insert schema")
	}

	println(res1.String())

	testObject := []byte(AddVehicle)
	res2, err2 := c.AddObject(object.NewRequestForAddObject(testObject))
	if err2 != nil {
		log.Fatalf("can't insert data")
	}

	println(res2.String())

	objectId := res2.ObjectID

	res3, err3 := c.GetObject(object.NewRequestForGetObject(objectId))
	if err3 != nil {
		log.Fatalf("can't load data object: " + objectId)
	}

	// we get raw JSON, so this needs to be unmarshaled into a struct
	println(string(res3.RawMessage))

	err4 := c.DeleteObject(object.NewRequestForDeleteObject(objectId))
	if err4 != nil {
		log.Fatalf("can't delete object: " + objectId)
	}

	classID := "FleetData.Customer"
	err5 := c.DeleteClassSchema(schema.NewRequestForDeleteClassSchema(classID))
	if err5 != nil {
		log.Fatalf("can't delete schema: " + classID)
	}
}


```

## Make reference 

```bash 
Setup: 
    make check                    Checks whether all project requirements are present.
    make graph-build-image        Builds a docker image containing infinite graph.
    make graph-create-container   Creates an infinite graph cocker container that can be started & stopped.
    make graph-delete-container   Deletes the infinite graph cocker container that can be started & stopped.
    make graph-connect-container   Deletes the infinite graph cocker container that can be started & stopped.

Graph: 
    make graph-start            Starts the local infinite graph.
    make graph-stop             Stops  the local infinite graph.
 
Dev: 
    make build                  Builds the code base incrementally (fast). Use for coding.
    make rebuild                Rebuilds all dependencies & the code base (slow). Use after go mod changes. 
    make stats                  Crunches & shows the latest project stats. 
```
