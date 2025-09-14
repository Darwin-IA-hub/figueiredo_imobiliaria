import React, { useState } from 'react';
import './TabbedNavigation.css';

export interface Tab {
  id: string;
  label: string;
  content: React.ReactNode;
}

interface TabbedNavigationProps {
  tabs: Tab[];
  defaultTab?: string;
  onTabChange?: (tabId: string) => void;
}

const TabbedNavigation: React.FC<TabbedNavigationProps> = ({
  tabs,
  defaultTab,
  onTabChange,
}) => {
  const [activeTab, setActiveTab] = useState(defaultTab || tabs[0]?.id || '');

  const handleTabClick = (tabId: string) => {
    setActiveTab(tabId);
    onTabChange?.(tabId);
  };

  const activeTabContent = tabs.find((tab) => tab.id === activeTab)?.content;

  if (tabs.length === 0) {
    return null;
  }

  return (
    <div className="tabbed-navigation">
      <div className="tabs-container">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={`tab-button ${activeTab === tab.id ? 'active' : ''}`}
            onClick={() => handleTabClick(tab.id)}
          >
            {tab.label}
          </button>
        ))}
      </div>
      <div className="tab-content">{activeTabContent}</div>
    </div>
  );
};

export default TabbedNavigation;
