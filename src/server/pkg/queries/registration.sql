-- name: InsertRegistration :exec
INSERT INTO ActivityRegistration (activityID, memberID, role, expectation, datetime)
VALUES (?, ?, ?, ?, CONVERT_TZ(NOW(), 'UTC', '+07:00'));

-- name: GetRegistrationStatus :one
SELECT COUNT(*) > 0 AS is_registered
FROM ActivityRegistration
WHERE activityID = ? AND memberID = ?;

-- name: ListActivityRegistration :many
SELECT Member.fname, Member.lname, role, Member.email, Member.phone, expectation, datetime
FROM ActivityRegistration
JOIN Member ON ActivityRegistration.memberID = Member.memberID
WHERE activityID = ?;

-- name: ListSubmittedMembers :many
SELECT 
    m.memberID, 
    m.fName, 
    m.lName, 
    m.email, 
    m.phone, 
    ar.role, 
    ar.expectation, 
    ar.datetime
    FROM 
        ActivityRegistration ar
    JOIN 
        Member m ON ar.memberID = m.memberID
    WHERE 
        ar.activityID = ?;

-- name: CheckProposer :one
SELECT COUNT(1) > 0 AS isProposer
    FROM Activity
    WHERE activityID = ? AND proposer = ?;


-- name: ListMemberActivities :many
SELECT 
    a.activityID, 
    a.title, 
    a.description, 
    ar.datetime, 
    a.proposer, 
    ar.role, 
    ar.expectation
    FROM 
        ActivityRegistration ar
    JOIN 
        Activity a ON ar.activityID = a.activityID
    WHERE 
        ar.memberID = ?;

