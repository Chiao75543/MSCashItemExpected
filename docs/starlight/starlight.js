// ============================================
// 星光錦囊 - 常數定義
// ============================================

// 每抽成本（點數）
const STARLIGHT_COST = 45;

// 星光錦囊主機率表 (%)
const STARLIGHT_RATES = {
    '靈魂艾爾達碎片交換券(10個)': 8.00,
    '靈魂艾爾達': 6.00,
    '永遠的輪迴星火': 14.40,
    '暗黑輪迴星火': 13.70,
    '特別附加潛在能力賦予卷軸': 7.80,
    '傳說潛在能力卷軸50%': 0.85,
    '傳說潛在能力卷軸100%': 0.55,
    '星力14星強化券': 15.00,
    '星力15星強化券': 10.00,
    '星力16星強化券': 7.00,
    '星力17星強化券': 3.40,
    '星力18星強化券': 1.50,
    '星力19星強化券': 0.60,
    '星力20星強化券': 0.40,
    '突破1星強化券100%(21星)': 0.45,
    '突破1星強化券100%(22星)': 0.20,
    '追加1星強化券30%(23星)': 0.15,
    '玲瓏星光': 10.00
};

// 玲瓏星光 -> 星光結晶體 機率表 (%)
const CRYSTAL_RATES = {
    '星力18星強化券': 18.00,
    '星力19星強化券': 12.00,
    '星力20星強化券': 6.00,
    '突破1星強化券30%(23星)': 10.00,
    '突破1星強化券50%(23星)': 4.00,
    '星光原石': 50.00
};

// 星光原石 機率表 (%)
const ROUGH_RATES = {
    '星力19星強化券': 10.00,
    '星力20星強化券': 8.00,
    '星力21星強化券': 2.00,
    '突破1星強化券30%(23星)': 8.00,
    '突破1星強化券50%(23星)': 6.00,
    '突破1星強化券100%(23星)': 5.00,
    '突破1星強化券30%(24星)': 7.00,
    '突破1星強化券50%(24星)': 4.00,
    '星光水晶': 50.00
};

// 星光水晶 機率表 (%)
const PURE_RATES = {
    '突破1星強化券50%(23星)': 20.00,
    '突破1星強化券100%(23星)': 15.00,
    '突破1星強化券30%(24星)': 8.00,
    '突破1星強化券50%(24星)': 4.00,
    '突破1星強化券100%(24星)': 2.00,
    '突破1星強化券30%(25星)': 0.70,
    '突破1星強化券50%(25星)': 0.30,
    '璀璨星光': 50.00
};

// 璀璨星光 機率表 (%)
const BRILLIANT_RATES = {
    '突破1星強化券30%(24星)': 29.00,
    '突破1星強化券50%(24星)': 19.00,
    '突破1星強化券100%(24星)': 14.00,
    '突破1星強化券30%(25星)': 20.00,
    '突破1星強化券50%(25星)': 9.00,
    '突破1星強化券100%(25星)': 4.00,
    '突破1星強化券30%(26星)': 3.00,
    '突破1星強化券50%(26星)': 2.00
};

// 價值為0的道具（不需要輸入）
const ZERO_VALUE_ITEMS = [
    '靈魂艾爾達碎片交換券(10個)',
    '靈魂艾爾達',
    '永遠的輪迴星火',
    '暗黑輪迴星火',
    '特別附加潛在能力賦予卷軸'
];

// 道具到輸入框ID的映射（只有第一階段道具）
const ITEM_INPUT_MAP = {
    '傳說潛在能力卷軸50%': 'sl-v-legend50',
    '傳說潛在能力卷軸100%': 'sl-v-legend100',
    '星力14星強化券': 'sl-v-star14',
    '星力15星強化券': 'sl-v-star15',
    '星力16星強化券': 'sl-v-star16',
    '星力17星強化券': 'sl-v-star17',
    '星力18星強化券': 'sl-v-star18',
    '星力19星強化券': 'sl-v-star19',
    '星力20星強化券': 'sl-v-star20',
    '突破1星強化券100%(21星)': 'sl-v-break21-100',
    '突破1星強化券100%(22星)': 'sl-v-break22-100',
    '追加1星強化券30%(23星)': 'sl-v-break23-30',
    '玲瓏星光': 'sl-v-crystal'
};

// 稀有道具列表（模擬器用）
const RARE_ITEMS = [
    '突破1星強化券100%(24星)',
    '突破1星強化券30%(25星)',
    '突破1星強化券50%(25星)',
    '突破1星強化券100%(25星)',
    '突破1星強化券30%(26星)',
    '突破1星強化券50%(26星)'
];

