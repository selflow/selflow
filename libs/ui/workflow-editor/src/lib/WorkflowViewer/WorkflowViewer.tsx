import ReactFlow, {
  Background,
  Controls,
  NodeMouseHandler,
  OnConnect,
  Panel,
} from 'reactflow';
import 'reactflow/dist/style.css';
import { WorkflowStepNode } from '../WorkflowStep/WorkflowStepNode';
import { FaBars, FaTimes } from 'react-icons/all';
import { useWorkflow } from '../Providers/WorkflowProvider';

export type WorkflowViewerProps = {
  isSideMenuOpen: boolean;
  setSideMenuOpen: (open: boolean) => void;
  viewOnly: boolean;
  onStepClick?: (nodeId: string) => void;
};

const nodeTypes = { workflowStep: WorkflowStepNode };

export const WorkflowViewer = ({
  isSideMenuOpen,
  setSideMenuOpen,
  viewOnly,
  onStepClick,
}: WorkflowViewerProps) => {
  const { nodes, edges, onEdgesChange, addDependency } = useWorkflow();

  const onConnect: OnConnect = (connection) => {
    if (!connection.source || !connection.target) return;
    addDependency(connection.target, connection.source);
  };

  const nodeClick: NodeMouseHandler = (_, { id }) =>
    onStepClick && onStepClick(id);

  return (
    <div className={'w-full h-full'}>
      <ReactFlow
        nodesConnectable={!viewOnly}
        nodes={nodes}
        edges={edges}
        nodesDraggable={false}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        nodeTypes={nodeTypes}
        onNodeClick={nodeClick}
      >
        {viewOnly ? null : (
          <Panel position={'top-right'}>
            <button
              onClick={() => setSideMenuOpen(!isSideMenuOpen)}
              className={
                'bg-orange-400 p-3 grid place-items-center rounded-full'
              }
            >
              {isSideMenuOpen ? (
                <FaTimes className={'fill-white'} size={24} />
              ) : (
                <FaBars className={'fill-white'} size={24} />
              )}
            </button>
          </Panel>
        )}
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
};
