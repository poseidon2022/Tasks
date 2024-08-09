# A Simple Task Manager API With MongoDB (Authentication added)

* This is a simple task manager api to post, update, visualize and delete tasks. 
* Update, delete and get tasks by ID
* Errors are handled in a very intuitive way.
* Two kinds of previleges : a user and an admin previlege
* A registered and logged-in user can visualize the tasks and also fetch a task by ID.
* an admin is allowed to post tasks, edit tasks and update tasks.
* an admin can escalate other users to an admin previlege.
* if there are no registered users, the first one to sign up will be given an admin previlege

## Folder Structure

* **Controllers:** folder to contain all the functions that invoke the task_services function up on an API call.
* **Data:** contains and implements task_Service functions to handle anything related to the database storage, such as manipulating it or reading from it. It also includes the initialization for the mongoDB client and all collections are accessed only by functions that are inside the task_services.go file.

**The code for the database client establishment**
```go
func GetClient() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error while connecting to the Database")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Database connecction fatally failed")
	}

	db := client.Database("task_management")
	fmt.Println("Database connection setup")
	return db
}

var collection = GetClient().Collection("tasks")
```

* The connection string inside the ApplyURI can be modified to another connection string to access and use another cluster.

* **Docs:** Description about the application and api documentation.
* **Models:** Struct description for how the stored fields should look like, we have two structs: one for the tasks and the other for users.
* **Router:** All the routes to handle api calls and invocations. There are three end-points added to the previous task: login, register and promote
* **main.go:** Entry point of our application.

### Check out the API documentation using this URL

https://documenter.getpostman.com/view/34371403/2sA3rxrZgh#5c36f5f7-3a33-4c96-92e4-3636a6411746