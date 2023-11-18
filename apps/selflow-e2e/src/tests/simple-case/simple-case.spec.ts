import { describe, expect } from 'vitest';
import { join } from 'path';
import { matchers } from '../../tools/trace';
import { withSelflowRunners } from '../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Workflow with dependencies', function () {
  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'should execute steps A and C before step B',
    join(__dirname, 'simple-case.yaml'),
    (trace) => {
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
  );
});
