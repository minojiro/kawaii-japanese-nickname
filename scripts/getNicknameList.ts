import fs from "fs";
import path from "path";

export const getNicknameList = (): string[] => {
  const csvPath = path.join(__dirname, "../nicknames.txt");
  const csvBuffer = fs.readFileSync(csvPath);
  return String(csvBuffer).split("\n").filter(Boolean);
};
