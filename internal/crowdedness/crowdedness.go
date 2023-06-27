package crowdedness

import "time"

type Crowdedness struct {
	Vervoerder                  string    `json:"vervoerder"`
	Treinnummer                 int       `json:"treinnummer"`
	StartStationUic             string    `json:"start_station_uic"`
	EindStationUic              string    `json:"eind_station_uic"`
	LogistiekMaterieeltypes     []string  `json:"logistiek_materieeltypes"`
	Vervoerstrajectpositie      int       `json:"vervoerstrajectpositie"`
	DruktePerTreinClassificatie int       `json:"drukte_per_trein_classificatie"`
	Fietsplaatsen               int       `json:"fietsplaatsen"`
	Zitplaatsen                 int       `json:"zitplaatsen"`
	VerkeersdatumAms            string    `json:"verkeersdatum_ams"`
	LaadmomentUtc               time.Time `json:"laadmoment_utc"`
	LaaddatumUtc                string    `json:"laaddatum_utc"`
}

type TrainResponse struct {
	Treinnummer     int           `json:"treinnummer"`
	DrukteBerichten []Crowdedness `json:"drukte_berichten"`
}

type Response []TrainResponse
