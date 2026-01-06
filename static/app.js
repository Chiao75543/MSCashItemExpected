// ============================================
// 常數定義
// ============================================

// 生肖機率 (%)
const ZODIAC_RATES = {
    '馬': 0.15,
    '羊': 0.20,
    '猴': 0.25,
    '雞': 0.80,
    '狗': 0.90,
    '豬': 1.00,
    '鼠': 2.50,
    '牛': 3.00,
    '虎': 3.50,
    '兔': 29.20,
    '龍': 29.25,
    '蛇': 29.25
};

// 心願箱組成（需湊齊各1個）
const BOX_REQUIREMENTS = {
    '小吉': ['兔', '龍', '蛇'],
    '中吉': ['兔', '龍', '蛇', '虎', '牛', '鼠'],
    '大吉': ['兔', '龍', '蛇', '虎', '牛', '鼠', '豬', '狗', '雞'],
    '超越': ['兔', '龍', '蛇', '虎', '牛', '鼠', '豬', '狗', '雞', '猴', '羊', '馬']
};

// 心願箱優先順序（高價值優先）
const BOX_PRIORITY = ['超越', '大吉', '中吉', '小吉'];

// 每抽成本（點數）
const COST_PER_DRAW = 27;

// 生肖順序（按機率從低到高）
const ZODIAC_ORDER = ['馬', '羊', '猴', '雞', '狗', '豬', '鼠', '牛', '虎', '兔', '龍', '蛇'];

// ============================================
// 計算函數
// ============================================

/**
 * 根據購買方式計算可得點數
 */
function calculatePoints(investment, method, discount) {
    switch (method) {
        case 'card':
            // 點卡 xx 折：投入金額 ÷ 折數（四捨五入）
            return discount > 0 ? Math.round(investment / discount) : investment;
        case 'cardreader':
            // 讀卡機 5% 回饋
            return investment * 1.05;
        case 'original':
            // 原價
            return investment;
        case 'gift':
            // 送禮 xx 折：投入金額 ÷ 折數
            return discount > 0 ? Math.round(investment / discount) : investment;
        default:
            return investment;
    }
}

/**
 * 計算期望獲得的氣息數量
 */
function calculateExpectedBreaths(drawCount) {
    const breaths = {};
    for (const [zodiac, rate] of Object.entries(ZODIAC_RATES)) {
        breaths[zodiac] = drawCount * (rate / 100);
    }
    return breaths;
}

/**
 * 計算期望心願箱數量（貪心算法，優先湊高價值）
 */
function calculateExpectedBoxes(breaths) {
    // 複製一份，避免修改原始資料
    const remaining = { ...breaths };
    const result = {
        '超越': 0,
        '大吉': 0,
        '中吉': 0,
        '小吉': 0
    };

    // 按優先順序湊箱
    for (const boxType of BOX_PRIORITY) {
        const requirements = BOX_REQUIREMENTS[boxType];

        // 找出可湊的箱數（取最小值）
        let minCount = Infinity;
        for (const zodiac of requirements) {
            if (remaining[zodiac] < minCount) {
                minCount = remaining[zodiac];
            }
        }

        if (minCount > 0 && minCount !== Infinity) {
            result[boxType] = minCount;
            // 扣除已使用的氣息
            for (const zodiac of requirements) {
                remaining[zodiac] -= minCount;
            }
        }
    }

    return result;
}

/**
 * 計算期望總價值
 */
function calculateExpectedValue(boxes, boxValues) {
    return boxes['小吉'] * boxValues.small +
           boxes['中吉'] * boxValues.medium +
           boxes['大吉'] * boxValues.large +
           boxes['超越'] * boxValues.super;
}

/**
 * 主計算函數
 */
