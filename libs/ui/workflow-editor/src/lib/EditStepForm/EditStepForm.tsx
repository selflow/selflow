import {
  Button,
  CodeEditor,
  Input,
  MultiSelect,
} from '@selflow/ui/components-kit';
import { Controller, useForm } from 'react-hook-form';
import { useWorkflow } from '../Providers/WorkflowProvider';
import { WorkflowStep } from '../types';
import { useEffect } from 'react';

type EditStepFormFields = WorkflowStep;

export type EditStepFormProps = {
  initialStep?: EditStepFormFields;
  viewOnly?: boolean;
  close?: () => void;
};
export const EditStepForm = ({
  initialStep,
  viewOnly,
  close,
}: EditStepFormProps) => {
  const { steps, setStep, addStep } = useWorkflow();

  const { register, handleSubmit, control, reset, formState } =
    useForm<EditStepFormFields>({
      defaultValues: initialStep,
    });

  const resetForm = (values?: Partial<WorkflowStep>) => {
    if (!values) {
      reset({
        id: '',
        needs: [],
        with: {
          image: '',
          commands: '',
        },
      });
    }
    reset(values);
  };

  useEffect(() => {
    resetForm(initialStep);
  }, [initialStep, reset]);

  const onSubmit = handleSubmit((data) => {
    if (initialStep) {
      setStep(initialStep, data);
    } else {
      addStep(data);
    }
    resetForm();
    close && close();
  });

  const title =
    viewOnly && initialStep
      ? `Viewing ${initialStep.id}`
      : initialStep
      ? `Editing ${initialStep.id}`
      : 'New Step';

  return (
    <>
      <h1 className={'text-2xl'}>{title}</h1>

      <form className={'mt-5'} onSubmit={onSubmit}>
        <Input
          type="text"
          label={'Step Id'}
          {...register('id', { required: true })}
          disabled={viewOnly}
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
              disabled={viewOnly}
            />
          )}
        />

        <h2 className={'text-xl mt-5'}>Container Configuration</h2>

        <Input
          type="text"
          label={'Docker Image'}
          {...register('with.image', { required: true })}
          disabled={viewOnly}
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
              contentEditable={!!viewOnly}
            />
          )}
        />

        {!viewOnly && (
          <div className={'w-full text-right my-5'}>
            <Button>Create</Button>
          </div>
        )}
      </form>
    </>
  );
};
