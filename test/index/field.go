package index

const testDeleteIndex = "myRentalContractIndex"

const testAddIndex = `
{
  "name": "myRentalContractIndex",
  "indexedClass": "FleetData.RentalContract",
  "isUnique": "true",
  "keyedAttributes":
    [
      {
        "attributeName": "trackingNumber"
      }    
    ]
}
`
