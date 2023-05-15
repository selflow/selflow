import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with dependencies', function () {
  it('should execute steps A and C before step B', async function () {
    const logs = await startRun(join(__dirname, 'simple-case.yaml'));
    const trace = parseLogs(logs);

    expect(trace).toHaveStep('step-a');
    expect(trace).toHaveStep('step-b');
    expect(trace).toHaveStep('step-c');

    expect(trace).toHaveStepLogged(['step-a', '##step-a##']);
    expect(trace).toHaveStepLogged(['step-b', '##step-b##']);
    expect(trace).toHaveStepLogged(['step-c', '##step-c##']);

    expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);
  });
});
