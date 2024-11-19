-- name: ListRequestingMembers :many
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE acceptDateTime IS NULL;

-- name: ListAcceptedMembers :many
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE acceptDateTime IS NOT NULL;

-- name: ListMember :one
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE memberID = ?;

-- name: ListMemberByEmail :one
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE email = ?;

-- name: InsertMember :exec
INSERT INTO MEMBER (fName, lName, email, phone, githubUrl, interest, reason) 
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: AcceptMember :exec
UPDATE MEMBER
SET acceptDateTime = NOW(),
    acceptAdmin = ? -- Include the admin responsible for the approval
WHERE memberID = ?;

