import { useRouter } from 'next/router';
import { trpc } from '../../utils/trpc';
import { FaPlay } from 'react-icons/fa';
import { WorkflowEditor, WorkflowStep } from '@selflow/ui/workflow-editor';
import { useEffect, useMemo, useState } from 'react';
import { Spinner } from '@selflow/ui/components-kit';

const mapDataToSteps = (data: any) =>
  Object.keys(data.state).map((stepId) => ({
    id: stepId,
    status: data.state[stepId],
    with: {
      image: '',
      commands: '',
    },
    needs: data.dependencies[stepId] ?? [],
  }));

export default function FollowRun() {
  const router = useRouter();
  const runId = router.query.runId as string;

  const [isRunTerminated, setIsRunTerminated] = useState(true);

  const { data } = trpc.status.useQuery(runId, {
    refetchInterval: isRunTerminated ? undefined : 1000,
  });

  useEffect(() => {
    if (!data) {
      return;
    }

    setIsRunTerminated(
      Object.keys(data.state).every((stepId) => data.state[stepId].name)
    );
  }, [data, setIsRunTerminated]);

  const steps = useMemo(() => (data ? mapDataToSteps(data) : []), [data]);

  if (!data) {
    return (
      <div className={'grid place-items-center h-screen'}>
        <Spinner size={'lg'} />
      </div>
    );
  }

  return (
    <div className={'h-screen w-screen overflow-hidden flex flex-col'}>
      <div className={'bg-blue-300 h-16 flex items-center p-5 gap-2'}>
        <button className={'rounded-full border-2 border-white p-3'}>
          <FaPlay className={'fill-white '} />
        </button>

        {!isRunTerminated && <Spinner size={'xs'} />}
      </div>
      <div className={'grow w-full'}>
        <WorkflowEditor steps={steps} viewOnly={true} />
      </div>
    </div>
  );
}
