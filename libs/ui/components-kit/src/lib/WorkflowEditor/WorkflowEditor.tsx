import ReactFlow, {
  addEdge,
  applyEdgeChanges,
  applyNodeChanges,
  Background,
  Controls,
  Edge,
  OnConnect,
  OnEdgesChange,
  OnNodesChange
} from 'reactflow';
import 'reactflow/dist/style.css';
import {useState} from "react";
import {WorkflowStepNode, WorkflowStepProps} from "./WorkflowStep/WorkflowStepNode";
import {Node} from "@reactflow/core/dist/esm/types";
import {WorkflowStep} from "./types";

export type WorkflowEditorProps = {
  steps: WorkflowStep[]
}

const ySpacing = 100;
const xSpacing = 300;
const yMargin = 50;
const xMargin = 50;

const getDepth = (stepId: string, stepMap: Record<string, WorkflowStep>): number => {
  if (!stepMap[stepId] || stepMap[stepId].dependencies.length === 0) return 0;
  return Math.max(...stepMap[stepId].dependencies.map(dependency => getDepth(dependency, stepMap) + 1))
}

const mapWorkflowStepToReactFlowNodeAndEdges = (steps: WorkflowStep[]): [Node<WorkflowStepProps>[], Edge[]] => {
  const stepsAsMap = steps.reduce((acc, step) => ({...acc, [step.id]: step}), {})

  let withDepth: Record<number, number> = {}

  let nodes: Node<WorkflowStepProps>[] = []
  let edges: Edge[] = []

  for (let step of steps) {
    const depth = getDepth(step.id, stepsAsMap)
    const lineIndex = withDepth[depth] === undefined ? 0 : withDepth[depth] + 1;
    nodes.push({
      id: step.id,
      type: 'workflowStep',
      position: {
        x: (depth * xSpacing) + xMargin,
        y: (lineIndex * ySpacing) + yMargin,
      },
      data: step
    })
    edges.push(...step.dependencies.map(dependency => ({
      id: `${step.id}_${dependency}`,
      source: dependency,
      target: step.id
    })))
    withDepth[depth] = lineIndex;
  }

  console.log([nodes, edges])

  return [nodes, edges]
}

const nodeTypes = {workflowStep: WorkflowStepNode};


export const WorkflowEditor = ({steps}: WorkflowEditorProps) => {
  const [initNodes, initEdges] = mapWorkflowStepToReactFlowNodeAndEdges(steps)
  const [nodes, setNodes] = useState<Node<WorkflowStepProps>[]>(initNodes);
  const [edges, setEdges] = useState<Edge[]>(initEdges);

  const onNodesChange: OnNodesChange = (changes) => setNodes((nds) => applyNodeChanges(changes, nds));
  const onEdgesChange: OnEdgesChange = (changes) => setEdges((eds) => applyEdgeChanges(changes, eds));
  const onConnect: OnConnect = (connection) => setEdges((eds) => addEdge(connection, eds));


  return <div className={"w-full h-full"}>
    <ReactFlow
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      onConnect={onConnect}
      nodeTypes={nodeTypes}
    >
      <Background/>
      <Controls/>
    </ReactFlow>
  </div>
}
