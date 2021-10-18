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

	if err == nil {
		t.Fatal("expected an error and got none")
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
		t.Error(err)
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
		t.Errorf("expected %d, got %d", m1.Viewers(), viewers)
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
		t.Errorf("expected %d, got %d", m1.Plays(), plays)
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
		t.Errorf("expected %d, got %d", m1.Viewers(), viewers)
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
		t.Errorf("expected %d, got %d", m1.Plays(), plays)
	}
}

func Test_Rating_1(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	err := m1.Rate(8.5)

	if err == nil {
		t.Fatal("expected an error and got none")
	}

	exp := 0.0
	act := m1.Rating()

	if act != exp {
		t.Errorf("expected, %.1f, got %.1f", exp, act)
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
		t.Fatal(err)
	}

	exp := 8.5
	act := m1.Rating()

	if act != exp {
		t.Errorf("expected, %.1f, got %.1f", exp, act)
	}
}

func Test_Rating_3(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)
	err := m1.Rate(8.3)
	if err != nil {
		t.Fatal(err)
	}

	m1.Play(200)
	err = m1.Rate(9.7)
	if err != nil {
		t.Fatal(err)
	}

	exp := 9.0
	act := m1.Rating()

	if act != exp {
		t.Errorf("expected, %.1f, got %.1f", exp, act)
	}
}

func Test_Rating_4(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)
	err := m1.Rate(8.3)
	if err != nil {
		t.Fatal(err)
	}

	m1.Play(200)
	err = m1.Rate(9.6)
	if err != nil {
		t.Fatal(err)
	}

	exp := 8.95
	act := m1.Rating()
	if act != exp {
		t.Errorf("expected, %.1f, got %.1f", exp, act)
	}
}

func Test_String_1(t *testing.T) {
	t.Parallel()

	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}

	m1.Play(100)

	err := m1.Rate(8.250)
	if err != nil {
		t.Fatal(err)
	}

	exp := "Avengers (185m) 82.5%"
	act := m1.String()
	if act != exp {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func Test_Theatre_Play_1(t *testing.T) {
	t.Parallel()

	t1 := Theatre{
		name: "Galaxy",
	}

	err := t1.Play(100)
	if err == nil {
		t.Fatal("expected an error and got none")
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
		t.Error(err)
	}
}

func Test_generateRandomNumber(t *testing.T) {
	t.Parallel()

	n := generateRandomNumber()
	if n < min || n > max {
		t.Errorf("n : %d is not in expected range{%d,%d}", n, min, max)
	}
}

func Test_critiqueFn_1(t *testing.T) {
	_, err := critiqueFn(nil)

	if err == nil {
		t.Fatal("expected an error and got none")
	}
}

func Test_critiqueFn_2(t *testing.T) {
	m1 := Movie{
		Name:   "Avengers",
		Length: 185,
	}
	n, err := critiqueFn(&m1)
	if err != nil {
		t.Fatal(err)
	}
	if n < min || n > max {
		t.Errorf("n : %f is not in expected range{%d,%d}", n, min, max)
	}
}

func Test_Critique_1(t *testing.T) {
	t1 := Theatre{
		name: "Galaxy",
	}

	var movies []*Movie
	err := t1.Critique(movies, critiqueFn)

	if err == nil {
		t.Fatal("expected an error and got none")
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

	movies := []*Movie{&m1, &m2}

	err := t1.Critique(movies, critiqueFn)

	if err != nil {
		t.Fatal(err)
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

	movies := []*Movie{&m1, &m2}

	critiqueFn = nil
	err := t1.Critique(movies, critiqueFn)

	if err == nil {
		t.Fatal("expected an error and got none")
	}
}
