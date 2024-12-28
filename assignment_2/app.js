// Import MongoClient from the mongodb module
const MongoClient = require('mongodb').MongoClient;

// Connection URL
const uri = "mongodb://127.0.0.1:27017";

// Database and Collection Name
const dbName = 'testDb';
const c_collection = 'customers';
const o_collection = 'orders';
const client = new MongoClient(uri);

console.log('Connecting to MongoDB server...');
async function connectToMongoDB() {
    try {
        // Connect to the MongoDB server
        await client.connect();
        console.log('Connected to MongoDB server');

        const db = client.db(dbName);
        const customer_collection = db.collection(c_collection);
        const order_collection = db.collection(o_collection);

        // Insert a single document
        // const result = await collection.insertOne({ name: 'John Doe' });
        // console.log('Inserted a document into the customers collection, with _id:', result.insertedId);

        // Insert multiple documents
        const customers = [
            {
              "name": "John Doe",
              "email": "johndoe@example.com",
              "address": {
                "street": "123 Main St",
                "city": "Springfield",
                "zipcode": "12345"
              },
              "phone": "555-1234",
              "registration_date": new Date("2023-01-01T12:00:00Z")
            },
            {
              "name": "Ankit Singh",
              "email": "abcdef@example.com",
              "address": {
                "street": "456 ABC St",
                "city": "Delhi",
                "zipcode": "12346"
              },
              "phone": "000-11111",
              "registration_date": new Date("2023-02-15T09:30:00Z")
            },
            {
              "name": "Vignesh Kumar",
              "email": "ghijkl@example.com",
              "address": {
                "street": "789 GHI St",
                "city": "Mumbai",
                "zipcode": "12347"
              },
              "phone": "555-8765",
              "registration_date": new Date("2023-03-06T14:00:00Z")
            },
            {
              "name": "Rahul Kumar",
              "email": "mnopqr@example.com",
              "address": {
                "street": "101 MNO St",
                "city": "Kolkata",
                "zipcode": "12348"
              },
              "phone": "555-4321",
              "registration_date": new Date("2023-04-20T16:00:00Z")
            },
            {
              "name": "Kartikey Gupta",
              "email": "stuvwx@email.com",
              "address": {
                "street": "101 STU St",
                "city": "Bangalore",
                "zipcode": "12349"
              },
              "phone": "555-6789",
              "registration_date": new Date("2023-05-15T14:00:00Z")
            }
        ];          

        const results_customers = await customer_collection.insertMany(customers);
        console.log('Inserted multiple documents into the customers collection, with _ids:', results_customers.insertedIds);

        orders = [
          {
            "order_id": "ORD123456",
            "customer_id": results_customers.insertedIds['1'], 
            "order_date": new Date("2023-05-15T14:00:00Z"),
            "status": "shipped",
            "items": [
              {
                "product_name": "Shirt",
                "quantity": 1,
                "price": 1500
              },
              {
                "product_name": "Shorts",
                "quantity": 2,
                "price": 25
              }
            ],
            "total_value": 1550
          },
          {
            "order_id": "ORD123457",
            "customer_id": results_customers.insertedIds['0'],
            "order_date": new Date("2023-06-01T10:30:00Z"),
            "status": "pending",
            "items": [
              {
                "product_name": "Tooth Brush",
                "quantity": 1,
                "price": 799
              },
              {
                "product_name": "Tooth Paste",
                "quantity": 1,
                "price": 20
              }
            ],
            "total_value": 819
          },
          {
            "order_id": "ORD123458",
            "customer_id": results_customers.insertedIds['3'],
            "order_date": new Date("2023-07-22T18:45:00Z"),
            "status": "delivered",
            "items": [
              {
                "product_name": "Badminton Racquet",
                "quantity": 2,
                "price": 400
              },
              {
                "product_name": "Shuttle Cock",
                "quantity": 1,
                "price": 50
              }
            ],
            "total_value": 850
          },
          {
            "order_id": "ORD123459",
            "customer_id": results_customers.insertedIds['2'],
            "order_date": new Date("2023-08-10T12:15:00Z"),
            "status": "shipped",
            "items": [
              {
                "product_name": "Pasta",
                "quantity": 1,
                "price": 300
              },
              {
                "product_name": "Noodles",
                "quantity": 1,
                "price": 30
              }
            ],
            "total_value": 330
          },
          {
            "order_id": "ORD123460",
            "customer_id": results_customers.insertedIds['2'],
            "order_date": new Date("2023-09-10T12:15:00Z"),
            "status": "shipped",
            "items": [
              {
                "product_name": "Book",
                "quantity": 1,
                "price": 250
              },
              {
                "product_name": "Pen",
                "quantity": 1,
                "price": 30
              }
            ],
            "total_value": 280
          }
        ]
        const results_orders = await order_collection.insertMany(orders);
        console.log('Inserted multiple documents into the orders collection, with _ids:', results_orders.insertedIds);


        // Query for a single document
        const query = { name: 'Jane Doe' };
        const result = await customer_collection.findOne(query);
        console.log('Found a document in the customers collection:', result);
      
    } catch (error) {
        console.error('Error connecting to MongoDB server', error);
    } finally {
        await client.close();
        console.log('Closed connection to MongoDB server');
    }
}

connectToMongoDB();

