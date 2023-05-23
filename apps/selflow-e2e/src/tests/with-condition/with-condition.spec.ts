import { startRun } from '../../tools/run';
import { expect } from 'vitest';
import { join } from 'path';
import { parseLogs } from '../../tools/logParser';
import { matchers } from '../../tools/trace';

expect.extend(matchers);

describe('Workflow with step conditions', function () {
  it('Step A should execute but not step B, C or D', async function () {
    const logs = await startRun(join(__dirname, 'with-condition.yaml'));
    const trace = parseLogs(logs);

    expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'CANCELLED']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-d', 'CANCELLED']);
  });

  it('Step A and C should execute but not step B', async function () {
    const logs = await startRun(
      join(__dirname, 'with-condition_template.yaml')
    );
    const trace = parseLogs(logs);

    expect(trace).toHaveStepTerminatedWithStatus(['step-a', 'SUCCESS']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-b', 'CANCELLED']);
    expect(trace).toHaveStepTerminatedWithStatus(['step-c', 'SUCCESS']);

    expect(trace).not.toHaveStepLogged(['step-b', '##step-b##']);
    expect(trace).toHaveStepLogged(['step-c', '##step-c##']);
  });
});
