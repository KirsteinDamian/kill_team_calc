package calculator

func CalculateShootingAttacks(attacker Attacker, defender Defender) (map[uint8]uint64, int) {
	var totalScenarios int
	attackerRolls := cartesianProduct(attacker.DiceCount)
	standardDefenderRolls := cartesianProduct(defender.DiceCount)
	APxDefenderRolls := GetAPxDefenderRolls(attacker, defender, standardDefenderRolls)
	PxDefenderRolls := GetPxDefenderRolls(attacker, defender, standardDefenderRolls)

	damageMap := make(map[uint8]uint64)

	for _, attRoll := range attackerRolls {
		var isAnyDmg bool = false
		var isAnyCrit bool = false
		for _, dice := range attRoll {
			if dice >= attacker.Balistic {
				isAnyDmg = true
				if dice == 6 {
					isAnyCrit = true
					break
				}
			}
		}
		scenarios := GetDefenderRolls(attacker, isAnyCrit, APxDefenderRolls, PxDefenderRolls)
		scenariosCount := len(scenarios)
		if isAnyDmg {
			for _, defRoll := range scenarios {
				var damage uint8 = CalcShootAttack(attacker, defender, attRoll, defRoll)
				var counter, ok = damageMap[damage]
				if ok {
					damageMap[damage] = counter + 1
				} else {
					damageMap[damage] = 1
				}
			}
		} else {
			var counter, ok = damageMap[0]
			if ok {
				damageMap[0] = counter + uint64(scenariosCount)
			} else {
				damageMap[0] = uint64(scenariosCount)
			}
		}
		totalScenarios += scenariosCount
	}
	return damageMap, totalScenarios
}

func CalcShootAttack(attacker Attacker, defdefender Defender, attackRoll []uint8, defenceRoll []uint8) uint8 {
	var attackSuccess uint8 = 0
	var attackCrit uint8 = 0
	var defenseSuccess uint8 = 0
	var defenseCrit uint8 = 0
	var attackNormal uint8
	var defenseNormal uint8
	var mortalWoundDmg uint8

	for _, i := range attackRoll {
		if i >= attacker.Balistic {
			attackSuccess += 1
		}
		if i == 6 {
			attackCrit += 1
		}
	}

	// Calcuate Mortal Wounds. This function assumes all modifiers like Lethal have already been applied.
	mortalWoundDmg = attackCrit * attacker.MWx

	for _, i := range defenceRoll {
		if i >= defdefender.Balistic {
			defenseSuccess += 1
		}
		if i == 6 {
			defenseCrit += 1
		}
	}
	attackNormal = attackSuccess - attackCrit
	defenseNormal = defenseSuccess - defenseCrit

	// Standard situation, critical hit is stronger or equal normal hit
	if attacker.Damage.Critical >= attacker.Damage.Basic {
		// crits parry crits/normals
		for defenseCrit > 0 {
			if attackCrit > 0 {
				attackCrit--
			} else if attackNormal > 0 {
				attackNormal--
			} else {
				break
			}
			defenseCrit--
		}
		// Example: normal 3, crit 7. It's better to parry 1 crit than 2 normals
		if attacker.Damage.Critical > attacker.Damage.Basic*2 {
			for defenseNormal > 0 {
				if attackCrit > 0 && defenseNormal > 1 {
					attackCrit--
					defenseNormal -= 2
				} else if attackNormal > 0 {
					attackNormal--
					defenseNormal--
				} else {
					break
				}
			}
			// Example: normal 3, crit 4. It's better to parry 2 normals than 1 crit
		} else {
			for defenseNormal > 0 {
				if attackCrit >= 1 && attackNormal == 1 && defenseNormal == 2 {
					attackCrit--
					defenseNormal -= 2
				} else if attackNormal > 0 {
					attackNormal--
					defenseNormal--
				} else if attackCrit > 0 && defenseNormal > 1 {
					attackCrit--
					defenseNormal -= 2
				} else {
					break
				}
			}
		}
		// Unusual situation when normal hit is stronger than critical hit (e.x. Meltagun with MW4 has normal 6 and crit 3)
	} else {
		// First we parry normals with normals, then crits with normals
		for defenseNormal > 0 {
			if attackNormal > 0 {
				attackNormal--
				defenseNormal--
			} else if attackCrit > 0 && defenseNormal > 1 {
				attackCrit--
				defenseNormal -= 2
			} else {
				break
			}
		}
		// First we parry normals with crits, then crits with crits
		for defenseCrit > 0 {
			if attackNormal > 0 {
				attackNormal--
			} else if attackCrit > 0 {
				attackCrit--
			} else {
				break
			}
			defenseCrit--
		}
	}

	return attackCrit*attacker.Damage.Critical + attackNormal*attacker.Damage.Basic + mortalWoundDmg
}

