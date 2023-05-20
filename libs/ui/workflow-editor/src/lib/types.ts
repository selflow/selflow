export interface WorkflowStepStatus {
  code: number;
  name: string;
  isFinished: boolean;
  isCancellable: boolean;
}

export interface WorkflowStep {
  id: string;
  status?: WorkflowStepStatus;
  needs: string[];
  with: {
    image: string;
    commands: string;
  };
}
