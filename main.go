package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// Import the Azure AD driver module (also imports the regular driver package)
	"github.com/microsoft/go-mssqldb/azuread"
)

func main() {
	db, err := ConnectWithMSI()
	if err != nil {
		panic(err)
	}

	_, err = db.QueryContext(context.Background(), "SELECT * FROM sys.tables")
	if err != nil {
		panic(err)
	}

	fmt.Println("query succeeded")
}

func ConnectWithMSI() (*sql.DB, error) {
	clientId, ok := os.LookupEnv("Azure_Client_ID")
	if !ok {
		return nil, fmt.Errorf("env Azure_Client_ID not set")
	}
	password, ok := os.LookupEnv("Azure_Client_Secret")
	if !ok {
		return nil, fmt.Errorf("env Azure_Client_Secret not set")
	}
	tenant, ok := os.LookupEnv("Azure_Client_Tenant")
	if !ok {
		return nil, fmt.Errorf("env Azure_Client_Tenant not set")
	}
	sqlServer, ok := os.LookupEnv("SQL_Server")
	if !ok {
		return nil, fmt.Errorf("env SQL_Server not set")
	}
	sqlDatabase, ok := os.LookupEnv("SQL_Database")
	if !ok {
		return nil, fmt.Errorf("env SQL_Database not set")
	}

	connectionString := fmt.Sprintf("sqlserver://%s?database=%s&fedauth=ActiveDirectoryServicePrincipal&user id=%s@%s&password=%s", sqlServer, sqlDatabase, clientId, tenant, password)

	return sql.Open(azuread.DriverName, connectionString)
}
