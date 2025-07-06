const fs = require('fs');
const path = require('path');

// Đọc contract artifacts sau khi deploy
const getDeployedContractAddress = () => {
  try {
    const buildPath = path.join(__dirname, '..', 'build', 'contracts', 'ContentStorage.json');
    const contractData = JSON.parse(fs.readFileSync(buildPath, 'utf8'));
    
    // Lấy address từ network 22988 (Hii Network)
    const networks = contractData.networks;
    const hiiNetwork = networks['22988'];
    
    if (hiiNetwork && hiiNetwork.address) {
      return hiiNetwork.address;
    }
    
    return null;
  } catch (error) {
    console.error('❌ Lỗi đọc contract address:', error.message);
    return null;
  }
};

// Cập nhật file .env trong Go backend
const updateGoBackendEnv = (contractAddress) => {
  try {
    // Đường dẫn đúng: từ backend/truffle/scripts lên 2 cấp rồi vào blockchain-demo-go
    const goBackendPath = path.join(__dirname, '..', '..', '..', 'blockchain-demo-go');
    
    // Kiểm tra và tạo thư mục nếu chưa tồn tại
    if (!fs.existsSync(goBackendPath)) {
      console.log(`❌ Không tìm thấy thư mục Go backend: ${goBackendPath}`);
      console.log('🔍 Kiểm tra lại cấu trúc thư mục...');
      return;
    }
    
    const envPath = path.join(goBackendPath, '.env');
    const envExamplePath = path.join(goBackendPath, '.env.example');
    
    // Tạo nội dung .env mới
    const envContent = `# Blockchain Demo Environment Configuration

# Hii Network Settings (Testnet)
NETWORK_URL=http://103.69.98.80:8545
CHAIN_ID=22988

# Contract Settings (Auto-updated from deployment)
CONTRACT_ADDRESS=${contractAddress}

# Private Key for transactions (SAME AS TRUFFLE DEPLOYMENT)
PRIVATE_KEY=e6b362527d55d4e0a8b8330eb32b0e8eafed357a5dc73e3f8024a4f5d424edd7

# Server Settings
PORT=8080
HOST=localhost
`;

    // Ghi file .env
    fs.writeFileSync(envPath, envContent);
    
    // Cập nhật .env.example
    const exampleContent = envContent.replace(
      /PRIVATE_KEY=.*/,
      'PRIVATE_KEY=your_private_key_here'
    ).replace(
      /CONTRACT_ADDRESS=0x.*/,
      'CONTRACT_ADDRESS=your_contract_address_here'
    );
    
    fs.writeFileSync(envExamplePath, exampleContent);
    
    console.log('✅ Đã cập nhật Go backend configuration');
    console.log(`📄 Contract Address: ${contractAddress}`);
    console.log(`📁 Updated files:`);
    console.log(`   - ${envPath}`);
    console.log(`   - ${envExamplePath}`);
    
  } catch (error) {
    console.error('❌ Lỗi cập nhật Go backend:', error.message);
  }
};

// Main function
const main = () => {
  console.log('🔄 Post-deployment: Updating Go backend configuration...');
  
  const contractAddress = getDeployedContractAddress();
  
  if (contractAddress) {
    console.log(`📋 Found deployed contract: ${contractAddress}`);
    updateGoBackendEnv(contractAddress);
    
    console.log('\n🎉 Configuration updated successfully!');
    console.log('\n📝 Next steps:');
    console.log('   1. cd ../../blockchain-demo-go');
    console.log('   2. go run cmd/server/main.go');
    console.log(`   3. API will connect to contract: ${contractAddress}`);
    
  } else {
    console.log('❌ Không tìm thấy contract address. Kiểm tra deployment.');
  }
};

main();
