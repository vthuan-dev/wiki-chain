// Blockchain API Service
const API_BASE_URL = 'http://localhost:8080/api/v1';

// Log API URL for debugging
console.log('🌐 API Base URL:', API_BASE_URL);

export interface CreateContentRequest {
  title: string;
  content: string;
  creator?: string;
}

export interface Content {
  id: string;
  title: string;
  content: string;
  creator: string;
  timestamp: string;
  verified: boolean;
  txHash?: string;
}

export interface Contest {
  id: string;
  name: string;
  description: string;
  start_date: string;  // Changed from startDate to match backend
  end_date: string;    // Changed from endDate to match backend
  organizer: string;
  active: boolean;
  image_url?: string;  // Changed from imageURL to match backend
  timestamp: string;
  tx_hash?: string;    // Changed from txHash to match backend
}

export interface CreateContestRequest {
  name: string;
  description: string;
  start_date: string;  // Changed from startDate to match backend
  end_date: string;    // Changed from endDate to match backend
  image_url?: string;  // Changed from imageURL to match backend
}

export interface BlockchainStats {
  totalContents: number;
  totalContests: number;
  totalContestants: number;
  totalSponsors: number;
  totalRegistrations: number;
}

class BlockchainApiService {
  // Health Check
  async healthCheck(): Promise<{ status: string; blockchain: string; message: string }> {
    console.log('🩺 Health check to:', `${API_BASE_URL}/health`);
    const response = await fetch(`${API_BASE_URL}/health`);
    console.log('🩺 Health check status:', response.status);
    
    if (!response.ok) {
      throw new Error('Health check failed');
    }
    return response.json();
  }

  // Content Operations
  async createContent(data: CreateContentRequest): Promise<{ success: boolean; message: string; txHash?: string; id: string }> {
    const response = await fetch(`${API_BASE_URL}/content`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to create content');
    }

    return response.json();
  }

  async getContent(id: string): Promise<{ success: boolean; data?: Content; message: string }> {
    const response = await fetch(`${API_BASE_URL}/content/${id}`);
    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to get content');
    }
    return response.json();
  }

  async getAllContents(): Promise<{ success: boolean; data: Content[]; total: number }> {
    const response = await fetch(`${API_BASE_URL}/contents`);
    if (!response.ok) {
      throw new Error('Failed to get contents');
    }
    return response.json();
  }

  // Contest Operations
  async createContest(data: CreateContestRequest): Promise<{ success: boolean; message: string; tx_hash?: string; id: string }> {
    console.log('🔄 Sending request to:', `${API_BASE_URL}/contests`); // Sửa từ contest sang contests (số nhiều)
    console.log('📋 Request data:', JSON.stringify(data, null, 2));
    
    const response = await fetch(`${API_BASE_URL}/contests`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    console.log('📡 Response status:', response.status);
    console.log('📡 Response headers:', Object.fromEntries(response.headers.entries()));
    
    const responseText = await response.text();
    console.log('📡 Raw response text:', responseText);

    if (!response.ok) {
      let error;
      try {
        error = JSON.parse(responseText);
      } catch (e) {
        console.error('❌ Failed to parse error response as JSON:', e);
        throw new Error(`HTTP ${response.status}: ${responseText || 'Unknown error'}`);
      }
      throw new Error(error.error || 'Failed to create contest');
    }

    try {
      const result = JSON.parse(responseText);
      console.log('✅ Parsed response:', result);
      return result;
    } catch (e) {
      console.error('❌ Failed to parse success response as JSON:', e);
      console.error('Raw response was:', responseText);
      console.error('Response status:', response.status);
      console.error('Response headers:', Object.fromEntries(response.headers.entries()));
      throw new Error(`Invalid JSON response from server: ${e.message}`);
    }
  }

  async getAllContests(): Promise<{ success: boolean; data: Contest[]; total: number }> {
    console.log('📋 Getting all contests from:', `${API_BASE_URL}/contests`);
    const response = await fetch(`${API_BASE_URL}/contests`);
    console.log('📋 Get contests status:', response.status);
    
    if (!response.ok) {
      throw new Error('Failed to get contests');
    }
    const result = await response.json();
    console.log('📋 Contests response:', result);
    return result;
  }

  async getContest(id: string): Promise<{ success: boolean; data?: Contest; message: string }> {
    const response = await fetch(`${API_BASE_URL}/contests/${id}`);
    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to get contest');
    }
    return response.json();
  }

  // Statistics
  async getStats(): Promise<{ success: boolean; data: BlockchainStats }> {
    const response = await fetch(`${API_BASE_URL}/stats`);
    if (!response.ok) {
      throw new Error('Failed to get statistics');
    }
    return response.json();
  }
}

export const blockchainApi = new BlockchainApiService();
