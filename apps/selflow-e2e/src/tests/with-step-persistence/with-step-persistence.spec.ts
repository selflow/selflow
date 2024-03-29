import { startCliRun, startRun } from '../../tools/run';
import { describe, expect, test } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with step persistence', function () {
  describe('Step B should access the file created by Step a', function () {
    const configFilePath = join(__dirname, 'with-step-persistence.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-b', 'Hello!']);
    };

    test('selflow-daemon', async () => {
      const logs = await startRun(configFilePath);
      verifyLogs(logs);
    });

    test('selflow-cli', async () => {
      const logs = await startCliRun(configFilePath);
      verifyLogs(logs);
    });
  });
});
