-- My first query works for counting all attended Exams per Student/Subject,
-- the solutions requires also adding in all combinations that arent attended (Cross-Joins)
-- SELECT
--     s.*,
--     ex.subject_name,
--     count(*) AS attended_exams
-- FROM
--     Examinations ex
--     LEFT JOIN Students s ON ex.student_id = s.student_id
-- GROUP BY
--     ex.student_id,
--     ex.subject_name,
--     ex.subject_name
-- ORDER BY
--     ex.student_id,
--     ex.subject_name
SELECT
    s.student_id,
    s.student_name,
    sub.subject_name,
    COUNT(ex.subject_name) AS attended_exams
FROM
    Students s
    CROSS JOIN Subjects sub
    LEFT JOIN Examinations ex ON s.student_id = ex.student_id
        AND ex.subject_name = sub.subject_name
GROUP BY
    s.student_id,
    s.student_name,
    sub.subject_name
ORDER BY
    s.student_id,
    sub.subject_name
