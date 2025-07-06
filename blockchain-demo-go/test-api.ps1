# PowerShell script để test API
# Chạy: .\test-api.ps1

$baseUrl = "http://localhost:8080/api/v1"

Write-Host "🚀 Testing Blockchain Event Management API" -ForegroundColor Green
Write-Host "Base URL: $baseUrl" -ForegroundColor Yellow
Write-Host ""

# Function to make HTTP requests
function Invoke-ApiRequest {
    param(
        [string]$Method,
        [string]$Url,
        [string]$Body = $null
    )
    
    try {
        $headers = @{
            "Content-Type" = "application/json"
        }
        
        if ($Body) {
            $response = Invoke-RestMethod -Uri $Url -Method $Method -Body $Body -Headers $headers
        } else {
            $response = Invoke-RestMethod -Uri $Url -Method $Method -Headers $headers
        }
        
        return $response
    }
    catch {
        Write-Host "❌ Error: $($_.Exception.Message)" -ForegroundColor Red
        return $null
    }
}

# Test 1: Health Check
Write-Host "1️⃣ Testing Health Check..." -ForegroundColor Cyan
$healthResponse = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/health"
if ($healthResponse) {
    Write-Host "✅ Health Status: $($healthResponse.status)" -ForegroundColor Green
}
Write-Host ""

# Test 2: Create Contest
Write-Host "2️⃣ Creating Contest..." -ForegroundColor Cyan
$contestData = @{
    name = "Road To ESSEN 2025"
    description = "Cuộc thi thiết kế board game Việt Nam lớn nhất năm 2025"
    start_date = "2025-06-24T00:00:00Z"
    end_date = "2025-07-20T23:59:59Z"
    image_url = "https://example.com/road-to-essen.jpg"
} | ConvertTo-Json

$contestResponse = Invoke-ApiRequest -Method "POST" -Url "$baseUrl/contests" -Body $contestData
if ($contestResponse -and $contestResponse.success) {
    $contestId = $contestResponse.id
    Write-Host "✅ Contest created with ID: $contestId" -ForegroundColor Green
} else {
    Write-Host "❌ Failed to create contest" -ForegroundColor Red
}
Write-Host ""

# Test 3: Create Contestant
Write-Host "3️⃣ Creating Contestant..." -ForegroundColor Cyan
$contestantData = @{
    name = "Nguyễn Văn A"
    details = "Sinh viên năm 4 chuyên ngành Thiết kế Game, có kinh nghiệm 2 năm"
    creator = "0x742d35cc6641c7b2b85ce462af7c9bb7a5db8b7a"
} | ConvertTo-Json

$contestantResponse = Invoke-ApiRequest -Method "POST" -Url "$baseUrl/contestants" -Body $contestantData
if ($contestantResponse -and $contestantResponse.success) {
    $contestantId = $contestantResponse.id
    Write-Host "✅ Contestant created with ID: $contestantId" -ForegroundColor Green
} else {
    Write-Host "❌ Failed to create contestant" -ForegroundColor Red
}
Write-Host ""

# Test 4: Create Sponsor
Write-Host "4️⃣ Creating Sponsor..." -ForegroundColor Cyan
$sponsorData = @{
    name = "VNG Corporation"
    contact_info = "sponsor@vng.com.vn | 028-1234-5678"
    sponsorship_amount = 50000000
} | ConvertTo-Json

$sponsorResponse = Invoke-ApiRequest -Method "POST" -Url "$baseUrl/sponsors" -Body $sponsorData
if ($sponsorResponse -and $sponsorResponse.success) {
    $sponsorId = $sponsorResponse.id
    Write-Host "✅ Sponsor created with ID: $sponsorId" -ForegroundColor Green
} else {
    Write-Host "❌ Failed to create sponsor" -ForegroundColor Red
}
Write-Host ""

