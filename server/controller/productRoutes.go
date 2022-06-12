package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"productManagement/database"
	"productManagement/model"
	"time"

	// "github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Actual controllers

func GetAllTheProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	var item model.Products
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&item)
	defer r.Body.Close()
	if decoderErr != nil {
		fmt.Println(decoderErr)
	}

	var result model.Users

	Token := r.Header.Get("Authorization")

	email, _ := VerifyToken(Token)
	if email == "" {
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Message = "Do not have authorization"
		returnErrorResponse(w, r, errorResponse)
		return
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = database.Collection1.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)
		defer cancel()
		if err != nil {
			errorResponse.Code = http.StatusUnprocessableEntity
			errorResponse.Message = "Email doesn't exist in the database"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {

			data, err := database.Collection2.Find(ctx, bson.M{"userId": result.ID})
			defer cancel()

			if err != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Cannot get data with this userId"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			var products []primitive.M

			for data.Next(context.Background()) {
				var product bson.M
				err := data.Decode(&product)
				if err != nil {
					log.Fatal(err)
				}
				products = append(products, product)
			}
			defer data.Close(context.Background())

			var successResponse = model.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You successfully fetched the products",
				Response: products,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "Problem in converting to json data"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			w.Write(successJSONResponse)
		}
	}
}

func AddOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var item model.Products
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&item)
	defer r.Body.Close()
	if decoderErr != nil {
		fmt.Println(decoderErr)
	}

	var result model.Users

	Token := r.Header.Get("Authorization")

	email, _ := VerifyToken(Token)
	if email == "" {
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Message = "Do not have authorization"
		returnErrorResponse(w, r, errorResponse)
		return
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = database.Collection1.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)
		defer cancel()
		if err != nil {
			errorResponse.Code = http.StatusUnprocessableEntity
			errorResponse.Message = "Email doesn't exist in the database"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {
			_, dberr := database.Collection2.InsertOne(ctx, bson.M{
				"title":        item.Title,
				"tagline":      item.Tagline,
				"price":        item.Price,
				"rating":       item.Rating,
				"totalRatings": item.TotalRatings,
				"productImage": item.ProductImage,
				"userId":       result.ID,
			})
			defer cancel()

			if dberr != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Cannot add the product into database"
				returnErrorResponse(w, r, errorResponse)
				return
			}

			var successResponse = model.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You have addded the product successfully",
				Response: result.Username,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Problem in converting to json data"
				returnErrorResponse(w, r, errorResponse)
			}
			w.Write(successJSONResponse)
		}
	}
}

func EditOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	var item model.Products
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&item)
	defer r.Body.Close()
	if decoderErr != nil {
		fmt.Println(decoderErr)
		// returnErrorResponse(w, r, errorResponse)
	}

	var result model.Users

	Token := r.Header.Get("Authorization")
	email, _ := VerifyToken(Token)
	fmt.Println(email)
	if email == "" {
		errorResponse.Code = http.StatusNotFound
		errorResponse.Message = "Please login first"
		returnErrorResponse(w, r, errorResponse)
		return
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = database.Collection1.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)
		defer cancel()
		if err != nil {
			errorResponse.Code = http.StatusUnprocessableEntity
			errorResponse.Message = "Email doesn't exist in the database"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {
			params := mux.Vars(r)
			id, err := primitive.ObjectIDFromHex(params["id"])
			if err != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Cannot take the params"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			filter := bson.M{"_id": id, "email": email}
			update := bson.M{"$set": bson.M{
				"title":        item.Title,
				"tagline":      item.Tagline,
				"price":        item.Price,
				"rating":       item.Rating,
				"totalRatings": item.TotalRatings,
				"productImage": item.ProductImage}}

			updated, dberr := database.Collection2.UpdateOne(ctx, filter, update)
			fmt.Println(updated)

			defer cancel()

			if dberr != nil {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "Failed to update the product"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			if updated.MatchedCount == 0 {
				errorResponse.Code = http.StatusUnauthorized
				errorResponse.Message = "Do not have authorization"
				returnErrorResponse(w, r, errorResponse)
				return
			}

			var successResponse = model.SuccessResponse{
				Code:     http.StatusBadRequest,
				Message:  "You have successfully updated the product",
				Response: result.Username,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Problem in converting to json data"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(successJSONResponse)
		}
	}

}

func DeleteAProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	var item model.Products
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&item)
	defer r.Body.Close()
	if decoderErr != nil {
		fmt.Println(decoderErr)
		// returnErrorResponse(w, r, errorResponse)
	}

	var result model.Users

	Token := r.Header.Get("Authorization")
	email, _ := VerifyToken(Token)
	if email == "" {
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Message = "Do not have authorization"
		returnErrorResponse(w, r, errorResponse)
		return
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = database.Collection1.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)
		defer cancel()
		if err != nil {
			errorResponse.Code = http.StatusUnprocessableEntity
			errorResponse.Message = "Email doesn't exist in the database"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {

			params := mux.Vars(r)
			id, err := primitive.ObjectIDFromHex(params["id"])
			if err != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Cannot take the params"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			filter := bson.M{"_id": id, "email": email}

			deleted, dberr := database.Collection2.DeleteOne(ctx, filter)

			defer cancel()

			if dberr != nil {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "Failed to delete the product"
				returnErrorResponse(w, r, errorResponse)
				return
			}

			if deleted.DeletedCount == 0 {
				errorResponse.Code = http.StatusUnauthorized
				errorResponse.Message = "Do not have authorization"
				returnErrorResponse(w, r, errorResponse)
				return
			}

			var successResponse = model.SuccessResponse{
				Code:     http.StatusBadRequest,
				Message:  "You have deleted the product",
				Response: result.Username,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Problem in converting to json data"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(successJSONResponse)
		}
	}

}
