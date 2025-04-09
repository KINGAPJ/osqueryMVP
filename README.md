# osqueryMVP
### 🛰️ Fetch Latest Data from API

You can query the latest system and installed apps data using the following `curl` command:

```bash
curl -v http://localhost:8080/latest_data -o output.txt
```

### What this does
Run migrations, and then run `go run .`
The code will generate the latest snapshot using `osqueryi` and upsert it to the mysql db.
The fields of the snapshot (InstalledProgram and SystemInfo) can be found in models/models.go
The endpoint data can be hit using GET /latest_data as shown above.

### Features which can be added
1. Periodic scans using osqueryd and storing the differentials.
2. Central logging server and not just logs in CLI.
3. Adding support for getting clientIDs to any number of machines connected to this server.