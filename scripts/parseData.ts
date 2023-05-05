import { parse } from "csv-parse";
import fs from "fs";
import path from "path";

const csvToArray = async (csvStr: string): Promise<string[][]> => {
  return new Promise((resolve, reject) => {
    const records: string[][] = [];
    const parser = parse();
    parser.on("readable", () => {
      let record: string[];
      while ((record = parser.read()) !== null) {
        records.push(record);
      }
      resolve(records);
    });
    parser.on("error", (err) => reject(err));
    parser.write(String(csvStr));
  });
};

export const parseData = async () => {
  const csvPath = path.join(__dirname, "../source.csv");
  const csvBuffer = fs.readFileSync(csvPath);
  return await csvToArray(String(csvBuffer));
};
