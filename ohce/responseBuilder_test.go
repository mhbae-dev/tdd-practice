package main

import "testing"

func TestNameResponseBuilder(t *testing.T) {
	t.Run("BuildResponse returns good morning user", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		clock := new(MockClock)

		clock.On("GetTime").Return(10)
		got := responseBuilder.BuildNameResponse("John", clock)
		want := "¡Good morning John!\n"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

	t.Run("BuildResponse returns good afternoon user", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		clock := new(MockClock)

		clock.On("GetTime").Return(13)
		got := responseBuilder.BuildNameResponse("John", clock)
		want := "¡Good afternoon John!\n"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

	t.Run("BuildResponse returns good evening user", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		clock := new(MockClock)

		clock.On("GetTime").Return(20)
		got := responseBuilder.BuildNameResponse("John", clock)
		want := "¡Good evening John!\n"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

	t.Run("BuildResponse returns good night user", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		clock := new(MockClock)

		clock.On("GetTime").Return(5)
		got := responseBuilder.BuildNameResponse("John", clock)
		want := "¡Good night John!\n"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

}

func TestNormalResponseBuilder(t *testing.T) {
	t.Run("returns reverse of input string", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		input := "hello"
		got := responseBuilder.BuildNormalResponse(input)
		want := "olleh"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

	t.Run("returns confirmation of palindrome", func(t *testing.T) {
		t.Parallel()
		responseBuilder := &ResponseBuilder{}
		input := "oto"
		got := responseBuilder.BuildNormalResponse(input)
		want := "oto\n" + "¡Nice palindrome!"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})

	//this might be a bad test and design as the test is dependent that BuildNameRespones has been called
	//prior to BuildNormalResponse to persist the name field
	t.Run("returns Goodbye if input is Stop!", func(t *testing.T) {
		t.Parallel()
		clock := new(MockClock)
		responseBuilder := &ResponseBuilder{}
		clock.On("GetTime").Return(10)
		input := "Stop!"
		name := "John"
		responseBuilder.BuildNameResponse(name, clock)
		got := responseBuilder.BuildNormalResponse(input)
		want := "Goodbye John"
		if got != want {
			t.Errorf("Expected BuildResponse to be %v, got %v", want, got)
		}
	})
}
