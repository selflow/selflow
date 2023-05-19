import {ComponentPropsWithoutRef} from "react";
import {Label} from "../Label/Label";


export type InputProps = ComponentPropsWithoutRef<'input'> & {
  label: string
}

export const Input = ({label, id, className, ...defaultInputProps}: InputProps) => {
  return <div className={`my-2 ${className}`}>
    <Label htmlFor={id}>{label}</Label>
    <input id={id}
           className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
           {...defaultInputProps}/>
  </div>
}
