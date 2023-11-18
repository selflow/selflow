import { describe, expect } from 'vitest';
import { join } from 'path';
import { matchers } from '../../../../tools/trace';
import { withSelflowRunners } from '../../../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Custom Images', function () {
  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'step-python should prompt python version and step-node the node version',
    join(__dirname, 'custom-image.yaml'),
    (trace, logs) => {
      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-python');
      expect(trace).toHaveStep('step-node');

      expect(trace).toHaveStepTerminatedWithStatus(['step-python', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-node', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-python', 'Python 3']);
      expect(trace).toHaveStepLogged(['step-node', 'v18']);
    }
  );
});
