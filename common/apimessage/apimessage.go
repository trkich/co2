package apimessage

const (
	Ok                      		= 99
	DatabaseError           		= 100
	InternalError           		= 101
	JsonParseError          		= 102
	AlreadyExist            		= 103
	UserNameOrPasswordWrong 		= 104
	ExceededLoginAttempts   		= 105
	TokenIsNotValid         		= 106
	NotFound                		= 107
	NotProcessed           		 	= 108
	IoTUnableToCreateThing   		= 201
	IoTUnableToCreateCertificate	= 202
	IoTUnableToCreatePolicy			= 203
	IoTUnableToAttachPolicy			= 204
	IoTUnableToAttachCertificate	= 205
	IoTUnableToDeleteThing			= 206
	IoTUnableToDeleteCertificate	= 207
)

var statusText = map[int]string{
	DatabaseError:           		"DATABASE_ERROR",
	InternalError:           		"INTERNAL_ERROR",
	JsonParseError:          		"JSON_PARSE_ERROR",
	AlreadyExist:            		"ALREADY_EXIST",
	ExceededLoginAttempts:   		"EXCEEDED_LOGIN_ATTEMPTS",
	UserNameOrPasswordWrong: 		"USERNAME_OR_PASSWORD_WRONG",
	Ok:                      		"OK",
	TokenIsNotValid:         		"TOKEN_IS_NOT_VALID",
	NotFound:                		"NOT_FOUND",
	NotProcessed:            		"NOT_PROCESSED",
	IoTUnableToCreateThing:			"IOT_UNABLE_TO_CREATE_THING",
	IoTUnableToCreateCertificate:	"IOT_UNABLE_TO_CREATE_CERTIFICATE",
	IoTUnableToCreatePolicy:		"IOT_UNABLE_TO_CREATE_POLICY",
	IoTUnableToAttachPolicy:		"IOT_UNABLE_TO_ATTACH_POLICY",
	IoTUnableToAttachCertificate:	"IOT_UNABLE_TO_ATTACH_CERTIFICATE",
	IoTUnableToDeleteThing: 		"IOT_UNABLE_TO_DELETE_THING",
	IoTUnableToDeleteCertificate: 	"IOT_UNABLE_TO_DELETE_CERTIFICATE",

}

func StatusText(code int) string {
	return statusText[code]
}
