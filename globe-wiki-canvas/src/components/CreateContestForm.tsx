import React, { useState } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, DialogClose } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Plus, Calendar, Image } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

interface CreateContestFormProps {
  children: React.ReactNode;
}

const CreateContestForm = ({ children }: CreateContestFormProps) => {
  const navigate = useNavigate();
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    id: '',
    name: '',
    description: '',
    startDate: '',
    endDate: '',
    imageURL: ''
  });

  const resetForm = () => {
    setFormData({
      id: '',
      name: '',
      description: '',
      startDate: '',
      endDate: '',
      imageURL: ''
    });
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Kiểm tra giá trị đầu vào
    if (!formData.startDate || !formData.endDate) {
      alert('Vui lòng nhập đầy đủ ngày bắt đầu và kết thúc!');
      return;
    }

    // Chuyển đổi sang ISO 8601 (RFC3339)
    const startDateISO = new Date(formData.startDate).toISOString();
    const endDateISO = new Date(formData.endDate).toISOString();

    // Log để kiểm tra
    console.log('startDate:', formData.startDate, '->', startDateISO);
    console.log('endDate:', formData.endDate, '->', endDateISO);

    const payload = {
      name: formData.name,
      description: formData.description,
      start_date: startDateISO,
      end_date: endDateISO,
      image_url: formData.imageURL,
    };

    try {
      const res = await fetch('http://localhost:8080/api/v1/contests', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
      const data = await res.json();
      if (res.ok && data.success) {
        alert('Tạo cuộc thi thành công! ID: ' + data.id);
        resetForm();
        setOpen(false); // Đóng dialog
        navigate(`/contest/${data.id}`); // Chuyển đến trang chi tiết contest
      } else {
        alert('Tạo cuộc thi thất bại: ' + (data.message || data.error));
      }
    } catch (err) {
      alert('Lỗi kết nối backend!');
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        {children}
      </DialogTrigger>
      <DialogContent className="sm:max-w-[600px] max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2 text-xl">
            <Plus className="w-5 h-5 text-blue-600" />
            Create New Contest
          </DialogTitle>
        </DialogHeader>
        
        <form onSubmit={handleSubmit} className="space-y-6 mt-4">
          {/* Contest Name */}
          <div className="space-y-2">
            <Label htmlFor="name" className="text-sm font-medium">
              Contest Name *
            </Label>
            <Input
              id="name"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              placeholder="e.g., Web3 Innovation Challenge"
              required
              className="w-full"
            />
          </div>

          {/* Contest Description */}
          <div className="space-y-2">
            <Label htmlFor="description" className="text-sm font-medium">
              Description *
            </Label>
            <Textarea
              id="description"
              name="description"
              value={formData.description}
              onChange={handleInputChange}
              placeholder="Describe your contest, rules, and objectives..."
              required
              className="w-full min-h-[100px]"
            />
          </div>

          {/* Start Date */}
          <div className="space-y-2">
            <Label htmlFor="startDate" className="text-sm font-medium flex items-center gap-2">
              <Calendar className="w-4 h-4" />
              Start Date *
            </Label>
            <Input
              id="startDate"
              name="startDate"
              type="datetime-local"
              value={formData.startDate}
              onChange={handleInputChange}
              required
              className="w-full"
            />
          </div>

          {/* End Date */}
          <div className="space-y-2">
            <Label htmlFor="endDate" className="text-sm font-medium flex items-center gap-2">
              <Calendar className="w-4 h-4" />
              End Date *
            </Label>
            <Input
              id="endDate"
              name="endDate"
              type="datetime-local"
              value={formData.endDate}
              onChange={handleInputChange}
              required
              className="w-full"
            />
          </div>

          {/* Image URL */}
          <div className="space-y-2">
            <Label htmlFor="imageURL" className="text-sm font-medium flex items-center gap-2">
              <Image className="w-4 h-4" />
              Contest Image URL
            </Label>
            <Input
              id="imageURL"
              name="imageURL"
              type="url"
              value={formData.imageURL}
              onChange={handleInputChange}
              placeholder="https://example.com/contest-image.jpg"
              className="w-full"
            />
            <p className="text-xs text-gray-500">Optional: URL to contest banner/logo image</p>
          </div>

          {/* Smart Contract Info */}
          <div className="bg-blue-50 p-4 rounded-lg border border-blue-200">
            <h4 className="font-medium text-blue-900 mb-2">Smart Contract Integration</h4>
            <p className="text-sm text-blue-700">
              This form will call your ContentStorage contract's <code className="bg-blue-100 px-1 rounded">createContest</code> function with the provided data.
            </p>
          </div>

          {/* Submit Button */}
          <div className="flex gap-3 pt-4">
            <Button 
              type="submit" 
              className="flex-1 bg-blue-600 hover:bg-blue-700 text-white"
            >
              Create Contest
            </Button>
            <Button 
              type="button" 
              variant="outline" 
              onClick={resetForm}
            >
              Reset Form
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export default CreateContestForm;
