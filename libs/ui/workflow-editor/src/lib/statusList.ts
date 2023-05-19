import {WorkflowStepStatus} from "./types";

export const statusList = [
  {
    name: 'SUCCESS',
    code: 0,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'ERROR',
    code: 1,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'CANCELLED',
    code: 2,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'RUNNING',
    code: 3,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'PENDING',
    code: 4,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'INITIALIZING',
    code: 5,
    isCancellable: false,
    isFinished: true
  } as const,
  {
    name: 'CREATED',
    code: 6,
    isCancellable: false,
    isFinished: true
  } as const
] as const

type StatusName = (typeof statusList)[number]['name']

export const statusMap = statusList.reduce<{[key in StatusName]: WorkflowStepStatus}>((acc, status) => ({...acc, [status.name]: status}), {} as any)
