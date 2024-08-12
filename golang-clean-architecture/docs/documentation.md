The public API request:

from the client an api call. public router and then controller, response from the controller to the client

The private API request:

JWT auth middleware = > from client to the protected router, then auth middleware and the controller. response from controller to the client

middleware for access token validation 

public and private routers and then middleware for the protected ones.

the same is true for the public ones, without the middleware

protectedRouter := gin.Group("")

protectedRouter.Use(someMiddleware)

RouterFunction(..., other arguments, prrouter)

the controller invokes functions on the usecases, so as the routers invoke functions on the controllers with some inputs.

so what does the domain have, the interfaces for the usecase and reposi
tory and also the struct for the user and tasks




first finish up eski the sign up and login. the unprotected routers malet nw
//the sign up will have two middle_ware functions....hashing the passwords
//the log in will have two middle_ware functions to compare the passwords and also generate a jwt token