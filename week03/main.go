package week03

import "fmt"

func main() {
	t1 := Theatre{
		name: "Galaxy",
	}

	t2 := Theatre{
		name: "InOrbit",
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
		fmt.Errorf(" Error : %s", err.Error())
	}

	fmt.Println(m1, m1.plays)
	fmt.Println(m2, m2.plays)

	err = t1.Play(50, &m1)
	if err != nil {
		return
	}

	err = t2.Play(100, &m2)
	if err != nil {
		return
	}

	fmt.Println(m1, m1.Plays(), m1.Viewers(), m1.Rating())

	fmt.Println(m2, m2.Plays(), m2.Viewers(), m2.Rating())
}