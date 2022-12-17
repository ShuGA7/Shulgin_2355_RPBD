package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Par struct {
	len int
	hp  int
	rep int
	fat int
}

var stats = Par{10, 100, 20, 30}
var newlen, newhp, newrep, newfat, vv int

func (p *Par) dighole(vv int) (int, int, int, int) {
	newlen := p.len
	newhp := p.hp
	newrep := p.rep
	newfat := p.fat
	fmt.Println("Копать быстро 1) Копать лениво 2).")
	fmt.Print("Ввод: ")
	fmt.Scanln(&vv)
	switch vv {
	case 1:
		newlen += 5
		newhp -= 30
	case 2:
		newlen += 2
		newhp -= 10
	}
	return newlen, newhp, newrep, newfat

}

func (p *Par) eatgrass(vv int) (int, int, int, int) {
	newlen := p.len
	newhp := p.hp
	newrep := p.rep
	newfat := p.fat
	fmt.Println("Какую траву есть? Жухлую 1) или зелёную (нужна репутация) 2)?")
	fmt.Print("Ваш ввод: ")
	fmt.Scanln(&vv)
	switch vv {
	case 1:
		fmt.Println("Здоровье + 10, вес + 15")
		newhp += 10
		newfat += 15
	case 2:
		if newrep < 30 {
			fmt.Println("Здоровье - 30")
			newhp -= 30
		} else {
			fmt.Println("Здоровье + 30, вес - 30")
			newfat += 30
			newhp -= 30
		}
	}
	return newlen, newhp, newrep, newfat
}

func (p *Par) fight(vv int) (int, int, int, int) {
	newlen := p.len
	newhp := p.hp
	newrep := p.rep
	newfat := p.fat
	var enemystrength, fullstrength int
	fmt.Println("С кем драться? со слабым 1), со средним 2) с сильным 3).")
	fmt.Print("Ввод: ")
	fmt.Scanln(&vv)
	switch vv {
	case 1:
		enemystrength = 30
	case 2:
		enemystrength = 50
	case 3:
		enemystrength = 70
	}
	fullstrength = enemystrength + newfat
	var chance float64 = float64(newfat) / float64(fullstrength)
	var random float64 = rand.Float64()
	rand.Seed(time.Now().UnixNano())
	if random <= chance {
		fmt.Println("Вы победили противника, + уважение")
		var winrep = enemystrength - newfat
		if winrep <= 0 {
			winrep = 10
		}
		newrep += winrep
		fmt.Println("Уважение:", newrep)
		fmt.Println("Got:", random)
		fmt.Println("Need<:", chance)
	} else {
		fmt.Println("Вы проиграли в бою, - здоровье")
		var losehp = enemystrength - newfat
		if losehp <= 0 {
			losehp += 10
		}
		newhp -= losehp
		fmt.Println("Здоровье:", newhp)
		fmt.Println("Got:", random)
		fmt.Println("Need<:", chance)
	}

	return newlen, newhp, newrep, newfat
}

func (p *Par) dead() (int, int, int, int) {
	newlen := p.len
	newhp := p.hp
	newrep := p.rep
	newfat := p.fat
	if (newlen <= 0) || (newhp <= 0) || (newrep <= 0) || (newfat <= 0) {
		fmt.Println("Вы проиграли")
		os.Exit(1)
	}
	return newlen, newhp, newrep, newfat
}

func (p *Par) night() (int, int, int, int) {
	newlen := p.len
	newhp := p.hp
	newrep := p.rep
	newfat := p.fat
	newlen -= 2
	newhp += 20
	newrep -= 2
	newfat -= 5
	return newlen, newhp, newrep, newfat

}

func (st *Par) win() int {
	newrep := st.rep
	if newrep >= 100 {
		fmt.Println("Вы победили, достигнув уважения в 100 единиц")
		os.Exit(1)
	}
	return newrep
}

func main() {
	for {
		fmt.Println("Что делать: копать нору 1), спать 2), поесть траву 3), драться 4).")
		fmt.Print("Ввод: ")
		fmt.Scanln(&vv)
		if vv <= 0 || vv > 4 {
			fmt.Println("Неверный ввод")
			continue
		}

		switch vv {
		case 1:
			newlen, newhp, newrep, newfat = stats.dighole(vv)
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			fmt.Println("Сегодня вы копали, характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			newlen, newhp, newrep, newfat = stats.night()
			fmt.Println("День закончился, вы пошли спать. Характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			main()
		case 2:
			newlen, newhp, newrep, newfat = stats.night()
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			fmt.Println("Вы спали весь день, характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			newlen, newhp, newrep, newfat = stats.night()
			fmt.Println("День закончился, вы пошли спать. Характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			main()
		case 3:
			newlen, newhp, newrep, newfat = stats.eatgrass(vv)
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			fmt.Println("Сегодня вы поели, характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			newlen, newhp, newrep, newfat = stats.night()
			fmt.Println("День закончился, вы пошли спать. Характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			main()
		case 4:
			newlen, newhp, newrep, newfat = stats.fight(vv)
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			fmt.Println("Сегодня вы подрались, характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			stats.win()
			newlen, newhp, newrep, newfat = stats.night()
			fmt.Println("День закончился, вы пошли спать. Характеристики:", "Длина норы:", newlen, "Здоровье:", newhp, "Уважение:", newrep, "Вес:", newfat)
			stats.dead()
			stats.len = newlen
			stats.hp = newhp
			stats.rep = newrep
			stats.fat = newfat
			main()
		}

	}
}
