import 'reactflow/dist/style.css';
import {useState} from 'react';
import {WorkflowStep} from './types';
import {WorkflowViewer} from './WorkflowViewer/WorkflowViewer';
import {RightSidePanel} from './RightSidePanel/RightSidePanel';
import {useWorkflow, WorkflowProvider} from './Providers/WorkflowProvider';
import {EditStepForm} from './EditStepForm/EditStepForm';

export type WorkflowEditorProps = {
  steps: WorkflowStep[];
};

export const WorkflowEditor = ({steps}: WorkflowEditorProps) => {
  return (
    <WorkflowProvider initialSteps={steps}>
      <WorkflowEditor$/>
    </WorkflowProvider>
  );
};

export const WorkflowEditor$ = () => {
  const [isRightSidePanelOpen, setIsRightSidePanelOpen] = useState(true);
  const [editedStep, setEditedStep] = useState<WorkflowStep | undefined>(
    undefined
  );

  const {steps} = useWorkflow();

  const onStepClick = (stepId: string) => {
    const step = steps.find((step) => step.id === stepId) ?? undefined;
    setEditedStep(step);
  };

  return (
    <div className={'w-full h-full flex overflow-hidden'}>
      <WorkflowViewer
        setSideMenuOpen={setIsRightSidePanelOpen}
        viewOnly={false}
        isSideMenuOpen={isRightSidePanelOpen}
        onStepClick={onStepClick}
      />
      <RightSidePanel isOpen={isRightSidePanelOpen}>
        <EditStepForm initialStep={editedStep}/>
      </RightSidePanel>
    </div>
  );
};
