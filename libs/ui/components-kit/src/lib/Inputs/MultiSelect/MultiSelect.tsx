import React, {FC, useState} from "react";
import {Combobox} from "@headlessui/react";
import {FaChevronDown, FaLink, FaTimes} from "react-icons/all";
import {Label} from "../Label/Label";


export type MultiSelectItem = { id: string, name: string }

export type MultiSelectProps = {
  items: MultiSelectItem[],
  placeholder?: string
  label: string
  initialSelectedItems?: MultiSelectItem[]
}


export const MultiSelect: FC<MultiSelectProps> = ({items, label, placeholder, initialSelectedItems}) => {
  const [selectedItems, setSelectedItems] = useState<MultiSelectItem[]>(initialSelectedItems ?? [])

  const removeItem = (itemId: MultiSelectItem['id']) => {
    const index = selectedItems.findIndex(item => item.id === itemId)
    setSelectedItems([...selectedItems.slice(0, index), ...selectedItems.slice(index + 1)])
  }

  return (
    <div className={"relative"}>
      {/* @ts-ignore */}
      <Combobox value={selectedItems} onChange={setSelectedItems} multiple>
        <Combobox.Label>
          <Label>{label}</Label>
        </Combobox.Label>

        {selectedItems.length > 0 && (
          <ul className={"flex flex-wrap gap-2 my-2"}>
            {selectedItems.map((selectedItem) => (
              <li key={selectedItem.id}
                  className={"inline-block py-1 px-2 text-white bg-blue-400 rounded cursor-pointer flex items-center gap-1"}
                  onClick={() => removeItem(selectedItem.id)}>
                <span>{selectedItem.name}</span>
                <FaTimes/>
              </li>
            ))}
          </ul>
        )}
        <Combobox.Button
          className={"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 flex items-center"}>
          <Combobox.Input placeholder={placeholder} className={"grow bg-transparent outline-none ring-none"}/>
          <FaChevronDown/>
        </Combobox.Button>
        <Combobox.Options className="absolute w-full mt-1 max-h-60 overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
          {items.map((item) => (
            <Combobox.Option key={item.id} value={item}>
              {({selected, active}) => (
                <div
                  className={`p-2 cursor-pointer bg-white hover:bg-gray-100 flex items-center ${active ? 'bg-gray-100 ' : ''}`}>
                <span className={"w-8 flex items-center justify-center"}>
                  {
                    selected && <FaLink className={"fill-sky-500"}/>
                  }
                </span>

                  <span>{item.name}</span>
                </div>
              )}
            </Combobox.Option>
          ))}
        </Combobox.Options>
      </Combobox>
    </div>
  )
}
