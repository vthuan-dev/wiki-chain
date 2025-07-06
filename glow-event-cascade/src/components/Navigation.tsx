import React, { useState } from 'react';
import { Search, ChevronDown, Menu, X } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

const Navigation = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isLanguageOpen, setIsLanguageOpen] = useState(false);

  // State cho search
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState<any[]>([]);
  const [showResults, setShowResults] = useState(false);

  // Hàm gọi API search
  const handleSearch = async (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && searchTerm.trim() !== '') {
      try {
        const res = await fetch(
          `http://localhost:8081/api/v1/contests/search?keyword=${encodeURIComponent(searchTerm)}`
        );
        const data = await res.json();
        setSearchResults(data.data || []);
        setShowResults(true);
      } catch (error) {
        setSearchResults([]);
        setShowResults(true);
      }
    }
  };

  // Đóng kết quả khi blur
  const handleBlur = () => {
    setTimeout(() => setShowResults(false), 200);
  };

  return (
    <nav className="sticky top-0 z-50 bg-slate-900/95 backdrop-blur-md border-b border-slate-700">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          {/* Left Section */}
          <div className="flex items-center space-x-8">
            {/* Logo/Brand */}
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-gradient-to-r from-yellow-400 to-orange-500 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold text-sm">E</span>
              </div>
              <span className="text-white font-bold text-lg hidden sm:block">EventHub</span>
            </div>

            {/* Desktop Navigation */}
            <div className="hidden md:flex items-center space-x-6">
              {/* Language Selector */}
              <div className="relative">
                <button
                  onClick={() => setIsLanguageOpen(!isLanguageOpen)}
                  className="flex items-center space-x-2 text-gray-300 hover:text-white transition-colors"
                >
                  <span className="text-sm">🇻🇳</span>
                  <span className="text-sm">Việt Nam</span>
                  <ChevronDown className="w-4 h-4" />
                </button>
                
                {isLanguageOpen && (
                  <div className="absolute top-full left-0 mt-2 w-32 bg-slate-800 rounded-md shadow-lg border border-slate-700 z-10">
                    <div className="py-1">
                      <button className="w-full px-3 py-2 text-left text-sm text-gray-300 hover:bg-slate-700 hover:text-white flex items-center space-x-2">
                        <span>🇻🇳</span>
                        <span>Việt Nam</span>
                      </button>
                      <button className="w-full px-3 py-2 text-left text-sm text-gray-300 hover:bg-slate-700 hover:text-white flex items-center space-x-2">
                        <span>🇺🇸</span>
                        <span>English</span>
                      </button>
                    </div>
                  </div>
                )}
              </div>

              {/* Navigation Links */}
              <a href="#" className="text-gray-300 hover:text-white transition-colors text-sm font-medium">
                TẤT CẢ SỰ KIỆN
              </a>
              <a href="#" className="text-gray-300 hover:text-white transition-colors text-sm font-medium">
                SỰ KIỆN SẮP TỚI
              </a>
            </div>
          </div>

          {/* Right Section */}
          <div className="flex items-center space-x-4">
            {/* Featured Posts Button */}
            <Button 
              variant="outline" 
              className="hidden sm:flex bg-transparent text-blue-300 border-blue-600 hover:bg-blue-600/20 hover:border-blue-500"
            >
              BÀI VIẾT NỔI BẬT
            </Button>

            {/* Search - Desktop */}
            <div className="hidden lg:block relative">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-4 h-4" />
                <Input
                  type="text"
                  placeholder="Tìm kiếm"
                  className="w-64 pl-10 bg-slate-800/50 border-slate-600 text-white placeholder-gray-400 focus:border-blue-400 focus:ring-blue-400/20"
                  value={searchTerm}
                  onChange={e => setSearchTerm(e.target.value)}
                  onKeyDown={handleSearch}
                  onFocus={() => setShowResults(true)}
                  onBlur={handleBlur}
                />
              </div>
              {/* Hiển thị kết quả tìm kiếm */}
              {showResults && searchResults.length > 0 && (
                <div className="absolute left-0 mt-2 w-96 bg-white rounded shadow-lg z-20 max-h-80 overflow-y-auto">
                  {searchResults.map((item, idx) => (
                    <div key={item.id || idx} className="p-3 border-b border-gray-200 hover:bg-gray-100 cursor-pointer">
                      <div className="font-semibold text-gray-800">{item.name}</div>
                      <div className="text-sm text-gray-500">{item.description}</div>
                    </div>
                  ))}
                </div>
              )}
              {showResults && searchResults.length === 0 && (
                <div className="absolute left-0 mt-2 w-96 bg-white rounded shadow-lg z-20 p-3 text-gray-500">
                  Không tìm thấy kết quả.
                </div>
              )}
            </div>

            {/* Mobile Menu Button */}
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="md:hidden text-gray-300 hover:text-white"
            >
              {isMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
            </button>
          </div>
        </div>

        {/* Mobile Menu */}
        {isMenuOpen && (
          <div className="md:hidden border-t border-slate-700 py-4">
            <div className="space-y-4">
              <a href="#" className="block text-gray-300 hover:text-white transition-colors text-sm font-medium">
                TẤT CẢ SỰ KIỆN
              </a>
              <a href="#" className="block text-gray-300 hover:text-white transition-colors text-sm font-medium">
                SỰ KIỆN SẮP TỚI
              </a>
              <Button 
                variant="outline" 
                className="w-full bg-transparent text-blue-300 border-blue-600 hover:bg-blue-600/20 hover:border-blue-500"
              >
                BÀI VIẾT NỔI BẬT
              </Button>
              
              {/* Mobile Search */}
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-4 h-4" />
                <Input
                  type="text"
                  placeholder="Tìm kiếm"
                  className="w-full pl-10 bg-slate-800/50 border-slate-600 text-white placeholder-gray-400 focus:border-blue-400 focus:ring-blue-400/20"
                />
              </div>
            </div>
          </div>
        )}
      </div>
    </nav>
  );
};

export default Navigation;
