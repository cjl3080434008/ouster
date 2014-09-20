package main

import (
    "bufio"
    "bytes"
    "fmt"
    "strconv"
)

var (
    BEGIN     = []byte("DIRECTIVE BEGIN")
    END       = []byte("DIRECTIVE END")
    CONDITION = []byte("(CONDITION:")
    ACTION    = []byte("(ACTION:")
)

func ReadLine(reader *bufio.Reader) (line []byte, err error) {
    line, err = reader.ReadBytes('\n')
    if err != nil {
        return
    }
    line = skipWhiteSpace(line)
    return
}

const input = `DIRECTIVE BEGIN
	(CONDITION:MasterSummonTiming)
	(ACTION:UseSkill,SKILL_SUMMON_MONSTERS,100)
DIRECTIVE END
DIRECTIVE BEGIN
	(CONDITION:MasterNotReady)
	(CONDITION:EnemyRangeMissile)
	(CONDITION:EnemyNotHallucination)
	(ACTION:UseSkill,SKILL_HALLUCINATION,20)
DIRECTIVE END
DIRECTIVE BEGIN
	(CONDITION:MasterNotReady)
	(ACTION:Wait)
DIRECTIVE END`

func main() {
    reader := bufio.NewReader(bytes.NewBufferString(input))

    set := &DirectiveSet{}
    for {
        line, err := ReadLine(reader)
        if err != nil {
            break
        }

        if bytes.HasPrefix(line, BEGIN) {
            set.parseBegin()
        } else if bytes.HasPrefix(line, CONDITION) {
            set.parseCondition(line)
        } else if bytes.HasPrefix(line, ACTION) {
            set.parseAction(line)
        } else if bytes.HasPrefix(line, END) {
            set.parseEnd()
        }
    }

    for _, v := range set.Directives {
        fmt.Printf("%#v\n", v)
    }
}

func skipWhiteSpace(line []byte) []byte {
    i := 0
    for {
        if line[i] == '\t' || line[i] == ' ' {
            i++
        } else {
            break
        }
    }
    return line[i:]
}

func (set *DirectiveSet) parseBegin() {
    set.Directives = append(set.Directives, &Directive{})
}

func (set *DirectiveSet) parseCondition(line []byte) {
    line = line[len(CONDITION):]
    idx := bytes.IndexByte(line, ')')
    if idx == -1 {
        panic("不对")
    }
    directive := set.Directives[len(set.Directives)-1]
    cond := string(line[:idx])
    directive.Conditions = append(directive.Conditions, DirectCond(cond))
    return
}

func (set *DirectiveSet) parseAction(line []byte) (action string, pamater string, ratio string) {
    directive := set.Directives[len(set.Directives)-1]
    line = line[len(ACTION):]
    idx := 0
    for {
        i := 0
        for ; i < len(line); i++ {
            if line[i] == ',' || line[i] == ')' {
                break
            }
        }
        token := line[:i]
        switch idx {
        case 0:
            action = string(token)
            directive.Action = DirectAction(action)
        case 1:
            pamater = string(token)
            directive.Parameter = skill2int(pamater)
        case 2:
            ratio = string(token)
            directive.Ratio, _ = strconv.Atoi(ratio)
        }
        idx++
        if line[i] == ')' {
            return
        }
        line = line[i+1:]
    }
}

func (set *DirectiveSet) parseEnd() {

}

const (
    DIRECTIVE_COND_ENEMY_RANGE_MELEE = iota
    DIRECTIVE_COND_ENEMY_RANGE_MISSILE
    DIRECTIVE_COND_ENEMY_RANGE_CLOSE
    DIRECTIVE_COND_ENEMY_RANGE_OUT_OF_SIGHT
    DIRECTIVE_COND_ENEMY_DYING
    DIRECTIVE_COND_ENEMY_NOT_BLOOD_DRAINED
    DIRECTIVE_COND_ENEMY_NOT_GREEN_POISONED
    DIRECTIVE_COND_ENEMY_NOT_YELLOW_POISONED
    DIRECTIVE_COND_ENEMY_NOT_DARKBLUE_POISONED
    DIRECTIVE_COND_ENEMY_NOT_GREEN_STALKERED
    DIRECTIVE_COND_ENEMY_NOT_PARALYZED
    DIRECTIVE_COND_ENEMY_NOT_DOOMED
    DIRECTIVE_COND_ENEMY_NOT_BLINDED
    DIRECTIVE_COND_ENEMY_NOT_IN_DARKNESS
    DIRECTIVE_COND_ENEMY_NOT_SEDUCTION
    DIRECTIVE_COND_IM_OK
    DIRECTIVE_COND_IM_DYING
    DIRECTIVE_COND_IM_DAMAGED
    DIRECTIVE_COND_IM_HIDING
    DIRECTIVE_COND_IM_WOLF
    DIRECTIVE_COND_IM_BAT
    DIRECTIVE_COND_IM_INVISIBLE
    DIRECTIVE_COND_IM_WALKING_WALL
    DIRECTIVE_COND_TIMING_BLOOD_DRAIN
    DIRECTIVE_COND_MASTER_SUMMON_TIMING
    DIRECTIVE_COND_MASTER_NOT_READY
    DIRECTIVE_COND_IM_IN_BAD_POSITION
    DIRECTIVE_COND_FIND_WEAK_ENEMY
    DIRECTIVE_COND_ENEMY_NOT_DEATH
    DIRECTIVE_COND_ENEMY_NOT_HALLUCINATION
    DIRECTIVE_COND_TIMING_MASTER_BLOOD_DRAIN
    DIRECTIVE_COND_TIMING_DUPLICATE_SELF
    DIRECTIVE_COND_ENEMY_RANGE_IN_MISSILE
    DIRECTIVE_COND_POSSIBLE_SUMMON_MONSTERS
    DIRECTIVE_COND_ENEMY_TILE_NOT_ACID_SWAMP
    DIRECTIVE_COND_ENEMY_ON_AIR
    DIRECTIVE_COND_ENEMY_ON_SAFE_ZONE
    DIRECTIVE_COND_CAN_ATTACK_THROWING_AXE
    DIRECTIVE_COND_MAX
)

