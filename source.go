package karigo

// Source ...
type Source interface {
	Get(Query) string
}

// // ResProvider ...
// type ResProvider interface {
// 	Get(Key) string
// }

// // ColProvider ...
// type ColProvider interface {
// }
