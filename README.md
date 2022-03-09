
# go-infinity
Golang client for infinity graph. 

Implements infinity graph v2021.3 REST API, v1, as per March, 2022. 

## Api implemented 

- [x] Server endpoint
- [x] FDB Endpoint 
- [x] Schema Endpoint
- [x] Object Endpoint
- [x] Query Endpoint
- [x] Index Endpoint
- [x] Transaction Endpoint

Doing:
- [ ] Admin Endpoint

Known issues:
* Transaction through the endpoint may not work. Test fails with server side error. 
* Class schema update doesn't work. Test fails for unknown reasons.
* Get Index & GetIndexClass *not* implemented because of server side error 
* Delete operations do not return anything so it's impossible to know at runtime if an delete actually succeeded.  

Workarounds:
* Write transactions, schema & all index operations as Do Query and send them through the query endpoint
* Delete only through Do query operations to circumvent the no-return bug. 

