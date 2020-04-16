package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Tenant is your comment
type Tenant struct {
	SocialSecurity string `json:"ss,omitempty" bson:"ss,omitempty"`
	FirstName      string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	CreatedBy      string `json:"createdby,omitempty" bson:"createdby,omitempty"`
}

// Property is yourcomment.
type Property struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
	Zipcode   string             `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Price     string             `json:"price,omitempty" bson:"price,omitempty"`
	Category  string             `json:"category,omitempty" bson:"category,omitempty"`
	CreatedBy string             `json:"createdby,omitempty" bson:"createdby,omitempty"`
}

// Item is yourcomment.
type Item struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Brand       string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Size        string             `json:"size,omitempty" bson:"size,omitempty"`
	Condition   string             `json:"condition,omitempty" bson:"condition,omitempty"`
	Sold        bool               `json:"sold,false" bson:"sold,false"`
	CreatedBy   string             `json:"createdby,omitempty" bson:"createdby,omitempty"`
}

// UpdateSoldStatus is your comment
func (i *Item) UpdateSoldStatus() {
	i.Sold = !i.Sold
}

// BearerToken is your comment
type BearerToken struct {
	IssuedTo         string `json:"issued_to,omitempty" bson:"issued_to,omitempty"`
	Audience         string `json:"audience,omitempty" bson:"audience,omitempty"`
	UserID           string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Scope            string `json:"scope,omitempty" bson:"scope,omitempty"`
	ExpiresIn        string `json:"expires_in,omitempty" bson:"expires_in,omitempty"`
	Email            string `json:"email,omitempty" bson:"email,omitempty"`
	VerifiedEmail    string `json:"verified_email,omitempty" bson:"verified_email,omitempty"`
	AccessType       string `json:"access_type,omitempty" bson:"access_type,omitempty"`
	Error            string `json:"error,omitempty" bson:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty" bson:"error_description,omitempty"`
}
