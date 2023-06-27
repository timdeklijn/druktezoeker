package bikes

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/timdeklijn/druktezoeker/internal/crowdedness"
	_ "github.com/timdeklijn/druktezoeker/internal/log"
)

// AggregatedBike is a struct used to store aggregate fietsplaatsen in before writing
// to the database.
type AggregatedBike struct {
	Treinnummer   int    `db:"treinnummer"`
	Station       string `db:"station"`
	Fietsplaatsen int    `db:"fietsplaatsen"`
	Date          string `db:"date"`
	LastUpdated   string `db:"last_updated"`
}

// AggregateBikes goes over the crowdedness response and sums all fietsplaatsen.
func AggregateBikes(responses crowdedness.Responses) (*AggregatedBike, error) {
	var aggregatedBike AggregatedBike
	for i, response := range responses {
		if i == 0 {
			aggregatedBike.Station = response.StartStationUic
			aggregatedBike.Treinnummer = response.Treinnummer
			aggregatedBike.Fietsplaatsen = response.Fietsplaatsen
			aggregatedBike.Date = response.VerkeersdatumAms
			continue
		}
		aggregatedBike.Fietsplaatsen += response.Fietsplaatsen
	}
	return &aggregatedBike, nil
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
func insertIntoTable(db *sql.DB, bikes *AggregatedBike) error {
	lastUpdated := time.Now().Format("2006-01-02 15:04:05")
	const insertIntoTable string = `
	INSERT INTO fietsplaatsen (
		treinnummer,
		station,
		fietsplaatsen,
		date,
		last_updated
	) VALUES (?, ?, ?, ?, ?);`

	if _, err := db.Exec(insertIntoTable, bikes.Treinnummer, bikes.Station, bikes.Fietsplaatsen, bikes.Date, lastUpdated); err != nil {
		return err
	}

	return nil
}

// WriteBikes creates a database connection and write results to the database.
func WriteBikes(bikes *AggregatedBike) error {
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
