import { FC, PropsWithChildren } from 'react';
import Image from 'next/image';
import Logo from './selflow-logo.png';
import Link from 'next/link';

export type NavbarProps = PropsWithChildren;

export const Navbar: FC<NavbarProps> = ({ children }) => {
  return (
    <div className={'shadow-lg h-16 px-3 flex'}>
      <Link href={'/'} className={'w-16'}>
        <div className={'h-full relative aspect-square'}>
          <Image fill={true} src={Logo} alt={'Selflow Logo'} />
        </div>
      </Link>
      <div className={'flex items-center p-5 gap-2'}>{children}</div>
    </div>
  );
};
