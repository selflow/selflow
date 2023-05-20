import {Meta, StoryFn} from '@storybook/react';
import {WorkflowEditor, WorkflowEditorProps} from './WorkflowEditor';
import {statusMap} from './statusList';
import {WorkflowStep} from './types';

const Story: Meta<typeof WorkflowEditor> = {
  component: WorkflowEditor,
  title: 'WorkflowEditor',
  parameters: {
    layout: 'fullscreen',
  },
};

Story.decorators = [
  (story) => <div className={'h-screen w-screen'}>{story()}</div>,
];

const stepWith: WorkflowStep['with'] = {
  image: 'alpine:3.18.0',
  commands: 'echo toto',
};

export default Story;

export const Template: StoryFn<WorkflowEditorProps> = (args) => (
  <WorkflowEditor {...args} />
);
Template.args = {
  steps: [
    {
      id: 'checkout',
      needs: [],
      status: statusMap.SUCCESS,
      with: stepWith,
    },
    {
      id: 'install-deps',
      needs: ['checkout'],
      status: statusMap.SUCCESS,
      with: stepWith,
    },
    {
      id: 'unit-tests',
      needs: ['install-deps'],
      status: statusMap.RUNNING,
      with: stepWith,
    },
    {
      id: 'linter',
      needs: ['install-deps'],
      status: statusMap.RUNNING,
      with: stepWith,
    },
    {
      id: 'build',
      needs: ['install-deps'],
      status: statusMap.SUCCESS,
      with: stepWith,
    },
    {
      id: 'e2e-tests',
      needs: ['build'],
      status: statusMap.RUNNING,
      with: stepWith,
    },
    {
      id: 'deploy',
      needs: ['e2e-tests', 'linter', 'unit-tests'],
      status: statusMap.PENDING,
      with: stepWith,
    },
  ],
};

export const Edition = {
  args: {
    steps: [],
  },
};
