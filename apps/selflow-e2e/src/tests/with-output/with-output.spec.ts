import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';
import { withSelflowRunners } from '../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Workflow with step output', function () {
  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'Step B should access the output of Step a',
    join(__dirname, 'with-output.yaml'),
    (trace) => {
      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-b', 'bar']);
    }
  );
});
