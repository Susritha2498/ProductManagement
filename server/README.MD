go get -u github.com/gorilla/mux
go get -u github.com/gorilla/handlers
go get go.mongodb.org/mongo-driver/mongo
go get golang.org/x/crypto/bcrypt
go get -u github.com/dgrijalva/jwt-go


"github.com/gorilla/mux"
"github.com/gorilla/handlers"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
"github.com/dgrijalva/jwt-go"
"golang.org/x/crypto/bcrypt"



go build main.go
go run main.go


{
    "username":"Susritha",
    "email":"sus123@gmail.com",
    "password":"12345",
    "age":24,
    "gender":"Female",
    "phone":9876543210,
    "city":"Vijayawada",
    "address":"48-6-6/5,Sri Sai Plaza"   
}


{
	"title":"Van Heusen",
	"tagline":"Men Grey Self Designed Custom Fit Formal Shirt",
	"rating":4.3,
	"totalRatings":"10.7k",
	"description":"100% Original Products
                    Pay on delivery might be available
                    Easy 30 days returns and exchanges
                    Try & Buy might be available",
	"size":"39 40 41 42",
	"productDetails":"Green and White printed casual shirt, has a spread collar, short sleeves, button placket, curved hem, and 1 patch pocket
    Size & Fit : Slim fit .The model (height 6') is wearing a size 40
    Material & Care: Material: Cotton , Machine Wash
    Specifications: Sleeve Length - Short Sleeves      Collar - Spread Collar
                    Fit - Slim Fit                     Print or Pattern Type - Floral
                    Occasion - Casual                  Length - Regular
                    Hemline - Curved                   Placket - Button Placket",
	"productImage":"https://cdn.shopify.com/s/files/1/0752/6435/products/PortoBeige_1d7ebe25-6d3c-47e3-829e-998ce1238782.jpg?v=1631124714"
}



{
	"title":"Van Heusen",
	"tagline":"Men Grey Self Designed Custom Fit Formal Shirt",
	"rating":"4.3",
	"totalRatings":"10.7k",
	"description":"100% Original Products.Pay on delivery might be available Easy 30 days returns and exchanges.Try & Buy might be available",
	"size":"39 40 41 42",
	"productDetails":"Green and White printed casual shirt, has a spread collar, short sleeves, button placket, curved hem, and 1 patch pocket. Size & Fit : Slim fit .The model (height 6') is wearing a size 40 .Material & Care: Material: Cotton , Machine Wash Specifications: Sleeve Length - Short Sleeves, Collar - Spread Collar,  Fit - Slim Fit , Print or Pattern Type - Floral Occasion - Casual, Length - Regular, Hemline - Curved, Placket - Button Placket",
	"productImage":"https://cdn.shopify.com/s/files/1/0752/6435/products/PortoBeige_1d7ebe25-6d3c-47e3-829e-998ce1238782.jpg?v=1631124714"
}