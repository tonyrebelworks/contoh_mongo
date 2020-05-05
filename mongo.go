package contact

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabase = "nimbly"

type MongoHandler struct {
	client   *mongo.Client
	database string
}

//MongoHandler Constructor
func NewHandler(address string) *MongoHandler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))
	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}
	return mh
}

// GetOne ...
func (mh *MongoHandler) GetOne(c *JourneyPlan, filter interface{}) error {
	//Will automatically create a collection if not available
	collection := mh.client.Database(mh.database).Collection("journey_plan")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(c)
	return err
}

//Get ...
func (mh *MongoHandler) Get(filter interface{}) []*JourneyPlan {
	collection := mh.client.Database(mh.database).Collection("journey_plan")

	// fmt.Println(collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result []*JourneyPlan
	for cur.Next(ctx) {
		journey := &JourneyPlan{}
		er := cur.Decode(journey)
		if er != nil {
			log.Fatal(er)
		}
		result = append(result, journey)
	}
	// fmt.Println(result)
	return result
}

func (mh *MongoHandler) AddOne(c *JourneyPlan) (*mongo.InsertOneResult, error) {
	collection := mh.client.Database(mh.database).Collection("journey_plan")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, c)
	return result, err
}

func (mh *MongoHandler) Update(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mh.client.Database(mh.database).Collection("journey_plan")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.UpdateMany(ctx, filter, update)
	return result, err
}

func (mh *MongoHandler) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mh.client.Database(mh.database).Collection("journey_plan")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}
