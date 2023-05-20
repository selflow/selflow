import {Button, CodeEditor, Input, MultiSelect,} from '@selflow/ui/components-kit';
import {Controller, useForm} from 'react-hook-form';
import {useWorkflow} from '../Providers/WorkflowProvider';
import {WorkflowStep} from '../types';
import {useEffect} from 'react';

type EditStepFormFields = WorkflowStep;

export type EditStepFormProps = {
  initialStep?: EditStepFormFields;
};
export const EditStepForm = ({ initialStep }: EditStepFormProps) => {
  const { steps, setStep, addStep } = useWorkflow();

  const { register, handleSubmit, control, reset } =
    useForm<EditStepFormFields>({
      defaultValues: initialStep,
    });

  useEffect(() => {
    reset(initialStep);
  }, [initialStep, reset]);

  const onSubmit = handleSubmit((data) => {
    if (initialStep) {
      setStep(initialStep, data);
    } else {
      addStep(data);
    }
  });

  return (
    <>
      <h1 className={'text-2xl'}>New Step</h1>

      <form className={'mt-5'} onSubmit={onSubmit}>
        <Input
          type="text"
          label={'Step Id'}
          {...register('id', { required: true })}
        />

        <Controller
          name={'needs'}
          control={control}
          defaultValue={[]}
          render={({ field }) => (
            <MultiSelect
              items={steps.map((step) => step.id)}
              placeholder={'Pick the step dependencies'}
              label={'Dependencies'}
              onChange={field.onChange}
              initialSelectedItems={field.value}
            />
          )}
        />

        <h2 className={'text-xl mt-5'}>Container Configuration</h2>

        <Input
          type="text"
          label={'Docker Image'}
          {...register('with.image', { required: true })}
        />

        <Controller
          name={'with.commands'}
          control={control}
          defaultValue={''}
          render={({ field }) => (
            <CodeEditor
              lang={'sh'}
              label={'Commands'}
              value={field.value}
              onChange={field.onChange}
            />
          )}
        />

        <div className={'w-full text-right my-5'}>
          <Button>Create</Button>
        </div>
      </form>
    </>
  );
};
