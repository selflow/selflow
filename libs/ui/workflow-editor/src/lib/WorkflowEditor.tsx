import 'reactflow/dist/style.css';
import { useState } from 'react';
import { WorkflowStep } from './types';
import { WorkflowViewer } from './WorkflowViewer/WorkflowViewer';
import { RightSidePanel } from './RightSidePanel/RightSidePanel';
import { useWorkflow, WorkflowProvider } from './Providers/WorkflowProvider';
import { EditStepForm } from './EditStepForm/EditStepForm';

type WorkflowEditorViewProps = {
  viewOnly?: boolean;
};

export type WorkflowEditorProps = WorkflowEditorViewProps & {
  steps: WorkflowStep[];
  onChange?: (steps: WorkflowStep[]) => void;
};

export const WorkflowEditor = ({
  steps,
  onChange,
  ...viewProps
}: WorkflowEditorProps) => {
  return (
    <WorkflowProvider initialSteps={steps} onChange={onChange ?? (() => null)}>
      <WorkflowEditor$ {...viewProps} />
    </WorkflowProvider>
  );
};

export const WorkflowEditor$ = ({ viewOnly }: WorkflowEditorViewProps) => {
  const [isRightSidePanelOpen, setIsRightSidePanelOpen] = useState(true);
  const [editedStep, setEditedStep] = useState<WorkflowStep | undefined>(
    undefined
  );

  const { steps } = useWorkflow();

  const onStepClick = (stepId: string) => {
    const step = steps.find((step) => step.id === stepId) ?? undefined;
    setEditedStep(step);
  };

  return (
    <div className={'w-full h-full flex overflow-hidden'}>
      <WorkflowViewer
        setSideMenuOpen={setIsRightSidePanelOpen}
        viewOnly={!!viewOnly}
        isSideMenuOpen={isRightSidePanelOpen}
        onStepClick={onStepClick}
      />
      <RightSidePanel isOpen={isRightSidePanelOpen}>
        <EditStepForm initialStep={editedStep} />
      </RightSidePanel>
    </div>
  );
};
