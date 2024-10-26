-- PostgreSQL
SELECT class
FROM Courses
GROUP BY class
HAVING count(DISTINCT student) > 4