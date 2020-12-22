package selectors

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack"
)

// AchievementSelectors represents all CSS selectors associated with character achievements.
type AchievementSelectors struct {
	List           SelectorInfo `json:"LIST"`
	Entry          SelectorInfo `json:"ENTRY"`
	ID             SelectorInfo `json:"ID"`
	Time           SelectorInfo `json:"TIME"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
}

// AttributeSelectors represents all CSS selectors associated with character attributes.
type AttributeSelectors struct {
	Strength            SelectorInfo `json:"STRENGTH"`
	Dexterity           SelectorInfo `json:"DEXTERITY"`
	Vitality            SelectorInfo `json:"VITALITY"`
	Intelligence        SelectorInfo `json:"INTELLIGENCE"`
	Mind                SelectorInfo `json:"MIND"`
	CriticalHitRate     SelectorInfo `json:"CRITICAL_HIT_RATE"`
	Determination       SelectorInfo `json:"DETERMINATION"`
	DirectHitRate       SelectorInfo `json:"DIRECT_HIT_RATE"`
	Defense             SelectorInfo `json:"DEFENSE"`
	MagicDefense        SelectorInfo `json:"MAGIC_DEFENSE"`
	AttackPower         SelectorInfo `json:"ATTACK_POWER"`
	SkillSpeed          SelectorInfo `json:"SKILL_SPEED"`
	AttackMagicPotency  SelectorInfo `json:"ATTACK_MAGIC_POTENCY"`
	HealingMagicPotency SelectorInfo `json:"HEALING_MAGIC_POTENCY"`
	SpellSpeed          SelectorInfo `json:"SPELL_SPEED"`
	Tenacity            SelectorInfo `json:"TENACITY"`
	Piety               SelectorInfo `json:"PIETY"`
	HP                  SelectorInfo `json:"HP"`
	MPGPCP              SelectorInfo `json:"MP_GP_CP"`
	MPGPCPParameterName SelectorInfo `json:"MP_GP_CP_PARAMETER_NAME"`
}

// CharacterSelectors represents all CSS selectors associated with basic character data.
type CharacterSelectors struct {
	ActiveClassJob      SelectorInfo `json:"ACTIVE_CLASSJOB"`
	ActiveClassJobLevel SelectorInfo `json:"ACTIVE_CLASSJOB_LEVEL"`
	Avatar              SelectorInfo `json:"AVATAR"`
	Bio                 SelectorInfo `json:"BIO"`
	FreeCompany         struct {
		Name       SelectorInfo `json:"NAME"`
		IconLayers struct {
			Bottom SelectorInfo `json:"BOTTOM"`
			Middle SelectorInfo `json:"MIDDLE"`
			Top    SelectorInfo `json:"TOP"`
		} `json:"ICON_LAYERS"`
	} `json:"FREE_COMPANY"`
	GrandCompany  SelectorInfo `json:"GRAND_COMPANY"`
	GuardianDeity SelectorInfo `json:"GUARDIAN_DEITY"`
	Name          SelectorInfo `json:"NAME"`
	Nameday       SelectorInfo `json:"NAMEDAY"`
	Portrait      SelectorInfo `json:"PORTRAIT"`
	PvPTeam       struct {
		Name       SelectorInfo `json:"NAME"`
		IconLayers struct {
			Bottom SelectorInfo `json:"BOTTOM"`
			Middle SelectorInfo `json:"MIDDLE"`
			Top    SelectorInfo `json:"TOP"`
		} `json:"ICON_LAYERS"`
	} `json:"PVP_TEAM"`
	RaceClanGender SelectorInfo `json:"RACE_CLAN_GENDER"`
	Server         SelectorInfo `json:"SERVER"`
	Title          SelectorInfo `json:"TITLE"`
	Town           SelectorInfo `json:"TOWN"`
}

// MinionSelectors represents all CSS selectors associated with character minions.
type MinionSelectors struct {
	List SelectorInfo `json:"LIST"`
	Name SelectorInfo `json:"NAME"`
}

// MountSelectors represents all CSS selectors associated with character mounts.
type MountSelectors struct {
	List SelectorInfo `json:"LIST"`
	Name SelectorInfo `json:"NAME"`
}

// GearSelectors represents the selectors associated with a character gear piece.
type GearSelectors struct {
	Name         SelectorInfo `json:"NAME"`
	DBLink       SelectorInfo `json:"DB_LINK"`
	MirageName   SelectorInfo `json:"MIRAGE_NAME"`
	MirageDBLink SelectorInfo `json:"MIRAGE_DB_LINK"`
	Stain        SelectorInfo `json:"STAIN"`
	Materia1     SelectorInfo `json:"MATERIA_1"`
	Materia2     SelectorInfo `json:"MATERIA_2"`
	Materia3     SelectorInfo `json:"MATERIA_3"`
	Materia4     SelectorInfo `json:"MATERIA_4"`
	Materia5     SelectorInfo `json:"MATERIA_5"`
	CreatorName  SelectorInfo `json:"CREATOR_NAME"`
}

// GearSetSelectors represents the selectors associated with a character gearset.
type GearSetSelectors struct {
	MainHand    GearSelectors `json:"MAINHAND"`
	OffHand     GearSelectors `json:"OFFHAND"`
	Head        GearSelectors `json:"HEAD"`
	Body        GearSelectors `json:"BODY"`
	Hands       GearSelectors `json:"HANDS"`
	Waist       GearSelectors `json:"WAIST"`
	Legs        GearSelectors `json:"LEGS"`
	Feet        GearSelectors `json:"FEET"`
	Earrings    GearSelectors `json:"EARRINGS"`
	Necklace    GearSelectors `json:"NECKLACE"`
	Bracelets   GearSelectors `json:"BRACELETS"`
	Ring1       GearSelectors `json:"RING1"`
	Ring2       GearSelectors `json:"RING2"`
	SoulCrystal struct {
		Name SelectorInfo `json:"NAME"`
	} `json:"SOULCRYSTAL"`
}

// OneClassJobSelectors represents all CSS selectors associated with a character's ClassJob.
type OneClassJobSelectors struct {
	Level       SelectorInfo `json:"LEVEL"`
	UnlockState SelectorInfo `json:"UNLOCKSTATE"`
	Exp         SelectorInfo `json:"EXP"`
}

// ClassJobSelectors represents all CSS selectors associated with all of a character's ClassJobs.
type ClassJobSelectors struct {
	Bozja struct {
		Level  SelectorInfo `json:"LEVEL"`
		Mettle SelectorInfo `json:"METTLE"`
		Name   SelectorInfo `json:"NAME"`
	} `json:"BOZJA"`
	Eureka struct {
		Level  SelectorInfo `json:"LEVEL"`
		Mettle SelectorInfo `json:"METTLE"`
		Name   SelectorInfo `json:"NAME"`
	} `json:"EUREKA"`
	Paladin       OneClassJobSelectors `json:"PALADIN"`
	Warrior       OneClassJobSelectors `json:"WARRIOR"`
	DarkKnight    OneClassJobSelectors `json:"DARKKNIGHT"`
	Gunbreaker    OneClassJobSelectors `json:"GUNBREAKER"`
	Monk          OneClassJobSelectors `json:"MONK"`
	Dragoon       OneClassJobSelectors `json:"DRAGOON"`
	Ninja         OneClassJobSelectors `json:"NINJA"`
	Samurai       OneClassJobSelectors `json:"SAMURAI"`
	WhiteMage     OneClassJobSelectors `json:"WHITEMAGE"`
	Scholar       OneClassJobSelectors `json:"SCHOLAR"`
	Astrologian   OneClassJobSelectors `json:"ASTROLOGIAN"`
	Bard          OneClassJobSelectors `json:"BARD"`
	Machinist     OneClassJobSelectors `json:"MACHINIST"`
	Dancer        OneClassJobSelectors `json:"DANCER"`
	BlackMage     OneClassJobSelectors `json:"BLACKMAGE"`
	Summoner      OneClassJobSelectors `json:"SUMMONER"`
	RedMage       OneClassJobSelectors `json:"REDMAGE"`
	BlueMage      OneClassJobSelectors `json:"BLUEMAGE"`
	Carpenter     OneClassJobSelectors `json:"CARPENTER"`
	Blacksmith    OneClassJobSelectors `json:"BLACKSMITH"`
	Armorer       OneClassJobSelectors `json:"ARMORER"`
	Goldsmith     OneClassJobSelectors `json:"GOLDSMITH"`
	Leatherworker OneClassJobSelectors `json:"LEATHERWORKER"`
	Weaver        OneClassJobSelectors `json:"WEAVER"`
	Alchemist     OneClassJobSelectors `json:"ALCHEMIST"`
	Culinarian    OneClassJobSelectors `json:"CULINARIAN"`
	Miner         OneClassJobSelectors `json:"MINER"`
	Botanist      OneClassJobSelectors `json:"BOTANIST"`
	Fisher        OneClassJobSelectors `json:"FISHER"`
}

// ProfileSelectors represents all CSS selectors associated with a character profile.
type ProfileSelectors struct {
	Achievements *AchievementSelectors
	Attributes   *AttributeSelectors
	Character    *CharacterSelectors
	ClassJob     *ClassJobSelectors
	GearSet      *GearSetSelectors
	Minion       *MinionSelectors
	Mount        *MountSelectors
}

// LoadProfileSelectors loads the profile selectors.
func LoadProfileSelectors() (*ProfileSelectors, error) {
	achievementsBytes, err := pack.Asset("profile/achievements.json")
	if err != nil {
		return nil, err
	}
	achievements := AchievementSelectors{}
	json.Unmarshal(achievementsBytes, &achievements)

	attributesBytes, err := pack.Asset("profile/attributes.json")
	if err != nil {
		return nil, err
	}
	attributes := AttributeSelectors{}
	json.Unmarshal(attributesBytes, &attributes)

	characterBytes, err := pack.Asset("profile/character.json")
	if err != nil {
		return nil, err
	}
	character := CharacterSelectors{}
	json.Unmarshal(characterBytes, &character)

	classJobBytes, err := pack.Asset("profile/classjob.json")
	if err != nil {
		return nil, err
	}
	classJob := ClassJobSelectors{}
	json.Unmarshal(classJobBytes, &classJob)

	gearsetBytes, err := pack.Asset("profile/gearset.json")
	if err != nil {
		return nil, err
	}
	gearset := GearSetSelectors{}
	json.Unmarshal(gearsetBytes, &gearset)

	minionBytes, err := pack.Asset("profile/minion.json")
	if err != nil {
		return nil, err
	}
	minion := MinionSelectors{}
	json.Unmarshal(minionBytes, &minion)

	mountBytes, err := pack.Asset("profile/mount.json")
	if err != nil {
		return nil, err
	}
	mount := MountSelectors{}
	json.Unmarshal(mountBytes, &mount)

	return &ProfileSelectors{
		Achievements: &achievements,
		Attributes:   &attributes,
		Character:    &character,
		ClassJob:     &classJob,
		GearSet:      &gearset,
		Minion:       &minion,
		Mount:        &mount,
	}, nil
}