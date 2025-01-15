package main

import "fmt"

type Orc struct {
	hp int
	name string
}

func (thrall *Orc) Punch() int {
	fmt.Printf("%s: Моя не умереть", thrall.name)
	return 10
}

func (thrall *Orc) IsFriendly() bool {
	return false
}

func (thrall *Orc) GetHpQuantity() int {
	return thrall.hp
}

func (thrall *Orc) GetName() string {
	return thrall.name
}

func (thrall *Orc) Greeting() {
	fmt.Printf("%s: Ыбымбе\n", thrall.GetName())
}

func (thrall *Orc) GetDamage(attack int) {
	thrall.hp = thrall.hp - attack
}

type Goblin struct {
	hp   int
	name string
}
   
func (goblin *Goblin) Punch() int {
	fmt.Printf("%s: Тыкать-тыкать!\n", goblin.name)
	return 5
}
   
func (goblin *Goblin) IsFriendly() bool {
	return false
}
   
func (goblin *Goblin) GetHpQuantity() int {
	return goblin.hp
}
   
func (goblin *Goblin) GetName() string {
	return goblin.name
}
   
func (goblin *Goblin) Greeting() {
	fmt.Printf("%s: Гыгыгы!\n", goblin.GetName())
}
   
func (goblin *Goblin) GetDamage(attack int) {
	goblin.hp = goblin.hp - attack
}
   
type Troll struct {
	hp   int
	name string
}
   
func (troll *Troll) Punch() int {
	fmt.Printf("%s: УАААААА!\n", troll.name)
	return 15
}
   
func (troll *Troll) IsFriendly() bool {
	return false
}
   
func (troll *Troll) GetHpQuantity() int {
	return troll.hp
}
   
func (troll *Troll) GetName() string {
	return troll.name
}
   
func (troll *Troll) Greeting() {
	fmt.Printf("%s: Гррррр!\n", troll.GetName())
}
   
func (troll *Troll) GetDamage(attack int) {
	troll.hp = troll.hp - attack
}

type Dragon struct {
	hp   int
	name string
}

func (dragon *Dragon) Punch() int {
	fmt.Printf("%s: ДЫШУ ОГНЕМ! ШШШШШШ!\n", dragon.name)
	return 100000
}

func (dragon *Dragon) IsFriendly() bool {
	return false
}

func (dragon *Dragon) GetHpQuantity() int {
	return dragon.hp
}

func (dragon *Dragon) GetName() string {
	return dragon.name
}

func (dragon *Dragon) Greeting() {
	fmt.Printf("%s: Я дракон, повелитель огня!\n", dragon.GetName())
}

func (dragon *Dragon) GetDamage(attack int) {
	dragon.hp = dragon.hp - attack
}