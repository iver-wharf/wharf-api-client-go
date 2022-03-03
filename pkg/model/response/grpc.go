package response

// CreatedLogsSummary contains a summary of the logs that was created after
// closing a CreateBuildLogStream.
type CreatedLogsSummary struct {
	LogsInserted uint
}
