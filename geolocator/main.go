package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Get MongoDB URL, DB Name, and Collection name from config

var (
	conn *mongo.Client
	ctx  = context.Background()
)

// Location is a GeoJSON type.
type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

// NewPoint returns a GeoJSON Point with longitude and latitude.
func NewPoint(long, lat float64) Location {
	return Location{
		"Point",
		[]float64{long, lat},
	}
}

// Point is a simple type with a location for geospatial
// queries.
type Point struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title"`
	Location Location           `json:"location"`
}

// createDBSession Create a new connection with the database.
func createDBSession() error {
	var err error
	conn, err = mongo.Connect(ctx, options.Client().
		ApplyURI(connString))
	if err != nil {
		return err
	}
	err = conn.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() error {
	ctx, cancel := context.
		WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := conn.Database(DBName)
	indexOpts := options.CreateIndexes().
		SetMaxTime(time.Second * 10)
	// Index to location 2dsphere type.
	pointIndexModel := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}
	pointIndexes := db.Collection(PointCollection).Indexes()
	_, err := pointIndexes.CreateOne(
		ctx,
		pointIndexModel,
		indexOpts,
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := createDBSession(); err != nil {
		fmt.Println(err)
		return
	}
	if err := createIndex(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected")
	p := Point{
		Title:    "Central Park",
		Location: NewPoint(-73.97, 40.77),
	}
	err := AddPoint(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	points, err := GetPointsByDistance(NewPoint(-73.97, 40.77), 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(points)
}

// AddPoint adds a new point to the collection.
func AddPoint(point Point) error {
	coll := conn.Database(DBName).Collection(PointCollection)
	point.ID = primitive.NewObjectID()
	insertResult, err := coll.InsertOne(ctx, point)
	if err != nil {
		fmt.Printf("Could not insert new Point. Id: %s\n", point.ID)
		return err
	}
	fmt.Printf("Inserted new Point. ID: %s\n", insertResult.InsertedID)
	return nil
}

// GetPointsByDistance gets all the points that are within the
// maximum distance provided in meters.
func GetPointsByDistance(location Location, distance int) ([]Point, error) {
	coll := conn.Database(DBName).Collection(PointCollection)
	var results []Point
	filter := bson.D{
		{"location",
			bson.D{
				{"$near", bson.D{
					{"$geometry", location},
					{"$maxDistance", distance},
				}},
			}},
	}
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return []Point{}, err
	}
	for cur.Next(ctx) {
		var p Point
		err := cur.Decode(&p)
		if err != nil {
			fmt.Println("Could not decode Point")
			return []Point{}, err
		}
		results = append(results, p)
	}
	return results, nil
}
