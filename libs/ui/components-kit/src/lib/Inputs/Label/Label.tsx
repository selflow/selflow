import {ComponentPropsWithoutRef, forwardRef} from "react";


export type LabelProps = Omit<ComponentPropsWithoutRef<'label'>, 'className'> & {}

export const Label = forwardRef<HTMLLabelElement, LabelProps>(
  (props, ref) => <label className={"block mb-2 text-sm font-medium text-gray-900"} ref={ref} {...props} />)
