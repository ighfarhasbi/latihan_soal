WITH tariff_usage AS (
    SELECT
        r.account_id,
        t.name AS tariff_name,
        COUNT(*) AS usage_count,
        RANK() OVER (PARTITION BY r.account_id ORDER BY COUNT(*) DESC) AS rnk
    FROM readings r
    JOIN tariffs t ON r.tariff_id = t.id
    GROUP BY r.account_id, t.name
)
SELECT
    a.username,
    a.email,
    tu.tariff_name AS most_frequent_tariff,
    SUM(r.amount) AS total_consumption,
    ROUND(SUM(t.cost)::numeric / COUNT(r.*), 2) AS average_cost_per_reading
FROM accounts a
JOIN readings r ON r.account_id = a.id
JOIN tariffs t ON r.tariff_id = t.id
JOIN tariff_usage tu ON tu.account_id = a.id AND tu.rnk = 1
GROUP BY a.username, a.email, tu.tariff_name
ORDER BY a.username ASC;