# Test 5: Register Contestant for Contest
if ($contestId -and $contestantId) {
    Write-Host "5️⃣ Registering Contestant for Contest..." -ForegroundColor Cyan
    $registrationData = @{
        contestant_id = $contestantId
    } | ConvertTo-Json

    $registrationResponse = Invoke-ApiRequest -Method "POST" -Url "$baseUrl/contests/$contestId/register" -Body $registrationData
    if ($registrationResponse -and $registrationResponse.success) {
        Write-Host "✅ Contestant registered successfully" -ForegroundColor Green
    } else {
        Write-Host "❌ Failed to register contestant" -ForegroundColor Red
    }
    Write-Host ""
}

# Test 6: Get Contest Details
if ($contestId) {
    Write-Host "6️⃣ Getting Contest Details..." -ForegroundColor Cyan
    $contestDetails = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/contests/$contestId"
    if ($contestDetails -and $contestDetails.success) {
        Write-Host "✅ Contest Name: $($contestDetails.data.name)" -ForegroundColor Green
        Write-Host "   Organizer: $($contestDetails.data.organizer)" -ForegroundColor Gray
    }
    Write-Host ""
}

# Test 7: Get Contestants in Contest
if ($contestId) {
    Write-Host "7️⃣ Getting Contestants in Contest..." -ForegroundColor Cyan
    $contestantsInContest = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/contests/$contestId/contestants"
    if ($contestantsInContest -and $contestantsInContest.success) {
        Write-Host "✅ Total contestants registered: $($contestantsInContest.total)" -ForegroundColor Green
        if ($contestantsInContest.contestants.Count -gt 0) {
            Write-Host "   First contestant: $($contestantsInContest.contestants[0].name)" -ForegroundColor Gray
        }
    }
    Write-Host ""
}

# Test 8: Get All Contests
Write-Host "8️⃣ Getting All Contests..." -ForegroundColor Cyan
$allContests = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/contests"
if ($allContests -and $allContests.success) {
    Write-Host "✅ Total contests: $($allContests.total)" -ForegroundColor Green
}
Write-Host ""

# Test 9: Get All Contestants
Write-Host "9️⃣ Getting All Contestants..." -ForegroundColor Cyan
$allContestants = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/contestants"
if ($allContestants -and $allContestants.success) {
    Write-Host "✅ Total contestants: $($allContestants.total)" -ForegroundColor Green
}
Write-Host ""

# Test 10: Get Statistics
Write-Host "🔟 Getting Statistics..." -ForegroundColor Cyan
$stats = Invoke-ApiRequest -Method "GET" -Url "$baseUrl/stats"
if ($stats -and $stats.success) {
    Write-Host "✅ Statistics:" -ForegroundColor Green
    Write-Host "   Contents: $($stats.data.total_contents)" -ForegroundColor Gray
    Write-Host "   Contests: $($stats.data.total_contests)" -ForegroundColor Gray
    Write-Host "   Contestants: $($stats.data.total_contestants)" -ForegroundColor Gray
    Write-Host "   Sponsors: $($stats.data.total_sponsors)" -ForegroundColor Gray
    Write-Host "   Registrations: $($stats.data.total_registrations)" -ForegroundColor Gray
}
Write-Host ""

# Test 11: Create Content
Write-Host "1️⃣1️⃣ Creating Content..." -ForegroundColor Cyan
$contentData = @{
    title = "Hướng dẫn tham gia cuộc thi Board Game"
    content = "Các bước cần thiết để tham gia cuộc thi thiết kế board game Việt Nam 2025. Bao gồm đăng ký, nộp bài và quy trình đánh giá..."
    creator = "admin"
} | ConvertTo-Json

$contentResponse = Invoke-ApiRequest -Method "POST" -Url "$baseUrl/content" -Body $contentData
if ($contentResponse -and $contentResponse.success) {
    Write-Host "✅ Content created with ID: $($contentResponse.id)" -ForegroundColor Green
}
Write-Host ""

Write-Host "🎉 API Testing Completed!" -ForegroundColor Green
Write-Host "Check the server logs for detailed transaction information." -ForegroundColor Yellow
