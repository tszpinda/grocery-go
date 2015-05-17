# grocery-go

[![wercker status](https://app.wercker.com/status/a7ae5d7b6ecbd81c0ad63780438d8392/m "wercker status")](https://app.wercker.com/project/bykey/a7ae5d7b6ecbd81c0ad63780438d8392)

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