// ============================================
// 星光錦囊 - 計算函數
// ============================================

/**
 * 計算星光錦囊期望道具數量（僅第一階段，不展開玲瓏星光）
 */
function calculateStarlightExpected(drawCount) {
    const items = {};

    // 只計算第一階段道具
    for (const [item, rate] of Object.entries(STARLIGHT_RATES)) {
        items[item] = drawCount * (rate / 100);
    }

    return items;
}

/**
 * 獲取道具價值
 */
function getItemValue(itemName, itemValues) {
    if (ZERO_VALUE_ITEMS.includes(itemName)) {
        return 0;
    }
    const inputId = ITEM_INPUT_MAP[itemName];
    if (inputId && itemValues[inputId] !== undefined) {
        return itemValues[inputId];
    }
    return 0;
}

/**
 * 計算星光錦囊期望總價值
 */
function calculateStarlightValue(items, itemValues) {
    let total = 0;
    for (const [item, count] of Object.entries(items)) {
        total += count * getItemValue(item, itemValues);
    }
    return total;
}

/**
 * 星光錦囊主計算函數
 */
function calculateStarlight(investment, method, discount, itemValues) {
    // 1. 計算可得點數
    const points = calculatePoints(investment, method, discount);

    // 2. 計算可抽次數
    const drawCount = points / STARLIGHT_COST;

    // 3. 計算每抽實際成本
    const costPerDraw = drawCount > 0 ? investment / drawCount : 0;

    // 4. 計算期望獲得各道具數量
    const expectedItems = calculateStarlightExpected(drawCount);

    // 5. 計算期望總價值
    const expectedValue = calculateStarlightValue(expectedItems, itemValues);

    // 6. 計算報酬率
    const roi = investment > 0 ? ((expectedValue - investment) / investment) * 100 : 0;

    return {
        points: points,
        draw_count: drawCount,
        cost_per_draw: costPerDraw,
        expected_items: expectedItems,
        expected_value: expectedValue,
        roi: roi
    };
}

// ============================================
// 星光錦囊 - 模擬器
// ============================================

/**
 * 根據機率表隨機抽取一個道具
 */
function drawFromRates(rates) {
    const roll = Math.random() * 100;
    let cumulative = 0;

    for (const [item, rate] of Object.entries(rates)) {
        cumulative += rate;
        if (roll < cumulative) {
            return item;
        }
    }

    // 應該不會到這裡，但以防萬一返回最後一個
    return Object.keys(rates)[Object.keys(rates).length - 1];
}

/**
 * 模式一：星光錦囊模擬器（不展開玲瓏星光）
 */
function simulateBagOnly(count) {
    const results = {};

    for (let i = 0; i < count; i++) {
        const item = drawFromRates(STARLIGHT_RATES);
        if (!results[item]) {
            results[item] = 0;
        }
        results[item]++;
    }

    return results;
}

/**
 * 模式二：玲瓏星光模擬器（分階段）
 * @param {number} crystalCount - 消耗的玲瓏星光數量
 * @returns {object} 包含各階段結果
 */
function simulateCrystalStages(crystalCount) {
    const results = {
        stage1: {},  // 星光結晶體階段
        stage2: {},  // 星光原石階段
        stage3: {},  // 星光水晶階段
        stage4: {},  // 璀璨星光階段
        roughCount: 0,
        pureCount: 0,
        brilliantCount: 0
    };

    // 輔助函數
    function addToStage(stage, item) {
        if (!results[stage][item]) {
            results[stage][item] = 0;
        }
        results[stage][item]++;
    }

    // 階段1：星光結晶體
    for (let i = 0; i < crystalCount; i++) {
        const item = drawFromRates(CRYSTAL_RATES);
        if (item === '星光原石') {
            results.roughCount++;
        } else {
            addToStage('stage1', item);
        }
    }

    // 階段2：星光原石
    for (let i = 0; i < results.roughCount; i++) {
        const item = drawFromRates(ROUGH_RATES);
        if (item === '星光水晶') {
            results.pureCount++;
        } else {
            addToStage('stage2', item);
        }
    }

    // 階段3：星光水晶
    for (let i = 0; i < results.pureCount; i++) {
        const item = drawFromRates(PURE_RATES);
        if (item === '璀璨星光') {
            results.brilliantCount++;
        } else {
            addToStage('stage3', item);
        }
    }

    // 階段4：璀璨星光
    for (let i = 0; i < results.brilliantCount; i++) {
        const item = drawFromRates(BRILLIANT_RATES);
        addToStage('stage4', item);
    }

    return results;
}

