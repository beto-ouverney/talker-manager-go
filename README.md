# Talker Manager GO LANG with SOLID #

Development of an API to simulate a management of a list of talkers in an event, using simple express command, and no database

# Other Versions of this project with other languages #

[Talker Manager TypeScript NodeJs SOLID Version](https://github.com/beto-ouverney/talker-manager-ts-node-solid)

[Talker Manager Node.Js SOLID Version](https://github.com/beto-ouverney/talker-manager-nodejs)

## Table of contents

- [General view](#general-view)
  - [The Challenge](#the-challenge)
- [The development process](#the-development-process)
  - [Tools used](#tools-used)
  - [Lessons learned](#lessons-learned)
- [Usage](#usage)
- [Author](#author)

## General view

### The challenge

Talker Manager GO LANG with SOLID is a project API to simulate a management of a list of talkers in an event, using simple express command, and no database

**The users must be capable of**

- endpoint GET /talker 
  -> The request must return status 200 and an array with all registered speakers
  -> If there is no registered speaker, the request must return status 200 and an empty array. 

- endpoint GET /talker/:id
  -> The request should return status 200 and a person speaking based on the route id.
  -> If no speaker is found based on the route id, the request must return status 404 with the following body:
  {
  "message": "Pessoa palestrante não encontrada"
  }

- endpoint POST /login
  -> should receive the email and password fields in the body of the request and return a random 16-character token.
  -> O endpoint deverá retornar um código de status 200 com o token gerado e o seguinte corpo:
{
  "token": "7mqaVRXJSp886CGr"
}
  -> O endpoint deve retornar um token aleatório a cada vez que for acessado.

- validations for the /login endpoint
  -> The fields received by the request must be validated and, if the values ​​are invalid, the endpoint must return status code 400 with the respective error message instead of the token.
  -> The validation rules are:

the email field is mandatory;
the email field must have a valid email address;
the password field is mandatory;
the password field must be at least 6 characters long.

- endpoint POST /talker
  -> should be able to add a new speaker person to your file;
  
- endpoint PUT /talker/:id
  -> should be able to edit a speaker person based on the route id, without changing the registered id.
  
- endpoint DELETE /talker/:id
  -> should be able to delete aa speaker person based on the route id.
  
- endpoint GET /talker/search?q=searchTerm
  -> should return an array of speakers that contain the term searched for in the URL's queryParam in their name. Should return status 200  

## The development process

### Tools used

#### Back-end

- Go Lang

### Lessons learned

In this project I could improve my knowledge in back-end, by:

- Make my own router without other frameworks.
- Using SOLID principles
- TDD and unit testing

## Usage

- You will have access to various scripts, that will help you achieving what you want to do.

  - To launch the application, run:
    ```bash
    go run main.go
    ```
## Test
   
    - To launch the tests application, run:
      ```bash
    go test -v ./test
    ```

## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
