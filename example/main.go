package main

import (
	"go-infinity/rest"
	"go-infinity/rest/object"
	"go-infinity/rest/schema"
	"log"
)

func main() {

	// we assume a DB to be present and the API server online. If not, please create one from the command line.
	// To do so:
	// 1) exec into the container:
	// make graph-connect-container

	// 2 Create DB:
	// objy CreateFd -fdNa DBNAME -fdDirPath /app/data

	// 3) Start REST API server
	// ensure the DB name is  in the restApiConfig.xml file
	// objy StartRESTServer -fdAlias kb -configFile restApiConfig.xml

	// check if the API server is running:
	//  objy CheckRESTServer -port 8185

	// When DB & APi server are up & running, exit container.

	// creates a client with default values i.e. localhost.
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
