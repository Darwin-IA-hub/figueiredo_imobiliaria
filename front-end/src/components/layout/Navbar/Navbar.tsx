import React, { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Icon } from '@iconify/react';
import { NAV_ITEMS, NavbarProps } from './Navbar.types';
import { useAuth } from '../../../hooks/useAuth';
import { useTheme } from '../../../contexts/ThemeContext';
import './Navbar.css';

const Navbar: React.FC<NavbarProps> = ({ className = '' }) => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [activeDropdown, setActiveDropdown] = useState<string | null>(null);
  const [userDropdownOpen, setUserDropdownOpen] = useState(false);
  const location = useLocation();
  const { user, logout } = useAuth();
  const { isDarkMode, toggleTheme } = useTheme();

  // Close mobile menu when route changes
  useEffect(() => {
    setIsMobileMenuOpen(false);
    setActiveDropdown(null);
  }, [location]);

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
    setActiveDropdown(null);
  };

  const toggleDropdown = (itemId: string) => {
    setActiveDropdown(activeDropdown === itemId ? null : itemId);
  };

  const toggleUserDropdown = () => {
    setUserDropdownOpen(!userDropdownOpen);
    setActiveDropdown(null);
  };

  const handleLogout = () => {
    logout();
    setUserDropdownOpen(false);
    setIsMobileMenuOpen(false);
  };

  const handleNavItemClick = (item: { href: string; isExternal?: boolean }) => {
    if (item.isExternal) {
      window.open(item.href, '_blank');
    }
    setIsMobileMenuOpen(false);
    setActiveDropdown(null);
  };

  const isActiveLink = (href: string) => {
    return (
      location.pathname === href || location.pathname.startsWith(`${href}/`)
    );
  };

  return (
    <nav
      className={`navbar ${className}`}
      role="navigation"
      aria-label="Main navigation"
    >
      <div className="navbar__container">
        {/* Logo and Brand */}
        <div className="navbar__brand">
          <Link
            to="/"
            className="navbar__logo-link"
            aria-label="Figueiredo Imóveis - Página inicial"
          >
            <img
              src="/src/assets/logo.png"
              alt="Figueiredo Imóveis"
              className="navbar__logo"
              width={40}
              height={40}
            />
            <span className="navbar__brand-name">Figueiredo Imóveis</span>
          </Link>
        </div>

        {/* Desktop Navigation */}
        <ul className="navbar__nav" role="menubar">
          {NAV_ITEMS.map((item) => (
            <li key={item.id} className="navbar__nav-item" role="none">
              {item.children ? (
                <div className="navbar__dropdown-wrapper">
                  <button
                    className={`navbar__nav-link ${isActiveLink(item.href) ? 'active' : ''}`}
                    onClick={() => toggleDropdown(item.id)}
                    aria-expanded={activeDropdown === item.id}
                    aria-haspopup="true"
                    role="menuitem"
                  >
                    {item.icon && (
                      <span className="navbar__icon">
                        <Icon icon={item.icon} width="18" height="18" />
                      </span>
                    )}
                    {item.label}
                    <span className="navbar__dropdown-arrow">▼</span>
                  </button>

                  {activeDropdown === item.id && (
                    <div className="navbar__dropdown" role="menu">
                      {item.children.map((child) => (
                        <Link
                          key={child.id}
                          to={child.href}
                          className={`navbar__dropdown-item ${isActiveLink(child.href) ? 'active' : ''}`}
                          role="menuitem"
                          onClick={() => handleNavItemClick(child)}
                        >
                          {child.label}
                        </Link>
                      ))}
                    </div>
                  )}
                </div>
              ) : (
                <Link
                  to={item.href}
                  className={`navbar__nav-link ${isActiveLink(item.href) ? 'active' : ''}`}
                  role="menuitem"
                  onClick={() => handleNavItemClick(item)}
                >
                  {item.icon && (
                    <span className="navbar__icon">
                      <Icon icon={item.icon} width="18" height="18" />
                    </span>
                  )}
                  {item.label}
                </Link>
              )}
            </li>
          ))}
        </ul>

        {/* User Actions - Desktop */}
        <div className="navbar__actions">
          {/* Theme Toggle */}
          <button
            className="navbar__theme-toggle"
            onClick={toggleTheme}
            aria-label={
              isDarkMode ? 'Switch to light mode' : 'Switch to dark mode'
            }
          >
            <Icon
              icon={
                isDarkMode ? 'line-md:sunny-outline-loop' : 'line-md:moon-loop'
              }
              width="20"
              height="20"
            />
          </button>

          {/* User Dropdown */}
          {user ? (
            <div className="navbar__user-dropdown-wrapper">
              <button
                className="navbar__user-toggle"
                onClick={toggleUserDropdown}
                aria-expanded={userDropdownOpen}
                aria-haspopup="true"
              >
                <span className="navbar__user-name">
                  {user.name.split(' ')[0]}
                </span>
                <Icon icon="line-md:account" width="20" height="20" />
              </button>

              {userDropdownOpen && (
                <div className="navbar__user-dropdown" role="menu">
                  <div className="navbar__user-info">
                    <span className="navbar__user-fullname">{user.name}</span>
                    <span className="navbar__user-email">{user.email}</span>
                  </div>
                  <button
                    className="navbar__logout-btn"
                    onClick={handleLogout}
                    role="menuitem"
                  >
                    <Icon icon="line-md:logout" width="18" height="18" />
                    Sair
                  </button>
                </div>
              )}
            </div>
          ) : (
            <Link to="/login" className="navbar__login-link">
              <span className="navbar__icon">
                <Icon icon="line-md:log-in" width="18" height="18" />
              </span>
              Entrar
            </Link>
          )}
        </div>

        {/* Mobile Menu Toggle */}
        <button
          className={`navbar__hamburger ${isMobileMenuOpen ? 'active' : ''}`}
          onClick={toggleMobileMenu}
          aria-label="Abrir menu de navegação"
          aria-expanded={isMobileMenuOpen}
        >
          <span className="navbar__hamburger-line"></span>
          <span className="navbar__hamburger-line"></span>
          <span className="navbar__hamburger-line"></span>
        </button>

        {/* Mobile Menu Overlay */}
        {isMobileMenuOpen && (
          <div
            className="navbar__mobile-overlay"
            onClick={toggleMobileMenu}
          ></div>
        )}

        {/* Mobile Navigation */}
        <div
          className={`navbar__mobile-menu ${isMobileMenuOpen ? 'active' : ''}`}
        >
          <div className="navbar__mobile-header">
            <img
              src="/src/assets/logo.png"
              alt="Figueiredo Imóveis"
              className="navbar__mobile-logo"
              width={32}
              height={32}
            />
            <button
              className="navbar__mobile-close"
              onClick={toggleMobileMenu}
              aria-label="Fechar menu"
            >
              <span className="navbar__close-icon">×</span>
            </button>
          </div>

          <ul className="navbar__mobile-nav" role="menubar">
            {NAV_ITEMS.map((item) => (
              <li key={item.id} className="navbar__mobile-item" role="none">
                {item.children ? (
                  <div className="navbar__mobile-dropdown">
                    <button
                      className={`navbar__mobile-link ${isActiveLink(item.href) ? 'active' : ''}`}
                      onClick={() => toggleDropdown(item.id)}
                      aria-expanded={activeDropdown === item.id}
                    >
                      {item.icon && (
                        <span className="navbar__icon">
                          <Icon icon={item.icon} width="18" height="18" />
                        </span>
                      )}
                      {item.label}
                      <span className="navbar__mobile-arrow">
                        {activeDropdown === item.id ? '▲' : '▼'}
                      </span>
                    </button>

                    {activeDropdown === item.id && (
                      <div
                        className="navbar__mobile-dropdown-content"
                        role="menu"
                      >
                        {item.children.map((child) => (
                          <Link
                            key={child.id}
                            to={child.href}
                            className={`navbar__mobile-dropdown-item ${isActiveLink(child.href) ? 'active' : ''}`}
                            role="menuitem"
                            onClick={() => handleNavItemClick(child)}
                          >
                            {child.label}
                          </Link>
                        ))}
                      </div>
                    )}
                  </div>
                ) : (
                  <Link
                    to={item.href}
                    className={`navbar__mobile-link ${isActiveLink(item.href) ? 'active' : ''}`}
                    role="menuitem"
                    onClick={() => handleNavItemClick(item)}
                  >
                    {item.icon && (
                      <span className="navbar__icon">
                        <Icon icon={item.icon} width="18" height="18" />
                      </span>
                    )}
                    {item.label}
                  </Link>
                )}
              </li>
            ))}
          </ul>

          {/* User Actions - Mobile */}
          <div className="navbar__mobile-actions">
            {/* Theme Toggle */}
            <button
              className="navbar__mobile-theme-toggle"
              onClick={toggleTheme}
              aria-label={
                isDarkMode ? 'Switch to light mode' : 'Switch to dark mode'
              }
            >
              <Icon
                icon={
                  isDarkMode
                    ? 'line-md:sunny-outline-loop'
                    : 'line-md:moon-loop'
                }
                width="20"
                height="20"
              />
              {isDarkMode ? 'Modo Claro' : 'Modo Escuro'}
            </button>

            {/* Logout Button */}
            {user && (
              <button
                className="navbar__mobile-logout-btn"
                onClick={handleLogout}
              >
                <Icon icon="line-md:logout" width="18" height="18" />
                Sair
              </button>
            )}
          </div>

          {/* Contact Info in Mobile Menu */}
          <div className="navbar__mobile-contact">
            <div className="navbar__contact-item">
              <span className="navbar__contact-icon">📞</span>
              <span>(11) 9999-9999</span>
            </div>
            <div className="navbar__contact-item">
              <span className="navbar__contact-icon">✉️</span>
              <span>contato@figueiredo.com</span>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
