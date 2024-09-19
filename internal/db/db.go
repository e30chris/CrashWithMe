package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type DB struct {
	conn *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	// Get environment variables
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create a new connector
	ctx := context.Background()
	d, err := cloudsqlconn.NewDialer(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %v", err)
	}

	// Use the Cloud SQL Go connector to create a connection to the database
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
		dbUser, dbPassword, dbName, d.Dial(instanceConnectionName))
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, fmt.Errorf("conn.Ping: %v", err)
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

func (db *DB) Insert(query string, args ...interface{}) (int64, error) {
	result, err := db.conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) CreateTables() error {
	createAirportsTable := `
	CREATE TABLE IF NOT EXISTS Airports (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		city VARCHAR(255) NOT NULL,
		country VARCHAR(255) NOT NULL,
		iata_code CHAR(3) NOT NULL UNIQUE,
		icao_code CHAR(4) NOT NULL UNIQUE,
		latitude DOUBLE PRECISION NOT NULL,
		longitude DOUBLE PRECISION NOT NULL,
		altitude INT NOT NULL,
		timezone VARCHAR(50) NOT NULL,
		dst CHAR(1) NOT NULL,
		tz_database_time_zone VARCHAR(50) NOT NULL
	);`

	createFlightsTable := `
	CREATE TABLE IF NOT EXISTS Flights (
		id SERIAL PRIMARY KEY,
		flight_number VARCHAR(10) NOT NULL,
		origin_airport_id INT NOT NULL,
		destination_airport_id INT NOT NULL,
		departure_time TIMESTAMP NOT NULL,
		arrival_time TIMESTAMP NOT NULL,
		FOREIGN KEY (origin_airport_id) REFERENCES Airports(id),
		FOREIGN KEY (destination_airport_id) REFERENCES Airports(id)
	);`

	createFlightCrewTable := `
	CREATE TABLE IF NOT EXISTS FlightCrew (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		user_name VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(20),
		whatsapp VARCHAR(20),
		work_email VARCHAR(255) NOT NULL UNIQUE,
		personal_email VARCHAR(255),
		facebook VARCHAR(255)
	);`

	createPropertyTable := `
	CREATE TABLE IF NOT EXISTS Property (
		id SERIAL PRIMARY KEY,
		property_name VARCHAR(255) NOT NULL,
		bed_size VARCHAR(50) NOT NULL,
		date_available DATE NOT NULL,
		price_per_night DECIMAL(10, 2) NOT NULL,
		distance_from_airport DOUBLE PRECISION NOT NULL
	);`

	createBookingsTable := `
	CREATE TABLE IF NOT EXISTS Bookings (
		id SERIAL PRIMARY KEY,
		flight_crew_id INT NOT NULL,
		property_id INT NOT NULL,
		booking_date DATE NOT NULL,
		FOREIGN KEY (flight_crew_id) REFERENCES FlightCrew(id),
		FOREIGN KEY (property_id) REFERENCES Property(id)
	);`

	_, err := db.conn.Exec(createAirportsTable)
	if err != nil {
		return fmt.Errorf("CreateTables: %v", err)
	}

	_, err = db.conn.Exec(createFlightsTable)
	if err != nil {
		return fmt.Errorf("CreateTables: %v", err)
	}

	_, err = db.conn.Exec(createFlightCrewTable)
	if err != nil {
		return fmt.Errorf("CreateTables: %v", err)
	}

	_, err = db.conn.Exec(createPropertyTable)
	if err != nil {
		return fmt.Errorf("CreateTables: %v", err)
	}

	_, err = db.conn.Exec(createBookingsTable)
	if err != nil {
		return fmt.Errorf("CreateTables: %v", err)
	}

	return nil
}

func (db *DB) Migrate() error {
	// TODO: Implement database migration logic here
	return nil
}

func (db *DB) AddProperty(property_name, bed_size string, date_available string, price_per_night float64, distance_from_airport float64) (int64, error) {
	query := `
	INSERT INTO Property (property_name, bed_size, date_available, price_per_night, distance_from_airport)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	var id int64
	err := db.conn.QueryRow(query, property_name, bed_size, date_available, price_per_night, distance_from_airport).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) AddBooking(flight_crew_id, property_id int, booking_date string) (int64, error) {
	query := `
	INSERT INTO Bookings (flight_crew_id, property_id, booking_date)
	VALUES ($1, $2, $3)
	RETURNING id;`

	var id int64
	err := db.conn.QueryRow(query, flight_crew_id, property_id, booking_date).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
