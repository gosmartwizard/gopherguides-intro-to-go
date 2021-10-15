package week03

import (
	"testing"
)

func Test_Rate_1(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	var rating float32 = 8.2
	err := m1.Rate(rating)
	e := "can't review a movie without watching it first"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}

func Test_Rate_2(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(play)
	var rating float32 = 9.5

	err := m1.Rate(rating)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if rating != m1.ratings {
		t.Fatalf("Actual rating : %f and Expected rating %f", rating, m1.ratings)
	}
}

func Test_Movie_Play_1(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 100
	plays := play
	m1.Play(viewers)

	if plays != m1.Plays() {
		t.Fatalf("Actual plays : %d and Expected plays %d", plays, m1.Plays())
	}
}

func Test_Viewers_1(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	viewers := 250
	plays := play
	m1.Play(viewers)

	if plays != m1.Plays() {
		t.Fatalf("Actual plays : %d and Expected plays %d", plays, m1.Plays())
	}
}

func Test_Rating_1(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)

	err := m1.Rate(8.5)
	if err != nil {
		t.Fatalf(err.Error())
	}

	act := 8.5
	exp := m1.Rating()

	if act != exp {
		t.Fatalf("Actual rating : %.1f and Expected rating %.1f", act, exp)
	}
}

func Test_Rating_2(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)
	err := m1.Rate(8.5)
	if err != nil {
		t.Fatalf(err.Error())
	}

	m1.Play(200)
	err = m1.Rate(9.5)
	if err != nil {
		t.Fatalf(err.Error())
	}

	act := 9.0
	exp := m1.Rating()
	if act != exp {
		t.Fatalf("Actual rating : %.1f and Expected rating %.1f", act, exp)
	}
}

func Test_String_1(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)

	err := m1.Rate(8.2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	act := "Avengers (185m) 8.2%"
	exp := m1.String()
	if act != exp {
		t.Fatalf("Actual rating : %s and Expected rating %s", act, exp)
	}
}

func Test_Theatre_Play_1(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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

func Test_generateRandomNumber(t *testing.T) {
	t.Parallel()

	n := generateRandomNumber()
	if n < min || n > max {
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
	if r < min || r > max {
		t.Errorf(err.Error())
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

func Test_Critique_2(t *testing.T) {
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

	critiqueFn = nil
	err := t1.Critique(movies, critiqueFn)
	e := "CritiqueFn is nil"
	if err.Error() != e {
		t.Fatalf(err.Error())
	}
}
