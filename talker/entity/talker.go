package talker

//Talk represents the talker date speak and yours rate
type Talk struct {
	WatchedAt string `json:"watchedAt"`
	Rate      int    `json:"rate"`
}

//Talker represents the talker
type Talker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
	Talk Talk   `json:"talk"`
}
