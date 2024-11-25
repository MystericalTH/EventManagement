-- Query to list members requesting approval, with decrypted fields
-- name: ListRequestingMembers :many
SELECT 
    memberID, 
    AES_DECRYPT(fName, SHA1('68299640939')) AS fName, 
    AES_DECRYPT(lName, SHA1('68299640939')) AS lName, 
    email, 
    AES_DECRYPT(phone, SHA1('68299640939')) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE acceptDateTime IS NULL;

-- Query to list accepted members, with decrypted fields
-- name: ListAcceptedMembers :many
SELECT 
    memberID, 
    AES_DECRYPT(fName,SHA1('68299640939')) AS fName, 
    AES_DECRYPT(lName,SHA1('68299640939')) AS lName, 
    email, 
    AES_DECRYPT(phone,SHA1('68299640939')) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE acceptDateTime IS NOT NULL;

-- Query to retrieve a specific member by ID, with decrypted fields
-- name: ListMember :one
SELECT 
    memberID, 
    AES_DECRYPT(fName, SHA1('68299640939')) AS fName, 
    AES_DECRYPT(lName, SHA1('68299640939')) AS lName, 
    email, 
    AES_DECRYPT(phone, SHA1('68299640939')) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE memberID = ?;

-- Query to retrieve a specific member by email, with decrypted fields
-- name: ListMemberByEmail :one
SELECT 
    memberID, 
    AES_DECRYPT(fName, SHA1('68299640939')) AS fName, 
    AES_DECRYPT(lName, SHA1('68299640939')) AS lName, 
    email, 
    AES_DECRYPT(phone, SHA1('68299640939')) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE email = ?;

-- Insert a new member with encrypted fields
-- name: InsertMember :exec
INSERT INTO MEMBER (
    fName, 
    lName, 
    email, 
    phone, 
    githubUrl, 
    interest, 
    reason
) VALUES (
    AES_ENCRYPT(?, SHA1('68299640939')), -- Encrypt fName
    AES_ENCRYPT(?, SHA1('68299640939')), -- Encrypt lName
    ?, 
    AES_ENCRYPT(?, SHA1('68299640939')), -- Encrypt phone
    ?, ?, ?
);

-- Approve a member by updating the acceptDateTime and acceptAdmin fields
-- name: AcceptMember :exec
UPDATE MEMBER
SET 
    acceptDateTime = CONVERT_TZ(NOW(), 'UTC', '+07:00'),
    acceptAdmin = ? -- Include the admin responsible for the approval
WHERE memberID = ?;

-- Delete a member from the database
-- name: DeleteMember :exec
DELETE FROM MEMBER
WHERE memberID = ?;
