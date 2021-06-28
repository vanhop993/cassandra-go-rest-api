package app

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	Keyspace = `masterdata`

	CreateKeyspace = `create keyspace if not exists masterdata with replication = {'class':'SimpleStrategy', 'replication_factor':1}`

	CreateTable = `
					create table if not exists users (
					id varchar,
					username varchar,
					email varchar,
					phone varchar,
					date_of_birth date,
					primary key (id)
	)`
)

func DBConnect() (*gocql.ClusterConfig, error) {
	cluster := gocql.NewCluster("127.0.0.1") //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	err = session.Query(CreateKeyspace).Exec()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// create table
	cluster.Keyspace = Keyspace
	session, err = cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	err = session.Query(CreateTable).Exec()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return cluster, nil
}