// ============================================
// 星光錦囊 - UI 邏輯
// ============================================

document.addEventListener('DOMContentLoaded', function() {
    // ============================================
    // 期望值計算機
    // ============================================

    const slCalculateBtn = document.getElementById('sl-calculate-btn');
    const slResultDiv = document.getElementById('sl-result');

    if (slCalculateBtn) {
        slCalculateBtn.addEventListener('click', function() {
            const investment = parseFloat(document.getElementById('sl-investment').value) || 0;
            const method = document.querySelector('input[name="sl-method"]:checked').value;

            let discount = 1;
            if (method === 'card') {
                discount = parseFloat(document.getElementById('sl-card-discount').value) || 1;
            } else if (method === 'gift') {
                discount = parseFloat(document.getElementById('sl-gift-discount').value) || 1;
            }

            // 收集所有道具價值
            const itemValues = {};
            for (const inputId of Object.values(ITEM_INPUT_MAP)) {
                const input = document.getElementById(inputId);
                if (input) {
                    itemValues[inputId] = parseFloat(input.value) || 0;
                }
            }

            if (investment <= 0) {
                alert('請輸入投入資金');
                return;
            }

            const result = calculateStarlight(investment, method, discount, itemValues);
            displayStarlightResult(result, itemValues);
        });
    }

    function displayStarlightResult(result, itemValues) {
        slResultDiv.style.display = 'block';

        // 基本資訊
        document.getElementById('sl-r-points').textContent = result.points.toFixed(0) + ' 點';
        document.getElementById('sl-r-draws').textContent = result.draw_count.toFixed(2) + ' 次';
        document.getElementById('sl-r-cost').textContent = result.cost_per_draw.toFixed(2) + ' 元';

        // 期望道具列表
        const itemsDiv = document.getElementById('sl-r-items');
        itemsDiv.innerHTML = '';

        // 按數量排序
        const sortedItems = Object.entries(result.expected_items)
            .sort((a, b) => b[1] - a[1]);

        for (const [itemName, count] of sortedItems) {
            if (count < 0.0001) continue; // 跳過數量太小的

            const value = getItemValue(itemName, itemValues);
            const totalValue = count * value;
            const isZeroValue = ZERO_VALUE_ITEMS.includes(itemName) || value === 0;

            const row = document.createElement('div');
            row.className = 'item-row' + (isZeroValue ? ' zero-value' : '');
            row.innerHTML = `
                <span class="name">${itemName}</span>
                <span class="count">${count.toFixed(4)}</span>
                <span class="value">${isZeroValue ? '-' : totalValue.toFixed(2) + '元'}</span>
            `;
            itemsDiv.appendChild(row);
        }

        // 期望總價值與報酬率
        document.getElementById('sl-r-value').textContent = result.expected_value.toFixed(2) + ' 元';

        const roiSpan = document.getElementById('sl-r-roi');
        const roiValue = result.roi.toFixed(2);
        roiSpan.textContent = (result.roi >= 0 ? '+' : '') + roiValue + '%';
        roiSpan.className = result.roi >= 0 ? 'positive' : 'negative';

        // 滾動到結果
        slResultDiv.scrollIntoView({ behavior: 'smooth' });
    }

    // ============================================
    // 模擬器
    // ============================================

    const simTypeRadios = document.querySelectorAll('input[name="sim-type"]');
    const simBagSettings = document.getElementById('sim-bag-settings');
    const simCrystalSettings = document.getElementById('sim-crystal-settings');
    const slSimulateBtn = document.getElementById('sl-simulate-btn');
    const slSimResultDiv = document.getElementById('sl-sim-result');
    const slCrystalResultDiv = document.getElementById('sl-crystal-result');

    // 模擬器類型切換
    if (simTypeRadios.length > 0) {
        simTypeRadios.forEach(radio => {
            radio.addEventListener('change', function() {
                if (this.value === 'bag') {
                    simBagSettings.style.display = 'block';
                    simCrystalSettings.style.display = 'none';
                } else {
                    simBagSettings.style.display = 'none';
                    simCrystalSettings.style.display = 'block';
                }
                // 隱藏結果
                slSimResultDiv.style.display = 'none';
                slCrystalResultDiv.style.display = 'none';
            });
        });
    }

    if (slSimulateBtn) {
        slSimulateBtn.addEventListener('click', async function() {
            const simType = document.querySelector('input[name="sim-type"]:checked').value;

            if (simType === 'bag') {
                // 星光錦囊模擬器
                await runBagSimulation();
            } else {
                // 玲瓏星光模擬器
                await runCrystalSimulation();
            }
        });
    }

    /**
     * 星光錦囊模擬器
     */
    async function runBagSimulation() {
        const count = parseInt(document.getElementById('sl-sim-count').value) || 100;

        if (count <= 0 || count > 10000) {
            alert('請輸入 1-10000 之間的抽數');
            return;
        }

        slSimResultDiv.style.display = 'block';
        slCrystalResultDiv.style.display = 'none';

        const slSimAnimation = document.getElementById('sl-sim-animation');
        const slSimItems = document.getElementById('sl-sim-items');

        slSimItems.innerHTML = '';
        slSimAnimation.textContent = '模擬中...';

        await new Promise(resolve => setTimeout(resolve, 50));

        const results = simulateBagOnly(count);

        slSimAnimation.textContent = `模擬 ${count} 次完成！`;

        const sortedResults = Object.entries(results)
            .sort((a, b) => b[1] - a[1]);

        for (const [itemName, itemCount] of sortedResults) {
            const isRare = RARE_ITEMS.includes(itemName) || itemName === '玲瓏星光';

            const row = document.createElement('div');
            row.className = 'sim-item' + (isRare ? ' rare' : '');
            row.innerHTML = `
                <span class="name">${itemName}</span>
                <span class="count">x${itemCount}</span>
            `;
            slSimItems.appendChild(row);
        }

        slSimResultDiv.scrollIntoView({ behavior: 'smooth' });
    }

    /**
     * 玲瓏星光模擬器
     */
    async function runCrystalSimulation() {
        const crystalCount = parseInt(document.getElementById('sl-crystal-count').value) || 10;

        if (crystalCount <= 0 || crystalCount > 1000) {
            alert('請輸入 1-1000 之間的玲瓏星光數量');
            return;
        }

        slSimResultDiv.style.display = 'none';
        slCrystalResultDiv.style.display = 'block';

        const slCrystalAnimation = document.getElementById('sl-crystal-animation');
        slCrystalAnimation.textContent = '模擬中...';

        await new Promise(resolve => setTimeout(resolve, 50));

        const results = simulateCrystalStages(crystalCount);

        slCrystalAnimation.textContent = `消耗 ${crystalCount} 顆玲瓏星光完成！`;

        // 顯示各階段結果
        displayStageResult('sl-crystal-stage1', results.stage1, `獲得 ${results.roughCount} 個星光原石`);
        displayStageResult('sl-crystal-stage2', results.stage2, `獲得 ${results.pureCount} 個星光水晶`, 'sl-stage2-section', results.roughCount > 0);
        displayStageResult('sl-crystal-stage3', results.stage3, `獲得 ${results.brilliantCount} 個璀璨星光`, 'sl-stage3-section', results.pureCount > 0);
        displayStageResult('sl-crystal-stage4', results.stage4, '', 'sl-stage4-section', results.brilliantCount > 0);

        slCrystalResultDiv.scrollIntoView({ behavior: 'smooth' });
    }

    /**
     * 顯示階段結果
     */
    function displayStageResult(containerId, items, extraInfo, sectionId, show) {
        const container = document.getElementById(containerId);
        container.innerHTML = '';

        // 控制區段顯示
        if (sectionId) {
            const section = document.getElementById(sectionId);
            section.style.display = show ? 'block' : 'none';
            if (!show) return;
        }

        // 顯示額外資訊（如獲得多少原石）
        if (extraInfo) {
            const infoDiv = document.createElement('div');
            infoDiv.className = 'sim-item rare';
            infoDiv.innerHTML = `<span class="name">${extraInfo}</span>`;
            container.appendChild(infoDiv);
        }

        // 顯示獎品
        const sortedItems = Object.entries(items)
            .sort((a, b) => b[1] - a[1]);

        for (const [itemName, itemCount] of sortedItems) {
            const isRare = RARE_ITEMS.includes(itemName);

            const row = document.createElement('div');
            row.className = 'sim-item' + (isRare ? ' rare' : '');
            row.innerHTML = `
                <span class="name">${itemName}</span>
                <span class="count">x${itemCount}</span>
            `;
            container.appendChild(row);
        }
    }
});
