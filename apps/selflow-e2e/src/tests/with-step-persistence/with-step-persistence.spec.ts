import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';
import { withSelflowRunners } from '../../tools/withSelflowRunners';

expect.extend(matchers);

describe('Workflow with step persistence', function () {
  withSelflowRunners(
    ['selflow-cli', 'selflow-daemon'],
    'Step B should access the file created by Step A',
    join(__dirname, 'with-step-persistence.yaml'),
    (trace) => {
      expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
      expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);

      expect(trace).toHaveStepLogged(['step-b', 'Hello!']);
    }
  );
});
