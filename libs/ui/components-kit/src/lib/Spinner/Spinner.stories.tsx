import {Meta, StoryFn} from "@storybook/react";
import {Spinner, SpinnerSizes} from "./Spinner";


const Story: Meta<typeof Spinner> = {
  component: Spinner,
  title: 'ui-kit/Spinner'
}

export default Story

export const Template: StoryFn = (args) => <Spinner {...args} />



export const Sizes = () => <div className={"grid grid-cols-2 gap-2"}>
  {
    SpinnerSizes.map(size => (
      <>
        <code className={"grid content-center text-center font-mono"}>{size}</code>
        <Spinner size={size} />
      </>
    ))
  }
</div>



