package crowdedness

import "time"

type Response struct {
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

type Responses []Response
