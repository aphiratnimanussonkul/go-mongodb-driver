package config

import (
  "context"
  "fmt"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
)

func GetMongoDB() (*mongo.Database, error) {
  uri := "mongodb+srv://admin:admin@cluster0-mpkxk.mongodb.net"
  clientOptions := options.Client().ApplyURI(uri)

  // Connect to MongoDB
  client, err := mongo.Connect(context.TODO(), clientOptions)

  if err != nil {
    log.Fatal(err)
  }

  // Check the connection
  err = client.Ping(context.TODO(), nil)

  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connected to MongoDB!")
  return  client.Database("CourseOnline"), nil
}
