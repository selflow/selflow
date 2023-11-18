import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';
import { withSelflowRunners } from '../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Workflow with step conditions', function () {
  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'Step A should execute but not step B, C or D',
    join(__dirname, 'with-condition.yaml'),
    (trace) => {
      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-d', 'CANCELLED']);
    }
  );

  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'Step A and C should execute but not step B',
    join(__dirname, 'with-condition_template.yaml'),
    (trace) => {
      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

      expect(trace).not.toHaveStepLogged(['step-b', '##step-b##']);
      expect(trace).toHaveStepLogged(['step-c', '##step-c##']);
    }
  );
});
