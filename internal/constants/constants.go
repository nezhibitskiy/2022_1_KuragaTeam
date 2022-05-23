package constants

import "errors"

var (
	ErrLetter                     = errors.New("at least one letter is required")
	ErrNum                        = errors.New("at least one digit is required")
	ErrCount                      = errors.New("at least eight characters long is required")
	ErrBan                        = errors.New("password uses unavailable symbols")
	ErrWrongData                  = errors.New("wrong data")
	ErrEmailIsNotUnique           = errors.New("email is not unique")
	ErrWrongToken                 = errors.New("wrong payment token")
	ErrWrongAmount                = errors.New("wrong amount")
	ErrWringCountPaymentsForToken = errors.New("wrong count payments for token")
	ErrNoSubscription             = errors.New("no subscription")
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
	NoRequestID            = "No RequestID in context"
	NoMovieID              = "No MovieID in context"
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
	AddLikeURL           = "/api/v1/like"
	RemoveLikeURL        = "/api/v1/dislike"
	LikesURL             = "/api/v1/likes"
	UserRatingURL        = "/api/v1/userRating"
	PaymentIsCreated     = "Payment is created"
	UnsupportedMediaType = "Unsupported media type"
	PaymentsTokenURL     = "/api/v1/payments/token"
	PaymentURL           = "/api/v1/payment"
	SubscribeURL         = "/api/v1/subscribe"
)

var (
	ImageTypes = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}
)

const (
	MoviesSearchLimit  = 3
	PersonsSearchLimit = 3
	Price              = 2
)
