package model

import (
	"fmt"
	"time"
	"log"


	"app/shared/database"

	"gopkg.in/mgo.v2/bson"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	ObjectID  bson.ObjectId `bson:"_id"`
	ID        uint32        `db:"user_id" bson:"id,omitempty"` // Don't use Id, use UserID() instead for consistency with MongoDB
	FirstName string        `db:"first_name" bson:"first_name"`
	LastName  string        `db:"last_name" bson:"last_name"`
	Email     string        `db:"username" bson:"email"`
	Password  string        `db:"password" bson:"password"`
	StatusID  uint8         `db:"salt" bson:"status_id"`
	CreatedAt time.Time     `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" bson:"updated_at"`
	Deleted   uint8         `db:"deleted" bson:"deleted"`
}

// UserStatus table contains every possible user status (active/inactive)
type UserStatus struct {
	ID        uint8     `db:"id" bson:"id"`
	Status    string    `db:"status" bson:"status"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   uint8     `db:"deleted" bson:"deleted"`
}

// UserID returns the user id
func (u *User) UserID() string {
	r := ""

	switch database.ReadConfig().Type {
	case database.TypeMySQL:
		r = fmt.Sprintf("%v", u.ID)
	case database.TypeMongoDB:
		r = u.ObjectID.Hex()
	case database.TypeBolt:
		r = u.ObjectID.Hex()
	}

	return r
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	var err error

	result := User{}

	switch database.ReadConfig().Type {
	case database.TypeMySQL:
		fmt.Printf("VINOD ROCKS")

		err = database.SQL.Get(&result, "SELECT * FROM vn_user WHERE email = ? LIMIT 1", email)
		fmt.Printf("%+v\n", result)

	default:
		err = ErrCode
	}

	log.Println(result)

	return result, standardizeError(err)
}

// UserCreate creates user
func UserCreate(firstName, lastName, email, password string) error {
	var err error

	now := time.Now()

	switch database.ReadConfig().Type {
	case database.TypeMySQL:
		_, err = database.SQL.Exec("INSERT INTO user (first_name, last_name, email, password) VALUES (?,?,?,?)", firstName,
			lastName, email, password)
	case database.TypeMongoDB:
		if database.CheckConnection() {
			session := database.Mongo.Copy()
			defer session.Close()
			c := session.DB(database.ReadConfig().MongoDB.Database).C("user")

			user := &User{
				ObjectID:  bson.NewObjectId(),
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				StatusID:  1,
				CreatedAt: now,
				UpdatedAt: now,
				Deleted:   0,
			}
			err = c.Insert(user)
		} else {
			err = ErrUnavailable
		}
	case database.TypeBolt:
		user := &User{
			ObjectID:  bson.NewObjectId(),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
			StatusID:  1,
			CreatedAt: now,
			UpdatedAt: now,
			Deleted:   0,
		}

		err = database.Update("user", user.Email, &user)
	default:
		err = ErrCode
	}

	return standardizeError(err)
}
