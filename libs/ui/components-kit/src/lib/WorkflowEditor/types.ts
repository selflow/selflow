

export interface WorkflowStepStatus {
  code: number
  name: string
  isFinished: boolean
  isCancellable: boolean
}


export interface WorkflowStep {
  id: string
  status: WorkflowStepStatus
  dependencies: string[]
}
