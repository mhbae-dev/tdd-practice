package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockClock struct {
	mock.Mock
}

func (c *MockClock) GetTime() int {
	arg := c.Called()
	return arg.Int(0)
}

type MockResponseBuilder struct {
	mock.Mock
}

func (rb *MockResponseBuilder) BuildNameResponse(input string, clock ClockI) string {
	arg := rb.Called(input, clock)
	return arg.String(0)
}

func (rb *MockResponseBuilder) BuildNormalResponse(input string) string {
	arg := rb.Called(input)
	return arg.String(0)
}

func TestOhce(t *testing.T) {
	t.Run("Should greet user ", func(t *testing.T) {
		t.Parallel()
		clock := new(MockClock)
		responseBuilder := new(MockResponseBuilder)
		ohce := Ohce{responseBuilder, clock}

		responseBuilder.On("BuildNameResponse", "John", clock).Return("¡Good morning John!")

		want := "¡Good morning John!"
		got := ohce.Greet("John")
		responseBuilder.AssertCalled(t, "BuildResponse", "John", clock)

		assert.Equal(t, want, got)
		responseBuilder.AssertExpectations(t)
	})

	t.Run("Should reverse input", func(t *testing.T) {
		t.Parallel()
		clock := new(MockClock)
		responseBuilder := new(MockResponseBuilder)
		ohce := Ohce{responseBuilder, clock}

		responseBuilder.On("BuildNormalResponse", "hello").Return("olleh")

		want := "olleh"
		got := ohce.Respond("hello")
		responseBuilder.AssertCalled(t, "BuildNormalResponse", "hello")

		assert.Equal(t, want, got)
		responseBuilder.AssertExpectations(t)
	})
}
