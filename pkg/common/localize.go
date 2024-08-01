package common

const (
	LocalizeHTTP                = "http"
	LocalizeInternalServerError = LocalizeHTTP + ".500"
	LocalizeValidation          = "validation"
	LocalizeEmailAlreadyUsed    = LocalizeValidation + ".email_already_used"
)

var mapLocalize = map[error]string{
	ErrEmailAlreadyUsed: LocalizeEmailAlreadyUsed,
}

func GetLocalizeTag(err error) string {
	if _, seen := mapLocalize[err]; !seen {
		return LocalizeInternalServerError
	}
	return mapLocalize[err]
}
