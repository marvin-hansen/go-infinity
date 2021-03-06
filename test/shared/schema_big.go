package shared

const TestBigSchema = `
[
  {
    "className": "FleetData.Customer",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "rewardPoints",
        "logicalType": "integer",
        "encoding": "unsigned",
        "storage": "b32"
      },
      {
        "attributeName": "firstName",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "lastName",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "address",
        "logicalType": "object",
        "class": "FleetData.Address"
      },	  
      {
        "attributeName": "userId",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "pw",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "rentals",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.RentalContract",
          "inverseAttribute": "customer"
        }
      }
    ]
  },
  {
    "className": "FleetData.Location",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "rateFactor",
        "logicalType": "real",
        "storage": "b32"
      },
      {
        "attributeName": "name",
        "logicalType": "string",
        "encoding": "utf16",
        "storage": "variable"
      },
      {
        "attributeName": "address",
        "logicalType": "object",
        "class": "FleetData.Address"
      },
      {
        "attributeName": "airportCode",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "pendingPickups",
        "logicalType": "set",
        "collectionTypeName": "TreeSetOfReferences",
        "elementSpecification": {
          "logicalType": "reference",
		  "referencedClass": "FleetData.Reservation"
        }
      },
      {
        "attributeName": "pendingReturns",
        "logicalType": "set",
        "collectionTypeName": "TreeSetOfReferences",
        "elementSpecification": {
          "logicalType": "reference",
		  "referencedClass": "FleetData.Reservation"
        }
      },
      {
        "attributeName": "stalls",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.Stall",
          "inverseAttribute": "location"
        }
      }
    ]
  },
  {
    "className": "FleetData.RentalEvent",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "timestamp",
        "logicalType": "datetime",
        "encoding": "objy",
		"storage": "datetimeoffset"
      },
      {
        "attributeName": "rental",
        "logicalType": "reference",
        "referencedClass": "FleetData.RentalContract"
      }
    ]
  },
  {
    "className": "FleetData.PickUp",
    "isReferenceable": true,
    "superClass": 
      {
        "className": "FleetData.RentalEvent"
      },
    "attributes": [
      {
        "attributeName": "completed",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "delta",
        "logicalType": "interval",
        "encoding": "objy"
      },
      {
        "attributeName": "atStall",
        "logicalType": "reference",
        "referencedClass": "FleetData.Stall",
        "inverseAttribute": "assignedTo"
      }
    ]
  },
  {
    "className": "FleetData.RentalCompany",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "name",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "corporateAddress",
        "logicalType": "object",
        "class": "FleetData.Address"
      },
      {
        "attributeName": "locations",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.Location"
        }
      },
      {
        "attributeName": "highVolumeDates",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "date",
          "encoding": "objy"
        }
      },
      {
        "attributeName": "categories",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.VehicleCategory"
        }
      },
      {
        "attributeName": "models",
        "logicalType": "list",
        "collectionTypeName": "TreeListOfReferences",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.VehicleModel"
        }
      },
      {
        "attributeName": "vehicles",
        "logicalType": "set",
        "collectionTypeName": "TreeSetOfReferences",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.Vehicle"
        }
      },
      {
        "attributeName": "customers",
        "logicalType": "map",
        "collectionTypeName": "NameToReferenceMap",
        "keySpecification": {
          "logicalType": "string",
          "encoding": "byte",
          "storage": "variable"
        },
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.Customer"
        }
      }
    ]
  },
  {
    "className": "FleetData.Address",
    "isReferenceable": false,
    "attributes": [
      {
        "attributeName": "street",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "city",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "state",
        "logicalType": "string",
        "encoding": "byte",
        "storage": "fixed",
        "fixedLength": 3
      },
      {
        "attributeName": "postalCode",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      }
    ]
  },
  {
    "className": "FleetData.RentalContract",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "customer",
        "logicalType": "reference",
        "referencedClass": "FleetData.Customer",
        "inverseAttribute": "rentals"
      },
      {
        "attributeName": "vehicle",
        "logicalType": "reference",
        "referencedClass": "FleetData.Vehicle",
        "inverseAttribute": "pastRentals"
      },
      {
        "attributeName": "trackingNumber",
        "logicalType": "integer",
        "encoding": "signed",
        "storage": "b16"
      },
      {
        "attributeName": "events",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.RentalEvent"
        }
      }
    ]
  },
  {
    "className": "FleetData.Reservation",
    "isReferenceable": true,
    "superClass":
      {
        "className": "FleetData.RentalEvent"
      },
    "attributes": [
      {
        "attributeName": "pickup",
        "logicalType": "datetime",
        "encoding": "objy",
		"storage": "datetime"
      },
      {
        "attributeName": "dropoff",
        "logicalType": "datetime",
        "encoding": "objy",
		"storage": "datetime"
      },
      {
        "attributeName": "reservedCategory",
        "logicalType": "reference",
        "referencedClass": "FleetData.VehicleCategory"
      },
      {
        "attributeName": "pickupLocation",
        "logicalType": "reference",
        "referencedClass": "FleetData.Location"
      },
      {
        "attributeName": "returnLocation",
        "logicalType": "reference",
        "referencedClass": "FleetData.Location"
      }
    ]
  },
  {
    "className": "FleetData.Return",
    "isReferenceable": true,
    "superClass":
      {
        "className": "FleetData.RentalEvent"
      },
    "attributes": [
      {
        "attributeName": "finalized",
        "logicalType": "date",
        "encoding": "objy"
      },
      {
        "attributeName": "timeOfDay",
        "logicalType": "time",
        "encoding": "objy"
      }
    ]
  },
  {
    "className": "FleetData.Stall",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "number",
        "logicalType": "integer",
        "encoding": "signed",
        "storage": "b16"
      },
      {
        "attributeName": "location",
        "logicalType": "reference",
        "referencedClass": "FleetData.Location",
        "inverseAttribute": "stalls"
      },
      {
        "attributeName": "occupiedBy",
        "logicalType": "reference",
        "referencedClass": "FleetData.Vehicle",
        "inverseAttribute": "atStall"
      },
      {
        "attributeName": "assignedTo",
        "logicalType": "reference",
        "referencedClass": "FleetData.PickUp",
        "inverseAttribute": "atStall"
      }
    ]
  },
  {
    "className": "FleetData.Vehicle",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "available",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "retired",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "license",
        "logicalType": "string",
        "encoding": "byte",
        "storage": "variable"
      },
      {
        "attributeName": "servRecs",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "string",
          "encoding": "utf16",
          "storage": "variable"
        }
      },
      {
        "attributeName": "currentRental",
        "logicalType": "reference",
        "referencedClass": "FleetData.RentalContract"
      },
      {
        "attributeName": "pastRentals",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.RentalContract",
          "inverseAttribute": "vehicle"
        }
      },
      {
        "attributeName": "model",
        "logicalType": "reference",
        "referencedClass": "FleetData.VehicleModel",
        "inverseAttribute": "vehicles"
      },
      {
        "attributeName": "atStall",
        "logicalType": "reference",
        "referencedClass": "FleetData.Stall",
        "inverseAttribute": "occupiedBy"
      }
    ]
  },
  {
    "className": "FleetData.VehicleCategory",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "rate",
        "logicalType": "real",
        "storage": "b32"
      },
      {
        "attributeName": "retired",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "name",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "models",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.VehicleModel",
          "inverseAttribute": "categories"
        }
      }
    ]
  },
  {
    "className": "FleetData.VehicleModel",
    "isReferenceable": true,
    "superClass": null,
    "attributes": [
      {
        "attributeName": "doors",
        "logicalType": "integer",
        "encoding": "signed",
        "storage": "b32"
      },
      {
        "attributeName": "seatingCapacity",
        "logicalType": "integer",
        "encoding": "unsigned",
        "storage": "b8"
      },
      {
        "attributeName": "powerWindows",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "automatic",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "airConditioning",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "cdPlayer",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "cruiseControl",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "retired",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      },
      {
        "attributeName": "modelName",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "brand",
        "logicalType": "string",
        "encoding": "utf8",
        "storage": "variable"
      },
      {
        "attributeName": "vehicles",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.Vehicle",
          "inverseAttribute": "model"
        }
      },
      {
        "attributeName": "categories",
        "logicalType": "list",
        "storage": "variable",
        "elementSpecification": {
          "logicalType": "reference",
          "referencedClass": "FleetData.VehicleCategory",
          "inverseAttribute": "models"
        }
      }
    ]
  }
]
`
