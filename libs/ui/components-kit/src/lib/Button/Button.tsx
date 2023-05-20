import {ComponentPropsWithoutRef, forwardRef} from 'react';

export type ButtonProps = Omit<
  ComponentPropsWithoutRef<'button'>,
  'className'
> & {};

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (props, ref) => (
    <button
      ref={ref}
      {...props}
      className={
        'bg-blue-400 rounded text-white py-3 px-5 hover:bg-blue-500 transition-colors'
      }
    />
  )
);
