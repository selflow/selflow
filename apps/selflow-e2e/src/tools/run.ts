import { execSync } from 'child_process';
import * as process from 'process';

export async function startRun(configFilePath: string): Promise<string> {
  const result = execSync(
    `go run github.com/selflow/selflow/apps/selflow-cli run \"${configFilePath}\"`,
    {
      env: {
        ...process.env,
        JSON_LOGS: 'TRUE',
      },
    }
  );

  return result.toString();
}
