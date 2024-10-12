WITH user_confirmation_counts AS (
  SELECT user_id, 
         COUNT(*) AS totalCount, 
         COUNT(CASE WHEN action = 'confirmed' THEN 1 END) AS confirmedCount
  FROM Confirmations
  GROUP BY user_id
),
user_confirmation_rate AS (
  SELECT user_id, ROUND((confirmedCount / NULLIF(totalCount, 0)),2) AS confirmation_rate
  FROM user_confirmation_counts
),
users_without_confirmations AS (
  SELECT user_id, 0 AS confirmation_rate
  FROM Signups
  WHERE user_id NOT IN (SELECT user_id FROM user_confirmation_rate)
)
SELECT * 
FROM user_confirmation_rate
UNION
SELECT * 
FROM users_without_confirmations;