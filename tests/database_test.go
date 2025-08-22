package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/alfianyulianto/go-room-managament/app"
)

func TestConnection(t *testing.T) {
	db := app.NewDB()
	assert.NotNil(t, db)
}
