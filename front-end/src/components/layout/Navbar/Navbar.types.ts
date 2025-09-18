export interface NavItem {
  id: string;
  label: string;
  href: string;
  icon?: string;
  children?: NavItem[];
  isExternal?: boolean;
}

export interface NavbarProps {
  className?: string;
  onNavigate?: (path: string) => void;
}

export interface MobileMenuProps {
  isOpen: boolean;
  onClose: () => void;
  items: NavItem[];
}

// Real estate specific navigation structure
export const NAV_ITEMS: NavItem[] = [
  {
    id: 'properties',
    label: 'Propriedades',
    href: '/properties',
    icon: 'line-md:home-twotone',
  },
  {
    id: 'clients',
    label: 'Clientes',
    href: '/clients',
    icon: 'line-md:person-twotone',
  },
  {
    id: 'leads',
    label: 'Leads',
    href: '/leads',
    icon: 'line-md:lightbulb-twotone',
  },
  {
    id: 'financing',
    label: 'Financiamentos',
    href: '/financing',
    icon: 'line-md:clipboard-minus',
  },
];
