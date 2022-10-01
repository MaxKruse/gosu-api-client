package gosuapiclient

import "time"

type BeatmapObject struct {
	BeatmapsetID     int         `json:"beatmapset_id"`
	DifficultyRating float64     `json:"difficulty_rating"`
	ID               int         `json:"id"`
	Mode             string      `json:"mode"`
	Status           string      `json:"status"`
	TotalLength      int         `json:"total_length"`
	UserID           int         `json:"user_id"`
	Version          string      `json:"version"`
	Accuracy         float64     `json:"accuracy"`
	Ar               float64     `json:"ar"`
	Bpm              int         `json:"bpm"`
	Convert          bool        `json:"convert"`
	CountCircles     int         `json:"count_circles"`
	CountSliders     int         `json:"count_sliders"`
	CountSpinners    int         `json:"count_spinners"`
	Cs               float64     `json:"cs"`
	DeletedAt        interface{} `json:"deleted_at"`
	Drain            float64     `json:"drain"`
	HitLength        int         `json:"hit_length"`
	IsScoreable      bool        `json:"is_scoreable"`
	LastUpdated      time.Time   `json:"last_updated"`
	ModeInt          int         `json:"mode_int"`
	Passcount        int         `json:"passcount"`
	Playcount        int         `json:"playcount"`
	Ranked           int         `json:"ranked"`
	URL              string      `json:"url"`
	Checksum         string      `json:"checksum"`
	Beatmapset       Beatmapset  `json:"beatmapset"`
	Failtimes        Failtimes   `json:"failtimes"`
	MaxCombo         int         `json:"max_combo"`
}
type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}
type Availability struct {
	DownloadDisabled int         `json:"download_disabled"`
	MoreInformation  interface{} `json:"more_information"`
}
type NominationsSummary struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}
type Beatmapset struct {
	Artist             string             `json:"artist"`
	ArtistUnicode      string             `json:"artist_unicode"`
	Covers             Covers             `json:"covers"`
	Creator            string             `json:"creator"`
	FavouriteCount     int                `json:"favourite_count"`
	Hype               interface{}        `json:"hype"`
	ID                 int                `json:"id"`
	Nsfw               bool               `json:"nsfw"`
	Offset             int                `json:"offset"`
	PlayCount          int                `json:"play_count"`
	PreviewURL         string             `json:"preview_url"`
	Source             string             `json:"source"`
	Spotlight          bool               `json:"spotlight"`
	Status             string             `json:"status"`
	Title              string             `json:"title"`
	TitleUnicode       string             `json:"title_unicode"`
	TrackID            int                `json:"track_id"`
	UserID             int                `json:"user_id"`
	Video              bool               `json:"video"`
	Availability       Availability       `json:"availability"`
	Bpm                int                `json:"bpm"`
	CanBeHyped         bool               `json:"can_be_hyped"`
	DiscussionEnabled  bool               `json:"discussion_enabled"`
	DiscussionLocked   bool               `json:"discussion_locked"`
	IsScoreable        bool               `json:"is_scoreable"`
	LastUpdated        time.Time          `json:"last_updated"`
	LegacyThreadURL    string             `json:"legacy_thread_url"`
	NominationsSummary NominationsSummary `json:"nominations_summary"`
	Ranked             int                `json:"ranked"`
	RankedDate         time.Time          `json:"ranked_date"`
	Storyboard         bool               `json:"storyboard"`
	SubmittedDate      time.Time          `json:"submitted_date"`
	Tags               string             `json:"tags"`
	Ratings            []int              `json:"ratings"`
}
type Failtimes struct {
	Fail []int `json:"fail"`
	Exit []int `json:"exit"`
}

type DifficultyAttribute struct {
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	StarRating           float64 `json:"star_rating"`
	MaxCombo             int     `json:"max_combo"`
	AimDifficulty        float64 `json:"aim_difficulty"`
	SpeedDifficulty      float64 `json:"speed_difficulty"`
	SpeedNoteCount       float64 `json:"speed_note_count"`
	FlashlightDifficulty int     `json:"flashlight_difficulty"`
	SliderFactor         float64 `json:"slider_factor"`
	ApproachRate         float64 `json:"approach_rate"`
	OverallDifficulty    float64 `json:"overall_difficulty"`
}

const (
	None              = 0
	NoFail            = 1
	Easy              = 2
	TouchDevice       = 4
	Hidden            = 8
	HardRock          = 16
	SuddenDeath       = 32
	DoubleTime        = 64
	Relax             = 128
	HalfTime          = 256
	Nightcore         = 512 // Only set along with DoubleTime. i.e: NC only gives 576
	Flashlight        = 1024
	Autoplay          = 2048
	SpunOut           = 4096
	Relax2            = 8192  // Autopilot
	Perfect           = 16384 // Only set along with SuddenDeath. i.e: PF only gives 16416
	Key4              = 32768
	Key5              = 65536
	Key6              = 131072
	Key7              = 262144
	Key8              = 524288
	FadeIn            = 1048576
	Random            = 2097152
	Cinema            = 4194304
	Target            = 8388608
	Key9              = 16777216
	KeyCoop           = 33554432
	Key1              = 67108864
	Key3              = 134217728
	Key2              = 268435456
	ScoreV2           = 536870912
	Mirror            = 1073741824
	KeyMod            = Key1 | Key2 | Key3 | Key4 | Key5 | Key6 | Key7 | Key8 | Key9 | KeyCoop
	FreeModAllowed    = NoFail | Easy | Hidden | HardRock | SuddenDeath | Flashlight | FadeIn | Relax | Relax2 | SpunOut | KeyMod
	ScoreIncreaseMods = Hidden | HardRock | DoubleTime | Flashlight | FadeIn
)
