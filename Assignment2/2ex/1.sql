SELECT name, COUNT(*)
FROM users
GROUP BY name
HAVING COUNT(*) > 1;

DELETE FROM users
WHERE id NOT IN (
    SELECT MIN(id)
    FROM users
    GROUP BY name
);
