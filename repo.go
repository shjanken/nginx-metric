package metric

// Repo is backend
type Repo interface {
	Insert(item *Item) error
	BatchInsert(item []Item) error
}
