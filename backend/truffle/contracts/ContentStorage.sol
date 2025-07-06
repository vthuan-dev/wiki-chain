// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ContentStorage {
    // Định nghĩa struct cho thí sinh
    struct Contestant {
        string id;
        string name;
        string details;
        address creator;
        uint256 timestamp;
        bool verified;
        bool exists;
    }
    
    // Định nghĩa struct cho cuộc thi
    struct Contest {
        string id;
        string name;
        string description;
        uint256 startDate;
        uint256 endDate;
        address organizer;
        bool active;
        bool exists;
        string imageURL;
    }
    
    // Định nghĩa struct cho nhà tài trợ
    struct Sponsor {
        string id;
        string name;
        string contactInfo;
        uint256 sponsorshipAmount;
        address walletAddress;
        bool exists;
    }
    
    // Định nghĩa struct cho nội dung wiki thông thường
    struct Content {
        string title;
        string content;
        address creator;
        uint256 timestamp;
        bool verified;
        bool exists;
    }
    
    // Mapping từ ID đến các loại dữ liệu
    mapping(string => Content) private contents;
    mapping(string => Contestant) private contestants;
    mapping(string => Contest) private contests;
    mapping(string => Sponsor) private sponsors;
    
    // Mapping để theo dõi thí sinh đăng ký tham gia cuộc thi
    mapping(string => mapping(string => bool)) private contestantRegistrations; // contestId => (contestantId => registered)
    
    // Mảng lưu tất cả ID để có thể lấy danh sách
    string[] private contentIds;
    string[] private contestantIds;
    string[] private contestIds;
    string[] private sponsorIds;
    
    // Events
    event ContentAdded(string indexed id, string title);
    event ContestantAdded(string indexed id, string name);
    event ContestCreated(string indexed id, string name);
    event SponsorAdded(string indexed id, string name);
    event ContestantRegistered(string indexed contestId, string indexed contestantId);
    
    constructor() {}
    
    // Lưu nội dung mới
    function storeContent(string memory id, string memory title, string memory contentText, bool verified) public {
        // Kiểm tra xem ID đã tồn tại chưa
        require(!contents[id].exists, "Content with this ID already exists");
        
        // Lưu nội dung mới
        contents[id] = Content({
            title: title,
            content: contentText,
            creator: msg.sender,
            timestamp: block.timestamp,
            verified: verified,
            exists: true
        });
        
        // Thêm ID vào danh sách
        contentIds.push(id);
        
        // Phát event
        emit ContentAdded(id, title);
    }
    
    // Lấy nội dung theo ID
    function getContent(string memory id) public view returns (
        string memory title,
        string memory content,
        address creator,
        uint256 timestamp,
        bool verified
    ) {
        // Kiểm tra xem nội dung có tồn tại không
        require(contents[id].exists, "Content does not exist");
        
        Content memory c = contents[id];
        return (c.title, c.content, c.creator, c.timestamp, c.verified);
    }
    
    // Lấy danh sách tất cả ID
    function getAllContentIds() public view returns (string[] memory) {
        return contentIds;
    }
    
    // QUẢN LÝ THÍ SINH
    
    // Thêm thí sinh mới
    function addContestant(
        string memory id, 
        string memory name, 
        string memory details, 
        bool verified
    ) public {
        require(!contestants[id].exists, "Contestant with this ID already exists");
        
        contestants[id] = Contestant({
            id: id,
            name: name,
            details: details,
            creator: msg.sender,
            timestamp: block.timestamp,
            verified: verified,
            exists: true
        });
        
        contestantIds.push(id);
        emit ContestantAdded(id, name);
    }
    
    // Lấy thông tin thí sinh theo ID
    function getContestant(string memory id) public view returns (
        string memory name,
        string memory details,
        address creator,
        uint256 timestamp,
        bool verified
    ) {
        require(contestants[id].exists, "Contestant does not exist");
        
        Contestant memory c = contestants[id];
        return (c.name, c.details, c.creator, c.timestamp, c.verified);
    }
    
    // Lấy danh sách tất cả ID thí sinh
    function getAllContestantIds() public view returns (string[] memory) {
        return contestantIds;
    }
    
    // QUẢN LÝ CUỘC THI
    
    // Thêm cuộc thi mới
    function createContest(
        string memory id,
        string memory name,
        string memory description,
        uint256 startDate,
        uint256 endDate,
        string memory imageURL
    ) public {
        require(!contests[id].exists, "Contest with this ID already exists");
        require(endDate > startDate, "End date must be after start date");
        
        contests[id] = Contest({
            id: id,
            name: name,
            description: description,
            startDate: startDate,
            endDate: endDate,
            organizer: msg.sender,
            active: true,
            exists: true,
            imageURL: imageURL
        });
        
        contestIds.push(id);
        emit ContestCreated(id, name);
    }
    
    // Lấy thông tin cuộc thi theo ID
    function getContest(string memory id) public view returns (
        string memory name,
        string memory description,
        uint256 startDate,
        uint256 endDate,
        address organizer,
        bool active,
        string memory imageURL
    ) {
        require(contests[id].exists, "Contest does not exist");
        
        Contest memory c = contests[id];
        return (c.name, c.description, c.startDate, c.endDate, c.organizer, c.active, c.imageURL);
    }
    
    // Lấy danh sách tất cả ID cuộc thi
    function getAllContestIds() public view returns (string[] memory) {
        return contestIds;
    }
    
    // QUẢN LÝ NHÀ TÀI TRỢ
    
    // Thêm nhà tài trợ mới
    function addSponsor(
        string memory id,
        string memory name,
        string memory contactInfo,
        uint256 sponsorshipAmount
    ) public {
        require(!sponsors[id].exists, "Sponsor with this ID already exists");
        
        sponsors[id] = Sponsor({
            id: id,
            name: name,
            contactInfo: contactInfo,
            sponsorshipAmount: sponsorshipAmount,
            walletAddress: msg.sender,
            exists: true
        });
        
        sponsorIds.push(id);
        emit SponsorAdded(id, name);
    }
    
    // Lấy thông tin nhà tài trợ theo ID
    function getSponsor(string memory id) public view returns (
        string memory name,
        string memory contactInfo,
        uint256 sponsorshipAmount,
        address walletAddress
    ) {
        require(sponsors[id].exists, "Sponsor does not exist");
        
        Sponsor memory s = sponsors[id];
        return (s.name, s.contactInfo, s.sponsorshipAmount, s.walletAddress);
    }
    
    // Lấy danh sách tất cả ID nhà tài trợ
    function getAllSponsorIds() public view returns (string[] memory) {
        return sponsorIds;
    }
    
    // QUẢN LÝ ĐĂNG KÝ THÍ SINH - CUỘC THI
    
    // Đăng ký thí sinh vào cuộc thi
    function registerContestant(string memory contestId, string memory contestantId) public {
        require(contests[contestId].exists, "Contest does not exist");
        require(contestants[contestantId].exists, "Contestant does not exist");
        require(contests[contestId].active, "Contest is not active");
        require(block.timestamp < contests[contestId].endDate, "Contest registration is closed");
        require(!contestantRegistrations[contestId][contestantId], "Contestant already registered for this contest");
        
        contestantRegistrations[contestId][contestantId] = true;
        emit ContestantRegistered(contestId, contestantId);
    }
    
    // Kiểm tra thí sinh đã đăng ký cuộc thi hay chưa
    function isContestantRegistered(string memory contestId, string memory contestantId) public view returns (bool) {
        return contestantRegistrations[contestId][contestantId];
    }
    
    // Lấy danh sách thí sinh đã đăng ký một cuộc thi
    function getContestantsInContest(string memory contestId) public view returns (string[] memory) {
        require(contests[contestId].exists, "Contest does not exist");
        
        // Đếm số thí sinh đã đăng ký
        uint count = 0;
        for (uint i = 0; i < contestantIds.length; i++) {
            if (contestantRegistrations[contestId][contestantIds[i]]) {
                count++;
            }
        }
        
        // Tạo mảng kết quả
        string[] memory result = new string[](count);
        uint index = 0;
        for (uint i = 0; i < contestantIds.length; i++) {
            if (contestantRegistrations[contestId][contestantIds[i]]) {
                result[index] = contestantIds[i];
                index++;
            }
        }
        
        return result;
    }
}
