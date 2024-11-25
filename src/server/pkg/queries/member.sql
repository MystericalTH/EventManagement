-- Set the encryption key (replace 'your_secure_key' with your actual key)
SET @encryption_key = 'your_secure_key';

-- Query to list members requesting approval, with decrypted fields
-- name: ListRequestingMembers :many
SELECT 
    memberID, 
    AES_DECRYPT(fName, @encryption_key) AS fName, 
    AES_DECRYPT(lName, @encryption_key) AS lName, 
    email, 
    AES_DECRYPT(phone, @encryption_key) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE acceptDateTime IS NULL;

-- Query to list accepted members, with decrypted fields
-- name: ListAcceptedMembers :many
SELECT 
    memberID, 
    AES_DECRYPT(fName, @encryption_key) AS fName, 
    AES_DECRYPT(lName, @encryption_key) AS lName, 
    email, 
    AES_DECRYPT(phone, @encryption_key) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE acceptDateTime IS NOT NULL;

-- Query to retrieve a specific member by ID, with decrypted fields
-- name: ListMember :one
SELECT 
    memberID, 
    AES_DECRYPT(fName, @encryption_key) AS fName, 
    AES_DECRYPT(lName, @encryption_key) AS lName, 
    email, 
    AES_DECRYPT(phone, @encryption_key) AS phone, 
    githubUrl, 
    interest, 
    reason 
FROM MEMBER
WHERE memberID = ?;

-- Query to retrieve a specific member by email, with decrypted fields
-- name: ListMemberByEmail :one
SELECT 
    memberID, 
    AES_DECRYPT(fName, @encryption_key) AS fName, 
    AES_DECRYPT(lName, @encryption_key) AS lName, 
    email, 
    AES_DECRYPT(phone, @encryption_key) AS phone, 
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
    AES_ENCRYPT(?, @encryption_key), -- Encrypt fName
    AES_ENCRYPT(?, @encryption_key), -- Encrypt lName
    ?, 
    AES_ENCRYPT(?, @encryption_key), -- Encrypt phone
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
