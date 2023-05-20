import type {Meta} from '@storybook/react';
import {CodeEditor} from './CodeEditor';

const Story: Meta<typeof CodeEditor> = {
  component: CodeEditor,
  title: 'inputs/CodeEditor',
};
export default Story;

Story.decorators = [(story) => <div className={'w-[600px]'}>{story()}</div>];

export const Primary = {
  args: {
    label: 'Some label',
    lang: 'sh',
    value: `
cowsay "Selflow is love, Selflow is Life ❤️"

# ____________________________________
#< Selflow is love, Selflow is Life ❤️ >
# ------------------------------------
#        \\   ^__^
#         \\  (oo)\\_______
#            (__)\\       )\\/\\
#                ||----w |
#                ||     ||
`,
  },
};
