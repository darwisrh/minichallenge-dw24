package main

import "fmt"

type Profil struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(name string) Profil {
	profile := Profil{}
	if name == "Goku" {
		profile = Profil{
			Name:   "Goku",
			Health: 400,
			Power:  300,
			Exp:    100,
		}
	} else if name == "Sasuke" {
		profile = Profil{
			Name:   "Sasuke",
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	} else if name == "Naruto" {
		profile = Profil{
			Name:   "Naruto",
			Health: 150,
			Power:  200,
			Exp:    50,
		}
	} else {
		fmt.Println("Profile not found")
	}
	return profile
}

func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name:", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp:", profile.Exp)
	fmt.Println("======================")
	profilee := MakeProfile("Naruto")
	fmt.Println("Name:", profilee.Name)
	fmt.Println("Health :", profilee.Health)
	fmt.Println("Power :", profilee.Power)
	fmt.Println("Exp:", profilee.Exp)
	fmt.Println("======================")
	profileee := MakeProfile("Sasuke")
	fmt.Println("Name:", profileee.Name)
	fmt.Println("Health :", profileee.Health)
	fmt.Println("Power :", profileee.Power)
	fmt.Println("Exp:", profileee.Exp)
}

func (profile *Profil) PowerUp(multiplier int) {
	profile.Power = profile.Power + (profile.Power * multiplier)
	profile.Health = profile.Health + (profile.Health * multiplier)
	profile.Exp = profile.Exp + (profile.Exp * multiplier)
}