func GetAPxDefenderRolls(attacker Attacker, defender Defender, standardDefenderRolls [][]uint8) [][]uint8 {
	if attacker.APx > 0 && defender.DiceCount >= attacker.APx {
		return cartesianProduct(defender.DiceCount - attacker.APx)
	} else if attacker.APx > 0 && defender.DiceCount < attacker.APx {
		return cartesianProduct(0)
	} else {
		return standardDefenderRolls
	}
}

func GetPxDefenderRolls(attacker Attacker, defender Defender, standardDefenderRolls [][]uint8) [][]uint8 {
	if attacker.Px > 0 && defender.DiceCount >= attacker.Px {
		return cartesianProduct(defender.DiceCount - attacker.Px)
	} else if attacker.Px > 0 && defender.DiceCount < attacker.Px {
		return cartesianProduct(0)
	} else {
		return standardDefenderRolls
	}
}

func GetDefenderRolls(attacker Attacker, isAnyCrit bool, apxDefenderRolls [][]uint8, pxDefenderRolls [][]uint8) [][]uint8 {
	// Get defender rolls based on Px and APX. if no Px and APx ability enable, then apxDefenderRolls has normal defender rolls
	if attacker.Px > attacker.APx && isAnyCrit {
		return pxDefenderRolls
	} else {
		return apxDefenderRolls
	}
}

type Attacker struct {
	DiceCount uint8
	Balistic  uint8
	Damage
	AbilitiesAttack
}

type Defender struct {
	DiceCount uint8
	Balistic  uint8
	Hp        uint8
}

type Damage struct {
	Basic    uint8
	Critical uint8
}

type AbilitiesAttack struct {
	MWx uint8 // Mortal wound
	APx uint8 // Armour penetration
	Px  uint8 // Piercing
	Reroll
	Rending            bool  // If you rolled a crit, you can retain one normal into crit
	Lethalx            uint8 // Value from which critical is success
	NormsToCritsx      uint8 // How many normals are auto convert to crits
	FailToNormalIfCrit bool  // Convert failed roll into normal if there is crit, e.x. Kommando strategic ploy "Dakka! Dakka! Dakka!"
	CloseAssault       bool  // If you have two or more sucesses, promote a fail to a normal success. Imperial Navy Breachers strategic ploy.
}

type Reroll struct {
	Ceaseless            bool //Reroll all 1
	Balanced             bool // Reroll one dice
	Relentless           bool // Reroll any or all dices
	CeaselessAndBalanced bool // Reroll 1 and then one other dice
	MostCommonFail       bool // Reroll most common fails
}

// cartesianProduct generates all possible combinations of dice rolls.
func cartesianProduct(repeat uint8) [][]uint8 {
	var size int = 1
	for i := 0; i < int(repeat); i++ {
		size *= 6 // 6 sides of a dice
	}

	results := make([][]uint8, size)
	for i := range results {
		results[i] = make([]uint8, repeat)
		tmp := i
		for j := range results[i] {
			// Determine the value for each position
			results[i][j] = uint8(tmp%6 + 1)
			tmp /= 6
		}
	}
	return results
}

// type roll struct{
// 	one uint8
// 	two uint8
// 	three uint8
// 	four uint8
// 	five uint8
// 	six uint8
// }
