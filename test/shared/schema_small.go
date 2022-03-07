package shared

const ClientError = "Client should not be nil"

const TestUpdateClassSchema = `
{
  "className": "FleetData.Customer",
  "attributesToDelete": [
    "rewardPoints",
  ],
}
`

const TestUpdateSchema = `
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

const TestAddSchema = `[
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
