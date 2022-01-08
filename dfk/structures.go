package dfk

import "time"

type Rarity int

const (
	Mythic    Rarity = 4
	Legendary Rarity = 3
	Rare      Rarity = 2
	Uncommon  Rarity = 1
	Common    Rarity = 0
)

type DataWrapHeroResponse struct {
	Hero HeroResponse `json:"hero"`
}

type HeroResponse struct {
	Active1             string `json:"active1"`
	Active2             string `json:"active2"`
	Agility             int    `json:"agility"`
	AgilityGrowthP      int    `json:"agilityGrowthP"`
	AgilityGrowthS      int    `json:"agilityGrowthS"`
	Dexterity           int    `json:"dexterity"`
	DexterityGrowthP    int    `json:"dexterityGrowthP"`
	DexterityGrowthS    int    `json:"dexterityGrowthS"`
	Endurance           int    `json:"endurance"`
	EnduranceGrowthP    int    `json:"enduranceGrowthP"`
	EnduranceGrowthS    int    `json:"enduranceGrowthS"`
	Fishing             int    `json:"fishing"`
	Foraging            int    `json:"foraging"`
	Gardening           int    `json:"gardening"`
	Generation          int    `json:"generation"`
	Hp                  int    `json:"hp"`
	HpLgGrowth          int    `json:"hpLgGrowth"`
	HpRgGrowth          int    `json:"hpRgGrowth"`
	HpSmGrowth          int    `json:"hpSmGrowth"`
	ID                  string `json:"id"`
	Intelligence        int    `json:"intelligence"`
	IntelligenceGrowthP int    `json:"intelligenceGrowthP"`
	IntelligenceGrowthS int    `json:"intelligenceGrowthS"`
	Level               int    `json:"level"`
	Luck                int    `json:"luck"`
	LuckGrowthP         int    `json:"luckGrowthP"`
	LuckGrowthS         int    `json:"luckGrowthS"`
	MainClass           string `json:"mainClass"`
	MaxSummons          int    `json:"maxSummons"`
	Mining              int    `json:"mining"`
	Mp                  int    `json:"mp"`
	MpLgGrowth          int    `json:"mpLgGrowth"`
	MpRgGrowth          int    `json:"mpRgGrowth"`
	MpSmGrowth          int    `json:"mpSmGrowth"`
	Owner               Owner  `json:"owner"`
	Passive1            string `json:"passive1"`
	Passive2            string `json:"passive2"`
	Profession          string `json:"profession"`
	Rarity              int    `json:"rarity"`
	Shiny               bool   `json:"shiny"`
	Stamina             int    `json:"stamina"`
	StaminaFullAt       string `json:"staminaFullAt"`
	StatBoost1          string `json:"statBoost1"`
	StatBoost2          string `json:"statBoost2"`
	StatGenes           string `json:"statGenes"`
	Status              string `json:"status"`
	Strength            int    `json:"strength"`
	StrengthGrowthP     int    `json:"strengthGrowthP"`
	StrengthGrowthS     int    `json:"strengthGrowthS"`
	SubClass            string `json:"subClass"`
	Summons             int    `json:"summons"`
	Vitality            int    `json:"vitality"`
	VitalityGrowthP     int    `json:"vitalityGrowthP"`
	VitalityGrowthS     int    `json:"vitalityGrowthS"`
	Wisdom              int    `json:"wisdom"`
	WisdomGrowthP       int    `json:"wisdomGrowthP"`
	WisdomGrowthS       int    `json:"wisdomGrowthS"`
	Xp                  string `json:"xp"`
}

type Owner struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type Hero struct {
	ID          string          `json:"id"`
	Stats       HeroStats       `json:"stats"`
	Professions HeroProfessions `json:"professions"`
	Actives     [2]string       `json:"actives"`
	Passives    [2]string       `json:"passives"`
	Generation  int             `json:"generation"`
	Level       int             `json:"level"`
	XP          string          `json:"xp"`
	Rarity      Rarity          `json:"rarity"`
	Shiny       bool            `json:"shiny"`
	Summons     HeroSummons     `json:"summons"`
	Owner       Owner           `json:"owner"`
	StatGenes   int             `json:"statGenes"`
	Status      string          `json:"status"`
}

type HeroClasses struct {
	Main string `json:"main"`
	Sub  string `json:"sub"`
}

type HeroStamina struct {
	Points        int   `json:"points"`
	FullTimestamp int64 `json:"full"`
}

func (h HeroStamina) SecondsLeft() int {
	return int(h.FullTimestamp - time.Now().Unix())
}

type HeroStatBoosts struct {
	Green string `json:"green"`
	Blue  string `json:"blue"`
}

func (h HeroStatBoosts) Purple() (*string, bool) {
	if h.Green == h.Blue {
		return &h.Green, true
	}
	return nil, false
}

type HeroSummons struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type HeroCombatStats struct {
	HP HeroCombatStat `json:"hp"`
	MP HeroCombatStat `json:"mp"`
}

type HeroCombatStat struct {
	Stat   string               `json:"stat"`
	Growth HeroCombatStatGrowth `json:"growth"`
}

type HeroCombatStatGrowth struct {
	Large    int `json:"large"`
	RgGrowth int `json:"recessiveGene"`
	Small    int `json:"small"`
}

type HeroProfessions struct {
	Gene      string `json:"gene"`
	Fishing   int    `json:"fishing"`
	Foraging  int    `json:"foraging"`
	Gardening int    `json:"gardening"`
	Mining    int    `json:"mining"`
}

type HeroStats struct {
	Agility      HeroStat `json:"agility"`
	Dexterity    HeroStat `json:"dexterity"`
	Endurance    HeroStat `json:"endurance"`
	Intelligence HeroStat `json:"intelligence"`
	Strength     HeroStat `json:"strength"`
	Vitality     HeroStat `json:"vitality"`
	Luck         HeroStat `json:"luck"`
	Wisdom       HeroStat `json:"wisdom"`
}

type HeroStat struct {
	Points string         `json:"points"`
	Stat   int            `json:"stat"`
	Growth HeroStatGrowth `json:"growth"`
}

type HeroStatGrowth struct {
	P int `json:"p"`
	S int `json:"s"`
}
