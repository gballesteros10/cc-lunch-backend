package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type LunchOrder struct {
	ID       primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID  `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OptionID *primitive.ObjectID `json:"option_id,omitempty" bson:"option_id,omitempty"`
	Day      *int                `json:"day,omitempty" bson:"day,omitempty"`
}
