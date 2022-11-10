# search-car


## routes

## Run

to start use this project in local enviroment download the project and using terminal do the follow command in the paste of project:
```sh
`go run main.go`
project wil be available at: http://http://localhost:1323
```

### validate user
    using local:
    POST: http://http://localhost:1323/login
    format data: 

   ```sh
   `{
        "username": "teste",
        "password":"testesenha"
    }`
    ```
    the rote will returnn the userId

### register
    using local:
    POST: http://http://localhost:1323/cadastro
    format data: 

   ```sh
   `{
        "name":"elisa",
        "username": "elisasouza",
        "password":"elisa",
        "email": "elisa@elisa.com",
        "telefone":""
    }`
    ```
    the rote will return the userId

### change permission
    using local:
    POST: http://http://localhost:1323/user/change-permission
    format data: 

   ```sh
   `{
        "userId":"ab2c-1d",
        "permission": "1",
    }`
    ```
    the rote will return "ok"

### get user info
    using local:
    POST: http://http://localhost:1323/user/getInformation
    format data: 

   ```sh
   `{
        "userId":"ab2c-1d",
    }`
    ```
    the rote will return all information of userId

### change permission user
    using local:
    POST: http://http://localhost:1323/user/change-permission
    format data: 

   ```sh
   `{
    "userId": "e30d9eef-3e35-4571-a7f0-9b9221c79b5f",
    "permission": "1"
    }`
    ```
    the rote will return "ok"

##ROUTES CAR

### create car
    using local:
    POST: http://http://localhost:1323/car/create
    format data: 

   ```sh
   `{
    "placa":"placateste1",
    "state":"SP",
    "municipio":"Campinas",
    "cor":"preto",
    "marcaEModelo":"FIAT HONDA",
    "anoDoCarro":"2002",
    "chassi" :"aleatorio",
    "renavam":"aleatorio",
    "nome": "Elisa Castro de Souza",
    "userId": "30d9eef-3e35-4571-a7f0-9b9221c79b5f"
    }`
    ```
    the rote will return "ok"

### Get car by state
    using local:
    POST: http://http://localhost:1323/car/state/:state

    where is state put the state like:
    http://http://localhost:1323/car/state/SP

    the rote will return the cars that are in SP 

### Get car by state
    using local:
    POST: http://http://localhost:1323/car/placa/:placa

    where is placa put the placa like:
    http://http://localhost:1323/car/placa/3010lia

    the rote will return the user that register the placa