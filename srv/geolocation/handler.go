package main

import (
	"context"
	"github.com/phuonghau98/walkie3/srv/geolocation/proto/geolocation"
	"github.com/phuonghau98/walkie3/srv/user/proto/user"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/metadata"
	"time"
)

func (s *Service) UpdateLocation (ctx context.Context, location *geolocation.Location) (*geolocation.Empty, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	count, err := s.repo.dbClient.Collection("geolocation").CountDocuments(ctx, bson.D{{"userID", foundUser.Id}})
	if count == 0 {
		_, err = s.repo.dbClient.Collection("geolocation").InsertOne(ctx, LocationDocument{
			UserID:    foundUser.Id,
			Longitude: location.Longitude,
			Latitude: location.Latitude,
			UpdatedAt: time.Now(),
			SharedIDs: []string{},
		})
		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.repo.dbClient.Collection("geolocation").UpdateOne(ctx, bson.D{{"userID", foundUser.Id}}, bson.D{
			{"$set", bson.M{
				"updatedAt": time.Now(),
				"longitude": location.Longitude,
				"latitude": location.Latitude,
			}},
		})
	}
	return &geolocation.Empty{}, nil
}

func (s *Service) GetLocation (ctx context.Context, empty *geolocation.Empty) (*geolocation.LocationResponse, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	var result geolocation.LocationResponse
	conditions := bson.D{{
		"$or", bson.A{
			bson.D{{
				"sharedIDs", bson.D{{
					"$in", bson.A{foundUser.Id},
				}},
			}},
			bson.D{{
				"userID", foundUser.Id,
			}},
		},
	}}
	cur, err := s.repo.dbClient.Collection("geolocation").Find(ctx, conditions)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var tmpLocation LocationDocument
		_ = cur.Decode(&tmpLocation)
		result.Locations = append(result.Locations, &geolocation.Location{
			Longitude: tmpLocation.Longitude,
			Latitude: tmpLocation.Latitude,
			UserID: tmpLocation.UserID,
			UpdatedAt: tmpLocation.UpdatedAt.Unix(),
			SharedIDs: tmpLocation.SharedIDs,
		})
	}
	return &result, nil
}

func (s *Service) ShareLocation(ctx context.Context, req *geolocation.LocationRequest) (*geolocation.Empty, error) {
	headers, _ := metadata.FromIncomingContext(ctx)
	foundUser, err := s.userClient.GetUserFromToken(ctx, &user.TokenInfo{Token:headers["authorization"][0]})
	if err != nil {
		return nil, err
	}
	var foundLocation *LocationDocument
	if req.UserID != "" {
		err := s.repo.dbClient.Collection("geolocation").FindOne(ctx, bson.D{{"userID", foundUser.Id}}).Decode(&foundLocation)
		if err != nil {
			return nil, err
		}
		if req.IsRemove == 1 {
			for i := 0; i < len(foundLocation.SharedIDs); i++ {
				if foundLocation.SharedIDs[i] == req.UserID {
					foundLocation.SharedIDs = append(foundLocation.SharedIDs[:i], foundLocation.SharedIDs[i+1:]...)
					i--
				}
				s.repo.dbClient.Collection("geolocation").FindOneAndUpdate(ctx, bson.D{{"userID", foundUser.Id}}, bson.M{
					"$set": bson.M{"sharedIDs": foundLocation.SharedIDs},
				})
			}
		} else {
			if Find(foundLocation.SharedIDs, req.UserID) == false {
				s.repo.dbClient.Collection("geolocation").FindOneAndUpdate(ctx, bson.D{{"userID", foundUser.Id}}, bson.M{
					"$set": bson.M{"sharedIDs": append(foundLocation.SharedIDs, req.UserID)},
				})
			}
		}
	}

	return &geolocation.Empty{}, nil
}

func Find (slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}