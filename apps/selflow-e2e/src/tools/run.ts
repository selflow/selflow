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

export async function startCliRun(configFilePath: string): Promise<string> {
  try {
    const result = execSync(
      `go run github.com/selflow/selflow/apps/selflow-cli exec \"${configFilePath}\"`,
      {
        env: {
          ...process.env,
          JSON_LOGS: 'TRUE',
          LOG_LEVEL: 'DEBUG',
          SELFLOW_DAEMON_BASE_DIRECTORY: `${process.env.PWD}`,
        },
      }
    );

    return result.toString();
  } catch (e) {
    console.log(e.stdout.toString());
    console.log(e.stderr.toString());
    throw e;
  }
}
