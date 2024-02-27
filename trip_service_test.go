package trip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldSeeFriendsTrips_WhenUserLoggedIn(t *testing.T) {
	service := Service{
		tripDao: &Dao{},
	}

	loggedUser := &User{}
	friends := []User{
		{},
	}
	tripsByUser := []Trip{}
	trips := []Trip{}

	tripsByLoggedUser, err := service.GetTripsByLoggedUser(loggedUser, friends, tripsByUser, nil, trips)

	assert.NoError(t, err)
	assert.Equal(t, []Trip{}, tripsByLoggedUser)
}
