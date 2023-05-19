import {Button} from "./Button";
import {Meta} from "@storybook/react";


const Story: Meta<typeof Button> = {
  component: Button,
  title: 'Button'
}

export default Story

export const Primary = {
  args: {
    children: 'Click Me !',
    onClick: () => alert('You clicked me !')
  }
}


