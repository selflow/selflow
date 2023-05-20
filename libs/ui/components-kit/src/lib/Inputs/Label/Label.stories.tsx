import type {Meta} from '@storybook/react';
import {Label} from './Label';

const Story: Meta<typeof Label> = {
  component: Label,
  title: 'inputs/Label',
};
export default Story;

export const Primary = {
  args: {
    children: 'Hello World !',
  },
};