function calculate(investment, method, discount, boxValues) {
    // 1. 計算可得點數
    const points = calculatePoints(investment, method, discount);

    // 2. 計算可抽次數
    const drawCount = points / COST_PER_DRAW;

    // 3. 計算每個氣息的實際成本
    const costPerBreath = drawCount > 0 ? investment / drawCount : 0;

    // 4. 計算期望獲得各氣息數量
    const expectedBreaths = calculateExpectedBreaths(drawCount);

    // 5. 計算期望可湊心願箱數量（貪心算法）
    const expectedBoxes = calculateExpectedBoxes(expectedBreaths);

    // 6. 計算期望總價值
    const expectedValue = calculateExpectedValue(expectedBoxes, boxValues);

    // 7. 計算報酬率
    const roi = investment > 0 ? ((expectedValue - investment) / investment) * 100 : 0;

    return {
        points: points,
        draw_count: drawCount,
        cost_per_breath: costPerBreath,
        expected_breaths: expectedBreaths,
        expected_boxes: expectedBoxes,
        expected_value: expectedValue,
        roi: roi
    };
}

// ============================================
// UI 邏輯
// ============================================

document.addEventListener('DOMContentLoaded', function() {
    const calculateBtn = document.getElementById('calculate-btn');
    const resultDiv = document.getElementById('result');

    calculateBtn.addEventListener('click', function() {
        const investment = parseFloat(document.getElementById('investment').value) || 0;
        const method = document.querySelector('input[name="method"]:checked').value;

        let discount = 1;
        if (method === 'card') {
            discount = parseFloat(document.getElementById('card-discount').value) || 1;
        } else if (method === 'gift') {
            discount = parseFloat(document.getElementById('gift-discount').value) || 1;
        }

        const boxValues = {
            small: parseFloat(document.getElementById('box-small').value) || 0,
            medium: parseFloat(document.getElementById('box-medium').value) || 0,
            large: parseFloat(document.getElementById('box-large').value) || 0,
            super: parseFloat(document.getElementById('box-super').value) || 0
        };

        if (investment <= 0) {
            alert('請輸入投入資金');
            return;
        }

        // 本地計算（不需要 API）
        const result = calculate(investment, method, discount, boxValues);
        displayResult(result);
    });

    function displayResult(result) {
        resultDiv.style.display = 'block';

        // 基本資訊
        document.getElementById('r-points').textContent = result.points.toFixed(0) + ' 點';
        document.getElementById('r-draws').textContent = result.draw_count.toFixed(2) + ' 次';
        document.getElementById('r-cost').textContent = result.cost_per_breath.toFixed(2) + ' 元';

        // 期望氣息
        const breathsDiv = document.getElementById('r-breaths');
        breathsDiv.innerHTML = '';
        ZODIAC_ORDER.forEach(zodiac => {
            const count = result.expected_breaths[zodiac] || 0;
            const item = document.createElement('div');
            item.className = 'breath-item';
            item.innerHTML = `
                <span class="name">${zodiac}</span>
                <span class="count">${count.toFixed(2)}</span>
            `;
            breathsDiv.appendChild(item);
        });

        // 期望心願箱
        document.getElementById('r-box-super').textContent = (result.expected_boxes['超越'] || 0).toFixed(4);
        document.getElementById('r-box-large').textContent = (result.expected_boxes['大吉'] || 0).toFixed(4);
        document.getElementById('r-box-medium').textContent = (result.expected_boxes['中吉'] || 0).toFixed(4);
        document.getElementById('r-box-small').textContent = (result.expected_boxes['小吉'] || 0).toFixed(4);

        // 期望總價值與報酬率
        document.getElementById('r-value').textContent = result.expected_value.toFixed(2) + ' 元';

        const roiSpan = document.getElementById('r-roi');
        const roiValue = result.roi.toFixed(2);
        roiSpan.textContent = (result.roi >= 0 ? '+' : '') + roiValue + '%';
        roiSpan.className = result.roi >= 0 ? 'positive' : 'negative';

        // 滾動到結果
        resultDiv.scrollIntoView({ behavior: 'smooth' });
    }
});
