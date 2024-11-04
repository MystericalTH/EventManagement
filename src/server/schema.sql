CREATE TABLE Admin (
    adminID INT PRIMARY KEY,
    email VARCHAR(320) NOT NULL
);

CREATE TABLE Developer (
    developerID INT PRIMARY KEY,
    email VARCHAR(320) NOT NULL
);

CREATE TABLE chatDevAd (
    messageID INT PRIMARY KEY,
    adminID INT NOT NULL,
    developerID INT NOT NULL,
    message TEXT NOT NULL,
    datetime DATETIME NOT NULL,
    FOREIGN KEY (adminID) REFERENCES Admin(adminID),
    FOREIGN KEY (developerID) REFERENCES Developer(developerID)
);

CREATE TABLE Member (
    memberID INT PRIMARY KEY,
    fName VARCHAR(255) NOT NULL,
    lName VARCHAR(255) NOT NULL,
    email VARCHAR(320) NOT NULL,
    phone VARCHAR(20),
    githubUrl VARCHAR(320),
    interest TEXT NOT NULL,
    reason TEXT NOT NULL,
    acceptDateTime DATETIME,
    acceptAdmin INT NOT NULL,
    FOREIGN KEY (acceptAdmin) REFERENCES Admin(adminID)
);

CREATE TABLE Activity (
    activityID INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    proposer INT NOT NULL,
    startDate DATE NOT NULL,
    endDate DATE NOT NULL,
    maxNumber INT,
    format VARCHAR(10) CHECK (format IN ('workshop', 'project')),
    description TEXT NOT NULL,
    proposeDateTime DATETIME NOT NULL,
    acceptAdmin INT NOT NULL,
    acceptDateTime DATETIME NOT NULL,
    applicationStatus VARCHAR(20) NOT NULL,
    FOREIGN KEY (acceptAdmin) REFERENCES Admin(adminID),
    FOREIGN KEY (proposer) REFERENCES Member(memberID)
);

CREATE TABLE Workshop (
    workshopID INT PRIMARY KEY,
    startTime TIME NOT NULL,
    endTime TIME NOT NULL,
    FOREIGN KEY (workshopID) REFERENCES Activity(activityID)
);

CREATE TABLE Project (
    projectID INT PRIMARY KEY,
    advisor VARCHAR(255),
    FOREIGN KEY (projectID) REFERENCES Activity(activityID)
);

CREATE TABLE ActivityRoles (
    activityID INT NOT NULL,
    activityRole VARCHAR(50) NOT NULL,
    PRIMARY KEY (activityID, activityRole),
    FOREIGN KEY (activityID) REFERENCES Activity(activityID)
);

CREATE TABLE Feedback (
    feedbackID INT PRIMARY KEY,
    activityID INT NOT NULL,
    memberID INT NOT NULL,
    feedbackMessage TEXT NOT NULL,
    feedbackDateTime DATETIME NOT NULL,
    FOREIGN KEY (activityID) REFERENCES Activity(activityID),
    FOREIGN KEY (memberID) REFERENCES Member(memberID)
);

CREATE TABLE wsRegist (
    workshopID INT NOT NULL,
    memberID INT NOT NULL,
    role VARCHAR(30) NOT NULL,
    expectation TEXT NOT NULL,
    datetime DATETIME NOT NULL,
    PRIMARY KEY (memberID, workshopID),
    FOREIGN KEY (memberID) REFERENCES Member(memberID),
    FOREIGN KEY (workshopID) REFERENCES Workshop(workshopID)
);

CREATE TABLE pjRegist (
    projectID INT NOT NULL,
    memberID INT NOT NULL,
    role VARCHAR(30) NOT NULL,
    expectation TEXT NOT NULL,
    datetime DATETIME NOT NULL,
    PRIMARY KEY (memberID, projectID),
    FOREIGN KEY (memberID) REFERENCES Member(memberID),
    FOREIGN KEY (projectID) REFERENCES Project(projectID)
);
