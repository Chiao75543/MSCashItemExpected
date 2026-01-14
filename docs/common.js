// ============================================
// 共用函數
// ============================================

/**
 * 根據購買方式計算可得點數
 * @param {number} investment - 投入金額
 * @param {string} method - 購買方式 (card, cardreader, original, gift)
 * @param {number} discount - 折扣數
 * @returns {number} 可得點數
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

// ============================================
// Tab 切換邏輯
// ============================================

document.addEventListener('DOMContentLoaded', function() {
    // 主 Tab 切換
    const tabBtns = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');

    tabBtns.forEach(btn => {
        btn.addEventListener('click', function() {
            const targetTab = this.dataset.tab;

            // 移除所有 active
            tabBtns.forEach(b => b.classList.remove('active'));
            tabContents.forEach(c => c.classList.remove('active'));

            // 添加 active
            this.classList.add('active');
            document.getElementById(`${targetTab}-tab`).classList.add('active');
        });
    });

    // 子 Tab 切換
    const subTabBtns = document.querySelectorAll('.sub-tab-btn');
    const subTabContents = document.querySelectorAll('.sub-tab-content');

    subTabBtns.forEach(btn => {
        btn.addEventListener('click', function() {
            const targetSubTab = this.dataset.subtab;

            // 移除所有 active
            subTabBtns.forEach(b => b.classList.remove('active'));
            subTabContents.forEach(c => c.classList.remove('active'));

            // 添加 active
            this.classList.add('active');
            document.getElementById(`${targetSubTab}-tab`).classList.add('active');
        });
    });
});
