import React from 'react';
import { WorkflowStepStatus } from '../types';
import { Spinner } from '@selflow/ui/components-kit';
import { FaBan, FaCheck, FaTimes } from 'react-icons/fa';

export type WorkflowStepStatusIndicatorProps = {
  status: WorkflowStepStatus;
};

const WorkflowStepStatusIndicatorByStatusName: Record<string, JSX.Element> = {
  SUCCESS: <FaCheck className={'fill-green-700'} size={20} />,
  ERROR: <FaTimes className={'fill-red-700'} size={20} />,
  CANCELLED: <FaBan className={'fill-gray-700'} size={20} />,
  RUNNING: <Spinner size={'xs'} />,
};

const PendingIcon = (
  <span className="relative flex h-3 w-3">
    <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-orange-400 opacity-75"></span>
    <span className="relative inline-flex rounded-full h-3 w-3 bg-orange-500"></span>
  </span>
);

export const WorkflowStepNodeStatusIndicator = ({
  status,
}: WorkflowStepStatusIndicatorProps) => {
  return WorkflowStepStatusIndicatorByStatusName[status.name] ?? PendingIcon;
};
