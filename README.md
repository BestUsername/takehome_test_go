# Takehome test

## Program Requirements
- Built in Go with one HTTP POST route
- Receive a listing of orders in XML (described in the example below)
- Process each record concurrently with goroutines (convert the "data" field in each Order to be uppercase)
- Return data in JSON (described in the example below)
- The service does all work in-memory and does not need to persist data
- The service calls should be blocking

### Full Example
​
#### Input Example (`input_example.xml`):
```
<orderList>	
	<order>
		<id>aeffb38f-a1a0-48e7-b7a8-2621a2678534</id>
		<data>first_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
	<order>
		<id>beffb38f-b1a0-58e7-c7a8-3621a2678534</id>
		<data>second_Order_Data</data>
		<createdAt>0001-01-01T00:00:00Z</createdAt>
		<updatedAt>0001-01-01T00:00:00Z</updatedAt>
	</order>
<orderList>
```
​
#### Example Request:
​
	curl -X POST <YOUR_JSON_SERVER>:8080/process -d @input_example.xml
​
#### Example Response:
```
{
	"orderList": [
		{
			"id": "aeffb38f-a1a0-48e7-b7a8-2621a2678534",
			"data": "FIRST_ORDER_DATA",
			"createdAt": "0001-01-01T00:00:00Z",
			"updatedAt": "0001-01-01T00:00:00Z"
		},
		{
			"id": "beffb38f-b1a0-58e7-c7a8-3621a2678534",
			"data": "SECOND_ORDER_DATA",
			"createdAt": "0001-01-01T00:00:00Z",
			"updatedAt": "0001-01-01T00:00:00Z"
		}
	]
}
```

## How to test
- Go tests:
    go test
- Bash tests:
    ./RUNME.sh

## How to run
go run main.go

## Some Notes
- I changed the output format because the given format wasn't valid JSON.
- Using channels wasn't required as each goroutine edited the Order objects in-memory, lowering system requirements.
