import type { Meta } from '@storybook/react';
import { Input } from './Inputs';

const Story: Meta<typeof Input> = {
  component: Input,
  title: 'inputs/Input',
};
export default Story;

export const Primary = {
  args: {
    label: 'Some label',
    placeholder: 'Some placeholder'
  },
};
