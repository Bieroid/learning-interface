package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	Hp int
	AttackDmg int
	Name string
}

type Mob interface {
	Punch() int
	IsFriendly() bool
	GetHpQuantity() int
	GetName() string
	Greeting()
	GetDamage(int)
}

type Yakut struct {
	hp int
	name string
}

func (ayal *Yakut) Punch() int {
	fmt.Printf("Якуты так просто не сдаются")
	return 100
}

func (ayal *Yakut) IsFriendly() bool {
	//fmt.Printf("Да я добрый якут")
	return true
}

func (ayal *Yakut) GetHpQuantity() int {
	//fmt.Printf("Ладно, скажу скока осталось")
	return ayal.hp
}

func (ayal *Yakut) GetName() string {
	//fmt.Printf("Ладно, скажу скока осталось")
	return ayal.name
}

func (ayal *Yakut) Greeting() {
	fmt.Printf("%s: Сразу говорю - бью больно\n", ayal.GetName())
}

func (ayal *Yakut) GetDamage(attack int) {
	ayal.hp = ayal.hp - attack
}

func (p player) Interact(mob Mob) {
	if mob.GetHpQuantity() > 0 {
		if mob.IsFriendly() {
			mob.Greeting()
			fmt.Printf("%s: Ну ладно, живи, %s\n", p.Name, mob.GetName())
		} else {
			p.Punch(mob)
		}
	}
}

func (p player) Punch(mob Mob) {
	fmt.Printf("%s: Получай плюху\n", p.Name)
	mob.GetDamage(p.AttackDmg)
	if mob.GetHpQuantity() > 0 {
		p.Hp = p.Hp - mob.Punch()
		if p.Hp < 0 {
			fmt.Printf("%s: Петя ну чо там когда го?\n", p.Name)
		}
	}
}

func main() {
	german := player{Hp: 10000, AttackDmg: 10, Name: "German"}
	mob := CreateRandomMob()

	for german.Hp > 0 || mob.GetHpQuantity() > 0 {
		german.Interact(mob)
	}
}

// Функция для создания случайного моба
func CreateRandomMob() Mob {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	mobType := rand.Intn(4)          // Случайное число от 0 до 3
   
	switch mobType {
	case 0:
	 return &Orc{hp: 100, name: "Тралл"}
	case 1:
	 return &Goblin{hp: 30, name: "Гоблин"}
	case 2:
	 return &Troll{hp: 150, name: "Тролль"}
	case 3:
	 return &Dragon{hp: 500, name: "Дракош"}
	default:
	 return &Yakut{hp: 5000, name: "Айал"}
	}
}

// Придумать что нибудь красивое (рандомное значение атаки(min = max), разнообразные реплики (смешные), добавить ульту (при накоплении энергии(добавить энергию), допилить генератор мобоу))
