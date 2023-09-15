import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import './UsageScenario.scss';

interface UsageScenarioProps {
  scenarioId: number;
}

const UsageScenario: React.FC<UsageScenarioProps> = ({ scenarioId }) => {
  const usageScenarios = useSelector((state: RootState) => state.api.usageScenarios);

  const scenario = usageScenarios.find((s) => s.id === scenarioId);

  if (!scenario) {
    return null;
  }

  return (
    <div className="usage-scenario">
      <img src={`/images/usage-scenarios/${scenario.image}`} alt={scenario.title} />
      <h3>{scenario.title}</h3>
      <p>{scenario.description}</p>
    </div>
  );
};

export default UsageScenario;