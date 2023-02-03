param location string = resourceGroup().location
param serverNamePrefix string = 'sql'
param sqlDBName string = 'testDB'
param administratorLogin string = 'mysqladmin'
param aadAdminObjectId string = ''
param aadAdminLogin string = ''

resource sqlServer 'Microsoft.Sql/servers@2022-05-01-preview' = {
  name: uniqueString(serverNamePrefix, resourceGroup().id, location)
  location: location
  properties: {
    administratorLogin: administratorLogin
    administratorLoginPassword: guid(subscription().subscriptionId, resourceGroup().id, serverNamePrefix, location)
    administrators: {
      administratorType: 'ActiveDirectory'
      login: aadAdminLogin
      sid: aadAdminObjectId
      tenantId: tenant().tenantId
    }
  }

  resource testDB 'databases' = {
    name: sqlDBName
    location: location
    sku: {
      name: 'GP_Gen5'
      capacity: 2
    }
    properties: {
      collation: 'SQL_Latin1_General_CP1_CI_AS'
      createMode: 'Default'
      maxSizeBytes: 268435456000
      zoneRedundant: false
    }
  }
}
