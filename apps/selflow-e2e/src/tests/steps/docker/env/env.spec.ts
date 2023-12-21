import { startCliRun, startRun } from '../../../../tools/run';
import { expect, test } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../../../tools/logParser';
import { matchers } from '../../../../tools/trace';

expect.extend(matchers);

describe('Environment variables', function () {
  describe('step-a and step-b should prompt the values of their environment variables ', function () {
    const configFilePath = join(__dirname, 'env.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);

      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-a');
      expect(trace).toHaveStep('step-b');

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-a', '##step-a##']);
      expect(trace).toHaveStepLogged(['step-b', '##step-b##']);
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

  describe('Dynamic docker environment variables', function () {
    const configFilePath = join(__dirname, 'env-dynamic.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);

      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-a');
      expect(trace).toHaveStep('step-b');
      expect(trace).toHaveStep('step-c');

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-a', 'my-env-value']);
      expect(trace).toHaveStepLogged(['step-c', 'bar']);
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
