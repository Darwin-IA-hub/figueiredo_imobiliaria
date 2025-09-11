'use client';

import { useEffect, useLayoutEffect, useRef, useState, useCallback } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import {
  Home,
  Users,
  Heart,
  BarChart3,
  TrendingUp,
  Rocket,
  Sun,
  Moon,
  LogOut,
} from 'lucide-react';

type NavItem = {
  name: string;
  href: string;
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>;
  color: string; // hex
  colorRgb?: string; // computed at runtime
};

const navItems: NavItem[] = [
  { name: 'Início', href: '/', icon: Home, color: '#3b82f6' },
  { name: 'Contatos', href: '/contacts', icon: Users, color: '#ef4444' },
  { name: 'Interesses', href: '/interests', icon: Heart, color: '#10b981' },
  { name: 'Estatísticas Darwin IA', href: '/darwin-ai-stats', icon: BarChart3, color: '#f59e0b' },
  { name: 'Estatísticas Gerais', href: '/general-stats', icon: TrendingUp, color: '#8b5cf6' },
  { name: 'Lançamentos', href: '/launches', icon: Rocket, color: '#ec4899' },
];

// helper: #rrggbb -> "r g b"
function hexToRgbChannelString(hex: string) {
  const n = hex.replace('#', '');
  const bigint = parseInt(n.length === 3 ? n.split('').map(c => c + c).join('') : n, 16);
  const r = (bigint >> 16) & 255;
  const g = (bigint >> 8) & 255;
  const b = bigint & 255;
  return `${r} ${g} ${b}`;
}

