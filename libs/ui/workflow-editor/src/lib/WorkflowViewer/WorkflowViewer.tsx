import ReactFlow, {
  Background,
  Controls,
  NodeMouseHandler,
  OnConnect,
  Panel,
} from 'reactflow';
import 'reactflow/dist/style.css';
import { WorkflowStepNode } from '../WorkflowStep/WorkflowStepNode';
import { FaBars, FaPlus, FaTimes } from 'react-icons/fa';
import { useWorkflow } from '../Providers/WorkflowProvider';

export type WorkflowViewerProps = {
  isSideMenuOpen: boolean;
  setSideMenuOpen: (open: boolean) => void;
  viewOnly: boolean;
  onStepClick?: (nodeId: string) => void;
  onAddClick?: () => void;
};

const nodeTypes = { workflowStep: WorkflowStepNode };

export const WorkflowViewer = ({
  isSideMenuOpen,
  setSideMenuOpen,
  viewOnly,
  onStepClick,
  onAddClick,
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
        <Panel position={'top-right'}>
          <div className={'flex flex-col gap-2'}>
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
            {!viewOnly && onAddClick ? (
              <button
                onClick={onAddClick}
                className={
                  'bg-blue-400 p-3 grid place-items-center rounded-full'
                }
              >
                <FaPlus className={'fill-white'} size={24} />
              </button>
            ) : null}
          </div>
        </Panel>
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
};
