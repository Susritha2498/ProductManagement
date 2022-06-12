package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"productManagement/database"
	"productManagement/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("jwt_secret_key")

// To create the JWT while signing in
func CreateJWT(email string) (w string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err == nil {
		return tokenString, nil
	}
	return "", err
}

//To Verify the JWT Token while using APIS
func VerifyToken(tokenString string) (email string, err error) {
	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if token != nil {
		return claims.Email, nil
	}
	return "", err
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest model.Credentials
	var result model.Users
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}

	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&loginRequest)
	defer r.Body.Close()

	if decoderErr != nil {
		fmt.Println(decoderErr)
		// returnErrorResponse(w, r, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if loginRequest.Email == "" {
			errorResponse.Code = http.StatusConflict
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else if loginRequest.Password == "" {
			errorResponse.Code = http.StatusConflict
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var err = database.Collection1.FindOne(ctx, bson.M{
				"email": loginRequest.Email,
			}).Decode(&result)

			defer cancel()

			if err != nil {
				errorResponse.Code = http.StatusBadRequest
				errorResponse.Message = "Email not found. Please register to login"
				returnErrorResponse(w, r, errorResponse)
				return
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(loginRequest.Password))
				if err != nil {
					errorResponse.Code = http.StatusForbidden
					errorResponse.Message = "Password did not match. Try again"
					returnErrorResponse(w, r, errorResponse)
					return
				}

				tokenString, _ := CreateJWT(loginRequest.Email)

				if tokenString == "" {
					errorResponse.Code = http.StatusForbidden
					errorResponse.Message = "Error while creating jwt token"
					returnErrorResponse(w, r, errorResponse)
					return
				}

				var successResponse = model.SuccessResponse{
					Code:    http.StatusOK,
					Message: "Log in successful",
					Response: model.SuccessfulLoginResponse{
						AuthToken: tokenString,
						Email:     loginRequest.Email,
						Username:  result.Username,
					},
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
}

// SignUpUser Used for Signing up the Users
func SignUpUser(w http.ResponseWriter, r *http.Request) {
	var registrationRequest model.Users
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}

	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&registrationRequest)
	defer r.Body.Close()

	if decoderErr != nil {
		fmt.Println(decoderErr)
		// returnErrorResponse(w, r, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if registrationRequest.Username == "" {
			errorResponse.Message = "Username can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else if registrationRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else if registrationRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else if registrationRequest.Phone == 0 {
			errorResponse.Message = "Phone number can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else if registrationRequest.City == "" {
			errorResponse.Message = "City can't be empty"
			returnErrorResponse(w, r, errorResponse)
			return
		} else {
			// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			userfound := database.Collection1.FindOne(context.Background(), bson.M{"email": registrationRequest.Email})

			if userfound != nil {
				var errorResponse = model.ErrorResponse{
					Code: http.StatusUnprocessableEntity, Message: "Email is already registered",
				}
				returnErrorResponse(w, r, errorResponse)
				return

			}
			// defer cancel()
			// mobilefound := database.Collection1.FindOne(ctx, bson.M{"phone": registrationRequest.Phone})
			// if mobilefound != nil {
			// 	var errorResponse = model.ErrorResponse{
			// 		Code: http.StatusUnprocessableEntity, Message: "Mobile number is already registered",
			// 	}
			// 	returnErrorResponse(w, r, errorResponse)
			// 	return
			// }

			var registrationResponse = model.SuccessfulLoginResponse{
				Email: registrationRequest.Email,
			}
			bytes, err := bcrypt.GenerateFromPassword([]byte(registrationRequest.Password), 14)
			if err != nil {
				returnErrorResponse(w, r, errorResponse)
				return
			}
			hashPassword := string(bytes)
			registrationRequest.Password = hashPassword

			_, databaseErr := database.Collection1.InsertOne(context.Background(), bson.M{
				"email":    registrationRequest.Email,
				"password": registrationRequest.Password,
				"username": registrationRequest.Username,
				"phone":    registrationRequest.Phone,
				"city":     registrationRequest.City,
			})

			if databaseErr != nil {
				errorResponse.Code = http.StatusExpectationFailed
				errorResponse.Message = "Error while inserting into database"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			// defer cancel()
			var successResponse = model.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You are registered, login again",
				Response: registrationResponse,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				errorResponse.Message = "Error when converting to json"
				returnErrorResponse(w, r, errorResponse)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(successResponse.Code)
			w.Write(successJSONResponse)
		}
	}
}

// GetUserDetails Used for getting the user details using user token

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var result model.Users
	var errorResponse = model.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Internal server error",
	}
	bearerToken := r.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		returnErrorResponse(w, r, errorResponse)
		return
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var err = database.Collection1.FindOne(ctx, bson.M{
			"email": email,
		}).Decode(&result)

		defer cancel()
		if err != nil {
			returnErrorResponse(w, r, errorResponse)
			return
		} else {
			var successResponse = model.SuccessResponse{
				Code:     http.StatusOK,
				Message:  "You are logged in successfully",
				Response: result.Username,
			}

			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(w, r, errorResponse)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(successJSONResponse)
		}
	}
}

func returnErrorResponse(w http.ResponseWriter, r *http.Request, errorMesage model.ErrorResponse) {
	httpResponse := model.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorMesage.Code)
	w.Write(jsonResponse)
}

/*
func getAllUsers() []primitive.D {
	data, err := database.Collection1.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// type persons map[string]string
	var persons []primitive.D

	for data.Next(context.Background()) {
		var person bson.D
		err := data.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}
	defer data.Close(context.Background())
	return persons

}

//Actual controllers

func GetAllTheUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allUsers := getAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func RegisterAUser(w http.ResponseWriter, r *http.Request) {
	var person model.Users
	json.NewDecoder(r.Body).Decode(&person)

	filter1 := bson.M{"email": person.Email}
	userfound, err := database.Collection1.Find(context.Background(), filter1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return

	}
	if userfound != nil {
		w.WriteHeader(http.StatusConflict)
		var err error
		err = errors.New("Mobile already in use")
		// message := json.RawMessage("Mobile already in use")
		json.NewEncoder(w).Encode(err)
		return

	}

	// filter2 := bson.M{"phone": person.Phone}
	// mobilefound, err := database.Collection1.Find(context.Background(), filter2)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	log.Fatal(err)
	// 	return
	// }

	// if mobilefound != nil {
	// 	w.WriteHeader(http.StatusConflict)
	// 	log.Fatal("This phone is already in use")
	// 	return
	// }

	var hashPassword string
	// var err error
	bytes, err := bcrypt.GenerateFromPassword([]byte(person.Password), 14)
	hashPassword = string(bytes)
	person.Password = hashPassword
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	inserted, err := database.Collection1.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The id of the document that is just inserted is : ", inserted.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	json.NewEncoder(w).Encode(person)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var person model.Credentials
	json.NewDecoder(r.Body).Decode(&person)
	// userF, err = checkLogin(person)
	json.NewEncoder(w).Encode(person)
}

*/
