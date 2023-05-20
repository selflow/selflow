import type { Meta } from '@storybook/react';
import { MultiSelect } from './MultiSelect';

const Story: Meta<typeof MultiSelect> = {
  component: MultiSelect,
  title: 'inputs/MultiSelect',
};
export default Story;

const items = [
  {id: 'toto', name: 'Toto'},
  {id: 'tata', name: 'Tata'},
  {id: 'titi', name: 'Titi'},
  {id: 'tutu', name: 'tutu'},
];

export const Primary = {
  args: {
    label: 'Some Label',
    items,
  },
};

export const WithInitialSelection = {
  args: {
    ...Primary.args,
    initialSelectedItems: [items[0], items[2]],
  },
};
