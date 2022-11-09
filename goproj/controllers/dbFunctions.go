package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	database "hello/db"
	"hello/models"
)

var UserInfoCollection *mongo.Collection = database.OpenCollection(database.Client, "UserInfo")

func FindAllPost(ctx context.Context) (u []models.Post, e error) {
	var postCollection *mongo.Collection = database.OpenCollection(database.Client, "post")
	result, err := postCollection.Find(ctx, bson.M{})
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find posts")
	}

	err = result.All(ctx, &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func FindAllUsers(ctx context.Context) (u []models.User, e error) {
	var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
	result, err := userCollection.Find(ctx, bson.M{})
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find posts")
	}

	err = result.All(ctx, &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func FindOne(ctx context.Context, title string) (u models.Post, e error) {
	result := postCollection.FindOne(ctx, bson.M{"title": title}).Decode(&u)
	fmt.Println(result)
	fmt.Println(u)
	return u, nil
}
func FindOneUser(ctx context.Context, username string) (u models.User, e error) {
	result := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&u)
	fmt.Println(result)
	return u, nil
}
func FindOneUserAndUnsubscribe(ctx context.Context, username string) (u models.User, e error) {
	result := userCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"issub", "foku"}}}}).Decode(&u)
	fmt.Println(result)
	return u, nil
}
func FindOneUserAndSubscribe(ctx context.Context, username string) (u models.User, e error) {
	result := userCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"issub", "subscribed"}}}}).Decode(&u)
	fmt.Println(result)
	return u, nil
}
func FindOneUserClone(ctx context.Context, username string) (u models.UserInfo, e error) {
	alredy, e := FindUserInfo(ctx, username)
	if e != nil {
		result, _ := FindOneUser(ctx, username)
		var res bool
		if result.IsSub == "subscribed" {
			res = true
		} else {
			res = false
		}
		user := models.UserInfo{
			Username: *result.Username,
			Email:    *result.Email,
			Phone:    *result.Phone,
			IsSub:    res,
			IsLux:    false,
		}
		a, _ := UserInfoCollection.InsertOne(ctx, user)
		fmt.Println(a)
		return user, nil
	} else {
		return alredy, nil
	}

}
func FindUserInfo(ctx context.Context, username string) (u models.UserInfo, e error) {
	result := UserInfoCollection.FindOne(ctx, bson.M{"username": username}).Decode(&u)
	fmt.Println(result)
	return u, nil
}
func UpdateUserInfoUnsub(ctx context.Context, username string) (u models.UserInfo, e error) {
	result := UserInfoCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"issub", false}}}})
	fmt.Println(result)
	return u, nil
}
func UpdateUserInfoSub(ctx context.Context, username string) (u models.UserInfo, e error) {
	result := UserInfoCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"issub", true}}}})
	fmt.Println(result)
	return u, nil
}
func ToLux(ctx context.Context, username string) (u models.UserInfo, e error) {
	result := UserInfoCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"islux", true}}}})
	fmt.Println(result)
	return u, nil
}
func UnLux(ctx context.Context, username string) (u models.UserInfo, e error) {
	result := UserInfoCollection.FindOneAndUpdate(ctx, bson.M{"username": username}, bson.D{{"$set", bson.D{{"islux", false}}}})
	fmt.Println(result)
	return u, nil
}
func FindUserPosts(ctx context.Context, username string) (u []models.Post, e error) {
	result, err := postCollection.Find(ctx, bson.M{"username": username})
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find posts")
	}
	err = result.All(ctx, &u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func FindAllSubs(ctx context.Context) (a []models.User, e error) {
	result, err := userCollection.Find(ctx, bson.M{"issub": "subscribed"})
	if result.Err() != nil {
		return a, fmt.Errorf("failed to find posts")
	}
	err = result.All(ctx, &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
