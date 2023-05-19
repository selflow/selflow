import {CodeEditor, Input, MultiSelect, Button} from "@selflow/ui/components-kit";

export type EditStepFormProps = {}

const sampleItems = [
  {id: "toto", name: "Toto"},
  {id: "tata", name: "Tata"},
  {id: "tutu", name: "Tutu"},
  {id: "titi", name: "Titi"},
]

export const EditStepForm = ({}: EditStepFormProps) => {
  return <>
    <h1 className={"text-2xl"}>New Step</h1>

    <form className={"mt-5"}>

      <Input type="text" label={"Step Id"}/>
      <Input type="text" label={"Id"}/>

      <CodeEditor lang={'sh'} label={"Commands"}/>

      <MultiSelect items={sampleItems} placeholder={"Pick the step dependencies"} label={"Dependencies"}/>

      <div className={"w-full text-right my-5"}>
        <Button>Create</Button>
      </div>
    </form>

  </>
}
