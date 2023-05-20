import {PropsWithChildren} from "react";


export type RightSidePanelProps = PropsWithChildren & {
  isOpen: boolean,
}

export const RightSidePanel = ({isOpen, children}: RightSidePanelProps) => {
  return <div className={`grid grid-cols-[0] ${isOpen ? 'grid-cols-[1fr]' : ''}`}>
    <div className={"w-[600px] h-full p-5 overflow-y-scroll"}>
      {children}
    </div>
  </div>
}
