import React, { useState, useEffect, useRef } from 'react';
import { Search, ChevronDown, Plus } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import CreateContestForm from './CreateContestForm';
import { motion, AnimatePresence } from 'framer-motion';

const HomePage = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState<any[]>([]);
  const [searching, setSearching] = useState(false);
  const [searchError, setSearchError] = useState('');
  const debounceRef = useRef<NodeJS.Timeout | null>(null);
  const navigate = useNavigate();

  const categories = [
    { name: 'Smart Contracts', articles: '1,205+', code: 'smart-contracts' },
    { name: 'DeFi Contests', articles: '856+', code: 'defi' },
    { name: 'NFT Challenges', articles: '945+', code: 'nft' },
    { name: 'Web3 Development', articles: '1,120+', code: 'web3' },
    { name: 'Security Contests', articles: '780+', code: 'security' },
    { name: 'Blockchain Games', articles: '650+', code: 'games' },
    { name: 'DAO Governance', articles: '430+', code: 'dao' },
    { name: 'Layer 2 Solutions', articles: '385+', code: 'layer2' },
    { name: 'Cross-chain Tech', articles: '290+', code: 'cross-chain' },
    { name: 'Zero Knowledge', articles: '245+', code: 'zk' },
  ];

  // Search as you type (debounce)
  useEffect(() => {
    if (!searchTerm.trim()) {
      setSearchResults([]);
      setSearchError('');
      return;
    }
    setSearching(true);
    setSearchError('');
    if (debounceRef.current) clearTimeout(debounceRef.current);
    debounceRef.current = setTimeout(async () => {
      try {
        const res = await fetch(`http://localhost:8080/api/v1/contests/search?keyword=${encodeURIComponent(searchTerm)}`);
        const data = await res.json();
        if (res.ok && data.success) {
          setSearchResults(data.data || []);
          if ((data.data || []).length === 0) setSearchError('No contests found.');
        } else {
          setSearchError(data.message || data.error || 'Search failed.');
        }
      } catch (err) {
        setSearchError('Cannot connect to backend.');
      }
      setSearching(false);
    }, 400); // debounce 400ms
    // eslint-disable-next-line
  }, [searchTerm]);

  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <header className="flex justify-between items-center p-4">
        <div className="flex items-center gap-4">
          <CreateContestForm>
            <Button 
              variant="outline" 
              size="default"
              className="flex items-center gap-2 hover:bg-blue-50 hover:border-blue-300 transition-all duration-200"
            >
              <Plus className="w-4 h-4" />
              Create New Contest
            </Button>
          </CreateContestForm>
        </div>
        <div className="flex gap-4">
          <button className="text-blue-600 hover:text-blue-800 font-medium">Log in</button>
          <button className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors">
            Sign up
          </button>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex flex-col items-center justify-center px-4 py-12">
        {/* Logo and Title */}
        <div className="text-center mb-12">
          <div className="w-32 h-32 mx-auto mb-6 relative">
            <div className="w-full h-full bg-gray-100 rounded-full flex items-center justify-center border-2 border-gray-200">
              <div className="text-6xl font-bold text-gray-400">C</div>
            </div>
          </div>
          <h1 className="text-4xl font-serif font-bold text-gray-900 mb-2">CONTEST WIKI</h1>
          <p className="text-lg text-gray-600">The Free Contest Encyclopedia</p>
        </div>

        {/* Languages Grid */}
        <div className="grid grid-cols-2 md:grid-cols-5 gap-8 mb-12 max-w-4xl w-full">
          {categories.map((category, index) => (
            <Link
              key={category.code}
              to="/article"
              className="text-center hover:bg-gray-50 p-4 rounded-lg transition-colors"
            >
              <div className="text-blue-600 hover:text-blue-800 font-medium text-lg mb-1">
                {category.name}
              </div>
              <div className="text-sm text-gray-500">{category.articles} articles</div>
            </Link>
          ))}
        </div>

        {/* Search Bar */}
        <div className="w-full max-w-2xl mb-8">
          <div className="flex rounded-lg border border-gray-300 overflow-hidden shadow-sm">
            <div className="relative">
              <button className="flex items-center gap-2 px-4 py-3 bg-gray-50 border-r border-gray-300 hover:bg-gray-100 transition-colors">
                <span className="text-sm font-medium">EN</span>
                <ChevronDown className="w-4 h-4" />
              </button>
            </div>
            <input
              type="text"
              placeholder="Search Contest Wiki"
              className="flex-1 px-4 py-3 outline-none"
              value={searchTerm}
              onChange={e => setSearchTerm(e.target.value)}
              onFocus={() => window.scrollTo({ top: 0, behavior: 'smooth' })}
            />
            <button
              className="px-6 py-3 bg-blue-600 text-white hover:bg-blue-700 transition-colors"
              disabled={searching}
            >
              <Search className="w-5 h-5" />
            </button>
          </div>
          {/* Hiển thị kết quả tìm kiếm với animation */}
          <AnimatePresence>
            {searching && (
              <motion.div
                key="searching"
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: 10 }}
                className="mt-4 text-blue-600"
              >
                Searching...
              </motion.div>
            )}
            {searchError && (
              <motion.div
                key="error"
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: 10 }}
                className="mt-4 text-red-600"
              >
                {searchError}
              </motion.div>
            )}
            {searchResults.length > 0 && (
              <motion.div
                key="results"
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: 10 }}
                className="mt-4 bg-gray-50 rounded-lg p-4 border border-gray-200"
              >
                <div className="font-semibold mb-2">Search Results:</div>
                <ul className="space-y-2">
                  {searchResults.slice(0, 5).map((contest, idx) => (
                    <motion.li
                      key={contest.id || idx}
                      initial={{ opacity: 0, x: 20 }}
                      animate={{ opacity: 1, x: 0 }}
                      exit={{ opacity: 0, x: 20 }}
                      className="p-2 bg-white rounded shadow-sm border border-gray-100 cursor-pointer hover:bg-blue-50 transition-colors"
                      onClick={() => navigate(`/contest/${contest.id}`)}
                    >
                      <div className="font-bold text-blue-700">{contest.name}</div>
                      <div className="text-sm text-gray-600">{contest.description}</div>
                      <div className="text-xs text-gray-400">Start: {contest.start_date || contest.startDate}</div>
                      <div className="text-xs text-gray-400">End: {contest.end_date || contest.endDate}</div>
                    </motion.li>
                  ))}
                </ul>
              </motion.div>
            )}
          </AnimatePresence>
        </div>

        {/* Language Selector */}
        <div className="mb-16">
          <button className="flex items-center gap-2 text-blue-600 hover:text-blue-800 font-medium">
            Filter by Blockchain Platform
            <ChevronDown className="w-4 h-4" />
          </button>
        </div>
      </main>

      {/* Footer */}
      <footer className="bg-gray-50 py-8 px-4">
        <div className="max-w-6xl mx-auto">
          <div className="flex flex-wrap justify-center gap-8 mb-6">
            <div className="flex items-center gap-3">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
                <div className="w-6 h-6 bg-blue-600 rounded"></div>
              </div>
              <div>
                <div className="font-medium text-blue-600">Contests</div>
                <div className="text-sm text-gray-600">Active competitions</div>
              </div>
            </div>
            <div className="flex items-center gap-3">
              <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center">
                <div className="w-6 h-6 bg-green-600 rounded"></div>
              </div>
              <div>
                <div className="font-medium text-green-600">Contestants</div>
                <div className="text-sm text-gray-600">Verified participants</div>
              </div>
            </div>
            <div className="flex items-center gap-3">
              <div className="w-12 h-12 bg-yellow-100 rounded-lg flex items-center justify-center">
                <div className="w-6 h-6 bg-yellow-600 rounded"></div>
              </div>
              <div>
                <div className="font-medium text-yellow-600">Sponsors</div>
                <div className="text-sm text-gray-600">Contest supporters</div>
              </div>
            </div>
          </div>
          <div className="text-center text-sm text-gray-500">
            Contest Wiki is powered by blockchain technology for transparent contest management.
          </div>
        </div>
      </footer>
    </div>
  );
};

export default HomePage;
