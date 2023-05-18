import React from "react";
import {Handle, NodeProps, Position} from "reactflow";
import {WorkflowStepStatus} from "../types";
import {WorkflowStepStatusIndicator} from "./WorkflowStepStatusIndicator";


export type WorkflowStepProps = {
  status: WorkflowStepStatus
}


export const WorkflowStep = ({id, data: {status}}: NodeProps<WorkflowStepProps>) => {
  return <>
    <Handle type="target" position={Position.Left}  />
    <div className={"h-[70px] w-[200px] bg-white p-5 border-2 border-gray-400 rounded flex items-center gap-2 font-mono"}>
      <WorkflowStepStatusIndicator status={status} />
      <span>{id}</span>
    </div>
    <Handle type="source" position={Position.Right} />
  </>
}
