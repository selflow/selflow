import { WorkflowStep } from '../types';
import { Node } from '@reactflow/core/dist/esm/types';
import { WorkflowStepProps } from '../WorkflowStep/WorkflowStepNode';
import { Edge, MarkerType } from 'reactflow';

const ySpacing = 100;
const xSpacing = 300;
const yMargin = 50;
const xMargin = 50;

const getDepth = (
  stepId: string,
  stepMap: Record<string, WorkflowStep>
): number => {
  if (!stepMap[stepId] || stepMap[stepId].needs.length === 0) return 0;
  return Math.max(
    ...stepMap[stepId].needs.map(
      (dependency) => getDepth(dependency, stepMap) + 1
    )
  );
};

export const mapWorkflowStepToReactFlowNodeAndEdges = (
  steps: WorkflowStep[]
): [Node<WorkflowStepProps>[], Edge[]] => {
  const stepsAsMap = steps.reduce(
    (acc, step) => ({ ...acc, [step.id]: step }),
    {}
  );

  const withDepth: Record<number, number> = {};

  const nodes: Node<WorkflowStepProps>[] = [];
  const edges: Edge[] = [];

  for (const step of steps) {
    const depth = getDepth(step.id, stepsAsMap);
    const lineIndex = withDepth[depth] === undefined ? 0 : withDepth[depth] + 1;
    nodes.push({
      id: step.id,
      type: 'workflowStep',
      position: {
        x: depth * xSpacing + xMargin,
        y: lineIndex * ySpacing + yMargin,
      },
      data: step,
    });
    edges.push(
      ...step.needs.map((dependency) => ({
        id: `${step.id}_${dependency}`,
        source: dependency,
        target: step.id,
        markerEnd: {
          type: MarkerType.Arrow,
        },
      }))
    );
    withDepth[depth] = lineIndex;
  }

  return [nodes, edges];
};
