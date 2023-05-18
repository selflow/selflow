import {Meta, StoryFn} from "@storybook/react";
import {WorkflowEditor} from "./WorkflowEditor";


const Story: Meta<typeof WorkflowEditor> = {
  component: WorkflowEditor,
  title: 'ui-kit/WorkflowEditor',
  parameters: {
    layout: 'fullscreen'
  }
}

export default Story

export const Template: StoryFn = (args) => <WorkflowEditor {...args} />

Template.decorators = [
  story => <div className={"h-screen w-screen"}>{story()}</div>
]




