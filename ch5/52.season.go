package main

import (
	"errors"
	"fmt"
)

func season(month int) (string, error) {

	var season string

	switch {
	case month > 0 && month <= 3:
		season = "Spring"
	case month > 3 && month <= 6:
		season = "Summer"
	case month > 6 && month <= 9:
		season = "Autumn"
	case month > 9 && month <= 12:
		season = "Winter"
	default:
		return season, errors.New("invalid month")
	}

	return season, nil
}

func season2(month int) (string, error) {

	var season string

	switch month{
	case 1, 2, 3:
		season = "Spring"
	case 4, 5, 6:
		season = "Summer"
	case 7, 8, 9:
		season = "Autumn"
	case 10, 11, 12:
		season = "Winter"
	default:
		return season, errors.New("invalid month")
	}

	return season, nil
}


func main ()  {

	r, err := season(1)
	fmt.Printf("r = %v, err = %v\n",r, err)
	r, err = season(3)
	fmt.Printf("r = %v, err = %v\n",r, err)

	r, err = season(4)
	fmt.Printf("r = %v, err = %v\n",r, err)

	r, err = season(7)
	fmt.Printf("r = %v, err = %v\n",r, err)

	r, err = season(12)
	fmt.Printf("r = %v, err = %v\n",r, err)

	r, err = season(0)
	fmt.Printf("r = %v, err = %v\n",r, err)

	r, err = season(15)
	fmt.Printf("r = %v, err = %v\n",r, err)
}