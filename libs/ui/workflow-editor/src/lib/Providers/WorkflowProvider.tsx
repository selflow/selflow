import {createContext, FC, PropsWithChildren, useContext, useEffect, useState} from "react";
import {WorkflowStep} from "../types";
import {mapWorkflowStepToReactFlowNodeAndEdges} from "./stepToNode";
import {Node} from "@reactflow/core/dist/esm/types";
import {WorkflowStepProps} from "../WorkflowStep/WorkflowStepNode";
import {addEdge, applyEdgeChanges, Edge, OnConnect, OnEdgesChange} from "reactflow";


export type WorkflowProviderState = {
  steps: WorkflowStep[],
  nodes: Node<WorkflowStepProps>[],
  edges: Edge[],
  addStep: (newStep: WorkflowStep) => void
  addDependency: (stepId: WorkflowStep['id'], dependency: WorkflowStep['id']) => void
  onEdgesChange: OnEdgesChange
  onConnect: OnConnect,
  setStep: (oldStep: WorkflowStep, newStep: WorkflowStep) => void
}

const WorkflowContext = createContext<WorkflowProviderState>({
  nodes: [],
  steps: [],
  addStep: () => null,
  edges: [],
  addDependency: () => null,
  onEdgesChange: () => null,
  onConnect: () => null,
  setStep: () => null
})

export type WorkflowProviderProps = PropsWithChildren & {
  initialSteps: WorkflowStep[]
}

export const WorkflowProvider: FC<WorkflowProviderProps> = ({children, initialSteps}) => {
  const [initialNodes, initialEdges] = mapWorkflowStepToReactFlowNodeAndEdges(initialSteps)
  const [steps, setSteps] = useState(initialSteps);

  const [nodes, setNodes] = useState(initialNodes)
  const [edges, setEdges] = useState(initialEdges)

  useEffect(() => {
    const [newNodes, newEdges] = mapWorkflowStepToReactFlowNodeAndEdges(steps)

    setNodes(newNodes)
    setEdges(newEdges)
  }, [steps, setNodes, setEdges])

  const onEdgesChange: OnEdgesChange = (changes) => setEdges((eds) => applyEdgeChanges(changes, eds));
  const onConnect: OnConnect = (connection) => setEdges((eds) => addEdge(connection, eds));

  const newDependency = (stepId: WorkflowStep['id'], dependency: WorkflowStep['id']) => {
    const index = steps.findIndex(step => step.id === stepId);
    setSteps([
      ...steps.slice(0, index),
      {
        ...steps[index],
        needs: [
          ...new Set([...steps[index].needs, dependency])
        ]
      },
      ...steps.slice(index + 1)
    ])
  }

  const addStep = (newStep: WorkflowStep) => setSteps(
    [
      ...steps,
      newStep
    ]
  )

  const setStep = (oldStep: WorkflowStep, newStep: WorkflowStep) => {
    const hasIdChanged = oldStep.id !== newStep.id
    if (hasIdChanged) {
      setSteps(steps.map(step => {
        if (step.id === oldStep.id) {
          return {...newStep}
        }
        const dependencyIndex = step.needs.indexOf(oldStep.id)
        if (dependencyIndex === -1) {
          console.log('dep not found for step ' + step.id)
          return {...step}
        }

        console.log('dep found for step ' + step.id)

        return {
          ...step,
          needs: [
            ...step.needs.slice(0, dependencyIndex),
            newStep.id,
            ...step.needs.slice(dependencyIndex + 1)
          ]
        }
      }))
    } else {
      const stepIndex = steps.findIndex(step => step.id === oldStep.id)
      setSteps([
        ...steps.slice(0, stepIndex),
        newStep,
        ...steps.slice(stepIndex + 1)
      ])
    }
  }

  return <WorkflowContext.Provider value={{
    steps,
    nodes,
    addStep,
    edges,
    addDependency: newDependency,
    onEdgesChange,
    onConnect,
    setStep
  }}>
    {children}
  </WorkflowContext.Provider>
}

export const useWorkflow = () => useContext(WorkflowContext)

