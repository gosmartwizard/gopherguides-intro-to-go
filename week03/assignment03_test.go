package week03

import (
	"testing"
)

func Test_Rate_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	var rating float32 = 82.0
	err := m1.Rate(rating)
	e := "can't review a movie without watching it first"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}

func Test_Rate_2(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(1)
	var rating float32 = 82.0

	err := m1.Rate(rating)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if rating != m1.ratings {
		t.Fatalf("Actual rating : %f and Expected rating %f", rating, m1.ratings)
	}
}

func Test_Movie_Play_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 50
	m1.Play(viewers)

	if viewers != m1.Viewers() {
		t.Fatalf("Actual viewers : %d and Expected viewers %d", viewers, m1.Viewers())
	}
}

func Test_Movie_Play_2(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 100
	m1.Play(viewers)

	if viewers != m1.Plays() {
		t.Fatalf("Actual viewers : %d and Expected viewers %d", viewers, m1.Plays())
	}
}

func Test_Viewers_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 150
	m1.Play(viewers)

	if viewers != m1.Viewers() {
		t.Fatalf("Actual viewers : %d and Expected viewers %d", viewers, m1.Viewers())
	}
}

func Test_Plays_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 250
	m1.Play(viewers)

	if viewers != m1.Plays() {
		t.Fatalf("Actual viewers : %d and Expected viewers %d", viewers, m1.Plays())
	}
}

func Test_Rating_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)

	err := m1.Rate(82.0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	act := float64(m1.ratings / float32(m1.Plays()))
	exp := m1.Rating()
	if act != exp {
		t.Fatalf("Actual rating : %f and Expected rating %f", act, exp)
	}
}

func Test_String_1(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)

	err := m1.Rate(82.0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	act := "Avengers (185m) 82.0%"
	exp := m1.String()
	if act != exp {
		t.Fatalf("Actual rating : %s and Expected rating %s", act, exp)
	}
}

func Test_Theatre_Play_1(t *testing.T) {

	t1 := Theatre{
		name: "Galaxy",
	}

	err := t1.Play(100)
	e := "no movies to play"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}

func Test_Theatre_Play_2(t *testing.T) {

	t1 := Theatre{
		name: "Galaxy",
	}

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	m2 := Movie{
		Name:   "Batman",
		Length: 165,
	}

	err := t1.Play(100, &m1, &m2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test_Critique_1(t *testing.T) {

	t1 := Theatre{
		name: "Galaxy",
	}

	var movies []*Movie
	err := t1.Critique(movies, critiqueFn)
	e := "no movies to Critique"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}

/*func Test_Critique_2(t *testing.T) {

	t1 := Theatre{
		name : "Galaxy",
	}

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	m2 := Movie{
		Name:   "Batman",
		Length: 165,
	}

	var movies []*Movie
	movies = append(movies, &m1, &m2)

	critiqueFn = nil
	err := t1.Critique(movies, critiqueFn)
	e := "CritiqueFn is nil"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}*/

func Test_Critique_3(t *testing.T) {

	t1 := Theatre{
		name: "Galaxy",
	}

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m2 := Movie{
		Name:   "Batman",
		Length: 165,
	}

	var movies []*Movie
	movies = append(movies, &m1, &m2)

	err := t1.Critique(movies, critiqueFn)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Test_generateRandomNumber(t *testing.T) {
	n := generateRandomNumber()
	if n < 1 || n > 100 {
		t.Errorf("n : %d is not in expected range{1,100}", n)
	}
}

func Test_critiqueFn_1(t *testing.T) {
	_, err := critiqueFn(nil)
	e := "no movie for rating"
	if err.Error() != e {
		t.Errorf(err.Error())
	}
}

func Test_critiqueFn_2(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	r, err := critiqueFn(&m1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if r < 1 || r > 100 {
		t.Errorf(err.Error())
	}
}
