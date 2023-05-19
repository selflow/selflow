import {Meta, StoryFn} from "@storybook/react";
import {WorkflowEditor, WorkflowEditorProps} from "./WorkflowEditor";
import {statusMap} from "./statusList";


const Story: Meta<typeof WorkflowEditor> = {
  component: WorkflowEditor,
  title: 'Workflow-Editor/WorkflowEditor',
  parameters: {
    layout: 'fullscreen'
  }
}

export default Story

export const Template: StoryFn<WorkflowEditorProps> = (args) => <WorkflowEditor {...args} />
Template.args = {
  steps: [
    {
      id: 'checkout',
      dependencies: [],
      status: statusMap.SUCCESS
    },
    {
      id: 'install-deps',
      dependencies: ['checkout'],
      status: statusMap.SUCCESS
    },
    {
      id: 'unit-tests',
      dependencies: ['install-deps'],
      status: statusMap.RUNNING
    },
    {
      id: 'linter',
      dependencies: ['install-deps'],
      status: statusMap.RUNNING
    },
    {
      id: 'build',
      dependencies: ['install-deps'],
      status: statusMap.SUCCESS
    },
    {
      id: 'e2e-tests',
      dependencies: ['build'],
      status: statusMap.RUNNING
    },
    {
      id: 'deploy',
      dependencies: ['e2e-tests', 'linter', 'unit-tests'],
      status: statusMap.PENDING
    }
  ]
}

Template.decorators = [
  story => <div className={"h-screen w-screen"}>{story()}</div>
]




