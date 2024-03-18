package calculator_test

import (
	"ktcalc/calculator"
	"testing"
)

func TestCalcShootAttack(t *testing.T) {
	test := []struct {
		name            string
		attacker        calculator.Attacker
		defender        calculator.Defender
		attRoll         []uint8
		defRoll         []uint8
		expected_damage uint8
	}{
		// Basic parameters
		{
			"0 dmg- all misses att",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 2, 1, 2, 2}, []uint8{4, 4, 6}, 0,
		},
		{
			"0 dmg- all misses att/def",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 2, 1, 2, 2}, []uint8{1, 1, 2}, 0,
		},
		{
			"0 dmg- 1 save normal by crit",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 4, 1, 2, 2}, []uint8{1, 6, 2}, 0,
		},
		{
			"0 dmg- 1 save normal by normal",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 4, 1, 2, 2}, []uint8{1, 4, 2}, 0,
		},
		{
			"0 dmg- 1 save crit by crit",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 6, 1, 2, 2}, []uint8{1, 6, 2}, 0,
		},
		{
			"0 dmg- 1 save crit by 2 normals",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 1, 2, 6}, []uint8{1, 4, 5}, 0,
		},
		{
			"0 dmg- 2 save crit by 2 normals and crit",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 6, 2, 6}, []uint8{6, 4, 5}, 0,
		},
		{
			"0 dmg- 2 save crit by 2 normals and crit",
			calculator.Attacker{DiceCount: 5, Balistic: 5, Damage: calculator.Damage{Basic: 13, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 5, Hp: 12},
			[]uint8{1, 1, 6, 2, 6}, []uint8{6, 5, 5}, 0,
		},
		{
			"0 dmg- combined",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 6, 5, 5}, []uint8{6, 4, 5}, 0,
		},
		{
			"0 dmg- small attack",
			calculator.Attacker{DiceCount: 2, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 5}, []uint8{6, 4, 5}, 0,
		},
		{
			"3 dmg- no save",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 2, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 2, 4, 3, 1}, []uint8{2, 3}, 3,
		},
		{
			"3 dmg- 1 normal save",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 2, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 2, 4, 5, 1}, []uint8{4, 3}, 3,
		},
		{
			"3 dmg- 1 crit save",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 2, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 2, 4, 6, 1}, []uint8{6, 3}, 3,
		},
		{
			"3 dmg- combined",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 5, 4, 6, 6}, []uint8{6, 6, 6}, 3,
		},
		{
			"4 dmg- combined",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 4, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 5, 4, 6, 6}, []uint8{6, 6, 6}, 4,
		},
		{
			"4 dmg- combined",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 4, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 4, 2, 6}, []uint8{5, 5, 1}, 4,
		},
		{
			"4 dmg- combined",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 4, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 4, 4, 6}, []uint8{5, 5, 1}, 4,
		},
		{
			"7 dmg- no save",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 1, 2, 3, 6}, []uint8{1, 2, 1}, 7,
		},
		{
			"7 dmg- 1 normal save",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 1, 4, 4, 6}, []uint8{1, 5, 1}, 7,
		},
		{
			"7 dmg- 1 crit save",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 1, 4, 6, 6}, []uint8{1, 6, 1}, 7,
		},
		{
			"7 dmg- 2 normal saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 5, 4, 4, 6}, []uint8{3, 5, 1}, 7,
		},
		{
			"7 dmg- 1 crit 1 normal saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 5, 4, 6, 6}, []uint8{3, 6, 1}, 7,
		},
		{
			"7 dmg- 2 crit saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 5, 6, 6, 6}, []uint8{6, 6, 1}, 7,
		},
		{
			"7 dmg- 2 normal 1 crit saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{6, 5, 4, 4, 6}, []uint8{3, 5, 6}, 7,
		},
		{
			"7 dmg- 2 crits 1 normal saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{4, 4, 6, 6, 6}, []uint8{4, 6, 6}, 7,
		},
		{
			"7 dmg- combined",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{4, 1, 6, 6, 6}, []uint8{2, 6, 6}, 7,
		},
		{
			"7 dmg- 3 normal saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{4, 4, 3, 5, 6}, []uint8{3, 4, 5}, 7,
		},
		{
			"7 dmg- 3 normal saves xD",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 4, 3, 6, 6}, []uint8{3, 4, 5}, 7,
		},
		{
			"7 dmg- 3 crit saves",
			calculator.Attacker{DiceCount: 5, Balistic: 3, Damage: calculator.Damage{Basic: 3, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{4, 6, 6, 6, 6}, []uint8{6, 6, 6}, 7,
		},
		{
			"Basic greater than crit",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 10, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 1, 5, 4, 6, 6}, []uint8{6, 6, 6}, 4,
		},
		{
			"Basic greater than crit",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 10, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 5, 5, 4, 6, 6}, []uint8{6, 6, 6}, 8,
		},
		{
			"Basic greater than crit",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 10, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 5, 2, 4, 3, 6}, []uint8{5, 5, 1}, 4,
		},
		{
			"Basic greater than crit",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 10, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 2, 2, 1, 3, 2}, []uint8{2, 2, 1}, 0,
		},
		{
			"Crit over 2x better than normal",
			calculator.Attacker{DiceCount: 6, Balistic: 4, Damage: calculator.Damage{Basic: 1, Critical: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 4, Hp: 12},
			[]uint8{1, 2, 2, 4, 4, 6}, []uint8{5, 5, 1}, 2,
		},
		// MWx tests
		{
			"MWx no crit 0 dmg",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4},
				AbilitiesAttack: calculator.AbilitiesAttack{MWx: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 3, 4, 4}, []uint8{4, 5, 1}, 0,
		},
		{
			"MWx no crit 3 dmg",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4},
				AbilitiesAttack: calculator.AbilitiesAttack{MWx: 4}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 5, 4, 4}, []uint8{4, 5, 1}, 3,
		},
		{
			"MWx only 3 MW",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4},
				AbilitiesAttack: calculator.AbilitiesAttack{MWx: 2}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 6, 6, 6}, []uint8{6, 6, 6}, 6,
		},
		{
			"MWx 2 crit 11 dmg",
			calculator.Attacker{DiceCount: 4, Balistic: 4, Damage: calculator.Damage{Basic: 3, Critical: 4},
				AbilitiesAttack: calculator.AbilitiesAttack{MWx: 2}},
			calculator.Defender{DiceCount: 3, Balistic: 3, Hp: 12},
			[]uint8{1, 4, 6, 6}, []uint8{1, 4, 5}, 11,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			var result uint8 = calculator.CalcShootAttack(tc.attacker, tc.defender, tc.attRoll, tc.defRoll)
			if result != tc.expected_damage {
				t.Errorf("Test '%v' failed, %v =/= %v", tc.name, result, tc.expected_damage)
			}
		})
	}
}
