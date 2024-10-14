WITH t AS (
    SELECT
        emp.employee_id,
        emp.name,
        (
            SELECT
                count(*)
            FROM
                Employees AS sub
            WHERE
                emp.employee_id = sub.reports_to) AS reports_count,
        (
            SELECT
                round(avg(age), 0)
            FROM
                Employees AS sub2
            WHERE
                emp.employee_id = sub2.reports_to) AS average_age
    FROM
        Employees emp
)
SELECT
    *
FROM
    t
WHERE
    reports_count > 0
ORDER BY
    employee_id
