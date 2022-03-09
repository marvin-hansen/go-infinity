package transaction

const testTransaction = `
[
  {
    "method": "post", "uri": "/v1/object",
    "body":
      { 
      "class": "FleetData.Vehicle",
      "attributes":
        { 
        "license": "XLT5663"
        }
      }, "result": "keep"
  },
  {
    "method": "post", "uri": "/v1/query",
    "body":
      {
        "query": "FROM FleetData.Customer WHERE rewardPoints > 25 RETURN *;",
        "language": "do"
      }, "result": "keep"
  }
]
`