export default function Navbar() {
  const pathname = usePathname();
  const [activeIndex, setActiveIndex] = useState(0);
  const [isDark, setIsDark] = useState(true);

  // measure & animate indicator
  const containerRef = useRef<HTMLDivElement | null>(null);
  const itemRefs = useRef<Array<HTMLAnchorElement | null>>([]);
  const [indicator, setIndicator] = useState({ left: 0, width: 44, beamLeft: 0, beamWidth: 56 });

  // Create a stable ref callback
  const setItemRef = useCallback((index: number) => (el: HTMLAnchorElement | null) => {
    itemRefs.current[index] = el;
  }, []);

  // initialize active index from route
  useEffect(() => {
    const i = navItems.findIndex((n) => n.href === pathname);
    if (i >= 0) setActiveIndex(i);
  }, [pathname]);

  // theme on mount
  useEffect(() => {
    const stored = typeof window !== 'undefined' ? localStorage.getItem('theme') : null;
    const dark = stored ? stored === 'dark' : true;
    setIsDark(dark);
    document.documentElement.classList.toggle('dark', dark);
  }, []);

  // recompute positions on resize/route/active changes
  useLayoutEffect(() => {
    const el = itemRefs.current[activeIndex];
    const wrap = containerRef.current;
    if (!el || !wrap) return;

    const cr = el.getBoundingClientRect();
    const wr = wrap.getBoundingClientRect();

    const tabCenter = cr.left - wr.left + cr.width / 2;
    const capsuleWidth = Math.min(64, Math.max(40, Math.round(cr.width * 0.45)));
    const capsuleLeft = tabCenter - capsuleWidth / 2;

    const beamWidth = Math.max(48, Math.round(cr.width * 0.5));
    const beamLeft = tabCenter - beamWidth / 2;

    setIndicator({
      left: capsuleLeft,
      width: capsuleWidth,
      beamLeft,
      beamWidth,
    });
  }, [activeIndex, pathname]);

  useEffect(() => {
    const onResize = () => {
      // force recompute with same dependencies
      setActiveIndex((i) => i);
    };
    window.addEventListener('resize', onResize);
    return () => window.removeEventListener('resize', onResize);
  }, []);

  const toggleTheme = () => {
    const next = !isDark;
    setIsDark(next);
    document.documentElement.classList.toggle('dark', next);
    localStorage.setItem('theme', next ? 'dark' : 'light');
  };

  // accent variables
  const accent = navItems[activeIndex]?.color ?? '#22c55e';
  const accentRgb = hexToRgbChannelString(accent);

  return (
    <nav 
      className="relative bg-[var(--nav-bg)] border-b border-[var(--nav-border)] backdrop-blur-md transition-colors duration-200"
      role="navigation"
      aria-label="Main navigation"
    >
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        {/* top bar */}
        <div className="flex h-16 items-center justify-between">
          {/* left: logo */}
          <Link href="/" className="flex items-center gap-2">
            <img src="/logo.png" alt="Figueiredo Imóveis" className="h-8 w-auto" />
          </Link>

          {/* center: pill navbar */}
          <div
            className="relative"
            style={
              {
                // accent variables for CSS
                ['--accent' as any]: accent,
                ['--accent-rgb' as any]: accentRgb,
              } as React.CSSProperties
            }
          >
            <div
              ref={containerRef}
              className="relative flex items-center gap-1 rounded-full bg-[var(--nav-item-bg)] px-1 py-1 shadow-lg dark:shadow-[0_2px_20px_rgba(0,0,0,.35)] shadow-[0_2px_20px_rgba(0,0,0,.1)]"
              role="tablist"
              aria-label="Primary"
            >
              {/* sliding capsule (top) */}
              <div
                className="absolute -top-1 h-2 rounded-full bg-[var(--accent)] transition-all duration-300 ease-out will-change-transform"
                style={{
                  left: indicator.left,
                  width: indicator.width,
                }}
              />

              {/* spotlight beam */}
              <div
                className="pointer-events-none absolute top-0 -translate-y-[2px] rounded-b-full opacity-60 transition-all duration-300 ease-out will-change-transform"
                style={{
                  left: indicator.beamLeft,
                  width: indicator.beamWidth,
                  height: '70%',
                  // gradient beam using accent → transparent
                  background:
                    'linear-gradient(to bottom, rgb(var(--accent-rgb)) 0%, rgba(var(--accent-rgb),0.35) 18%, rgba(var(--accent-rgb),0.12) 40%, transparent 78%)',
                  filter: 'blur(0.5px)',
                }}
              />

              {navItems.map((item, i) => (
                <Link
                  key={item.name}
                  href={item.href}
                  ref={setItemRef(i)}
                  onClick={() => setActiveIndex(i)}
                  role="tab"
                  aria-selected={activeIndex === i}
                  className={[
                    'relative z-10 flex items-center gap-2 rounded-full px-4 py-2 text-sm font-medium transition-colors',
                    activeIndex === i
                      ? 'text-[var(--accent)] text-shadow:0_0_10px_var(--accent)'
                      : 'text-[var(--nav-text)] hover:text-[var(--nav-text-hover)]',
                  ].join(' ')}
                  style={
                    activeIndex === i
                      ? ({
                          '--accent': item.color,
                          '--accent-rgb': hexToRgbChannelString(item.color),
                        } as React.CSSProperties)
                      : undefined
                  }
                >
                  <item.icon
                    className={['h-4 w-4', activeIndex === i ? 'text-[var(--accent)] drop-shadow-[0_0_10px_var(--accent)]' : 'text-current'].join(
                      ' ',
                    )}
                    // for screen readers
                    aria-hidden="true"
                  />
                  {item.name}
                </Link>
              ))}
            </div>
          </div>

          {/* right: actions */}
          <div className="flex items-center gap-3">
            <button
              onClick={toggleTheme}
              className="text-[var(--nav-text)] hover:text-[var(--nav-text-hover)] transition-colors"
              aria-label={isDark ? 'Alternar para tema claro' : 'Alternar para tema escuro'}
            >
              {isDark ? <Sun className="h-5 w-5" /> : <Moon className="h-5 w-5" />}
            </button>
            <button
              className="text-[var(--nav-text)] hover:text-[var(--nav-text-hover)] transition-colors"
              aria-label="Sair"
              // onClick={() => signOut()} // hook your auth here
            >
              <LogOut className="h-5 w-5" />
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}
