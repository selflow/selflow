import { trpc } from '../utils/trpc';
import { WorkflowEditor, WorkflowStep } from '@selflow/ui/workflow-editor';
import { FaPlay } from 'react-icons/fa';
import { useState } from 'react';
import { useRouter } from 'next/router';

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
      <div className={'bg-blue-300 h-16 flex items-center p-5 gap-2'}>
        <button
          className={'rounded-full border-2 border-white p-3'}
          onClick={onPlay}
        >
          <FaPlay className={'fill-white '} />
        </button>
        <p>{JSON.stringify(newSteps)}</p>
      </div>
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
