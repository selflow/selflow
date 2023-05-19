import ReactFlow, {Background, Controls, Edge, OnConnect, OnEdgesChange, OnNodesChange, Panel} from 'reactflow';
import 'reactflow/dist/style.css';
import {WorkflowStepNode, WorkflowStepProps} from "../WorkflowStep/WorkflowStepNode";
import {Node} from "@reactflow/core/dist/esm/types";
import {FaBars, FaTimes} from "react-icons/all";

export type WorkflowViewerProps = {
  nodes: Node<WorkflowStepProps>[],
  edges: Edge[],
  onNodesChange: OnNodesChange,
  onEdgesChange: OnEdgesChange,
  onConnect: OnConnect,
  isSideMenuOpen: boolean
  setSideMenuOpen: (open: boolean) => void
  viewOnly: boolean
}


const nodeTypes = {workflowStep: WorkflowStepNode};


export const WorkflowViewer = ({
                                 nodes,
                                 edges,
                                 onNodesChange,
                                 onEdgesChange,
                                 onConnect,
                                 isSideMenuOpen,
                                 setSideMenuOpen,
                                 viewOnly
                               }: WorkflowViewerProps) => {

  return <div className={"w-full h-full"}>
    <ReactFlow
      nodesConnectable={!viewOnly}
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      onConnect={onConnect}
      nodeTypes={nodeTypes}
    >
      {viewOnly ? null : <Panel position={"top-right"}>
        <button onClick={() => setSideMenuOpen(!isSideMenuOpen)}
                className={"bg-orange-400 p-3 grid place-items-center rounded-full"}>
          {isSideMenuOpen ? <FaTimes className={'fill-white'} size={24}/> :
            <FaBars className={'fill-white'} size={24}/>}
        </button>
      </Panel>}
      <Background/>
      <Controls/>
    </ReactFlow>
  </div>
}
