# How to run this project
1. make database on mysql with name orders_by
2. import file orders_by.sql into your database
3. copy .env.example to .env
4. change .env config with your computer config
5. run with go run main.go  

# list api
1. List orders
```json
GET http://localhost:8090/api/orders

```

2. Create orders
```json
POST http://localhost:8090/api/orders

// body json
{
    "orderedAt": "2019-11-09T21:21:46+00:00",
    "customerName": "Tom Jerry",
    "items": [
        {
            "itemCode": "123",
            "description": "IPhone 10X",
            "quantity": 1
        }
    ]
}
```
3. Detail Orders
```json
GET http://localhost:8090/api/orders/1
```

4. Update Orders
```json
PUT http://localhost:8090/api/orders/1

//body json
{
    "orderId": 1,
    "customerName": "Spike Tyke",
    "orderedAt": "2019-11-09T21:21:46Z",
    "items": [
        {
            "lineItemId": 1,
            "itemCode": "123",
            "description": "IPhone 10X",
            "quantity": 10
        }
    ]
}
```
4. DELETE orders
```json
DELETE http://localhost:8090/api/orders/1
```