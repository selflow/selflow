import React from 'react';
import { WorkflowStepStatus } from '../types';

export type WorkflowStepStatusIndicatorProps = {
  status: WorkflowStepStatus;
};

const RunningIcon = (
  <>
    <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-orange-400 opacity-75"></span>
    <span className="relative inline-flex rounded-full h-3 w-3 bg-orange-500"></span>
  </>
);

const WorkflowStepStatusIndicatorByStatusName: Record<string, JSX.Element> = {
  SUCCESS: (
    <span className="relative inline-flex rounded-full h-3 w-3 bg-green-600" />
  ),
  ERROR: (
    <span className="relative inline-flex rounded-full h-3 w-3 bg-red-700" />
  ),
  CANCELLED: (
    <span className="relative inline-flex rounded-full h-3 w-3 bg-gray-700" />
  ),
  RUNNING: RunningIcon,
};

export const WorkflowStepNodeStatusIndicator = ({
  status,
}: WorkflowStepStatusIndicatorProps) => {
  return (
    <span className="relative flex h-3 w-3">
      {WorkflowStepStatusIndicatorByStatusName[status.name] ?? null}
    </span>
  );
};
