import fs from "fs";
import path from "path";
import { getNicknameList } from "./getNicknameList";
import { toHebon } from "jaconv";

(async () => {
  const distPath = path.join(__dirname, "../data.json");
  const japaneseList = await getNicknameList();
  const arr = japaneseList.map((j) => [j, toHebon(j).toLowerCase()]).sort();
  const json = JSON.stringify(arr, null, "\t");
  fs.writeFileSync(distPath, json);
  console.log(`done!🎉\n${distPath}`);
})();
