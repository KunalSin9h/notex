package data

/*
	Repository Method to make our handlers testable by mocking database.

	MongoDBRepository ----------
	                             \__ Repository
						         /
	MongoDBTestRepository-------

	MongoDBTestRepository will also define these methods but they will be dummy and return
	data without query the database, this is used for integration testing
*/

type Repository interface{}
