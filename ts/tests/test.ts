import path from "path"
import { runSqldef } from "../src/lib"
import { expect } from "chai"

describe("runSqldef", function() {
  it("should run the schema tool successfully", function() {
    const dbType = "postgres"
    const dbName = process.env.DB_NAME || ""
    const user = process.env.DB_USER || ""
    const password = process.env.DB_PASSWORD || ""
    const host = process.env.DB_HOST || ""
    const port = process.env.DB_PORT || "5432"
    const schemaFile = path.resolve(__dirname, process.env.SCHEMA_FILE_PATH || "")
    const enableDropTable = process.env.ENABLE_DROP_TABLE ? parseInt(process.env.ENABLE_DROP_TABLE, 10) : 0

    const result = runSqldef(dbType, dbName, user, password, host, port, schemaFile, enableDropTable)
    expect(result).to.equal(0)
  })
})
