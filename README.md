# Process-receipt-backend

This is a webservice that fulfils the documented API ([api.yml](https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml)). Unit test was not done due to time constraint (48 hours). However, if done in industrial setting, unit test would be must. 
> Note: This is my first time using Go. Although accepted in other language, I thought this would be good oppotunity to learn a new language and since the problem statement said that the reviewer already has go install I also thought it would be helpful for the reviewer

## How to run this webservice?

### Prerequisites

- <strong> Go </strong>- You need to have the Go programming language installed on your system. If you haven't installed it yet, you can get it from the official Go website.

- <strong> Go Dependencies </strong> - The code uses external packages which need to be fetched:

  - ```github.com/google/uuid```
  - ```github.com/gorilla/mux```

### Steps to Run the Code

- <strong> Clone or Download the Repository </strong> - If you have this code in a repository, clone it. Otherwise, save the provided code into a file named main.go in a new directory.

- <strong> Fetch Dependencies </strong> - Navigate to the directory containing the main.go file and run:
```
go get -u github.com/google/uuid
go get -u github.com/gorilla/mux
```
- <strong> Run the Service </strong>- Still in the same directory, run:
```
go run *.go
```
This will start the service on port 8080. If everything goes well, you won't see any output immediately – this just means the server is running.
- <strong> Testing the Service </strong> - You can use tools like curl or Postman to test the service endpoints:
  - Processing a Receipt:

        Endpoint: http://localhost:8080/receipts/process
        Method: POST
        Body: [A JSON receipt] (see example below)

  - Fetching Points for a Receipt:

        Endpoint: http://localhost:8080/receipts/{id}/points 
        Method: GET 
        Replace {id} with the id received when processing a receipt.

-  <strong> Stopping the Service </strong>- If you want to stop the service, go to the terminal where the server is running and press CTRL + C.

## Example

Receipt Json for POST request:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```
Total Points for above JSON in GET request
```
{"points":28}
``````

## Read More
Read More [here](https://github.com/fetch-rewards/receipt-processor-challenge)
