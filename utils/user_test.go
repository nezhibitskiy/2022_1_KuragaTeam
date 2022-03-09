package utils

import (
	"errors"
	"log"
	"myapp/models"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		pass string
		err  error
	}{
		{
			name: "NoCharacterAtAll",
			pass: "",
			err:  upErr,
		},
		{
			name: "JustEmptyStringAndWhitespace",
			pass: " \n\t\r\v\f ",
			err:  banErr,
		},
		{
			name: "MixtureOfEmptyStringAndWhitespace",
			pass: "U u\n1\t?\r1\v2\f34",
			err:  banErr,
		},
		{
			name: "MissingUpperCaseString",
			pass: "uu1?1234",
			err:  upErr,
		},
		{
			name: "MissingLowerCaseString",
			pass: "UU1?1234",
			err:  lowErr,
		},
		{
			name: "MissingNumber",
			pass: "Uua?aaaa",
			err:  numErr,
		},
		{
			name: "LessThanRequiredMinimumLength",
			pass: "Uu1?123",
			err:  countErr,
		},
		{
			name: "ValidPassword",
			pass: "Uu1?1234",
			err:  nil,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := ValidatePassword(c.pass)

			assert.Equal(t, c.err, err)
		})
	}
}


func TestUser(t *testing.T) {
	tests := []struct {
		name string
		pass models.User
		err  error
	}{
		{
			name: "NoFieldsAtAll",
			pass: models.User{
				ID:       0,
				Name:     "",
				Email:    "",
				Password: "",
				Salt:     "",
			},
			err: validator.ErrorMap{"Email": validator.ErrorArray{validator.TextErr{Err: errors.New("regular expression mismatch")}},
				"Name":     validator.ErrorArray{validator.TextErr{Err: errors.New("zero value")}},
				"Password": validator.ErrorArray{validator.TextErr{Err: errors.New("less than min")}}},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateUser(&c.pass)

			assert.Equal(t, c.err, err)
		})
	}
}

func TestUserExists(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "email", "password", "salt"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"
	password, _ := HashAndSalt(pwd, salt.String())

	pgxRows := pgxpoolmock.NewRows(columns).AddRow(int64(1), "Ilias@mail.ru", password, salt.String()).ToPgxRows()

	user := models.User{
		ID:       1,
		Name:     "Ilias",
		Email:    "Ilias@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}
	mockPool.EXPECT().Query(gomock.Any(), `SELECT id, email, password, salt FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)

	userPool := &UserPool{
		Pool: mockPool,
	}

	userID, result, err := userPool.IsUserExists(user)

	assert.Equal(t, int64(1), userID)
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestUserNotExists(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "email", "password", "salt"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"

	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	user := models.User{
		ID:       1,
		Name:     "Ilias",
		Email:    "Ilias@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}
	mockPool.EXPECT().Query(gomock.Any(), `SELECT id, email, password, salt FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)

	userPool := &UserPool{
		Pool: mockPool,
	}

	userID, result, err := userPool.IsUserExists(user)

	assert.Equal(t, int64(0), userID)
	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func TestUserUnique(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "name", "email", "password", "salt"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"

	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	user := models.User{
		ID:       2,
		Name:     "Danya",
		Email:    "Danya@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}

	mockPool.EXPECT().Query(gomock.Any(), `SELECT * FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)

	userPool := &UserPool{
		Pool: mockPool,
	}

	result, err := userPool.IsUserUnique(user)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestUserNotUnique(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"id", "name", "email", "password", "salt"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"
	password, _ := HashAndSalt(pwd, salt.String())

	pgxRows := pgxpoolmock.NewRows(columns).AddRow(int64(2), "Danya", "Danya@mail.ru", password, salt.String()).ToPgxRows()

	user := models.User{
		ID:       2,
		Name:     "Danya",
		Email:    "Danya@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}
	mockPool.EXPECT().Query(gomock.Any(), `SELECT * FROM users WHERE email=$1`, user.Email).Return(pgxRows, nil)

	userPool := &UserPool{
		Pool: mockPool,
	}

	result, err := userPool.IsUserUnique(user)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func TestGetUserName(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
	columns := []string{"username"}

	salt, _ := uuid.NewV4()
	pwd := "Pass123321"

	pgxRows := pgxpoolmock.NewRows(columns).AddRow("Ivan").ToPgxRows()

	user := models.User{
		ID:       3,
		Name:     "Ivan",
		Email:    "Ivan@mail.ru",
		Password: pwd,
		Salt:     salt.String(),
	}

	expectedResult := ""

	mockPool.EXPECT().QueryRow(gomock.Any(), `SELECT username FROM users WHERE id=$1`, user.ID).Return(pgxRows)

	if pgxRows.Next() {
		err := pgxRows.Scan(&expectedResult)
		if err != nil {
			log.Fatal(err)
		}
	}
	userPool := &UserPool{
		Pool: mockPool,
	}

	result, errFunc := userPool.GetUserName(user.ID)

	assert.Equal(t, user.Name, result)
	assert.Nil(t, errFunc)
}
