
import React from 'react';
import { useParams, Link } from 'react-router-dom';
import { ArrowLeft, Calendar, MapPin, Users, Facebook, Twitter, Instagram } from 'lucide-react';
import { Button } from '@/components/ui/button';
import Navigation from '@/components/Navigation';

const EventDetail = () => {
  const { id } = useParams();

  // Sample event data - in a real app, this would come from an API or state management
  const eventData = {
    1: {
      title: "THE MC FACE",
      subtitle: "THANH ÂM KHỞI SỰ, NGÂN VANG MỘT HÀNH TRÌNH",
      description: "THE MC FACE là sân chơi bổ ích và thiết thực dành riêng cho sinh viên Trường Đại học Văn Lang, nhằm tìm kiếm và phát triển tài năng người dẫn chương trình. Chương trình công kết quả trình học hội kiến thức và kỹ năng dẫn chương trình của các thành viên CLB, đồng thời thúc đẩy tinh thần thi đua học tập, khuyến khích sự nỗ lực và cống hiến của các cá nhân trong các hoạt động chung. Đây cũng là dịp để vinh danh các cá nhân, tập thể có thành tích nổi bật, đóng góp tạo tin thần nồi lực và cổng hiến trong năm học.",
      image: "/lovable-uploads/c4fee725-14fe-4194-8f65-59c6621ca00c.png",
      date: "Từ 29.06.2025",
      location: "Trường Đại học Văn Lang",
      category: "Competition",
      organizer: "CLB MC Văn Lang"
    },
    // Default fallback for other event IDs
    default: {
      title: "Cuộc thi thiết kế Board Game Việt Nam",
      subtitle: "SÁNG TẠO - ĐỔI MỚI - PHÁT TRIỂN",
      description: "Cuộc thi Road To ESSEN 2025 - Cuộc thi thiết kế Boardgame Việt Nam. Là sân chơi dành riêng cho các nhà thiết kế game tại Việt Nam, nhằm tìm kiếm và phát triển những ý tưởng sáng tạo trong lĩnh vực board game. Cuộc thi không chỉ là nơi thể hiện tài năng mà còn là cơ hội để các nhà thiết kế trẻ được giao lưu, học hỏi và phát triển kỹ năng chuyên môn.",
      image: "/lovable-uploads/2b37505b-5ad0-4227-8483-401b77d6f335.png",
      date: "Từ 24.06.2025",
      location: "Hà Nội",
      category: "Competition",
      organizer: "Ban tổ chức ESSEN Vietnam"
    }
  };

  const event = eventData[id as keyof typeof eventData] || eventData.default;

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-800">
      <Navigation />
      
      {/* Back Button */}
      <div className="max-w-7xl mx-auto px-4 pt-8">
        <Link to="/" className="inline-flex items-center text-blue-300 hover:text-blue-200 transition-colors mb-6">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Quay lại
        </Link>
      </div>

      {/* Hero Section with Event Banner */}
      <section className="relative py-16 px-4">
        <div className="max-w-7xl mx-auto">
          <div className="relative h-96 rounded-lg overflow-hidden mb-8">
            <img
              src={event.image}
              alt={event.title}
              className="w-full h-full object-cover"
            />
            <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>
            <div className="absolute bottom-0 left-0 right-0 p-8">
              <div className="text-center">
                <h1 className="text-4xl md:text-6xl font-bold text-white mb-4 tracking-tight">
                  {event.title}
                </h1>
                <p className="text-xl text-blue-200 font-light">
                  {event.subtitle}
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Event Details Section */}
      <section className="py-12 px-4">
        <div className="max-w-4xl mx-auto">
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-12">
            {/* Main Content */}
            <div className="lg:col-span-2">
              <div className="bg-slate-800/50 rounded-lg p-8 border border-slate-700">
                <h2 className="text-2xl font-bold text-white mb-6">GIỚI THIỆU</h2>
                <p className="text-gray-300 leading-relaxed text-lg">
                  {event.description}
                </p>
                
                <div className="mt-8">
                  <h3 className="text-xl font-bold text-blue-300 mb-4">THU GỌN</h3>
                  <Button 
                    variant="outline" 
                    className="bg-transparent text-blue-300 border-blue-600 hover:bg-blue-600/20 hover:border-blue-500"
                  >
                    Xem thêm thông tin
                  </Button>
                </div>
              </div>
            </div>

            {/* Sidebar */}
            <div className="space-y-6">
              {/* Event Info Card */}
              <div className="bg-slate-800/50 rounded-lg p-6 border border-slate-700">
                <h3 className="text-xl font-bold text-blue-300 mb-4">THÔNG TIN SỰ KIỆN</h3>
                <div className="space-y-4">
                  <div className="flex items-center text-gray-300">
                    <Calendar className="w-5 h-5 mr-3 text-blue-400" />
                    <span>{event.date}</span>
                  </div>
                  <div className="flex items-center text-gray-300">
                    <MapPin className="w-5 h-5 mr-3 text-blue-400" />
                    <span>{event.location}</span>
                  </div>
                  <div className="flex items-center text-gray-300">
                    <Users className="w-5 h-5 mr-3 text-blue-400" />
                    <span>{event.organizer}</span>
                  </div>
                </div>
              </div>

              {/* Organization Card */}
              <div className="bg-slate-800/50 rounded-lg p-6 border border-slate-700">
                <h3 className="text-xl font-bold text-blue-300 mb-4">BAN TỔ CHỨC</h3>
                <p className="text-gray-300 mb-4">{event.organizer}</p>
                <div className="flex space-x-3">
                  <Button size="sm" className="bg-blue-600 hover:bg-blue-700 text-white">
                    <Facebook className="w-4 h-4 mr-2" />
                    Facebook
                  </Button>
                </div>
              </div>

              {/* Action Buttons */}
              <div className="space-y-3">
                <Button className="w-full bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold py-3">
                  Đăng ký tham gia
                </Button>
                <Button 
                  variant="outline" 
                  className="w-full bg-transparent text-blue-300 border-blue-600 hover:bg-blue-600/20 hover:border-blue-500"
                >
                  Chia sẻ sự kiện
                </Button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export default EventDetail;
