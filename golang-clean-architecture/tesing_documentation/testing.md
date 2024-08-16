# Task 08 - Testing the Simple task management API and Conclusive Remarks

## Progress Analysis and Overall completion

* This whole api practice application started off with a simple task management application using in memory storage just to understand the basic concepts of golang and in order to familiarize ones self with the functionalities and frameworks that Golang provides.

* Then we incorporated the usage of databases (specifically, mongodb) to our application and tweaked the implementation a bit to accomodate simple CRUD operations to and from the DB.

* What is an application without a proper authentication and authorization? Just a mere mockery of what the language offers. We added a simple authentication and authorization functionality for users along with data validation and proper user information storage.

* It doesn't end here. we refined our folder structure and logic using the concept of clean architecture (trying to adhere to industrial best practices).

* We then wrote the test script to test the functionality of the application and verify that the application behaves as intended under different conditions. Which is the main topic of analysis for this particular documentation.

## About the Testing Stage of the Application

### Repository Testing

**Assumption:** The database setup (using mongo db for this particular application) cannot be a point of failure and is assumed to always work fine.

* I used dependencies such as *suite* and *testing* for this stage.
* And also a temporary database and collection that is tear down at the end of the repository test automatically.
* In the setup stage of the suite, the database along with a collection is created and a connection is made to the mongodb. 
* different operations (both user and task operations) are made on the test functions and behaviour is analysed. *Test dependency* was not faced at this stage. it will be discussed later on the use cases.

### Usecase Testing

**Assumption:** The repository functionalities are now assumed to be tested and work fine. I used the *mockery* library to setup the mock functions that are used in place of the real repository functions that used the database connection for any operation. At this stage of testing, using mock functions will get rid of an external dependency and results in a replicable and fast outcomes when running the test.

**Mock Repository and Mock Usecase from an Interface**, using the command, ```go install mockery``` and
```mockery --all``` first navigating to the domain folder. this creates mock functions for all the interfaces defined in the domain.go file.

* I used dependencies such as *suite* and *testing* for this stage.
* called the mock functions with a simulated response whenever a tested use case function involves calling a repository function
* Faced a test dependency issue when trying to run the test, so I used SetupTest instead of a SetupSuite function, the former setup runs each time when a new test function is about to be run.
* Both positive and negative tests were run in a particular format and the outputs are asserted to return a certain response.

**Example Test Functions: Just a snapshot**

```go
func (suite *UserTestSuite) TestUserRegister_DatabaseError() {

	user := &domain.User{Email: "test@example.com", Password: "password123", Role: "admin"}
	suite.mockRepo.On("Register", user).Return(errors.New("database error"))
	suite.mockRepo.On("VerifyFirst", user).Return(nil)
	suite.mockRepo.On("UserExists", user).Return(nil)
	err := suite.useCase.Register(user)
	suite.Error(err, "error when creating a user")
	suite.Equal(err.Error(), "database error")
	suite.mockRepo.AssertCalled(suite.T(), "Register", user)
}

func (suite *UserTestSuite) TestUserRegister_UserAlreadyExists() {
	user := &domain.User{Email: "test@example.com", Password: "password123", Role: "admin"}
	suite.mockRepo.On("Register", user).Return(nil)
	suite.mockRepo.On("VerifyFirst", user).Return(nil)
	suite.mockRepo.On("UserExists", user).Return((errors.New("user already exists")))
	err := suite.useCase.Register(user)
	suite.Error(err, "expected error when user already exists")
	suite.Equal(err.Error(), "user already exists")
	suite.mockRepo.AssertCalled(suite.T(), "UserExists", user)
}
```
### Controller testing

* A similar setup format as the usecase, using the *suite* module from *testify*
* routes were defined to simulate the request sending and receiving process. some functionalities were protected to only be done with an admin previlege. Below is the initialization/setup for the controller suite.

```go 
func (suite *ControllerTestSuite) SetupTest() {
	suite.router = gin.Default()
	suite.mockUserUseCase = new(mocks.UserUseCase)
	suite.mockTaskUseCase = new(mocks.TaskUseCase)
	suite.userController = &controllers.UserController{
		UserUseCase : suite.mockUserUseCase,
	}
	suite.taskController = &controllers.TaskController {
		TaskUseCase : suite.mockTaskUseCase,
	}

	suite.TaskGroup = []*domain.Task{
		{
			Title : "Title 1",
			Description : "this is title 1",
			DueDate : time.Now(),
			Status : "pending",
		},
		{
			Title : "Title 2",
			Description : "this is title 2",
			DueDate : time.Now(),
			Status : "pending",
		},
	}

	suite.SingleTask = domain.Task{Title : "Title 1", Description : "this is title 1",Status : "pending",}
	suite.router.POST("/register", suite.userController.Register())
	suite.router.POST("/login", suite.userController.Login())
	suite.router.PUT("/promote/:id", infrastructure.AuthMiddleWare(), suite.userController.PromoteUser())
	suite.router.GET("/tasks", infrastructure.AuthMiddleWare(), suite.taskController.GetTasks())
	suite.router.POST("/tasks", infrastructure.AuthMiddleWare(), suite.taskController.PostTask())
	suite.router.DELETE("/tasks/:id", infrastructure.AuthMiddleWare(), suite.taskController.DeleteTask())
	suite.router.PUT("/tasks/:id", infrastructure.AuthMiddleWare(), suite.taskController.UpdateTask())
	suite.router.GET("/tasks/:id", infrastructure.AuthMiddleWare(), suite.taskController.GetTask())
}
```
* For this case, as is the repository a mock for the usecase, the usecase is a mock for the controller. And only those mock functions are used during the testing process.

### Infrastructure testing

* For this case the functions are directly called inside each test and no mock functions were used.
* You can check out each test file for further analysis. Would be a bit redundant if mentioned here.

**As usual the API documentation is given below**

https://documenter.getpostman.com/view/34371403/2sA3rxrZgh#5c36f5f7-3a33-4c96-92e4-3636a6411746
