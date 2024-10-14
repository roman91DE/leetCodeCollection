"" "
This query works for counting all attended Exams per Student/Subject,
the solutions requires also adding in all combinations that arent attended (Cross-Joins)
"""
SELECT
    s.*,
    ex.subject_name,
    count(*) AS attended_exams
FROM
    Examinations ex
    LEFT JOIN Students s ON ex.student_id = s.student_id
GROUP BY
    ex.student_id,
    ex.subject_name,
    ex.subject_name
ORDER BY
    ex.student_id,
    ex.subject_name
