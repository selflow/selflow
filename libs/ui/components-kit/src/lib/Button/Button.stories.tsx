import {Button} from "./Button";
import {Meta, StoryFn} from "@storybook/react";


const Story: Meta<typeof Button> = {
  component: Button,
  title: 'ui-kit/Button'
}

export default Story

export const Template: StoryFn = (args) => <Button {...args} />



