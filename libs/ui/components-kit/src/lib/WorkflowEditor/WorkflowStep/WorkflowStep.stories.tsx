import {Meta, StoryFn} from "@storybook/react";
import {WorkflowStep, WorkflowStepProps} from "./WorkflowStep";
import ReactFlow, {NodeProps} from "reactflow";
import {statusList, statusMap} from "../statusList";


const Story: Meta<typeof WorkflowStep> = {
  component: WorkflowStep,
  title: 'ui-kit/WorkflowStep',
  parameters: {
    layout: 'fullscreen'
  }
}

export default Story

export const Template: StoryFn<NodeProps<WorkflowStepProps>> = (args) => <WorkflowStep {...args} />

Template.args = {
  id: 'some-step',
  data: {
    status: statusMap.SUCCESS
  }
}

Template.decorators = [
  story =>
    <ReactFlow>
    {story()}
    </ReactFlow>
]

const nodeTypes = {workflowStep: WorkflowStep};

const colCount = 3;

const stepStatusNodes = statusList.map((status, index) => ({
  id: `status-${status.name.toLowerCase()}`,
  type: 'workflowStep',
  position: {x: (index % colCount) * 300, y: Math.floor(index / colCount) * 100},
  data: {status}
}))

export const StepStatus = () =>
  <div className={"h-screen w-screen p-10"}>
  <ReactFlow
    nodes={stepStatusNodes}
    nodeTypes={nodeTypes}

  >
  </ReactFlow>
  </div>

