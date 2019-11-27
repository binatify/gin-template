package currentuser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assertion := assert.New(t)

	role := UserRole(1)
	user := NewUser("id", "code", role)
	assertion.Equal("id", user.ID)
	assertion.Equal("code", user.Code)
	assertion.Equal(role, user.Role)

	user.SetExtraField("foo", "value")
	assertion.Equal(user.GetExtraField("foo"), "value")
}

func TestUser_InRoles(t *testing.T) {
	assertion := assert.New(t)
	user := NewUser("id", "code", UserRole(1))
	assertion.True(user.InRoles([]UserRole{1}))
}
