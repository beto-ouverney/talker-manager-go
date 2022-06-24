package errorschema

//StatusMsgError represents the status and message of the error response
type StatusMsgError struct {
	Status  int
	Message string
}

//ErrorResponse represents all errors response from the API
var ErrorResponse = map[string]StatusMsgError{
	"readFails": {
		Status:  500,
		Message: "internal failure to read data",
	},
	"writeFails": {
		Status:  500,
		Message: "internal failure to write data",
	},
	"talkerNotFound": {
		Status:  404,
		Message: "Speaker not found",
	},
	"invalidId": {
		Status:  400,
		Message: "ID must be a number",
	},
	"emailIsRequired": {
		Status:  400,
		Message: `The "email" field is required`,
	},
	"invalidEmail": {
		Status:  400,
		Message: `The "email" must have the format "email@email.com"`,
	},
	"passwordIsRequired": {
		Status:  400,
		Message: `The password is required`,
	},
	"invalidPassword": {
		Status:  400,
		Message: `The "password" must be at least 6 characters long`,
	},
	"tokenNotFound": {
		Status:  401,
		Message: "Token not found",
	},
	"invalidToken": {
		Status:  401,
		Message: "Token invalid",
	},
	"nameIsRequired": {
		Status:  400,
		Message: `The "name" field is required`,
	},
	"invalidName": {
		Status:  400,
		Message: `The "name" must be at least 3 caracters long`,
	},
	"ageIsRequired": {
		Status:  400,
		Message: `The "age" field is required`,
	},
	"invalidAge": {
		Status:  400,
		Message: "The speaker must be of legal age",
	},
	"talkIsRequired": {
		Status:  400,
		Message: `The "talk" field is mandatory and "watchedAt" and "rate" cannot be empty`,
	},
	"invalidWathedAt": {
		Status:  400,
		Message: `The "watchedAt" field must have the format "dd/mm/yyyy"`,
	},
	"invalidRate": {
		Status:  400,
		Message: `The "rate" field must be an integer from 1 to 5`,
	},
}
