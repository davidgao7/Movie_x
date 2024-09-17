package model

type Movie struct {
	Title     string
	Year      string
	Length    uint32
	RateLevel string
	Review    float32
	Genre     *string
	Stars     *string
}
