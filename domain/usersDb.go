package domain

import (
	"log"

	"github.com/gocql/gocql"
)

type Cassandra struct {
	Cluster *gocql.ClusterConfig
}

func NewCassandraDb(db *gocql.ClusterConfig) Cassandra {
	return Cassandra{Cluster: db}
}

func (c Cassandra) GetAllDb() ([]UserStruct, error) {
	session, err := c.Cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	query := "select id, username, email, phone, date_of_birth from users"
	rows := session.Query(query).Iter()
	var result []UserStruct
	var user UserStruct
	for rows.Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth) {
		result = append(result, user)
	}
	defer session.Close()
	return result, nil
}

func (c Cassandra) InserDb(user *UserStruct) (string, error) {
	session, err := c.Cluster.CreateSession()
	if err != nil {
		return "", err
	}
	query := `INSERT INTO users (id, username, email, phone, date_of_birth) VALUES (?, ?, ?, ?, ? ) ;`
	er1 := session.Query(query, user.Id, user.Username, user.Email, user.Phone, user.DateOfBirth).Exec()
	if er1 != nil {
		return "", er1
	}
	defer session.Close()
	resultString := "Create success"
	return resultString, nil
}

func (c Cassandra) GetByIdDB(id string) (*UserStruct, error) {
	session, err := c.Cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	query := "select id, username, email, phone, date_of_birth from users where id = ?"
	var user UserStruct
	er1 := session.Query(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
	if er1 != nil {
		return nil, err
	}
	defer session.Close()
	return &user, nil
}

func (c Cassandra) UpdateDb(user *UserStruct) (string, error) {
	session, er0 := c.Cluster.CreateSession()
	log.Println(er0)
	if er0 != nil {
		return "", er0
	}
	query := "update users set username = ?, email = ?, phone = ?, date_of_birth = ? where id = ?"
	er1 := session.Query(query, user.Username, user.Email, user.Phone, user.DateOfBirth, user.Id).Exec()
	if er1 != nil {
		return "", er1
	}
	resultString := "Update data success"
	defer session.Close()
	return resultString, nil
}

func (c Cassandra) DeleteUserDB(id string) (string, error) {
	session, err := c.Cluster.CreateSession()
	if err != nil {
		return "", err
	}
	query := "delete from users where id = ?"
	er1 := session.Query(query, id).Exec()
	if er1 != nil {
		return "", er1
	}
	resultString := "Delete success"
	return resultString, nil
}
