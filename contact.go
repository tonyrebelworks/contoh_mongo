package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

//JourneyPlan ...
type JourneyPlan struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	JourneyName string             `json:"journeyName" bson:"journey_name"`
}
