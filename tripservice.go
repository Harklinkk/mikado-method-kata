package trip

import (
	"github.com/pkg/errors"
)

var session *UserSession

type Trip struct {
}

type Service struct {
	tripDao TripDao
}

type TripDao interface {
	FindTripByUser(user *User) ([]Trip, error)
}

func (service *Service) GetTripByUser(user *User) ([]Trip, error) {
	tripByUser, err := service.tripDao.FindTripByUser(user)

	var trips []Trip
	friends, err := user.Friends()
	if err != nil {
		return trips, err
	}
	loggedUser, err := session.GetLoggedUser()
	if err != nil {
		return trips, err
	}
	return service.GetTripsByLoggedUser(loggedUser, friends, tripByUser, err, trips)
}

func (service *Service) GetTripsByLoggedUser(loggedUser *User, friends []User, tripByUser []Trip, err error, trips []Trip) ([]Trip, error) {
	var isFriend bool
	if loggedUser != nil {
		for _, friend := range friends {
			if *loggedUser == friend {
				isFriend = true
				break
			}
		}
		if isFriend {
			return tripByUser, err
		}
		return trips, err
	} else {
		return trips, errors.New("user not logged in")
	}
}

type UserSession struct {
}

func (userSession *UserSession) GetLoggedUser() (*User, error) {
	return nil, errors.New("UserSession.GetLoggedUser() should not be called in an unit test")
}

type User struct {
}

func (user *User) Friends() ([]User, error) {
	var friends []User
	return friends, nil
}

type Dao struct {
}

func (dao *Dao) FindTripByUser(user *User) ([]Trip, error) {
	return nil, errors.New("TripDAO should not be invoked on an unit test.")
}
