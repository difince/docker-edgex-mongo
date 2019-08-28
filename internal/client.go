package internal

import (
	"fmt"

	"github.com/globalsign/mgo"

	"github.com/edgexfoundry/docker-edgex-mongo/internal/pkg"
)

var databaseCollectionsMap = map[string]func(db *mgo.Database){
	//TODO Diana: Do I need empty database ?!
	"authorization":   nil,
	"admin":           creatAdminCollections,
	"metadata":        createMetadataCollections,
	"coredata":        createCoredataCollections,
	"rules_engine_db": nil,
	"notifications":   createNotificationCollections,
	"scheduler":       createSchedulerCollections,
	"logging":         createLoggingCollections,
	"exportclient":    createExpertClientCollections,
}

type DBInitClient struct {
	Configuration *pkg.Configuration
}

func (client *DBInitClient) PopulateDatabase() (err error) {
	session, err := pkg.GetSession(client.Configuration)

	if err != nil {
		return
	}

	defer session.Close()

	for dbName, createCollectionsFunc := range databaseCollectionsMap {
		client.createDatabase(session, dbName, createCollectionsFunc)
	}
	return
}

func (client *DBInitClient) createDatabase(session *mgo.Session, dbName string, createCollectionsFunc func(db *mgo.Database)) {
	pkg.LoggingClient.Debug(fmt.Sprintf("Settting up %v database", dbName))
	db := mgo.Database{
		Session: session,
		Name:    dbName,
	}

	err := db.UpsertUser(&mgo.User{
		Username: client.Configuration.Credentials[dbName].Username,
		Password: client.Configuration.Credentials[dbName].Password,
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})

	if err != nil {
		pkg.LoggingClient.Error("Error during execution: " + err.Error())
	}
	if createCollectionsFunc != nil {
		createCollectionsFunc(&db)
	}
}
