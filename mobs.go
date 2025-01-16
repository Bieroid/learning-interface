package main

import (
	"fmt"
	"math/rand"
)

type Yakut struct {
	hp int
	name string
	attack int
	exp int
}

func (enemy *Yakut) Punch() int {
	fmt.Printf("%s: Якуты так просто не сдаются, ", enemy.name)
	return enemy.GetAttack()
}

func (enemy *Yakut) IsFriendly() bool {
	i := rand.Intn(2)
	if i == 0 {
		return false
	} else {
		return true
	}
}

func (enemy *Yakut) GetAttack() int {
	minAttack := float64(enemy.attack) * 0.5
	maxAttack := float64(enemy.attack) * 1.5
	damage := minAttack + rand.Float64()*(maxAttack-minAttack)
	fmt.Printf("получай %d урона\n", int(damage))
	return int(damage)
}

func (enemy *Yakut) GetHpQuantity() int {
	return enemy.hp
}

func (enemy *Yakut) GetExp() int {
	return enemy.exp
}

func (enemy *Yakut) GetName() string {
	return enemy.name
}

func (enemy *Yakut) Greeting() {
	fmt.Printf("%s: Здарова, я добрый, на подгон\nВы рады и ваше хп увеличилось на 20\n\n", enemy.GetName())
}

func (enemy *Yakut) GetDamage(attack int, energy int, name string) {
	if energy < 20 {
		fmt.Printf("%s: Получай плюху на %d\n", name, attack)
		enemy.hp = enemy.hp - attack
	} else {
		fmt.Printf("%s: Я получаю силу от будды, я чувствую ее, получай божетсвенной ладонью на %d\n", name, attack + energy)
		enemy.hp = enemy.hp - attack - energy
	}
}

type Kalmyk struct {
	hp int
	name string
	attack int
	exp int
}

func (enemy *Kalmyk) Punch() int {
	fmt.Printf("%s: Якуты так просто не сдаются, ", enemy.name)
	return enemy.GetAttack()
}

func (enemy *Kalmyk) IsFriendly() bool {
	i := rand.Intn(2)
	if i == 0 {
		return false
	} else {
		return true
	}
}

func (enemy *Kalmyk) GetAttack() int {
	minAttack := float64(enemy.attack) * 0.5
	maxAttack := float64(enemy.attack) * 1.5
	damage := minAttack + rand.Float64()*(maxAttack-minAttack)
	fmt.Printf("получай %d урона\n", int(damage))
	return int(damage)
}

func (enemy *Kalmyk) GetHpQuantity() int {
	return enemy.hp
}

func (enemy *Kalmyk) GetExp() int {
	return enemy.exp
}

func (enemy *Kalmyk) GetName() string {
	return enemy.name
}

func (enemy *Kalmyk) Greeting() {
	fmt.Printf("Наран Петрович: Менд, брат, не узнал, это же я, Наран Петрович, а ты думал я %s, пошли в Синнабон лучше\nВы рады и ваше хп увеличилось на 20\n\n", enemy.GetName())
}

func (enemy *Kalmyk) GetDamage(attack int, energy int, name string) {
	if energy < 20 {
		fmt.Printf("%s: Получай плюху на %d\n", name, attack)
		enemy.hp = enemy.hp - attack
	} else {
		fmt.Printf("%s: Я получаю силу от будды, я чувствую ее, получай божетсвенной ладонью на %d\n", name, attack + energy)
		enemy.hp = enemy.hp - attack - energy
	}
}

type Tatarin struct {
	hp int
	name string
	attack int
	exp int
}

func (enemy *Tatarin) Punch() int {
	fmt.Printf("%s: Нет лучше подарочка, чем жена татарочка, ", enemy.name)
	return enemy.GetAttack()
}

func (enemy *Tatarin) IsFriendly() bool {
	i := rand.Intn(2)
	if i == 0 {
		return false
	} else {
		return true
	}
}

func (enemy *Tatarin) GetAttack() int {
	minAttack := float64(enemy.attack) * 0.5
	maxAttack := float64(enemy.attack) * 1.5
	damage := minAttack + rand.Float64()*(maxAttack-minAttack)
	fmt.Printf("получай %d урона\n", int(damage))
	return int(damage)
}

func (enemy *Tatarin) GetHpQuantity() int {
	return enemy.hp
}

func (enemy *Tatarin) GetExp() int {
	return enemy.exp
}

func (enemy *Tatarin) GetName() string {
	return enemy.name
}

func (enemy *Tatarin) Greeting() {
	fmt.Printf("%s: Калайсын, вместо драки давай лучше калик бахнем брат\nВы рады и ваше хп увеличилось на 20\n\n", enemy.GetName())
}

func (enemy *Tatarin) GetDamage(attack int, energy int, name string) {
	if energy < 20 {
		fmt.Printf("%s: Получай плюху на %d\n", name, attack)
		enemy.hp = enemy.hp - attack
	} else {
		fmt.Printf("%s: Я получаю силу от будды, я чувствую ее, получай божетсвенной ладонью на %d\n", name, attack + energy)
		enemy.hp = enemy.hp - attack - energy
	}
}