const (
    DIRECTIVE_ACTION_APPROACH = iota
    DIRECTIVE_ACTION_FLEE
    DIRECTIVE_ACTION_USE_SKILL
    DIRECTIVE_ACTION_FORGET
    DIRECTIVE_ACTION_CHANGE_ENEMY
    DIRECTIVE_ACTION_MOVE_RANDOM
    DIRECTIVE_ACTION_WAIT
    DIRECTIVE_ACTION_FAST_FLEE
    DIRECTIVE_ACTION_SAY
    DIRECTIVE_ACTION_MAX
)

type Directive struct {
    Conditions []int
    Action     int
    Parameter  int
    Ratio      int
    Weight     int
}

type DirectiveSet struct {
    Directives     []*Directive
    DeadDirectives []*Directive
    Name           string

    bAttackAir   bool
    bSeeSafeZone bool
}

func DirectAction(action string) int {
    for i := 0; i < len(DirectiveAction2String); i++ {
        if DirectiveAction2String[i] == action {
            return i
        }
    }
    return DIRECTIVE_ACTION_MAX
}

func DirectCond(cond string) int {
    for i := 0; i < len(DirectiveCondition2String); i++ {
        if DirectiveCondition2String[i] == cond {
            return i
        }
    }
    return DIRECTIVE_COND_MAX
}

var DirectiveAction2String = []string{
    "Approach",
    "Flee",
    "UseSkill",
    "Forget",
    "ChangeEnemy",
    "MoveRandom",
    "Wait",
    "FastFlee",
    "Say",
    "ActionMAX",
}

var DirectiveCondition2String = []string{
    "EnemyRangeMelee",
    "EnemyRangeMissile",
    "EnemyRangeClose",
    "EnemyRangeOutOfSight",
    "EnemyDying",
    "EnemyNotBloodDrained",
    "EnemyNotGreenPoisoned",
    "EnemyNotYellowPoisoned",
    "EnemyNotDarkbluePoisoned",
    "EnemyNotGreenStalkered",
    "EnemyNotParalyzed",
    "EnemyNotDoomed",
    "EnemyNotBlinded",
    "EnemyNotInDarkness",
    "EnemyNotSeduction",
    "ImOK",
    "ImDying",
    "ImDamaged",
    "ImHiding",
    "ImWolf",
    "ImBat",
    "ImInvisible",
    "ImWalkingWall",
    "TimingBloodDrain",
    "MasterSummonTiming",
    "MasterNotReady",
    "ImInBadPosition",
    "FindWeakEnemy",
    "EnemyNotDeath",
    "EnemyNotHallucination",
    "TimingMasterBloodDrain",
    "TimingDuplicateSelf",
    "EnemyRangeInMissile",
    "PossibleSummonMonsters",
    "EnemyTileNotAcidSwamp",
    "EnemyOnAir",
    "EnemyOnSafeZone",
    "CanAttackThrowingAxe",
    "ConditionMAX",
}

