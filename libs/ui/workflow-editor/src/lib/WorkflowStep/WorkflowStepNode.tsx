import React from 'react';
import { Handle, NodeProps, Position } from 'reactflow';
import { WorkflowStepStatus } from '../types';
import { WorkflowStepNodeStatusIndicator } from './WorkflowStepNodeStatusIndicator';

export type WorkflowStepProps = {
  status?: WorkflowStepStatus;
};

export const WorkflowStepNode = ({
  id,
  data: { status },
}: NodeProps<WorkflowStepProps>) => {
  return (
    <>
      <Handle type="target" position={Position.Left} />

      <div
        className={
          'h-[70px] w-[200px] bg-white p-5 border-2 border-gray-400 rounded-sm flex items-center gap-2 font-mono relative'
        }
      >
        {status ? (
          <div className={'absolute -top-2 -right-2'}>
            <WorkflowStepNodeStatusIndicator status={status} />
          </div>
        ) : null}

        <span>{id}</span>
      </div>
      <Handle type="source" position={Position.Right} />
    </>
  );
};
