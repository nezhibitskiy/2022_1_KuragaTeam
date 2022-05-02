// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"myapp/internal"
	compilations "myapp/internal/microservices/compilations/proto"
	movie "myapp/internal/microservices/movie/proto"
	"myapp/internal/persons"
	"sync"
)

// Ensure, that MockPersonsStorage does implement persons.Storage.
// If this is not the case, regenerate this file with moq.
var _ persons.Storage = &MockPersonsStorage{}

// MockPersonsStorage is a mock implementation of persons.Storage.
//
// 	func TestSomethingThatUsesStorage(t *testing.T) {
//
// 		// make and configure a mocked persons.Storage
// 		mockedStorage := &MockPersonsStorage{
// 			FindPersonFunc: func(text string) (*compilations.PersonCompilation, error) {
// 				panic("mock out the FindPerson method")
// 			},
// 			FindPersonByPartialFunc: func(text string) (*compilations.PersonCompilation, error) {
// 				panic("mock out the FindPersonByPartial method")
// 			},
// 			GetByMovieIDFunc: func(id int) ([]*movie.PersonInMovie, error) {
// 				panic("mock out the GetByMovieID method")
// 			},
// 			GetByPersonIDFunc: func(id int) (*internal.Person, error) {
// 				panic("mock out the GetByPersonID method")
// 			},
// 		}
//
// 		// use mockedStorage in code that requires persons.Storage
// 		// and then make assertions.
//
// 	}
type MockPersonsStorage struct {
	// FindPersonFunc mocks the FindPerson method.
	FindPersonFunc func(text string) (*compilations.PersonCompilation, error)

	// FindPersonByPartialFunc mocks the FindPersonByPartial method.
	FindPersonByPartialFunc func(text string) (*compilations.PersonCompilation, error)

	// GetByMovieIDFunc mocks the GetByMovieID method.
	GetByMovieIDFunc func(id int) ([]*movie.PersonInMovie, error)

	// GetByPersonIDFunc mocks the GetByPersonID method.
	GetByPersonIDFunc func(id int) (*internal.Person, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindPerson holds details about calls to the FindPerson method.
		FindPerson []struct {
			// Text is the text argument value.
			Text string
		}
		// FindPersonByPartial holds details about calls to the FindPersonByPartial method.
		FindPersonByPartial []struct {
			// Text is the text argument value.
			Text string
		}
		// GetByMovieID holds details about calls to the GetByMovieID method.
		GetByMovieID []struct {
			// ID is the id argument value.
			ID int
		}
		// GetByPersonID holds details about calls to the GetByPersonID method.
		GetByPersonID []struct {
			// ID is the id argument value.
			ID int
		}
	}
	lockFindPerson          sync.RWMutex
	lockFindPersonByPartial sync.RWMutex
	lockGetByMovieID        sync.RWMutex
	lockGetByPersonID       sync.RWMutex
}

// FindPerson calls FindPersonFunc.
func (mock *MockPersonsStorage) FindPerson(text string) (*compilations.PersonCompilation, error) {
	if mock.FindPersonFunc == nil {
		panic("MockPersonsStorage.FindPersonFunc: method is nil but Storage.FindPerson was just called")
	}
	callInfo := struct {
		Text string
	}{
		Text: text,
	}
	mock.lockFindPerson.Lock()
	mock.calls.FindPerson = append(mock.calls.FindPerson, callInfo)
	mock.lockFindPerson.Unlock()
	return mock.FindPersonFunc(text)
}

// FindPersonCalls gets all the calls that were made to FindPerson.
// Check the length with:
//     len(mockedStorage.FindPersonCalls())
func (mock *MockPersonsStorage) FindPersonCalls() []struct {
	Text string
} {
	var calls []struct {
		Text string
	}
	mock.lockFindPerson.RLock()
	calls = mock.calls.FindPerson
	mock.lockFindPerson.RUnlock()
	return calls
}

// FindPersonByPartial calls FindPersonByPartialFunc.
func (mock *MockPersonsStorage) FindPersonByPartial(text string) (*compilations.PersonCompilation, error) {
	if mock.FindPersonByPartialFunc == nil {
		panic("MockPersonsStorage.FindPersonByPartialFunc: method is nil but Storage.FindPersonByPartial was just called")
	}
	callInfo := struct {
		Text string
	}{
		Text: text,
	}
	mock.lockFindPersonByPartial.Lock()
	mock.calls.FindPersonByPartial = append(mock.calls.FindPersonByPartial, callInfo)
	mock.lockFindPersonByPartial.Unlock()
	return mock.FindPersonByPartialFunc(text)
}

// FindPersonByPartialCalls gets all the calls that were made to FindPersonByPartial.
// Check the length with:
//     len(mockedStorage.FindPersonByPartialCalls())
func (mock *MockPersonsStorage) FindPersonByPartialCalls() []struct {
	Text string
} {
	var calls []struct {
		Text string
	}
	mock.lockFindPersonByPartial.RLock()
	calls = mock.calls.FindPersonByPartial
	mock.lockFindPersonByPartial.RUnlock()
	return calls
}

// GetByMovieID calls GetByMovieIDFunc.
func (mock *MockPersonsStorage) GetByMovieID(id int) ([]*movie.PersonInMovie, error) {
	if mock.GetByMovieIDFunc == nil {
		panic("MockPersonsStorage.GetByMovieIDFunc: method is nil but Storage.GetByMovieID was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockGetByMovieID.Lock()
	mock.calls.GetByMovieID = append(mock.calls.GetByMovieID, callInfo)
	mock.lockGetByMovieID.Unlock()
	return mock.GetByMovieIDFunc(id)
}

// GetByMovieIDCalls gets all the calls that were made to GetByMovieID.
// Check the length with:
//     len(mockedStorage.GetByMovieIDCalls())
func (mock *MockPersonsStorage) GetByMovieIDCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGetByMovieID.RLock()
	calls = mock.calls.GetByMovieID
	mock.lockGetByMovieID.RUnlock()
	return calls
}

// GetByPersonID calls GetByPersonIDFunc.
func (mock *MockPersonsStorage) GetByPersonID(id int) (*internal.Person, error) {
	if mock.GetByPersonIDFunc == nil {
		panic("MockPersonsStorage.GetByPersonIDFunc: method is nil but Storage.GetByPersonID was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockGetByPersonID.Lock()
	mock.calls.GetByPersonID = append(mock.calls.GetByPersonID, callInfo)
	mock.lockGetByPersonID.Unlock()
	return mock.GetByPersonIDFunc(id)
}

// GetByPersonIDCalls gets all the calls that were made to GetByPersonID.
// Check the length with:
//     len(mockedStorage.GetByPersonIDCalls())
func (mock *MockPersonsStorage) GetByPersonIDCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockGetByPersonID.RLock()
	calls = mock.calls.GetByPersonID
	mock.lockGetByPersonID.RUnlock()
	return calls
}
