const { execSync } = require("child_process");
const os = require("os");
const path = require("path");
const fs = require("fs");

const platform = os.platform();
const goDir = path.join(__dirname, "go");
const outputDir = path.join(__dirname, "./");
let outputFile;

if (platform === "darwin") {
  outputFile = path.join(outputDir, "libsqldef.dylib");
} else if (platform === "linux") {
  outputFile = path.join(outputDir, "libsqldef.so");
} else {
  console.error("Unsupported platform:", platform);
  process.exit(1);
}

if (!fs.existsSync(goDir)) {
  console.error("Go directory not found:", goDir);
  process.exit(1);
}

try {
  execSync(`cd ${goDir} && go build -o ${outputFile} -buildmode=c-shared .`, { stdio: "inherit" });
  console.log(`Successfully built ${outputFile}`);
} catch (error) {
  console.error("Failed to build the Go library:", error);
  process.exit(1);
}
