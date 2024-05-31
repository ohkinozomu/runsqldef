# runsqldef

Embedding [sqldef](https://github.com/sqldef/sqldef) into a program using FFI

## Node.js

```bash
npm install runsqldef
```

```javascript
import path from "path";
import { fileURLToPath } from 'url';
import { dirname } from 'path';
import { runSqldef } from "runsqldef";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const dbType = "postgres";
const dbName = "your_db_name";
const user = "your_user";
const password = "your_password";
const host = "your_host";
const port = "5432";
const schemaFile = path.resolve(__dirname, "path/to/your/schema.sql");
const enableDropTable = 0;

const result = runSqldef(dbType, dbName, user, password, host, port, schemaFile, enableDropTable);
if (result !== 0) {
  console.error("Error running schema tool");
} else {
  console.log("Schema tool ran successfully");
}
```