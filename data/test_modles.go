package data

type MongoDBTestRepository struct{}

func init() {
	// Check if MongoDBTestRepository fully implements the Repository interface
	var _ Repository = (*MongoDBTestRepository)(nil)
}
