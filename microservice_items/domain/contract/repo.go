package contract

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Item() ItemRepo
}

// ItemRepo defines the data set for items
type ItemRepo interface {
}
