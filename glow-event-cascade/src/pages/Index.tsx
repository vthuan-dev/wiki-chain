import { useState, useEffect } from 'react';
import { Calendar, Search, ChevronDown, Star, Plus, Activity } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { toast, Toaster } from 'sonner';
import EventCard from '@/components/EventCard';
import Navigation from '@/components/Navigation';
import { blockchainApi, Contest, CreateContestRequest, BlockchainStats } from '@/services/blockchainApi';

const Index = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [activeFilter, setActiveFilter] = useState('Tất cả');
  
  // Blockchain states
  const [contests, setContests] = useState<Contest[]>([]);
  const [stats, setStats] = useState<BlockchainStats | null>(null);
  const [loading, setLoading] = useState(true);
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [dataInitialized, setDataInitialized] = useState(false);
  
  // Form states for creating contest
  const [formData, setFormData] = useState<CreateContestRequest>({
    name: '',
    description: '',
    start_date: '',
    end_date: '',
    image_url: ''
  });

  // Load data from blockchain
  useEffect(() => {
    if (!dataInitialized) {
      loadBlockchainData();
      setDataInitialized(true);
    }
  }, [dataInitialized]);

  const loadBlockchainData = async () => {
    try {
      setLoading(true);
      
      // Load contests and stats in parallel
      const [contestsResponse, statsResponse] = await Promise.all([
        blockchainApi.getAllContests(),
        blockchainApi.getStats()
      ]);

      let dataLoaded = false;  // Bắt đầu là false, chỉ set true khi thực sự có dữ liệu

      if (contestsResponse.success && contestsResponse.data) {
        console.log('📊 Got contests data:', contestsResponse.data);
        setContests(contestsResponse.data || []);
        dataLoaded = true;
      }

      if (statsResponse.success) {
        setStats(statsResponse.data);
        dataLoaded = true;
      }

      // Chỉ hiển thị thông báo khi thực sự có dữ liệu được tải
      // if (dataLoaded) {
      //   toast.success('✅ Đã tải dữ liệu từ blockchain thành công!', {
      //     id: 'blockchain-data-loaded', // Thêm id để tránh trùng lặp
      //   });
      // }
    } catch (error) {
      console.error('Error loading blockchain data:', error);
      toast.error('❌ Lỗi khi tải dữ liệu từ blockchain', {
        id: 'blockchain-data-error', // Thêm id để tránh trùng lặp
      });
    } finally {
      setLoading(false);
    }
  };

  const handleCreateContest = async () => {
    try {
      if (!formData.name || !formData.description || !formData.start_date || !formData.end_date) {
        toast.error('Vui lòng điền đầy đủ thông tin bắt buộc');
        return;
      }

      // Validate dates
      const startDate = new Date(formData.start_date);
      const endDate = new Date(formData.end_date);
      
      if (endDate <= startDate) {
        toast.error('Ngày kết thúc phải sau ngày bắt đầu');
        return;
      }

      // Convert to ISO string format
      const contestData: CreateContestRequest = {
        ...formData,
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
      };

      const response = await blockchainApi.createContest(contestData);
      
      if (response.success) {
        // toast.success(`✅ Tạo cuộc thi thành công! TX: ${response.tx_hash}`);
        setIsCreateDialogOpen(false);
        setFormData({
          name: '',
          description: '',
          start_date: '',
          end_date: '',
          image_url: ''
        });
        
        // Reload data
        await loadBlockchainData();
      }
    } catch (error) {
      console.error('Error creating contest:', error);
      toast.error(`❌ Lỗi tạo cuộc thi: ${error instanceof Error ? error.message : 'Unknown error'}`);
    }
  };

  // Convert blockchain contests to display format
  const displayEvents = contests && contests.length > 0 ? contests.map(contest => ({
    id: parseInt(contest.id.slice(-6), 16), // Use last 6 chars of hex ID as number
    title: contest.name,
    description: contest.description,
    image: contest.image_url || "/placeholder.svg",
    date: new Date(contest.start_date).toLocaleDateString('vi-VN'),
    category: contest.active ? "Active Contest" : "Ended Contest"
  })) : [];

  const filters = ['Tất cả', 'Gần đây nhất', 'Phổ biến nhất', 'Mới nhất'];

  // Chỉ hiển thị dữ liệu blockchain nếu có, không còn card mock/demo
  const allEvents = displayEvents.length > 0 
    ? [...displayEvents]
    : [];

  const filteredEvents = allEvents.filter(event =>
    event.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
    event.description.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-800">
      {/* Navigation */}
      <Navigation />

      {/* Hero Section */}
      <section className="relative py-24 px-4 text-center">
        <div className="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 backdrop-blur-sm"></div>
        <div className="relative z-10 max-w-4xl mx-auto">
          <div className="mb-8">
            <div className="w-16 h-16 mx-auto mb-6 bg-gradient-to-r from-yellow-400 to-orange-500 rounded-lg flex items-center justify-center">
              <Star className="w-8 h-8 text-white" />
            </div>
          </div>
          <h1 className="text-5xl md:text-6xl font-bold text-white mb-4 tracking-tight">
            TẤT CẢ SỰ KIỆN
          </h1>
          <p className="text-xl text-blue-200 mb-8 font-light">
            Nơi kết nối, lan toả và toả sáng
          </p>
          
          {/* Blockchain Stats */}
          {stats && (
            <div className="grid grid-cols-2 md:grid-cols-5 gap-4 mb-8 max-w-2xl mx-auto">
              <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4">
                <div className="text-2xl font-bold text-white">{stats.totalContests}</div>
                <div className="text-blue-200 text-sm">Cuộc thi</div>
              </div>
              <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4">
                <div className="text-2xl font-bold text-white">{stats.totalContestants}</div>
                <div className="text-blue-200 text-sm">Thí sinh</div>
              </div>
              <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4">
                <div className="text-2xl font-bold text-white">{stats.totalSponsors}</div>
                <div className="text-blue-200 text-sm">Nhà tài trợ</div>
              </div>
              <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4">
                <div className="text-2xl font-bold text-white">{stats.totalRegistrations}</div>
                <div className="text-blue-200 text-sm">Đăng ký</div>
              </div>
              <div className="bg-white/10 backdrop-blur-sm rounded-lg p-4">
                <div className="text-2xl font-bold text-white">{stats.totalContents}</div>
                <div className="text-blue-200 text-sm">Nội dung</div>
              </div>
            </div>
          )}
          
          {/* Create Contest Button */}
          <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white px-8 py-3 rounded-lg font-medium transition-all duration-300 hover:scale-105">
                <Plus className="w-5 h-5 mr-2" />
                Tạo cuộc thi mới trên Blockchain
              </Button>
            </DialogTrigger>
            <DialogContent className="bg-slate-900 border-slate-700 text-white max-w-2xl">
              <DialogHeader>
                <DialogTitle className="text-2xl font-bold text-center mb-4">
                  🎯 Tạo cuộc thi mới trên Blockchain
                </DialogTitle>
              </DialogHeader>
              
              <div className="space-y-6">
                <div className="grid grid-cols-1 gap-4">
                  <div>
                    <Label htmlFor="name" className="text-blue-200">Tên cuộc thi *</Label>
                    <Input
                      id="name"
                      value={formData.name}
                      onChange={(e) => setFormData(prev => ({ ...prev, name: e.target.value }))}
                      placeholder="Nhập tên cuộc thi..."
                      className="bg-slate-800 border-slate-600 text-white"
                    />
                  </div>
                  
                  <div>
                    <Label htmlFor="description" className="text-blue-200">Mô tả *</Label>
                    <Textarea
                      id="description"
                      value={formData.description}
                      onChange={(e) => setFormData(prev => ({ ...prev, description: e.target.value }))}
                      placeholder="Mô tả chi tiết về cuộc thi..."
                      className="bg-slate-800 border-slate-600 text-white min-h-[100px]"
                    />
                  </div>
                  
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <Label htmlFor="start_date" className="text-blue-200">Ngày bắt đầu *</Label>
                      <Input
                        id="start_date"
                        type="datetime-local"
                        value={formData.start_date}
                        onChange={(e) => setFormData(prev => ({ ...prev, start_date: e.target.value }))}
                        className="bg-slate-800 border-slate-600 text-white"
                      />
                    </div>
                    
                    <div>
                      <Label htmlFor="end_date" className="text-blue-200">Ngày kết thúc *</Label>
                      <Input
                        id="end_date"
                        type="datetime-local"
                        value={formData.end_date}
                        onChange={(e) => setFormData(prev => ({ ...prev, end_date: e.target.value }))}
                        className="bg-slate-800 border-slate-600 text-white"
                      />
                    </div>
                  </div>
                  
                  <div>                      <Label htmlFor="image_url" className="text-blue-200">URL hình ảnh</Label>
                      <Input
                        id="image_url"
                        value={formData.image_url}
                        onChange={(e) => setFormData(prev => ({ ...prev, image_url: e.target.value }))}
                        placeholder="https://example.com/image.jpg"
                        className="bg-slate-800 border-slate-600 text-white"
                      />
                  </div>
                </div>
                
                <div className="flex gap-4 pt-4">
                  <Button
                    onClick={() => setIsCreateDialogOpen(false)}
                    className="flex-1 bg-slate-700 hover:bg-slate-600"
                  >
                    Hủy
                  </Button>
                  <Button
                    onClick={handleCreateContest}
                    className="flex-1 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700"
                  >
                    <Activity className="w-4 h-4 mr-2" />
                    Tạo trên Blockchain
                  </Button>
                </div>
              </div>
            </DialogContent>
          </Dialog>
        </div>
      </section>

      {/* Events Section */}
      <section className="py-16 px-4">
        <div className="max-w-7xl mx-auto">
          {/* Section Header */}
          <div className="text-center mb-12">
            <div className="flex items-center justify-center gap-4 mb-8">
              <div className="h-px bg-gradient-to-r from-transparent via-blue-400 to-transparent flex-1"></div>
              <h2 className="text-2xl font-bold text-blue-300 px-4">
                10 sự kiện đang diễn ra
              </h2>
              <div className="h-px bg-gradient-to-r from-transparent via-blue-400 to-transparent flex-1"></div>
            </div>

            {/* Search Bar */}
            <div className="max-w-md mx-auto mb-8">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                <Input
                  type="text"
                  placeholder="Tìm kiếm"
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="pl-10 bg-slate-800/50 border-slate-600 text-white placeholder-gray-400 focus:border-blue-400 focus:ring-blue-400/20"
                />
              </div>
            </div>

            {/* Filters */}
            <div className="flex flex-wrap justify-center gap-2 mb-8">
              {filters.map((filter) => (
                <Button
                  key={filter}
                  onClick={() => setActiveFilter(filter)}
                  className={`
                    ${activeFilter === filter 
                      ? 'bg-blue-600 text-white border-blue-600' 
                      : 'bg-transparent text-blue-300 border-slate-600 hover:bg-blue-600/20 hover:border-blue-500'
                    }
                    transition-all duration-200 border
                  `}
                >
                  {filter}
                </Button>
              ))}
            </div>
          </div>

          {/* Events Grid */}
          {loading ? (
            <div className="flex justify-center items-center py-20">
              <div className="text-center">
                <Activity className="w-12 h-12 text-blue-400 animate-spin mx-auto mb-4" />
                <p className="text-blue-200">Đang tải dữ liệu từ blockchain...</p>
              </div>
            </div>
          ) : (
            allEvents.length === 0 ? (
              <div className="text-blue-200 text-lg mt-8">Chưa có sự kiện nào trên blockchain.</div>
            ) : (
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                {filteredEvents.map((event) => (
                  <EventCard key={event.id} event={event} />
                ))}
                {filteredEvents.length === 0 && (
                  <div className="col-span-full text-center py-20">
                    <p className="text-blue-200 text-lg">Không tìm thấy sự kiện nào</p>
                  </div>
                )}
              </div>
            )
          )}
        </div>
      </section>
      <Toaster richColors position="top-right" />
    </div>
  );
};

export default Index;
