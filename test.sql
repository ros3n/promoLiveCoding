SELECT DISTINCT ON (current.UserID, previous.UserID) (current.MetricValue - previous.MetricValue) AS 'result'
FROM nice_table AS current
INNER JOIN nice_table AS previous
ON current.UserID = previous.UserID
WHERE previous.Date > current.Date - 7.DAYS
ORDER BY previous.Date ASC, current.Date DESC
