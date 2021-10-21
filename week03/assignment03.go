package week03

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	min  = 1
	max  = 10
	play = 1
)

type Movie struct {
	Length  int
	Name    string
	viewers int
	plays   int
	ratings []float32
}

type Theatre struct {
	name string
}

func (m *Movie) Rate(rating float32) error {
	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}

	m.ratings = append(m.ratings, rating)

	return nil
}

func (m *Movie) Play(viewers int) {
	m.viewers += viewers
	m.plays += play
}

func (m Movie) Viewers() int {
	return m.viewers
}

func (m Movie) Plays() int {
	return m.plays
}

func (m Movie) Rating() float64 {
	if m.Plays() == 0 {
		return 0.00
	}

	var rs float32
	for _, rating := range m.ratings {
		rs += rating
	}

	r := float64(rs / float32(m.Plays()))
	r = math.Round(r*100) / 100

	return r
}

func (m Movie) String() string {
	r := m.Rating() * 10
	return fmt.Sprintf("%s (%dm) %.1f%%", m.Name, m.Length, r)
}

func (t *Theatre) Play(viewers int, movies ...*Movie) error {
	if len(movies) == 0 {
		return fmt.Errorf("no movies to play")
	}

	for _, m := range movies {
		m.Play(viewers)
	}

	return nil
}

func (t Theatre) Critique(movies []*Movie, cf CritiqueFn) error {
	if len(movies) == 0 {
		return fmt.Errorf("no movies to critique")
	}

	if cf == nil {
		return fmt.Errorf("critique function is nil")
	}

	for _, m := range movies {
		m.Play(play)

		r, err := cf(m)
		if err != nil {
			return err
		}

		err = m.Rate(r)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min

	return random
}

type CritiqueFn = func(*Movie) (float32, error)

var critiqueFn = func(movie *Movie) (float32, error) {
	if movie == nil {
		return 0.0, fmt.Errorf("no movie for rating")
	}

	rating := generateRandomNumber()

	return float32(rating), nil
}
