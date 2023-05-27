import { useRouter } from 'next/router';
import { trpc } from '../../utils/trpc';
import { FaPlay } from 'react-icons/fa';
import { WorkflowEditor, WorkflowStep } from '@selflow/ui/workflow-editor';
import { useEffect, useMemo, useState } from 'react';
import { Navbar, Spinner } from '@selflow/ui/components-kit';

const mapDataToSteps = (data: any) =>
  Object.keys(data.state).map((stepId) => ({
    if: data.steps[stepId].If,
    kind: data.steps[stepId].Kind,
    id: stepId,
    status: data.state[stepId],
    with: {
      ...data.steps[stepId].With,
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
      <div
        className={'grid place-items-center h-screen'}
        id={'full-screen-loader'}
      >
        <Spinner size={'lg'} />
      </div>
    );
  }

  return (
    <div className={'h-screen w-screen overflow-hidden flex flex-col'}>
      <Navbar>
        <h1 className={' text-xl'}>Run results</h1>
        {!isRunTerminated && <Spinner size={'xs'} />}
      </Navbar>

      <div className={'grow w-full'}>
        <WorkflowEditor steps={steps} viewOnly={true} />
      </div>
    </div>
  );
}
