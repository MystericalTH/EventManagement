-- name: InsertChat :exec
INSERT INTO chatDevAd (adminID, developerID, role, message, datetime) 
VALUES (?, ?, ?, ?, NOW());

-- name: ListChat :many
SELECT 
    c.messageid, 
    a.email AS admin_email, 
    d.email AS developer_email, 
    c.message, 
    c.datetime 
FROM 
    chatDevAd c
LEFT JOIN 
    Admin a ON c.adminid = a.adminID
LEFT JOIN 
    Developer d ON c.developerid = d.developerID;