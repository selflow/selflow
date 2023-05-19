import {EditStepForm} from "../EditStepForm/EditStepForm";


export type RightSidePanelProps = {
  isOpen: boolean,
  close: () => void
}

export const RightSidePanel = ({isOpen}: RightSidePanelProps) => {
  return <div className={`grid grid-cols-[0] ${isOpen ? 'grid-cols-[1fr]' : ''}`}>
    <div className={"w-[600px] h-full p-5 overflow-y-scroll"}>
      <EditStepForm/>
    </div>
  </div>
}
