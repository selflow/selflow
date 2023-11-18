import { describe, expect } from 'vitest';
import { join } from 'path';
import { matchers } from '../../../tools/trace';
import { withSelflowRunners } from '../../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Local Command', function () {
  withSelflowRunners(
    ['selflow-cli'],
    'should echo in step A and step B',
    join(__dirname, 'local-command.yaml'),
    (trace, logs) => {
      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-a');
      expect(trace).toHaveStep('step-b');

      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-a', '##step-a##']);
      expect(trace).toHaveStepLogged(['step-b', '##step-b##']);
    }
  );

  withSelflowRunners(
    ['selflow-cli'],
    'step-error should fail',
    join(__dirname, 'local-command-error.yaml'),
    (trace, logs) => {
      expect(logs).not.toEqual('');

      expect(trace).toHaveStep('step-error');
      expect(trace).toHaveStepTerminatedWithStatus(['step-error', 'ERROR']);
      expect(trace).toHaveStepLogged(['step-error', '##step-error##']);
    }
  );
});
