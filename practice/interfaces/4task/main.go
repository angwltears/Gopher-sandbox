package main

func main() {
	W1 := []Hero{
		Warrior{
			Weapon: "Excalibur",
			Damage: 1500,
		},
		Mage{
			Mana:        500,
			SpellDamage: 5000,
		},
		Archer{
			Bow:    "XXL",
			Damage: 1200,
		},
	}
	AllAttackDamage(W1)
}
