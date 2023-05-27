import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Learn how Selflow works',
    Svg: require('@site/static/img/undraw_lost.svg').default,
    description: (
      <>Want to try Selflow but got lost ? Let's get you back on rails !</>
    ),
  },
  {
    title: 'Start building your own workflows today',
    Svg: require('@site/static/img/undraw_start_building.svg').default,
    description: (
      <>Follow the Getting Started Guide and you will be setup in no time !</>
    ),
  },
  {
    title: 'Join us !',
    Svg: require('@site/static/img/undraw_join.svg').default,
    description: (
      <>
        Join the Selflow community and help us improve our tooling layer by
        providing ideas, feedbacks or even integrate your own modules !
      </>
    ),
  },
];

function Feature({ title, Svg, description }: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
