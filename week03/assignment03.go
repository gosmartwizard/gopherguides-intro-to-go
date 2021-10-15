package week03

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min int = 1
	max int = 100
	play int = 1
)

type Movie struct {
	Length  int
	Name    string
	viewers int
	plays   int
	ratings float32
}

type Theatre struct {
	name string
}

func (m *Movie) Rate(rating float32) error {
	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}
	m.ratings += rating
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
	return float64(m.ratings / float32(m.Plays()))
}

func (m Movie) String() string {
	s := fmt.Sprintf("%s (%dm) %.1f", m.Name, m.Length, m.ratings) + "%"
	return s
}

func (t *Theatre) Play(viewers int, movies ...*Movie) error {

	if 0 == len(movies) {
		return fmt.Errorf("no movies to play")
	}

	for _, m := range movies {
		m.Play(viewers)
	}

	return nil
}

func (t Theatre) Critique(movies []*Movie, cf CritiqueFn) error {

	if 0 == len(movies) {
		return fmt.Errorf("no movies to Critique")
	}

	if nil == cf {
		return fmt.Errorf("CritiqueFn is nil")
	}

	for _, m := range movies {
		m.Play(1)

		r, err := cf(m)
		if err != nil {
			return fmt.Errorf(" Error : %s", err.Error())
		}

		err = m.Rate(r)
		if err != nil {
			return fmt.Errorf(" Error : %s", err.Error())
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
