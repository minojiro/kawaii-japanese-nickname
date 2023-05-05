import fs from "fs";
import path from "path";
import { parseData } from "./parseData";

(async () => {
  const distPath = path.join(__dirname, "../data.json");
  const arr = await parseData();
  const json = JSON.stringify(arr, null, "\t");
  fs.writeFileSync(distPath, json);
  console.log(`done!🎉\n${distPath}`);
})();
