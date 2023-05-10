import {expect} from "vitest";
import {Event, LogLine} from "./logParser";

export type WorkflowExecutionTrace = {
  logs: LogLine[]
  stepLogs: Record<string, string[]>;
  events: Event[]
}

export const matchers = {
  toHaveStep(received: WorkflowExecutionTrace, expected: string) {
    const {isNot} = this
    return {
      pass: received.events.some(e => e.stepId === expected),
      message: () => `workflow execution does${isNot ? ' not' : ''} contains step ${expected}`
    }
  },

  toHaveStepStoppedBefore(received: WorkflowExecutionTrace, [firstStep, secondStep]: [string, string]) {
    const {isNot} = this

    const firstStepStopEventOrder = received.events.filter(step => step.stepId === firstStep && step.eventType === 'stop')[0].order
    const secondStepStartEventOrder = received.events.filter(step => step.stepId === secondStep && step.eventType === 'start')[0].order

    return {
      pass: firstStepStopEventOrder < secondStepStartEventOrder,
      message: () => `step ${firstStep} did${isNot ? ' not' : ''} stopped before ${secondStep} started`
    }
  },

  toHaveStepLogged(received: WorkflowExecutionTrace, [stepId, expectedLog]: [string, string]) {
    const {isNot} = this

    const stepLogs = received.stepLogs[stepId]
    if (!stepLogs) {
      return {
        pass: false,
        message: () => `no logs found for step with id ${stepId}`
      }
    }

    return {
      pass: stepLogs.some(log => log.includes(expectedLog)),
      message: () => `step ${stepId} did${isNot ? ' not' : ''} logged ${expectedLog}`
    }
  },

  toHaveStepTerminatedWithStatus(received: WorkflowExecutionTrace, [stepId, expectedStatus]: [string, string]) {
    const {isNot} = this

    if (!received.events.some(e => e.stepId === stepId)) {
      return {
        pass: false,
        message: () => `no step found with id ${stepId}`
      }
    }

    const exitStatus = received.events.filter(e => e.eventType === 'stop' && e.stepId === stepId)[0].status

    return {
      pass: exitStatus === expectedStatus,
      message: () => `step ${stepId} did${isNot ? ' not' : ''} exit with status ${exitStatus} and not ${expectedStatus}`
    }
  }
} as const

const _: Parameters<(typeof expect)["extend"]>[0] = matchers

expect.extend(matchers);

