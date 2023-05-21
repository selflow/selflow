import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with step output', function () {
  it('Step B should access the output of Step a', async function () {
    const logs = await startRun(join(__dirname, 'with-output.yaml'));
    const trace = parseLogs(logs);

    expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

    expect(trace).toHaveStepLogged(['step-b', 'bar']);
  });
});
