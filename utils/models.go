package utils

type Request struct {
	Ev  string `json:"ev"`
	Et  string `json:"et"`
	ID  string `json:"id"`
	UID string `json:"uid"`
	MID string `json:"mid"`
	T   string `json:"t"`
	P   string `json:"p"`
	L   string `json:"l"`
	SC  string `json:"sc"`
}

type Responce struct {
	Ev    string                       `json:"event"`
	Et    string                       `json:"event_type"`
	ID    string                       `json:"app_id"`
	UID   string                       `json:"user_id"`
	MID   string                       `json:"message_id"`
	T     string                       `json:"page_title"`
	P     string                       `json:"page_url"`
	L     string                       `json:"browser_language"`
	SC    string                       `json:"screen_size"`
	Atks  map[string]map[string]string `json:"attributes"`
	Uatks map[string]map[string]string `json:"traits"`
}
