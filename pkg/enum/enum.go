package enum

const (
	SkuNormal int = iota
	SkuPuzzle
)

const (
	InvalidPathVariable   = "invalid path variable"
	UncorrectPassword     = "password are not correct"
	UserNotFound          = "user not found"
	RegisterFailed        = "register failed"
	RegisterFailedSuggest = "register failed, perhaps user exists ?"
	DataNotFound          = "data not found"
	CreateFailed          = "create data failed"
	UpdateFailed          = "update data failed"
	RemoveFailed          = "remove data failed"
	SaveFailed            = "save data failed"
	AvailableImageSuffix  = "available image suffix only contains ['jpg', 'png', 'jpeg']\nyour suffix: "
	ResetPwdFailed        = "reset password failed"
	CantEditGoogleUser    = "can't edit user's profile, cause by user is google user"
	Unknown               = "unknown"
	IsAdmin               = "you are not admin"
	NoGoogle              = "google user can not access this api"
	InvalidParams         = "invalid params"
)
