package common

const (
	LocalizeHTTP                = "http"
	LocalizeInternalServerError = LocalizeHTTP + ".500"
	LocalizeValidation          = "validation"
	LocalizeEmailAlreadyUsed    = LocalizeValidation + ".email_already_used"
	LocalizeRecordNotFound      = LocalizeValidation + ".not_found"
)

var mapLocalize = map[error]string{
	ErrEmailAlreadyUsed: LocalizeEmailAlreadyUsed,
	ErrRecordNotFound:   LocalizeRecordNotFound,
}

func GetLocalizeTag(err error) string {
	if _, seen := mapLocalize[err]; !seen {
		return err.Error()
	}
	return mapLocalize[err]
}
