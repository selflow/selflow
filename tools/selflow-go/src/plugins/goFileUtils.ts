import { readFile } from 'node:fs/promises';
import { PACKAGE_REGEX } from './constants';

import { GoProject } from './types';

export function groupGoFileDirectories(files: string[]): Map<string, string[]> {
  return files.reduce((acc, file) => {
    const dir = file.split('/').slice(0, -1).join('/');
    if (!acc.has(dir)) {
      acc.set(dir, []);
    }
    acc.get(dir)?.push(file);
    return acc;
  }, new Map<string, string[]>());
}

export async function parseGoFiles(
  dirs: Map<string, string[]>
): Promise<GoProject[]> {
  return Promise.all(
    [...dirs.entries()].map(async ([dir, files]) => {
      files.sort();
      const firstFile = files[0];
      const fileContent = await readFile(firstFile);
      const packageName = PACKAGE_REGEX.exec(fileContent.toString())?.[1];
      return { packageName, dir, sourceFile: firstFile };
    })
  );
}
