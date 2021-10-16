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
	ratings float32
}

type Theatre struct {
	name string
}

func (m *Movie) Rate(rating float32) error {
	if 0 == m.plays {
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
	if 0 == m.Plays() {
		return 0.00
	}

	r := float64(m.ratings / float32(m.Plays()))
	r = math.Round(r*100) / 100

	return r
}

func (m Movie) String() string {
	r := m.Rating() * 10
	return fmt.Sprintf("%s (%dm) %.1f", m.Name, m.Length, r) + "%"
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
		m.Play(play)

		r, err := cf(m)
		if nil != err {
			return fmt.Errorf(" Error : %s", err.Error())
		}

		err = m.Rate(r)
		if nil != err {
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
	if nil == movie {
		return 0.0, fmt.Errorf("no movie for rating")
	}

	rating := generateRandomNumber()

	return float32(rating), nil
}
