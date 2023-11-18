import { describe, test } from 'vitest';
import { startCliRun, startRun } from './run';
import { parseLogs } from './logParser';
import { WorkflowExecutionTrace } from './trace';

export type SelflowRunner = (configFilePath: string) => Promise<string>;

const selflowRunnerMap = {
  'selflow-cli': startCliRun,
  'selflow-daemon': startRun,
} as const;

type SelflowRunnerName = keyof typeof selflowRunnerMap;

export function withSelflowRunners(
  runners: SelflowRunnerName[],
  testSuite: string,
  configfilePath: string,
  asserts: (trace: WorkflowExecutionTrace, logs: string) => void
) {
  describe.concurrent(testSuite, () => {
    for (let runner of runners) {
      test(runner, async (testContext) => {
        const logs = await selflowRunnerMap[runner](configfilePath);
        const trace = parseLogs(logs);
        asserts(trace, logs);
      });
    }
  });
}
