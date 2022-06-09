package entitie

//Talk represents the talker date speak and yours rate
type Talk struct {
	WatchedAt string
	Rate      int
}

//Talker represents the talker
type Talker struct {
	Name string
	Age  int
	ID   int
	Talk Talk
}
