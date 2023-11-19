import {startCliRun, startRun} from '../../tools/run';
import {describe, expect, test} from 'vitest';
import {join} from 'path';
import {parseLogs} from '../../tools/logParser';
import {matchers} from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with step conditions', function () {
  describe('Step A should execute but not step B, C or D', function () {
    const configFilePath = join(__dirname, 'with-condition.yaml');
    const verifyLogs = (logs: string) => {

      const trace = parseLogs(logs);

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-d', 'CANCELLED']);
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

  describe('Step A and C should execute but not step B', function () {
    const configFilePath = join(__dirname, 'with-condition_template.yaml')
    const verifyLogs = (logs: string) => {

      const trace = parseLogs(logs);

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

      expect(trace).not.toHaveStepLogged(['step-b', '##step-b##']);
      expect(trace).toHaveStepLogged(['step-c', '##step-c##']);
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
