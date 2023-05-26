import { trpc } from '../utils/trpc';
import { WorkflowEditor, WorkflowStep } from '@selflow/ui/workflow-editor';
import { FaPlay } from 'react-icons/fa';
import { useState } from 'react';
import { useRouter } from 'next/router';
import { Navbar } from '@selflow/ui/components-kit';

export function Index() {
  const [newSteps, setNewSteps] = useState<WorkflowStep[]>([]);
  const router = useRouter();

  const { mutate } = trpc.run.useMutation({
    onSuccess(runId) {
      router.push('/run/' + runId).then(null);
    },
  });

  const onPlay = () => {
    mutate({
      workflow: {
        timeout: '5m',
        steps: newSteps.reduce(
          (acc, step) => ({
            ...acc,
            [step.id]: {
              kind: 'docker',
              needs: step.needs,
              with: step.with,
            },
          }),
          {}
        ),
      },
    });
  };

  return (
    <div className={'h-screen w-screen overflow-hidden flex flex-col'}>
      <Navbar>
        <h1 className={'text-xl'}>Build your workflow</h1>

        <button
          className={'rounded-full border-4 border-blue-600 p-3'}
          onClick={onPlay}
        >
          <FaPlay className={'fill-blue-600'} />
        </button>
      </Navbar>

      <div className={'grow w-full'}>
        <WorkflowEditor
          steps={newSteps}
          viewOnly={false}
          onChange={setNewSteps}
        />
      </div>
    </div>
  );
}

export default Index;
