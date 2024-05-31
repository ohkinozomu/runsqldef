import { DataType, open, close, define } from "ffi-rs"
import path from "path"
import os from "os"

export function runSqldef(dbType: string, dbName: string, user: string, password: string, host: string, port: string, schemaFilePath: string, enableDropTable: number): number {
  let libPath: string
  if (os.platform() === "darwin") {
    libPath = path.resolve(__dirname, "../../libsqldef.dylib")
  } else if (os.platform() === "linux") {
    libPath = path.resolve(__dirname, "../../libsqldef.so")
  } else {
    throw new Error(`Unsupported platform: ${os.platform()}`)
  }

  open({
    library: "libsqldef",
    path: libPath
  })
  const libsqldef = define({
    Run: {
      library: "libsqldef",
      retType: DataType.I32,
      paramsType: [DataType.String, DataType.String, DataType.String, DataType.String, DataType.String, DataType.String, DataType.String, DataType.I32]
    }
  })
  const result = libsqldef.Run([dbType, dbName, user, password, host, port, schemaFilePath, enableDropTable])
  close("libsqldef")
  return result
}
