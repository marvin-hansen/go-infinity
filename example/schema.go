package main

const TestSchema = `[
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
