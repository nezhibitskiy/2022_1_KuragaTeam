package constants

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"myapp/internal/models"
	"net/http"
)

var (
	ErrLetter                  = errors.New("at least one letter is required")
	ErrNum                     = errors.New("at least one digit is required")
	ErrCount                   = errors.New("at least eight characters long is required")
	ErrBan                     = errors.New("password uses unavailable symbols")
	ErrWrongData               = errors.New("wrong data")
	EmailIsNotUnique           = errors.New("email is not unique")
	WrongToken                 = errors.New("wrong payment token")
	WrongAmount                = errors.New("wrong amount")
	WringCountPaymentsForToken = errors.New("wrong count payments for token")
	NoSubscription             = errors.New("no subscription")
)

const (
	UserObjectsBucketName = "avatars"
	DefaultImage          = "default_avatar.webp"

	UserCanBeLoggedIn      = "User can be logged in"
	UserCreated            = "User created"
	SessionRequired        = "Session required"
	UserIsUnauthorized     = "User is unauthorized"
	UserIsLoggedOut        = "User is logged out"
	FileTypeIsNotSupported = "File type is not supported"
	ProfileIsEdited        = "Profile is edited"
	LikeIsEdited           = "Like is edited"
	LikeIsRemoved          = "Like is removed"
	NoRequestId            = "No RequestID in context"
	NoMovieId              = "No MovieID in context"
)

const (
	SignupURL            = "/api/v1/signup"
	LoginURL             = "/api/v1/login"
	LogoutURL            = "/api/v1/logout"
	ProfileURL           = "/api/v1/profile"
	EditURL              = "/api/v1/edit"
	AvatarURL            = "/api/v1/avatar"
	CsrfURL              = "/api/v1/csrf"
	AuthURL              = "/api/v1/auth"
	CheckURL             = "/api/v1/check"
	AddLikeUrl           = "/api/v1/like"
	RemoveLikeUrl        = "/api/v1/dislike"
	LikesUrl             = "/api/v1/likes"
	UserRatingUrl        = "/api/v1/userRating"
	PaymentIsCreated     = "Payment is created"
	UnsupportedMediaType = "Unsupported media type"
	PaymentsTokenURL     = "/api/v1/payments/token"
	PaymentURL           = "/api/v1/payment"
	SubscribeURL         = "/api/v1/subscribe"
)

var (
	IMAGE_TYPES = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}
)

const (
	MoviesSearchLimit  = 3
	PersonsSearchLimit = 3
	Price              = 2
)

func RespError(ctx echo.Context, logger *zap.SugaredLogger, requestID, errorMsg string, status int) error {
	logger.Error(
		zap.String("ID", requestID),
		zap.String("ERROR", errorMsg),
		zap.Int("ANSWER STATUS", status),
	)
	resp, err := easyjson.Marshal(&models.Response{
		Status:  status,
		Message: errorMsg,
	})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSONBlob(status, resp)
}
