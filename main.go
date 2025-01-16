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
	Energy int
	Exp int
	Lvl int
}

type Mob interface {
	Punch() int
	IsFriendly() bool
	GetHpQuantity() int
	GetName() string
	Greeting()
	GetDamage(int, int, string)
	GetAttack() int
	GetExp() int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	german := player{Hp: 400, AttackDmg: 10, Name: "German", Energy: 0}
	tmblrWin := false

	startText()
	for german.Hp > 0 && !tmblrWin {
		locNumb := locatOutput()
		mob := CreateRandomMob(locNumb)
		if mob.IsFriendly() {
			mob.Greeting()
			german.Hp = german.Hp + 20
		} else {
			for german.Hp > 0 && mob.GetHpQuantity() > 0 {
				german.Battle(mob)
			}
		}
		if german.Exp >= 10 {
			lvlUp(&german)
			if german.Lvl == 3 {
				tmblrWin = true
			}
		}
	}
	fmt.Println("Поздравляю, вы достигли 3 уровня и прошли бета-версию игры, приходите позднее!!!")
}

func (p *player) Battle(mob Mob) {
	if mob.GetHpQuantity() > 0 {
		p.Punch(mob)
		infBattleOut(p, mob)
	}
	time.Sleep(500 * time.Millisecond)
}

func (p *player) Punch(mob Mob) {
	p.Energy = p.Energy + 1
	if p.Energy > 20 {
		p.Energy = 0
	}
	minAttack := float64(p.AttackDmg) * 0.5
	maxAttack := float64(p.AttackDmg) * 1.5
	damage := minAttack + rand.Float64()*(maxAttack-minAttack)
	mob.GetDamage(int(damage), p.Energy, p.Name)
	if mob.GetHpQuantity() > 0 {
		p.Hp = p.Hp - mob.Punch()
	}
}

func infBattleOut(p *player, mob Mob) {
	if (mob.GetHpQuantity()) > 0 {
		fmt.Printf("[Враг] %s: Здоровье - %d\n", mob.GetName(), mob.GetHpQuantity())
	} else {
		p.Exp = p.Exp + mob.GetExp()
		fmt.Printf("\nВы одержали славную победу над %s, получено %d опыта. Текущий опыт - %d, Текущий уровень - %d\n", mob.GetName(), mob.GetExp(), p.Exp, p.Lvl)
	}
	if p.Hp > 0 {
		fmt.Printf("[Игрок] %s: Здоровье - %d, Средняя атака - %d, Энергия - %d\n\n", p.Name, p.Hp, p.AttackDmg, p.Energy)
	} else {
		fmt.Printf("\n%s: Вы погибли бесславной смертью, враг плюет на ваш труп и уходит\n", p.Name)
	}
}

func locatOutput() int {
	fmt.Printf("1. Якутия\n2. Калмыкия\n3. Казань\n\n")
	var locNumb int

	for {
		fmt.Println("Выберите локацию:")
		_, err := fmt.Scan(&locNumb)
		if err != nil  || locNumb > 3 || locNumb <= 0 {
			fmt.Println("Ошибка: введите корректные данные")
			var discard string
			fmt.Scanln(&discard)
		} else {
			break
		}
	}
	return locNumb
}

func lvlUp(p *player) {
	p.Lvl++
	p.Exp = 0
	fmt.Printf("Ваш уровень повысился до %d, выберите улучшение:\n1. Увеличить атаку\n2. Увеличить здоровье\n\n", p.Lvl)
	var locNumb int
	for {
		_, err := fmt.Scan(&locNumb)
		if err != nil  || locNumb > 1 || locNumb <= 0 {
			fmt.Println("Ошибка: введите корректные данные")
			var discard string
			fmt.Scanln(&discard)
		} else {
			break
		}
	}
	if locNumb == 1 {
		p.AttackDmg += 10
		fmt.Printf("Ваша атака увеличена и равна %d\n\n", p.AttackDmg)
	} else {
		p.Hp += p.Hp/2
		fmt.Printf("Ваша здоровье увеличено и равно %d\n\n", p.Hp)
	}
}

func startText() {
	fmt.Printf("Добро пожаловать в бета-версию игры, ваша цель сражаться с врагами и увеличивать свой уровень. Для прохождения игры достигните 3 уровня, удачи!\n\n\n")
}

func CreateRandomMob(locNumb int) Mob {
	mobType := rand.Intn(2)
	
	switch locNumb {
		case 1: 
			createYakut(mobType)
		case 2:
			createKalmyk(mobType)
		case 3:
			createTatarin(mobType)
	}
	return nil
}

func createYakut(mobType int) Mob {
    switch mobType {
    case 0:
        fmt.Println("Вы приезжаете в город и перед вами появляется Айал:")
        return &Yakut{hp: 100, name: "Айал", attack: 10, exp: 5}
    default:
        fmt.Println("Вы приезжаете в город и перед вами появляется Айала батя:")
        return &Yakut{hp: 200, name: "Айала батя", attack: 10, exp: 10}
    }
}

func createKalmyk(mobType int) Mob {
    switch mobType {
	case 0:
		fmt.Println("Вы приезжаете в город и перед вами появляется Петя:")
		return &Kalmyk{hp: 100, name: "Петя", attack: 10, exp: 5}
	default:
		fmt.Println("Вы приезжаете в город и перед вами появляется Скутер:")
		return &Kalmyk{hp: 200, name: "Скутер", attack: 20, exp: 10}
}
}

func createTatarin(mobType int) Mob {
    switch mobType {
	case 0:
		fmt.Println("Вы приезжаете в город и перед вами появляется Закир:")
		return &Tatarin{hp: 100, name: "Закир", attack: 10, exp: 5}
	default:
		fmt.Println("Вы приезжаете в город и перед вами появляются Универсамовские:")
		return &Tatarin{hp: 200, name: "Универсамовские", attack: 20, exp: 10}
}
}