const (
    SKILL_ATTACK_MELEE = 0
    SKILL_ATTACK_ARMS
    SKILL_SELF
    SKILL_TILE
    SKILL_OBJECT

    SKILL_DOUBLE_IMPACT
    SKILL_TRIPLE_SLASHER
    SKILL_RAINBOW_SLASHER
    SKILL_THUNDER_SPARK
    SKILL_DANCING_SWORD
    SKILL_CROSS_COUNTER
    SKILL_FLASH_SLIDING
    SKILL_LIGHTNING_HAND
    SKILL_SNAKE_COMBO
    SKILL_SWORD_WAVE
    SKILL_DRAGON_RISING
    SKILL_FIVE_STORM_CRASH
    SKILL_HEAVENS_SWORD

    SKILL_SINGLE_BLOW
    SKILL_SPIRAL_SLAY
    SKILL_TRIPLE_BREAK
    SKILL_WILD_SMASH
    SKILL_GHOST_BLADE
    SKILL_POTENTIAL_EXPLOSION
    SKILL_SHADOW_WALK
    SKILL_CHARGING_POWER
    SKILL_HURRICANE_COMBO
    SKILL_TORNADO_SEVER
    SKILL_ARMAGEDDON_SLASHER
    SKILL_SOUL_SHOCK
    SKILL_SAINT_BLADE

    SKILL_FAST_RELOAD
    SKILL_QUICK_FIRE
    SKILL_SMG_MASTERY
    SKILL_MULTI_SHOT
    SKILL_HEAD_SHOT
    SKILL_AR_MASTERY
    SKILL_PIERCING
    SKILL_SNIPING
    SKILL_SG_MASTERY
    SKILL_REVEALER
    SKILL_CREATE_BOMB
    SKILL_SR_MASTERY
    SKILL_DISARM_MINE
    SKILL_INSTALL_MINE
    SKILL_CREATE_MINE

    SKILL_CREATE_HOLY_WATER
    SKILL_LIGHT
    SKILL_DETECT_HIDDEN
    SKILL_AURA_BALL
    SKILL_BLESS
    SKILL_CONTINUAL_LIGHT
    SKILL_FLARE
    SKILL_PURIFY
    SKILL_AURA_RING
    SKILL_STRIKING
    SKILL_DETECT_INVISIBILITY
    SKILL_IDENTIFY
    SKILL_AURA_SHIELD
    SKILL_ENCHANT
    SKILL_VISIBLE
    SKILL_CHAIN_AURA
    SKILL_SAINT_AURA

    SKILL_CURE_LIGHT_WOUNDS
    SKILL_CURE_POISON
    SKILL_PROTECTION_FROM_POISON
    SKILL_CAUSE_LIGHT_WOUNDS
    SKILL_CURE_SERIOUS_WOUNDS
    SKILL_REMOVE_CURSE
    SKILL_PROTECTION_FROM_CURSE
    SKILL_CAUSE_SERIOUS_WOUNDS
    SKILL_CURE_CRITICAL_WOUNDS
    SKILL_PROTECTION_FROM_ACID
    SKILL_SACRIFICE
    SKILL_CAUSE_CRITICAL_WOUNDS
    SKILL_CURE_ALL
    SKILL_REGENERATION
    SKILL_MASS_CURE
    SKILL_MASS_HEAL

    SKILL_BLOOD_DRAIN

    SKILL_POISONOUS_HANDS
    SKILL_GREEN_POISON
    SKILL_YELLOW_POISON
    SKILL_DARKBLUE_POISON
    SKILL_GREEN_STALKER

    SKILL_ACID_TOUCH
    SKILL_ACID_BOLT
    SKILL_ACID_BALL
    SKILL_ACID_SWAMP

    SKILL_PARALYZE
    SKILL_DOOM
    SKILL_HALLUCINATION
    SKILL_DEATH

    SKILL_BLOODY_NAIL
    SKILL_BLOODY_KNIFE
    SKILL_BLOODY_BALL
    SKILL_BLOODY_WALL
    SKILL_BLOODY_SPEAR

    SKILL_HIDE
    SKILL_DARKNESS
    SKILL_INVISIBILITY
    SKILL_TRANSFORM_TO_WOLF
    SKILL_TRANSFORM_TO_BAT

    SKILL_SUMMON_WOLF
    SKILL_SUMMON_CASKET
    SKILL_RAISING_DEAD
    SKILL_SUMMON_SERVANT

    SKILL_UN_BURROW
    SKILL_UN_TRANSFORM
    SKILL_UN_INVISIBILITY
    SKILL_THROW_HOLY_WATER

    SKILL_EAT_CORPSE
    SKILL_HOWL

    SKILL_RESTORE

    SKILL_BLOODY_MARKER
    SKILL_BLOODY_TUNNEL

    SKILL_SEDUCTION
    SKILL_WIND_DIVIDER
    SKILL_EARTHQUAKE

    SKILL_RESURRECT
    SKILL_PRAYER
    SKILL_MEDITATION

    SKILL_THUNDER_BOLT
    SKILL_EXPANSION
    SKILL_MIRACLE_SHIELD
    SKILL_THUNDER_FLASH
    SKILL_MASSACRE
    SKILL_INVINCIBLE

    SKILL_BERSERKER
    SKILL_MOONLIGHT_SEVER
    SKILL_SHADOW_DANCING
    SKILL_TYPHOON
    SKILL_PSYCHOKINESIS
    SKILL_EXTERMINATION

    SKILL_MIND_CONTROL
    SKILL_REVOLVING
    SKILL_FATALITY
    SKILL_BLITZ

    SKILL_ACTIVATION
    SKILL_PEACE
    SKILL_ENERGY_DROP
    SKILL_EXORCISM

    SKILL_SANCTUARY
    SKILL_REFLECTION
    SKILL_ARMAGEDDON

    SKILL_POISON_STRIKE
    SKILL_POISON_STORM
    SKILL_ACID_STRIKE
    SKILL_ACID_STORM
    SKILL_BLOODY_STRIKE
    SKILL_BLOODY_STORM

    // ÀÓ½Ã·Î Ãß°¡ 2002.2.22
    SKILL_SUMMON_BAT
    SKILL_CHARM
    SKILL_POLYMORPH
    SKILL_MEPHISTO
    SKILL_HYPNOSIS
    SKILL_TRANSFUSION
    SKILL_EXTREME
    SKILL_BLOODY_WAVE
    SKILL_THROW_BOMB

    SKILL_DOUBLE_SHOT
    SKILL_TRIPLE_SHOT
    SKILL_CURE_EFFECT
    SKILL_CRITICAL_EFFECT
    SKILL_CRITICAL_GROUND
    SKILL_VIGOR_DROP

    // by sigi. 2002.6.7
    SKILL_SWORD_MASTERY
    SKILL_SHIELD_MASTERY
    SKILL_THUNDER_STORM
    SKILL_CONCENTRATION
    SKILL_EVASION
    SKILL_HOLY_BLAST
    SKILL_HYMN
    SKILL_MENTAL_SWORD
    SKILL_OBSERVING_EYE
    SKILL_REFLECTION_EFFECT

    // 2002.6.21
    SKILL_TEMP
    SKILL_OPEN_CASKET

    // 2002.9.2
    SKILL_SUMMON_MONSTERS
    SKILL_GROUND_ATTACK

    // 2002.9.13
    SKILL_METEOR_STRIKE

    // 2002.9.23
    SKILL_DUPLICATE_SELF

    // 2002.9.24
    SKILL_BLOODY_MASTER_WAVE

    // 2002.10.1
    SKILL_BLOODY_WARP
    SKILL_BLOODY_SNAKE

    // 2002.10.22
    SKILL_SOUL_CHAIN

    // 2002.11.18
    SKILL_LIVENESS

    // 2002.11.20
    SKILL_DARKNESS_WIDE
    SKILL_POISON_STORM_WIDE
    SKILL_ACID_STORM_WIDE

    // 2002.12.26
    SKILL_SHARP_SHIELD
    SKILL_WIDE_LIGHTNING
    SKILL_AIR_SHIELD
    SKILL_POWER_OF_LAND
    SKILL_BULLET_OF_LIGHT
    SKILL_GUN_SHOT_GUIDANCE
    SKILL_REBUKE
    SKILL_SPIRIT_GUARD
    SKILL_TURN_UNDEAD

    SKILL_HANDS_OF_WISDOM
    SKILL_LIGHT_BALL
    SKILL_HOLY_ARROW

    // 2003.02.26
    SKILL_BLOODY_BREAKER
    SKILL_RAPID_GLIDING

    SKILL_MAGIC_ELUSION
    SKILL_POISON_MESH
    SKILL_ILLUSION_OF_AVENGE
    SKILL_WILL_OF_LIFE

    SKILL_DENIAL_MAGIC
    SKILL_REQUITAL
    SKILL_CONCEALMENT
    SKILL_SWORD_RAY
    SKILL_MULTI_AMPUTATE
    SKILL_NAIL_MASTERY

    SKILL_HIT_CONVERT
    SKILL_WILD_TYPHOON
    SKILL_ULTIMATE_BLOW
    SKILL_ILLENDUE
    SKILL_LIGHTNESS

    SKILL_FLOURISH
    SKILL_EVADE
    SKILL_SHARP_ROUND
    SKILL_HIDE_SIGHT
    SKILL_BACK_STAB
    SKILL_BLUNTING
    SKILL_GAMMA_CHOP
    SKILL_CROSS_GUARD
    SKILL_FIRE_OF_SOUL_STONE
    SKILL_ICE_OF_SOUL_STONE
    SKILL_SAND_OF_SOUL_STONE
    SKILL_BLOCK_HEAD
    SKILL_KASAS_ARROW
    SKILL_HANDS_OF_FIRE
    SKILL_PROMINENCE
    SKILL_RING_OF_FLARE
    SKILL_BLAZE_BOLT
    SKILL_ICE_FIELD
    SKILL_WATER_BARRIER
    SKILL_HANDS_OF_NIZIE
    SKILL_NYMPH_RECOVERY
    SKILL_LIBERTY
    SKILL_TENDRIL
    SKILL_GNOMES_WHISPER
    SKILL_STONE_AUGER
    SKILL_REFUSAL_ETHER
    SKILL_EARTHS_TEETH
    SKILL_ABSORB_SOUL
    SKILL_SUMMON_SYLPH
    SKILL_DRIFTING_SOUL
    SKILL_CRITICAL_MAGIC

    SKILL_EMISSION_WATER
    SKILL_BEAT_HEAD

    SKILL_DIVINE_SPIRITS

    SKILL_BLITZ_SLIDING
    SKILL_BLAZE_WALK
    SKILL_JABBING_VEIN
    SKILL_GREAT_HEAL
    SKILL_DIVINE_GUIDANCE
    SKILL_BLOODY_ZENITH

    SKILL_REDIANCE
    SKILL_LAR_SLASH

    SKILL_HEART_CATALYST
    SKILL_ARMS_MASTERY_1
    SKILL_VIVID_MAGAZINE
    SKILL_TRIDENT
    SKILL_ARMS_MASTERY_2
    SKILL_MOLE_SHOT
    SKILL_ETERNITY
    SKILL_PROTECTION_FROM_BLOOD

    SKILL_INSTALL_TRAP
    SKILL_CREATE_HOLY_POTION
    SKILL_MERCY_GROUND
    SKILL_HOLY_ARMOR

    SKILL_TRANSFORM_TO_WERWOLF
    SKILL_STONE_SKIN
    SKILL_ACID_ERUPTION
    SKILL_TALON_OF_CROW
    SKILL_GRAY_DARKNESS
    SKILL_BITE_OF_DEATH

    SKILL_WIDE_GRAY_DARKNESS

    SKILL_TELEPORT

    SKILL_FIRE_PIERCING
    SKILL_SUMMON_FIRE_ELEMENTAL
    SKILL_MAGNUM_SPEAR
    SKILL_HELLFIRE

    SKILL_ICE_LANCE
    SKILL_FROZEN_ARMOR
    SKILL_SUMMON_WATER_ELEMENTAL
    SKILL_EXPLOSION_WATER
    SKILL_SOUL_REBIRTH
    SKILL_SOUL_REBIRTH_MASTERY

    SKILL_REACTIVE_ARMOR
    SKILL_GROUND_BLESS
    SKILL_SUMMON_GROUND_ELEMENTAL
    SKILL_METEOR_STORM

    SKILL_SHARP_CHAKRAM
    SKILL_SHIFT_BREAK
    SKILL_WATER_SHIELD
    SKILL_DESTRUCTION_SPEAR
    SKILL_BLESS_FIRE
    SKILL_FATAL_SNICK
    SKILL_SAND_CROSS
    SKILL_DUCKING_WALLOP
    SKILL_CHARGING_ATTACK
    SKILL_DISTANCE_BLITZ

    SKILL_FABULOUS_SOUL
    SKILL_WILL_OF_IRON

    // Áúµå·¹ ½ºÅ³
    SKILL_WIDE_ICE_FIELD
    SKILL_GLACIER_1
    SKILL_GLACIER_2
    SKILL_ICE_AUGER
    SKILL_ICE_HAIL
    SKILL_WIDE_ICE_HAIL
    SKILL_ICE_WAVE

    SKILL_LAND_MINE_EXPLOSION
    SKILL_CLAYMORE_EXPLOSION
    SKILL_PLEASURE_EXPLOSION

    SKILL_DELEO_EFFICIO // 317	// DELETE EFFECT
    SKILL_REPUTO_FACTUM

    SKILL_SWORD_OF_THOR
    SKILL_BURNING_SOL_CHARGING
    SKILL_BURNING_SOL_LAUNCH
    SKILL_INSTALL_TURRET
    SKILL_TURRET_FIRE
    SKILL_SWEEP_VICE_1
    SKILL_SWEEP_VICE_3
    SKILL_SWEEP_VICE_5
    SKILL_WHITSUNTIDE
    SKILL_VIOLENT_PHANTOM
    SKILL_SUMMON_GORE_GLAND
    SKILL_GORE_GLAND_FIRE
    SKILL_DESTRUCTION_SPEAR_MASTERY
    SKILL_FATAL_SNICK_MASTERY
    SKILL_MAGNUM_SPEAR_MASTERY
    SKILL_ICE_LANCE_MASTERY
    SKILL_REACTIVE_ARMOR_MASTERY

    SKILL_THROWING_AXE
    SKILL_CHOPPING_FIREWOOD  // 337 ÀåÀÛÆÐ±â
    SKILL_CHAIN_THROWING_AXE // 338 µµ³¢ ¼¼°³ ´øÁö±â
    SKILL_MULTI_THROWING_AXE // 339 µµ³¢  ""
    SKILL_PLAYING_WITH_FIRE  // 340 ºÒÀå³­

    SKILL_INFINITY_THUNDERBOLT
    SKILL_SPIT_STREAM
    SKILL_PLASMA_ROCKET_LAUNCHER
    SKILL_INTIMATE_GRAIL
    SKILL_BOMBING_STAR
    SKILL_SET_AFIRE
    SKILL_NOOSE_OF_WRAITH

    SKILL_SHARP_HAIL
    SKILL_SUMMON_MIGA        // 349	// ¾Æ¿ì½ºÅÍÁî°¡ ¾²´Â ½ºÅ³
    SKILL_SUMMON_MIGA_ATTACK // 350	// ¼ÒÈ¯µÈ³ðÀÌ ¾²´Â ½ºÅ³
    SKILL_ICE_HORIZON
    SKILL_FURY_OF_GNOME

    SKILL_CANNONADE        // 353	// Æ÷°Ý
    SKILL_SELF_DESTRUCTION // 354	// ÀÚÆø°ø°Ý

    SKILL_AR_ATTACK      // 355	// ¸ó½ºÅÍ¿ë
    SKILL_SMG_ATTACK     // 356	// ¸ó½ºÅÍ¿ë
    SKILL_GRENADE_ATTACK // 357	// ¸ó½ºÅÍ¿ë

    SKILL_DRAGON_TORNADO
    SKILL_BIKE_CRASH
    SKILL_HARPOON_BOMB
    SKILL_PASSING_HEAL
    SKILL_ROTTEN_APPLE
    SKILL_WILD_WOLF
    SKILL_ABERRATION
    SKILL_HALO
    SKILL_DESTINIES
    SKILL_FIERCE_FLAME
    SKILL_SHADOW_OF_STORM
    SKILL_HEAL_PASS // 369 // ¿Å°Ü°¡´Â Èú

    SKILL_TRASLA_ATTACK
    SKILL_PUSCA_ATTACK
    SKILL_NOD_COPILA_ATTACK
    SKILL_NOD_COPILA_ATTACK_2

    SKILL_UNTERFELDWEBEL_FIRE
    SKILL_FELDWEBEL_FIRE

    SKILL_MAX
)

