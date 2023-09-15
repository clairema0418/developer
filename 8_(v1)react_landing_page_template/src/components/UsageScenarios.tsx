import React from 'react';
import clsx from 'clsx';
import { SceneSchema } from '../lib/schemas';
import '../styles/usage-scenarios.scss';

interface UsageScenariosProps {
  sceneData: SceneSchema[];
}

const UsageScenarios: React.FC<UsageScenariosProps> = ({ sceneData }) => {
  return (
    <section id="scene" className="usage-scenarios">
      <h2>產品使用場景和示例</h2>
      <div className="scene-container">
        {sceneData.map((scene, index) => (
          <div key={index} className={clsx('scene-item', `scene-item-${index}`)}>
            <img src={scene.image} alt={scene.title} />
            <h3>{scene.title}</h3>
            <p>{scene.description}</p>
          </div>
        ))}
      </div>
    </section>
  );
};

export default UsageScenarios;