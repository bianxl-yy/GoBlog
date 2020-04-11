package model

type navItem struct {
	Order int
	Text  string
	Title string
	Link  string
}

var (
	settings   map[string]string
	navigators []*navItem
)