func skill2int(skill string) int {
    for i, v := range SkillTypes2String {
        if v == skill {
            return i
        }
    }
    return SKILL_MAX
}

var SkillTypes2String = []string{
    "SKILL_ATTACK_MELEE",
    "SKILL_ATTACK_ARMS",
    "SKILL_SELF",
    "SKILL_TILE",
    "SKILL_OBJECT",
    "SKILL_DOUBLE_IMPACT",
    "SKILL_TRIPLE_SLASHER",
    "SKILL_RAINBOW_SLASHER",
    "SKILL_THUNDER_SPARK",
    "SKILL_DANCING_SWORD",
    "SKILL_CROSS_COUNTER",
    "SKILL_FLASH_SLIDING",
    "SKILL_LIGHTNING_HAND",
    "SKILL_SNAKE_COMBO",
    "SKILL_SWORD_WAVE",
    "SKILL_DRAGON_RISING",
    "SKILL_FIVE_STORM_CRASH",
    "SKILL_HEAVENS_SWORD",
    "SKILL_SINGLE_BLOW",
    "SKILL_SPIRAL_SLAY",
    "SKILL_TRIPLE_BREAK",
    "SKILL_WILD_SMASH",
    "SKILL_GHOST_BLADE",
    "SKILL_POTENTIAL_EXPLOSION",
    "SKILL_SHADOW_WALK",
    "SKILL_CHARGING_POWER",
    "SKILL_HURRICANE_COMBO",
    "SKILL_TORNADO_SEVER",
    "SKILL_ARMAGEDDON_SLASHER",
    "SKILL_SOUL_SHOCK",
    "SKILL_SAINT_BLADE",
    "SKILL_FAST_RELOAD",
    "SKILL_QUICK_FIRE",
    "SKILL_SMG_MASTERY",
    "SKILL_MULTI_SHOT",
    "SKILL_HEAD_SHOT",
    "SKILL_AR_MASTERY",
    "SKILL_PIERCING",
    "SKILL_SNIPING",
    "SKILL_SG_MASTERY",
    "SKILL_DETECT_MINE",
    "SKILL_CREATE_BOMB",
    "SKILL_SR_MASTERY",
    "SKILL_DISARM_MINE",
    "SKILL_INSTALL_MINE",
    "SKILL_CREATE_MINE",
    "SKILL_CREATE_HOLY_WATER",
    "SKILL_LIGHT",
    "SKILL_DETECT_HIDDEN",
    "SKILL_AURA_BALL",
    "SKILL_BLESS",
    "SKILL_CONTINUAL_LIGHT",
    "SKILL_FLARE",
    "SKILL_PURIFY",
    "SKILL_AURA_RING",
    "SKILL_STRIKING",
    "SKILL_DETECT_INVISIBILITY",
    "SKILL_IDENTIFY",
    "SKILL_AURA_SHIELD",
    "SKILL_ENCHANT",
    "SKILL_VISIBLE",
    "SKILL_CHAIN_AURA",
    "SKILL_SAINT_AURA",
    "SKILL_CURE_LIGHT_WOUNDS",
    "SKILL_CURE_POISON",
    "SKILL_PROTECTION_FROM_POISON",
    "SKILL_CAUSE_LIGHT_WOUNDS",
    "SKILL_CURE_SERIOUS_WOUNDS",
    "SKILL_REMOVE_CURSE",
    "SKILL_PROTECTION_FROM_CURSE",
    "SKILL_CAUSE_SERIOUS_WOUNDS",
    "SKILL_CURE_CRITICAL_WOUNDS",
    "SKILL_PROTECTION_FROM_ACID",
    "SKILL_SACRIFICE",
    "SKILL_CAUSE_CRITICAL_WOUNDS",
    "SKILL_CURE_ALL",
    "SKILL_REGENERATION",
    "SKILL_MASS_CURE",
    "SKILL_MASS_HEAL",
    "SKILL_BLOOD_DRAIN",
    "SKILL_POISONOUS_HANDS",
    "SKILL_GREEN_POISON",
    "SKILL_YELLOW_POISON",
    "SKILL_DARKBLUE_POISON",
    "SKILL_GREEN_STALKER",
    "SKILL_ACID_TOUCH",
    "SKILL_ACID_BOLT",
    "SKILL_ACID_BALL",
    "SKILL_ACID_SWAMP",
    "SKILL_PARALYZE",
    "SKILL_DOOM",
    "SKILL_HALLUCINATION",
    "SKILL_DEATH",
    "SKILL_BLOODY_NAIL",
    "SKILL_BLOODY_KNIFE",
    "SKILL_BLOODY_BALL",
    "SKILL_BLOODY_WALL",
    "SKILL_BLOODY_SPEAR",
    "SKILL_HIDE",
    "SKILL_DARKNESS",
    "SKILL_INVISIBILITY",
    "SKILL_TRANSFORM_TO_WOLF",
    "SKILL_TRANSFORM_TO_BAT",
    "SKILL_SUMMON_WOLF",
    "SKILL_SUMMON_CASKET",
    "SKILL_RAISING_DEAD",
    "SKILL_SUMMON_SERVANT",
    "SKILL_UN_BURROW",
    "SKILL_UN_TRANSFORM",
    "SKILL_UN_INVISIBILITY",
    "SKILL_THROW_HOLY_WATER",
    "SKILL_EAT_CORPSE",
    "SKILL_HOWL",
    "SKILL_RESTORE",
    "SKILL_BLOODY_MARKER",
    "SKILL_BLOODY_TUNNEL",
    "SKILL_SEDUCTION",
    "SKILL_WIND_DIVIDER",
    "SKILL_EARTHQUAKE",
    "SKILL_RESURRECT",
    "SKILL_PRAYER",
    "SKILL_MEDITATION",
    "SKILL_THUNDER_BOLT",
    "SKILL_EXPANSION",
    "SKILL_MIRACLE_SHIELD",
    "SKILL_THUNDER_FLASH",
    "SKILL_MASSACRE",
    "SKILL_INVINCIBLE",
    "SKILL_BERSERKER",
    "SKILL_MOONLIGHT_SEVER",
    "SKILL_SHADOW_DANCING",
    "SKILL_TYPHOON",
    "SKILL_PSYCHOKINESIS",
    "SKILL_EXTERMINATION",
    "SKILL_MIND_CONTROL",
    "SKILL_REVOLVING",
    "SKILL_FATALITY",
    "SKILL_BLITZ",
    "SKILL_ACTIVATION",
    "SKILL_PEACE",
    "SKILL_ENERGY_DROP",
    "SKILL_EXORCISM",
    "SKILL_SANCTUARY",
    "SKILL_REFLECTION",
    "SKILL_ARMAGEDDON",
    "SKILL_POISON_STRIKE",
    "SKILL_POISON_STORM",
    "SKILL_ACID_STRIKE",
    "SKILL_ACID_STORM",
    "SKILL_BLOODY_STRIKE",
    "SKILL_BLOODY_STORM",
    "SKILL_SUMMON_BAT",
    "SKILL_CHARM",
    "SKILL_POLYMORPH",
    "SKILL_MEPHISTO",
    "SKILL_HYPNOSIS",
    "SKILL_TRANSFUSION",
    "SKILL_EXTREME",
    "SKILL_BLOODY_WAVE",
    "SKILL_THROW_BOMB",
    "SKILL_DOUBLE_SHOT",
    "SKILL_TRIPLE_SHOT",
    "SKILL_CURE_EFFECT",
    "SKILL_CRITICAL_EFFECT",
    "SKILL_CRITICAL_GROUND",
    "SKILL_VIGOR_DROP",
    "SKILL_SWORD_MASTERY",
    "SKILL_SHIELD_MASTERY",
    "SKILL_THUNDER_STORM",
    "SKILL_CONCENTRATION",
    "SKILL_EVASION",
    "SKILL_HOLY_BLAST",
    "SKILL_HYMN",
    "SKILL_MENTAL_SWORD",
    "SKILL_OBSERVING_EYE",
    "SKILL_REFLECTION_EFFECT",
    "SKILL_TEMP",
    "SKILL_OPEN_CASKET",
    "SKILL_SUMMON_MONSTERS",
    "SKILL_GROUND_ATTACK",
    "SKILL_METEOR_STRIKE",
    "SKILL_DUPLICATE_SELF",
    "SKILL_BLOODY_MASTER_WAVE",
    "SKILL_BLOODY_WARP",
    "SKILL_BLOODY_SNAKE",
    "SKILL_SOUL_CHAIN",
    "SKILL_LIVENESS",
    "SKILL_DARKNESS_WIDE",
    "SKILL_POISON_STORM_WIDE",
    "SKILL_ACID_STORM_WIDE",
    "SKILL_SHARP_SHIELD",
    "SKILL_WIDE_LIGHTNING",
    "SKILL_AIR_SHIELD",
    "SKILL_POWER_OF_LAND",
    "SKILL_BULLET_OF_LIGHT",
    "SKILL_GUN_SHOT_GUIDANCE",
    "SKILL_REBUKE",
    "SKILL_SPIRIT_GUARD",
    "SKILL_TURN_UNDEAD",
    "SKILL_HANDS_OF_WISDOM",
    "SKILL_LIGHT_BALL",
    "SKILL_HOLY_ARROW",
    "SKILL_BLOODY_BREAKER",
    "SKILL_RAPID_GLIDING",
    "SKILL_MAGIC_ELUSION",
    "SKILL_POISON_MESH",
    "SKILL_ILLUSION_OF_AVENGE",
    "SKILL_WILL_OF_LIFE",
    "SKILL_DENIAL_MAGIC",
    "SKILL_REQUITAL",
    "SKILL_CONCEALMENT",
    "SKILL_SWORD_RAY",
    "SKILL_MULTI_AMPUTATE",
    "SKILL_NAIL_MASTERY",
    "SKILL_HIT_CONVERT",
    "SKILL_WILD_TYPHOON",
    "SKILL_ULTIMATE_BLOW",
    "SKILL_ILLENDUE",
    "SKILL_LIGHTNESS",
    "SKILL_FLOURISH",
    "SKILL_EVADE",
    "SKILL_SHARP_ROUND",
    "SKILL_HIDE_SIGHT",
    "SKILL_BACK_STAB",
    "SKILL_BLUNTING",
    "SKILL_GAMMA_CHOP",
    "SKILL_CROSS_GUARD",
    "SKILL_FIRE_OF_SOUL_STONE",
    "SKILL_ICE_OF_SOUL_STONE",
    "SKILL_SAND_OF_SOUL_STONE",
    "SKILL_BLOCK_HEAD",
    "SKILL_KASAS_ARROW",
    "SKILL_HANDS_OF_FIRE",
    "SKILL_PROMINENCE",
    "SKILL_RING_OF_FLARE",
    "SKILL_BLAZE_BOLT",
    "SKILL_ICE_FIELD",
    "SKILL_WATER_BARRIER",
    "SKILL_HANDS_OF_NIZIE",
    "SKILL_NYMPH_RECOVERY",
    "SKILL_LIBERTY",
    "SKILL_TENDRIL",
    "SKILL_GNOMES_WHISPER",
    "SKILL_STONE_AUGER",
    "SKILL_REFUSAL_ETHER",
    "SKILL_EARTHS_TEETH",
    "SKILL_ABSORB_SOUL",
    "SKILL_SUMMON_SYLPH",
    "SKILL_DRIFTING_SOUL",
    "SKILL_CRITICAL_MAGIC",
    "SKILL_EMISSION_WATER",
    "SKILL_BEAT_HEAD",
    "SKILL_DIVINE_SPIRITS",
    "SKILL_BLITZ_SLIDING",
    "SKILL_BLAZE_WALK",
    "SKILL_JABBING_VEIN",
    "SKILL_GREAT_HEAL",
    "SKILL_DIVINE_GUIDANCE",
    "SKILL_BLOODY_ZENITH",
    "SKILL_REDIANCE",
    "SKILL_LAR_SLASH",
    "SKILL_HEART_CATALYST",
    "SKILL_ARMS_MASTERY_1",
    "SKILL_VIVID_MAGAZINE",
    "SKILL_TRIDENT",
    "SKILL_ARMS_MASTERY_2",
    "SKILL_MOLE_SHOT",
    "SKILL_ETERNITY",
    "SKILL_PROTECTION_FROM_BLOOD",
    "SKILL_INSTALL_TRAP",
    "SKILL_CREATE_HOLY_POTION",
    "SKILL_MERCY_GROUND",
    "SKILL_HOLY_ARMOR",
    "SKILL_TRANSFORM_TO_WERWOLF",
    "SKILL_STONE_SKIN",
    "SKILL_ACID_ERUPTION",
    "SKILL_TALON_OF_CROW",
    "SKILL_GRAY_DARKNESS",
    "SKILL_BITE_OF_DEATH",
    "SKILL_WIDE_GRAY_DARKNESS",
    "SKILL_TELEPORT",
    "SKILL_FIRE_PIERCING",
    "SKILL_SUMMON_FIRE_ELEMENTAL",
    "SKILL_MAGNUM_SPEAR",
    "SKILL_HELLFIRE",
    "SKILL_ICE_LANCE",
    "SKILL_FROZEN_ARMOR",
    "SKILL_SUMMON_WATER_ELEMENTAL",
    "SKILL_EXPLOSION_WATER",
    "SKILL_SOUL_REBIRTH",
    "SKILL_SOUL_REBIRTH_MASTERY",
    "SKILL_REACTIVE_ARMOR",
    "SKILL_GROUND_BLESS",
    "SKILL_SUMMON_GROUND_ELEMENTAL",
    "SKILL_METEOR_STORM",
    "SKILL_SHARP_CHAKRAM",
    "SKILL_SHIFT_BREAK",
    "SKILL_WATER_SHIELD",
    "SKILL_DESTRUCTION_SPEAR",
    "SKILL_BLESS_FIRE",
    "SKILL_FATAL_SNICK",
    "SKILL_SAND_CROSS",
    "SKILL_DUCKING_WALLOP",
    "SKILL_CHARGING_ATTACK",
    "SKILL_DISTANCE_BLITZ",
    "SKILL_FABULOUS_SOUL",
    "SKILL_WILL_OF_IRON",
    "SKILL_WIDE_ICE_FIELD",
    "SKILL_GLACIER_1",
    "SKILL_GLACIER_2",
    "SKILL_ICE_AUGER",
    "SKILL_ICE_HAIL",
    "SKILL_WIDE_ICE_HAIL",
    "SKILL_ICE_WAVE",
    "SKILL_LAND_MINE_EXPLOSION",
    "SKILL_CLAYMORE_EXPLOSION",
    "SKILL_PLEASURE_EXPLOSION",
    "SKILL_DELEO_EFFICIO",
    "SKILL_REPUTO_FACTUM",
    "SKILL_SWORD_OF_THOR",
    "SKILL_BURNING_SOL_CHARGING",
    "SKILL_BURNING_SOL_LAUNCH",
    "SKILL_INSTALL_TURRET",
    "SKILL_TURRET_FIRE",
    "SKILL_SWEEP_VICE_1",
    "SKILL_SWEEP_VICE_3",
    "SKILL_SWEEP_VICE_5",
    "SKILL_WHITSUNTIDE",
    "SKILL_VIOLENT_PHANTOM",
    "SKILL_SUMMON_GORE_GLAND",
    "SKILL_GORE_GLAND_FIRE",
    "SKILL_DESTRUCTION_SPEAR_MASTERY",
    "SKILL_FATAL_SNICK_MASTERY",
    "SKILL_MAGNUM_SPEAR_MASTERY",
    "SKILL_ICE_LANCE_MASTERY",
    "SKILL_REACTIVE_ARMOR_MASTERY",
    "SKILL_THROWING_AXE",
    "SKILL_CHOPPING_FIREWOOD",
    "SKILL_CHAIN_THROWING_AXE",
    "SKILL_MULTI_THROWING_AXE",
    "SKILL_PLAYING_WITH_FIRE",
    "SKILL_INFINITY_THUNDERBOLT",
    "SKILL_SPIT_STREAM",
    "SKILL_PLASMA_ROCKET_LAUNCHER",
    "SKILL_INTIMATE_GRAIL",
    "SKILL_BOMBING_STAR",
    "SKILL_SET_AFIRE",
    "SKILL_NOOSE_OF_WRAITH",
    "SKILL_SHARP_HAIL",
    "SKILL_SUMMON_MIGA",
    "SKILL_SUMMON_MIGA_ATTACK",
    "SKILL_ICE_HORIZON",
    "SKILL_FURY_OF_GNOME",
    "SKILL_CANNONADE",
    "SKILL_SELF_DESTRUCTION",
    "SKILL_AR_ATTACK",
    "SKILL_SMG_ATTACK",
    "SKILL_GRENADE_ATTACK",
    "SKILL_DRAGON_TORNADO",
    "SKILL_BIKE_CRASH",
    "SKILL_HARPOON_BOMB",
    "SKILL_PASSING_HEAL",
    "SKILL_ROTTEN_APPLE",
    "SKILL_WILD_WOLF",
    "SKILL_ABERRATION",
    "SKILL_HALO",
    "SKILL_DESTINIES",
    "SKILL_FIERCE_FLAME",
    "SKILL_SHADOW_OF_STORM",
    "SKILL_HEAL_PASS",
    "SKILL_TRASLA_ATTACK",
    "SKILL_PUSCA_ATTACK",
    "SKILL_NOD_COPILA_ATTACK",
    "SKILL_NOD_COPILA_ATTACK_2",
    "SKILL_UNTERFELDWEBEL_FIRE",
    "SKILL_FELDWEBEL_FIRE",
    "SKILL_MAX",
}