package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOhceAcceptance(t *testing.T) {

	clock := new(MockClock)
	responseBuilder := &ResponseBuilder{}

	ohce := Ohce{responseBuilder, clock}
	clock.On("GetTime").Return(6)
	inputName := "John"

	want := "¡Good morning John!\n"
	got := ohce.Greet(inputName)
	assert.Equal(t, want, got)

	got1 := ohce.Respond("hello")
	want1 := "olleh"
	assert.Equal(t, want1, got1)

	got2 := ohce.Respond("oto")
	want2 := "oto\n¡Nice palindrome!"
	assert.Equal(t, want2, got2)

	got3 := ohce.Respond("stop")
	want3 := "pots"
	assert.Equal(t, want3, got3)

	got4 := ohce.Respond("Stop!")
	want4 := "Goodbye John"
	assert.Equal(t, want4, got4)
}
