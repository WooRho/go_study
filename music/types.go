package music

type MusicList []*Music
type Music struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Artist  []string `json:"artist"`
	Album   string   `json:"album"`
	PicId   string   `json:"pic_id"`
	UrlId   int      `json:"url_id"`
	LyricId int      `json:"lyric_id"`
	Source  string   `json:"source"`
}

type MusicUrl struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
	Br   int    `json:"br"`
}
