import {startCliRun, startRun} from '../../tools/run';
import {expect, test} from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with dependencies', function () {
  describe('should execute steps A and C before step B', function () {
    const configFilePath = join(__dirname, 'simple-case.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);

      expect(trace).toHaveStep('step-a');
      expect(trace).toHaveStep('step-b');
      expect(trace).toHaveStep('step-c');

      expect(trace).toHaveStepLogged(['step-a', '##step-a##']);
      expect(trace).toHaveStepLogged(['step-b', '##step-b##']);
      expect(trace).toHaveStepLogged(['step-c', '##step-c##']);

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);
    }

    test("selflow-daemon", async () => {
      const logs = await startRun(configFilePath);
      verifyLogs(logs)
    })

    test("selflow-cli", async () => {
      const logs = await startCliRun(configFilePath);
      verifyLogs(logs)
    })
  });
});
