package error

//StatusMsgError represents the status and message of the error response
type StatusMsgError struct {
	Status  int
	Message string
}

//ErrorResponse represents all errors response from the API
var ErrorResponse = map[string]StatusMsgError{
	"readFails": {
		Status:  500,
		Message: "falha interna na leitura de dados",
	},
	"writeFails": {
		Status:  500,
		Message: "falha interna na escrita dos dados",
	},
	"talkerNotFound": {
		Status:  404,
		Message: "Pessoa palestrante não encontrada",
	},
	"invalidId": {
		Status:  400,
		Message: "Id deve ser um número",
	},
	"emailIsRequired": {
		Status:  400,
		Message: `O campo "email" é obrigatório`,
	},
	"invalidEmail": {
		Status:  400,
		Message: `O "email" deve ter o formato "email@email.com"`,
	},
	"passwordIsRequired": {
		Status:  400,
		Message: `O campo "password" é obrigatório`,
	},
	"invalidPassword": {
		Status:  400,
		Message: `O "password" deve ter pelo menos 6 caracteres`,
	},
	"tokenNotFound": {
		Status:  401,
		Message: "Token não encontrado",
	},
	"invalidToken": {
		Status:  401,
		Message: "Token inválido",
	},
	"nameIsRequired": {
		Status:  400,
		Message: `O campo "name" é obrigatório`,
	},
	"invalidName": {
		Status:  400,
		Message: `O "name" deve ter pelo menos 3 caracteres`,
	},
	"ageIsRequired": {
		Status:  400,
		Message: `O campo "age" é obrigatório`,
	},
	"invalidAge": {
		Status:  400,
		Message: "A pessoa palestrante deve ser maior de idade",
	},
	"talkIsRequired": {
		Status:  400,
		Message: `O campo "talk" é obrigatório e "watchedAt" e "rate" não podem ser vazios`,
	},
	"invalidWathedAt": {
		Status:  400,
		Message: `O campo "watchedAt" deve ter o formato "dd/mm/aaaa"`,
	},
	"invalidRate": {
		Status:  400,
		Message: `O campo "rate" deve ser um inteiro de 1 à 5`,
	},
}
