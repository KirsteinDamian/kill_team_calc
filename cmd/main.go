package main

import (
	"fmt"
	"ktcalc/calculator"
	"sort"
	"time"
)

func main() {
	t0 := time.Now()
	att := calculator.Attacker{
		DiceCount: 6, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 7},
		AbilitiesAttack: calculator.AbilitiesAttack{MWx: 0, APx: 1, Px: 2},
	}
	def := calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12}
	damageMap := calculator.CalculateShootingAttacks(att, def)

	damageValues := make([]uint8, 0, len(damageMap))
	for i := range damageMap {
		damageValues = append(damageValues, i)
	}
	sort.Slice(damageValues, func(i, j int) bool { return damageValues[i] < damageValues[j] })

	damageProbability := make(map[uint8]float64)
	for damage, count := range damageMap {
		damageProbability[damage] = float64(count) // / float64(totalScenarios)
	}
	var avgDamage float64 = 0
	var killChance float64 = 0
	for dmg, prob := range damageProbability {
		avgDamage += float64(dmg) * prob
		if dmg >= def.Hp {
			killChance += prob
		}
	}
	for _, dmg := range damageValues {
		fmt.Printf("%v Damage: %.2f%%\n", dmg, damageProbability[dmg]*100)
	}
	fmt.Printf("Avarage damage: %.2f\n", avgDamage)
	fmt.Printf("Kill chance: %.2f%%\n", killChance*100)
	fmt.Printf("Time: %v\n", time.Since(t0))
}
