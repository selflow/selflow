import './Spinner.css';

export const SpinnerSizes = ['xs', 'sm', 'md', 'lg'] as const;

export type SpinnerSize = typeof SpinnerSizes[number];

export type SpinnerProps = {
  size?: SpinnerSize;
};

export const Spinner = ({ size = 'md' }: SpinnerProps) => (
  <div
    aria-label="Blue hamster running in a metal wheel"
    role="img"
    className={`wheel-and-hamster ${size}`}
  >
    <div className="wheel"></div>
    <div className="hamster">
      <div className="hamster__body">
        <div className="hamster__head">
          <div className="hamster__horn"></div>
          <div className="hamster__ear"></div>
          <div className="hamster__eye"></div>
          <div className="hamster__nose"></div>
        </div>
        <div className="hamster__limb hamster__limb--fr"></div>
        <div className="hamster__limb hamster__limb--fl"></div>
        <div className="hamster__limb hamster__limb--br"></div>
        <div className="hamster__limb hamster__limb--bl"></div>
        <div className="hamster__tail"></div>
      </div>
    </div>
    <div className="spoke"></div>
  </div>
);
