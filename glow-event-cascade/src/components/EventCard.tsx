
import React from 'react';
import { Calendar } from 'lucide-react';
import { Card, CardContent } from '@/components/ui/card';
import { Link } from 'react-router-dom';

interface Event {
  id: number;
  title: string;
  description: string;
  image: string;
  date: string;
  category: string;
}

interface EventCardProps {
  event: Event;
}

const EventCard: React.FC<EventCardProps> = ({ event }) => {
  return (
    <Card className="group bg-slate-800/50 border-slate-700 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-2 hover:shadow-2xl hover:shadow-blue-500/20 overflow-hidden">
      <div className="aspect-video overflow-hidden">
        <img
          src={event.image}
          alt={event.title}
          className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        />
      </div>
      
      <CardContent className="p-6">
        <div className="mb-3">
          <span className="inline-block px-3 py-1 text-xs font-medium bg-blue-600/20 text-blue-300 rounded-full">
            {event.category}
          </span>
        </div>
        
        <h3 className="text-xl font-bold text-white mb-3 group-hover:text-blue-300 transition-colors line-clamp-2">
          {event.title}
        </h3>
        
        <p className="text-gray-400 text-sm mb-4 line-clamp-2 leading-relaxed">
          {event.description}
        </p>
        
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2 text-gray-400">
            <Calendar className="w-4 h-4" />
            <span className="text-sm">{event.date}</span>
          </div>
          
          <Link 
            to={`/event/${event.id}`}
            className="text-blue-400 hover:text-blue-300 text-sm font-medium transition-colors"
          >
            Xem chi tiết →
          </Link>
        </div>
      </CardContent>
    </Card>
  );
};

export default EventCard;
