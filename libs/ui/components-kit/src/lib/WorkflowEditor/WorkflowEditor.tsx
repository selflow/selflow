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
import {WorkflowStep, WorkflowStepProps} from "./WorkflowStep/WorkflowStep";
import {Node} from "@reactflow/core/dist/esm/types";

export type WorkflowEditorProps = {}

const initialNodes: Node<WorkflowStepProps>[] = [
  {
    id: 'node-1', type: 'workflowStep', position: {x: 0, y: 0}, data: {
      status: {
        name: 'SUCCESS',
        code: 1,
        isCancellable: true,
        isFinished: true
      }
    }
  },
  {id: 'node-2', type: 'workflowStep', position: {x: 0, y: 200}, data:{
      status: {
        name: 'SUCCESS',
        code: 1,
        isCancellable: true,
        isFinished: true
      }
    }},
]

const nodeTypes = {workflowStep: WorkflowStep};


export const WorkflowEditor = ({}: WorkflowEditorProps) => {
  const [nodes, setNodes] = useState<Node<WorkflowStepProps>[]>(initialNodes);
  const [edges, setEdges] = useState<Edge[]>([]);

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
