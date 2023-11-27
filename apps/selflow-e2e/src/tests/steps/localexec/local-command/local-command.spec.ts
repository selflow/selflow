import { describe, expect, test } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../../../tools/logParser';
import { startCliRun } from '../../../../tools/run';
import { matchers } from '../../../../tools/trace';

expect.extend(matchers);

describe('Local Command', function () {
  describe('should echo in step A and step B', () => {
    const configFilePath = join(__dirname, 'local-command.yaml');

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

    test('selflow-cli', async () => {
      const logs = await startCliRun(configFilePath);
      return verifyLogs(logs);
    });
  });

  describe('step-error should fail', () => {
    const configFilePath = join(__dirname, 'local-command-error.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);
      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-error');
      expect(trace).toHaveStepTerminatedWithStatus(['step-error', 'ERROR']);
      expect(trace).toHaveStepLogged(['step-error', '##step-error##']);
    };

    test('selflow-cli', async () => {
      const logs = await startCliRun(configFilePath);
      return verifyLogs(logs);
    });
  });
});
