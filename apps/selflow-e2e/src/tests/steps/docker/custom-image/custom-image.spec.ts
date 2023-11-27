import { startCliRun, startRun } from '../../../../tools/run';
import { expect, test } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../../../tools/logParser';
import { matchers } from '../../../../tools/trace';

expect.extend(matchers);

describe('Custom Images', function () {
  describe('step-python should prompt python version and step-node the node version ', function () {
    const configFilePath = join(__dirname, 'custom-image.yaml');

    const verifyLogs = (logs: string) => {
      const trace = parseLogs(logs);

      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-python');
      expect(trace).toHaveStep('step-node');

      expect(trace).toHaveStepTerminatedWithStatus(['step-python', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-node', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-python', 'Python 3']);
      expect(trace).toHaveStepLogged(['step-node', 'v18']);
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
