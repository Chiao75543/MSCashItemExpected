document.addEventListener('DOMContentLoaded', function() {
    const calculateBtn = document.getElementById('calculate-btn');
    const resultDiv = document.getElementById('result');

    // 生肖順序（按機率從低到高）
    const zodiacOrder = ['馬', '羊', '猴', '雞', '狗', '豬', '鼠', '牛', '虎', '兔', '龍', '蛇'];

    calculateBtn.addEventListener('click', async function() {
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

        try {
            const response = await fetch('/api/calculate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    investment: investment,
                    method: method,
                    discount: discount,
                    box_values: boxValues
                })
            });

            if (!response.ok) {
                throw new Error('計算失敗');
            }

            const result = await response.json();
            displayResult(result);
        } catch (error) {
            alert('計算發生錯誤：' + error.message);
        }
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
        zodiacOrder.forEach(zodiac => {
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
