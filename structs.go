package csgsi

type State struct {
	Provider   *provider
	Map        *csmap
	Round      *round
	Player     *Player
	AllPlayers map[string]*Player // allplayers_*: steamid64 ...
	Previously *State
	Added      *State
	Auth       *auth
}

// provider
type provider struct {
	Name      string
	AppId     int
	Version   int
	SteamId   string
	Timestamp float32
}

// map
type csmap struct {
	Name    string
	Phase   string
	Round   int
	Team_ct *team
	Team_t  *team
}

// round
type round struct {
	Phase    string
	Win_team string
	Bomb     string
}

// player_id
type Player struct {
	SteamId     string
	Name        string
	Team        string
	Activity    string
	State       *PlayerState
	Weapons     map[string]*Weapon
	Match_stats *PlayerMatchStats
}

// win_team
type team struct {
	Score int
}

// player_state
type PlayerState struct {
	Health       int
	Armor        int
	Helmet       bool
	Flashed      int
	Smoked       int
	Burning      int
	Money        int
	Round_kills  int
	Round_killhs int
}

// player_weapons: weapon_0, weapon_1, weapon_2 ...
type Weapon struct {
	Name          string
	PaintKit      string
	Type          string
	State         string
	Ammo_clip     int
	Ammo_clip_max int
	Ammo_reserve  int
}

// player_match_stats
type PlayerMatchStats struct {
	Kills   int
	Assists int
	Deaths  int
	Mvps    int
	Score   int
}

type auth struct {
	Token string
}
