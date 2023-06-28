package bikes

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timdeklijn/druktezoeker/internal/crowdedness"
	_ "github.com/timdeklijn/druktezoeker/internal/log"
)

// BikeItem is a struct used to store aggregate fietsplaatsen in before writing
// to the database.
type BikeItem struct {
	Treinnummer   int    `db:"treinnummer"`
	Station       string `db:"station"`
	Fietsplaatsen int    `db:"fietsplaatsen"`
	Date          string `db:"date"`
	LastUpdated   string `db:"last_updated"`
}

// CollectBikes goes over the crowdedness response and sums all fietsplaatsen.
func CollectBikes(responses crowdedness.Response) (*[]BikeItem, error) {
	lastUpdated := time.Now().Format("2006-01-02 15:04:05")
	var bikeItems []BikeItem
	for _, response := range responses[0].DrukteBerichten {
		bike := BikeItem{
			Treinnummer:   response.Treinnummer,
			Station:       response.StartStationUic,
			Fietsplaatsen: response.Fietsplaatsen,
			Date:          response.VerkeersdatumAms,
			LastUpdated:   lastUpdated,
		}
		bikeItems = append(bikeItems, bike)

	}
	return &bikeItems, nil
}

// createTable will create the 'fietsplaatsen' table if it does not exist.
func createTable(db *sql.DB) error {
	const createTable string = `
	CREATE TABLE IF NOT EXISTS fietsplaatsen (
		id            INTEGER  NOT NULL PRIMARY KEY,
		treinnummer   INTEGER  NOT NULL,
		station       TEXT     NOT NULL,
		fietsplaatsen INTEGER  NOT NULL,
		date          DATETIME NOT NULL,
		last_updated  DATETIME NOT NULL
	);`

	if _, err := db.Exec(createTable); err != nil {
		return err
	}

	return nil
}

// insertIntoTable will write a query result into the 'fietsplaatsen' table.
func insertIntoTable(db *sql.DB, bikes *[]BikeItem) error {
	const insertIntoTable string = `
	INSERT INTO fietsplaatsen (
		treinnummer,
		station,
		fietsplaatsen,
		date,
		last_updated
	) VALUES (?, ?, ?, ?, ?);`

	for _, b := range *bikes {
		if _, err := db.Exec(insertIntoTable, b.Treinnummer, b.Station, b.Fietsplaatsen, b.Date, b.LastUpdated); err != nil {
			return err
		}
	}

	return nil
}

// WriteBikesToDB creates a database connection and write results to the database.
func WriteBikesToDB(bikes *[]BikeItem) error {
	db, err := sql.Open("sqlite3", "druktezoeker.db")
	if err != nil {
		return err
	}
	if err := createTable(db); err != nil {
		return err
	}

	if err := insertIntoTable(db, bikes); err != nil {
		return err
	}

	return nil
}
