# grocery-go

[![wercker status](https://app.wercker.com/status/db78d475e3bd8f5d6cb36169b53919d9/m "wercker status")](https://app.wercker.com/project/bykey/db78d475e3bd8f5d6cb36169b53919d9)

## An example app using:
- go / golang
- tigertonic / tiger tonic
- gorm
- heroku


## Contains:
- rest api

# heroku
You can see running app on heroku:
### health check
- https://tsz-grocery-go.herokuapp.com/healthCheck/database

### available rest apis
- GET: https://tsz-grocery-go.herokuapp.com/fruits
- GET: https://tsz-grocery-go.herokuapp.com/fruits/search?name={name}
- @POST: https://tsz-grocery-go.herokuapp.com/fruits/{id}/price
- @PUT: https://tsz-grocery-go.herokuapp.com/fruits/addOrOverride
