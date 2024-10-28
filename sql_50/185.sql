WITH ranked AS (
    SELECT
        d.name AS Department,
        e.name AS Employee,
        e.salary AS Salary,
        dense_rank() OVER (PARTITION BY d.name ORDER BY e.salary DESC) AS salary_rank
    FROM
        Employee e
        JOIN Department d ON e.departmentId = d.id
)
SELECT
    Department,
    Employee,
    Salary
FROM
    ranked
WHERE
    salary_rank <= 3
