package main

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
