package model

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("Adding product and also user models")
}

type Users struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Phone    int                `json:"phone"`
	City     string             `json:"city"`
}

type Products struct {
	ID           string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string `json:"title"`
	Tagline      string `json:"tagline"`
	Price        string `json:"price"`
	Rating       string `json:"rating"`
	TotalRatings string `json:"totalRatings"`
	// Description    string  `json:"description"`
	// Size           string  `json:"size"`
	// ProductDetails string  `json:"productDetails"`
	ProductImage string `json:"productImage"`
	UserId       string `json:"userId"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Code    int
	Message string
}

type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

type Claims struct {
	Email string
	jwt.StandardClaims
}

type SuccessfulLoginResponse struct {
	Email     string
	AuthToken string
	Username  string
}
