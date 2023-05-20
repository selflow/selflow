import React from 'react';
import {Handle, NodeProps, Position} from 'reactflow';
import {WorkflowStepStatus} from '../types';
import {WorkflowStepNodeStatusIndicator} from './WorkflowStepNodeStatusIndicator';

export type WorkflowStepProps = {
  status?: WorkflowStepStatus;
};

export const WorkflowStepNode = ({
                                   id,
                                   data: {status},
                                 }: NodeProps<WorkflowStepProps>) => {
  return (
    <>
      <Handle type="target" position={Position.Left}/>
      {status ? (
        <div
          className={
            'h-[70px] w-[200px] bg-white p-5 border-2 border-gray-400 rounded flex items-center gap-2 font-mono'
          }
        >
          <WorkflowStepNodeStatusIndicator status={status}/>
          <span>{id}</span>
        </div>
      ) : null}
      <Handle type="source" position={Position.Right}/>
    </>
  );
};
