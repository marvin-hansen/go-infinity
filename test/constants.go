package test

const clientError = "Client should not be nil"

const testUpdateSchema = `
[
  {
  "className": "FleetData.Customer",
  "attributes": [
      {
        "attributeName": "goldStatus",
        "logicalType": "boolean",
        "encoding": "binary",
        "storage": "byte"
      }
    ]
  }
]
`

const testAddSchema = `[
  {
    "className":"FleetData.Customer",
    "isReferenceable":true,
    "superClass":null,
    "attributes":[
      {
        "attributeName":"rewardPoints",
        "logicalType":"integer",
        "encoding":"unsigned",
        "storage":"b32"
      },
      {
        "attributeName":"firstName",
        "logicalType":"string",
        "encoding":"utf8",
        "storage":"variable"
      },
      {
        "attributeName":"lastName",
        "logicalType":"string",
        "encoding":"utf8",
        "storage":"variable"
      }
    ]
  }
]
`
