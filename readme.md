# Summary

Removing SQL AAD Admin could break running AAD based applications, cause an outage.

# Reproduce steps
1. create SQL azure with bicep
```bash
az deployment group create -g <your ResourceGroup> --template-file azure-sql.bicep -n sql1 --parameters location=eastus aadAdminObjectId=<SQL AAD Admin ObjectId> aadAdminLogin=<SQL AAD Admin Email>
```

2. grant SPN access to this Database (azure portal)
  - maybe tune firewall rule.
  - generate sid for given SPN clientID 
```sql
DECLARE @sid nvarchar(100)
set @sid=CONVERT(VARCHAR(1000), CAST(CAST('<SPN Client ID>' AS UNIQUEIDENTIFIER) AS varbinary(16)), 1)
select @sid
```
  - copy that 0x hex string, paste into the following command, and execute

```sql
CREATE USER aadspn WITH SID = 0x......, TYPE=E
exec sp_addrolemember db_owner, aadspn 
```

3. run program, it should allow the query.
```bash
Azure_Client_ID=<SPNClientID> Azure_Client_Secret=<SPNClientSecret> Azure_Client_Tenant=<SPNTenantID> go run main.go
```
4. Disable AAD admin

5. run program, now SQL Azure reject you 

```bash
Azure_Client_ID=<SPNClientID> Azure_Client_Secret=<SPNClientSecret> Azure_Client_Tenant=<SPNTenantID> go run main.go
```

panic: mssql: login error: Login failed for user '<token-identified principal>'. The server is not currently configured to accept this token.