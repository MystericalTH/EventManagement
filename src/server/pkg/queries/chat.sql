-- name: InsertAdminDevChat :exec
INSERT INTO chatDevAd (adminID, developerID, sender, message, timesent) 
VALUES (?, ?, ?, ?, NOW());

-- name: ListAdminDevChat :many
SELECT developerID, message, sender, timesent FROM chatDevAd 
WHERE adminID = ? AND developerID = ?;

-- name: ListInitialAdminChatToDev :many
SELECT 
    d.developerID AS 'developerID',
    d.fname AS developer_fname, 
    d.lname AS developer_lname, 
    c.message, 
    c.timesent 
FROM developer d 
LEFT JOIN 
    (SELECT 
         developerid, 
         MAX(timesent) AS latest_time
     FROM
         chatDevAd ch
     WHERE ch.adminid=?
     GROUP BY
         developerid, adminid
    ) latest
ON d.developerID = latest.developerID

LEFT JOIN chatDevAd c
ON 
     c.timesent = latest.latest_time
AND c.developerid = latest.developerid;

-- name: ListInitialDevChatToAdmin :many
SELECT 
    a.adminid AS 'adminID',
    a.fname AS admin_fname,
    a.lname AS admin_lname,
    c.message, 
    c.timesent 
FROM admin a
LEFT JOIN 
    (SELECT 
         adminid, 
         MAX(timesent) AS latest_time
     FROM
         chatDevAd ch
     WHERE ch.developerid=?
     GROUP BY
         adminid, developerid
    ) latest
ON a.adminID = latest.adminID
LEFT JOIN chatDevAd c
ON 
     c.timesent = latest.latest_time
AND c.adminID = latest.adminID;

