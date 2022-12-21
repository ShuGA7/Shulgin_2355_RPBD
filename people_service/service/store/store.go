package store

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

// NewStore creates new database connection

func (s *Store) ListPeople() ([]People, error) {
	rows, err := s.conn.Query(context.Background(), "SELECT id, name FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var peopleArray []People

	for rows.Next() {
		var man People

		if err = rows.Scan(&man.ID, &man.Name); err != nil {
			log.Fatal(err)
		}
		peopleArray = append(peopleArray, man)
	}
	if err = rows.Err(); err != nil {
		log.Fatal()
	}

	return peopleArray, err
}

func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// make migration

	// Read migrations from /home/mattes/migrations and connect to a local postgres database.
	m, err := migrate.New("file://migrations",
		connString)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("migration is done")

	return &Store{
		conn: conn,
	}
}

func (s *Store) GetPeopleByID(id string) (People, error) {
	var name string
	err := s.conn.QueryRow(context.Background(),
		"SELECT name FROM people where id=$1", id).Scan(&name)
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return People{Name: name, ID: i}, err
}

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int
	Name string
}
