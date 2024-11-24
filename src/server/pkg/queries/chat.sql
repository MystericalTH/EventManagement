-- name: InsertAdminDevChat :exec
INSERT INTO chatDevAd (adminID, developerID, sender, message, timesent) 
VALUES (?, ?, ?, ?, NOW());

-- name: ListAdminDevChat :many
SELECT developerID, message, sender, timesent FROM chatDevAd 
WHERE adminID = ? AND developerID = ?;

-- name: ListInitialAdminChatToDev :many
SELECT 
    d.fname AS developer_fname, 
    d.lname AS developer_lname, 
    c.message, 
    c.timesent 
FROM 
    chatDevAd c
JOIN 
    (SELECT 
         developerid, 
         MAX(timesent) AS latest_time
     FROM
         chatDevAd
     GROUP BY
         developerid, adminid) latest
ON 
    c.developerid = latest.developerid 
    AND c.timesent = latest.latest_time
JOIN 
    developer d ON c.developerid = d.developerid
WHERE 
    c.adminid = ?;

-- name: ListInitialDevChatToAdmin :many
SELECT 
    a.fname AS admin_fname,
    a.lname AS admin_lname,
    c.message, 
    c.timesent 
FROM 
    chatDevAd c
JOIN 
  (SELECT 
        adminid,
        MAX(timesent) AS latest_time
    FROM 
        chatDevAd
    GROUP BY 
        adminid, developerid
  ) latest
ON 
  c.adminid = latest.adminid 
  AND c.timesent = latest.latest_time
JOIN 
  admin a ON c.adminid = a.adminid
WHERE 
  c.developerid = ?;

