# Figueiredo Imóveis - Real Estate Management System

A modern React + TypeScript application for real estate property management with a professional component architecture.

## 🏗️ Project Structure

```
src/
├── components/                # Component-based architecture
│   ├── ui/                    # Shared UI components (reusable)
│   ├── layout/                # Layout components
│   │   └── Navbar/            # Navigation component
│   └── features/              # Feature-specific components
│       └── properties/        # Property management
│           └── PropertyCard/  # Property display component
├── hooks/                     # Custom React hooks
├── utils/                     # Utility functions
├── types/                     # Global TypeScript types
├── services/                  # API services
├── App.tsx                    # Main application
├── main.tsx                   # Entry point
└── vite-env.d.ts              # Vite environment types
```

## 🚀 Features

### **Implemented Components**

- **Navbar**: Responsive navigation with authentication state
- **PropertyCard**: Feature-rich property display with actions (view, edit, delete)
- **TypeScript**: Full type safety across the application
- **Styled Components**: Modern CSS-in-JS styling
- **Path Aliases**: Clean import paths using `@` prefix

### **Component Architecture**

- **Nested Structure**: Components organized by type and feature
- **Self-contained**: Each component has its own folder with types, styles, and tests
- **Reusable**: Shared UI components and feature-specific business components
- **Type-safe**: Comprehensive TypeScript interfaces

## 📦 Component Documentation

### **Navbar Component**

```typescript
import { Navbar } from '@layout/Navbar';

<Navbar
  isAuthenticated={true}
  userName="John Doe"
  onLogin={() => setAuthenticated(true)}
  onLogout={() => setAuthenticated(false)}
/>
```

**Props:**

- `isAuthenticated: boolean` - Authentication state
- `userName?: string` - User's display name
- `onLogin: () => void` - Login handler
- `onLogout: () => void` - Logout handler

### **PropertyCard Component**

```typescript
import { PropertyCard, Property } from '@features/properties/PropertyCard';

<PropertyCard
  property={{
    id: '1',
    title: 'Modern Apartment',
    price: 850000,
    location: 'Downtown, São Paulo',
    bedrooms: 3,
    bathrooms: 2,
    area: 120,
    imageUrl: '/property.jpg'
  }}
  onViewDetails={(id) => console.log('View:', id)}
  onEdit={(id) => console.log('Edit:', id)}
  onDelete={(id) => console.log('Delete:', id)}
/>
```

**Property Interface:**

```typescript
interface Property {
  id: string;
  title: string;
  price: number;
  location: string;
  bedrooms: number;
  bathrooms: number;
  area: number;
  imageUrl: string;
}
```

**Props:**

- `property: Property` - Property data
- `onViewDetails: (id: string) => void` - View details handler
- `onEdit: (id: string) => void` - Edit handler
- `onDelete: (id: string) => void` - Delete handler

## 🔧 TypeScript Path Aliases

Clean import paths using the `@` prefix:

```typescript
// Instead of messy relative paths
import Navbar from '../../../../components/layout/Navbar';

// Use clean aliases
import { Navbar } from '@layout/Navbar';
import { PropertyCard } from '@features/properties/PropertyCard';
import { useAuth } from '@hooks/useAuth';
import { Property } from '@types/Property';
```

**Available Aliases:**

- `@components/*` - `src/components/*`
- `@features/*` - `src/components/features/*`
- `@ui/*` - `src/components/ui/*`
- `@layout/*` - `src/components/layout/*`
- `@hooks/*` - `src/hooks/*`
- `@utils/*` - `src/utils/*`
- `@types/*` - `src/types/*`
- `@services/*` - `src/services/*`
- `@assets/*` - `src/assets/*`

## 🔐 Authentication Flow

The application implements a simple authentication system:

1. **Login**: Click "Login" button → `isAuthenticated = true`
2. **Logout**: Click "Logout" button → `isAuthenticated = false`
3. **State Management**: Uses React `useState` for auth state
4. **UI Updates**: Navbar shows user info when authenticated

**Example:**

```typescript
const [isAuthenticated, setIsAuthenticated] = useState(false);
const [userName, setUserName] = useState('');

const handleLogin = () => {
  setIsAuthenticated(true);
  setUserName('John Doe');
};

const handleLogout = () => {
  setIsAuthenticated(false);
  setUserName('');
};
```

## 🛠️ Development Setup

### **Prerequisites**

- Node.js 16+
- npm or yarn

### **Installation**

```bash
# Install dependencies
npm install

# Or with yarn
yarn install
```

### **Dependencies**

```json
{
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "styled-components": "^6.1.0"
  },
  "devDependencies": {
    "@types/styled-components": "^6.1.0",
    "@types/node": "^20.0.0",
    "typescript": "^5.0.0",
    "vite": "^4.0.0",
    "@vitejs/plugin-react-swc": "^3.0.0"
  }
}
```

### **Running the Application**

```bash
# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Run tests (if configured)
npm test
```

### **Development Server**

- **URL**: `http://localhost:5173`
- **Hot Module Replacement**: Enabled
- **TypeScript**: Full type checking
- **ESLint**: Code quality checks

## 🧪 Testing

Component tests are located in the same folder as the component:

```
src/components/layout/Navbar/Navbar.test.tsx
src/components/features/properties/PropertyCard/PropertyCard.test.tsx
```

**Example Test:**

```typescript
import { render, screen } from '@testing-library/react';
import { Navbar } from './Navbar';

test('renders login button when not authenticated', () => {
  render(<Navbar isAuthenticated={false} onLogin={() => {}} onLogout={() => {}} />);
  expect(screen.getByText('Login')).toBeInTheDocument();
});
```

## 🎨 Styling

- **Styled Components**: CSS-in-JS for component styling
- **Responsive Design**: Mobile-friendly layouts
- **Theme Support**: Easy to extend with themes
- **Component Isolation**: Styles are scoped to components

## 📚 Component Development Guidelines

### **Creating New Components**

1. Create folder: `src/components/features/[feature]/[ComponentName]/`
2. Create files:
   - `[ComponentName].tsx` - Main component
   - `[ComponentName].types.ts` - TypeScript interfaces
   - `[ComponentName].styles.ts` - Styled components
   - `[ComponentName].test.tsx` - Tests
   - `index.ts` - Public API exports

### **Best Practices**

- **PascalCase** for component files: `PropertyCard.tsx`
- **kebab-case** for folders: `property-card/`
- **TypeScript interfaces** for all props
- **Index files** for clean imports
- **Test coverage** for all components

## 🚀 Future Roadmap

### **Next Features to Implement**

- [ ] Lead management components
- [ ] Client management components
- [ ] Contact management components
- [ ] Financing calculator components
- [ ] API integration with backend
- [ ] Authentication with JWT
- [ ] Property search and filtering
- [ ] Image gallery for properties
- [ ] Admin dashboard

### **Additional Components**

```bash
# Lead management
src/components/features/leads/LeadForm/
src/components/features/leads/LeadList/

# Client management
src/components/features/clients/ClientCard/
src/components/features/clients/ClientForm/

# Contact management
src/components/features/contacts/ContactForm/
src/components/features/contacts/ContactList/

# Shared UI
src/components/ui/Button/
src/components/ui/Input/
src/components/ui/Modal/
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/new-component`
3. Commit your changes: `git commit -m 'Add new component'`
4. Push to the branch: `git push origin feature/new-component`
5. Open a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ for Figueiredo Imóveis - Your Real Estate Management Solution**
