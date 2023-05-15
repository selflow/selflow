import { WorkflowExecutionTrace } from './logParser';

interface WorkflowTraceExecutionMatchers<R = unknown> {
  toHaveStep(stepId: string): R;
  toHaveStepStoppedBefore(steps: [string, string]): R;
  toHaveStepLogged(steps: [string, string]): R;
  toHaveStepTerminatedWithStatus(steps: [string, string]): R;
}

declare module 'vitest' {
  interface Assertion<T = any> extends WorkflowTraceExecutionMatchers<T> {}
  interface AsymmetricMatchersContaining
    extends WorkflowTraceExecutionMatchers {}
}
