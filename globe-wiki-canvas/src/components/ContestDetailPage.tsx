import React, { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { Badge } from '@/components/ui/badge';

const formatDate = (iso: string) => {
  if (!iso) return '';
  const d = new Date(iso);
  return d.toLocaleString('vi-VN', { dateStyle: 'medium', timeStyle: 'short' });
};

const ContestDetailPage = () => {
  const { id } = useParams();
  const [contest, setContest] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  
  // Check if contest is active based on end date
  const isContestActive = (endDate: string) => {
    const currentDate = new Date();
    const contestEndDate = new Date(endDate);
    return currentDate < contestEndDate;
  };

  useEffect(() => {
    const fetchContest = async () => {
      setLoading(true);
      setError('');
      try {
        const res = await fetch(`http://localhost:8080/api/v1/contests/${id}`);
        const data = await res.json();
        if (res.ok && data.success && data.data) {
          setContest(data.data);
        } else {
          setError(data.message || data.error || 'Not found');
        }
      } catch (err) {
        setError('Cannot connect to backend.');
      }
      setLoading(false);
    };
    fetchContest();
  }, [id]);

  if (loading) return <div className="p-8 text-blue-600">Loading...</div>;
  if (error) return <div className="p-8 text-red-600">{error}</div>;
  if (!contest) return null;

  return (
    <div className="min-h-screen bg-white text-gray-900">
      <header className="bg-white border-b border-gray-200 px-4 py-3">
        <div className="max-w-7xl mx-auto flex justify-between items-center">
          <Link to="/" className="flex items-center gap-2 group">
            <div className="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
              <span className="text-sm font-bold text-white">C</span>
            </div>
            <span className="font-serif font-bold text-lg group-hover:text-blue-600">CONTEST WIKI</span>
          </Link>
          <span className="text-sm text-gray-500">Decentralized Contest Platform</span>
        </div>
      </header>
      <div className="max-w-3xl mx-auto p-8">
        <Link to="/" className="text-blue-600 hover:underline">&larr; Quay lại trang chủ</Link>
        <div className="mt-6 bg-white rounded-lg shadow-lg p-6">
          {/* Banner */}
          {contest.image_url && (
            <img
              src={contest.image_url}
              alt={contest.name}
              className="w-full max-w-xl h-64 object-cover rounded-lg mb-6 mx-auto"
            />
          )}
          {/* Title & Status */}
          <div className="flex items-center gap-3 mb-2">
            <h1 className="text-2xl font-bold flex-1">{contest.name}</h1>
            <Badge className="bg-white border border-gray-200" variant="outline">
              <span className={isContestActive(contest.end_date || contest.endDate) ? 'text-green-600' : 'text-gray-600'}>
                {isContestActive(contest.end_date || contest.endDate) ? 'Đang mở' : 'Đã kết thúc'}
              </span>
            </Badge>
          </div>
          {/* Description */}
          <p className="mb-4 text-gray-700">{contest.description}</p>
          {/* Info Table */}
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm mb-4">
            <div>
              <span className="font-medium">Mã cuộc thi:</span> <span className="break-all">{contest.id}</span>
            </div>
            <div>
              <span className="font-medium">Tổ chức bởi:</span> <span className="break-all">{contest.organizer}</span>
            </div>
            <div>
              <span className="font-medium">Bắt đầu:</span> {formatDate(contest.start_date)}
            </div>
            <div>
              <span className="font-medium">Kết thúc:</span> {formatDate(contest.end_date)}
            </div>
            <div>
              <span className="font-medium">Timestamp:</span> {formatDate(contest.timestamp)}
            </div>
            <div>
              <span className="font-medium">Trạng thái:</span> {contest.active ? 'Active' : 'Inactive'}
            </div>
            {contest.tx_hash && (
              <div className="col-span-2">
                <span className="font-medium">Transaction: </span>
                <a 
                  href={`https://explorer.testnet.hii.network/tx/${contest.tx_hash}`}
                  target="_blank"
                  rel="noopener noreferrer" 
                  className="text-blue-600 hover:text-blue-800 hover:underline break-all"
                >
                  {contest.tx_hash}
                </a>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ContestDetailPage;
