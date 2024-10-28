WITH filtered AS (
    SELECT
        s.*,
        rank() OVER (PARTITION BY s.product_id ORDER BY s.year ASC) AS year_rank
    FROM
        Sales s
)
SELECT
    product_id,
    year AS first_year,
    quantity,
    price
FROM
    filtered
WHERE
    year_rank = 1
