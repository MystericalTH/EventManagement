CREATE DATABASE IF NOT EXISTS ClubManagement;
USE ClubManagement;
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%';

-- Table structure for Admin
CREATE TABLE Admin (
  adminID int(11) NOT NULL AUTO_INCREMENT,
  fName varchar(255) NOT NULL,
  lName varchar(255) NOT NULL,
  email varchar(320) NOT NULL,
  PRIMARY KEY (adminID)
);

-- Table structure for Member
CREATE TABLE Member (
  memberID int(11) NOT NULL AUTO_INCREMENT,
  fName varchar(255) NOT NULL,
  lName varchar(255) NOT NULL,
  email varchar(320) UNIQUE NOT NULL,
  phone varchar(20) UNIQUE NOT NULL,
  githubUrl varchar(320) UNIQUE NOT NULL,
  interest text NOT NULL,
  reason text NOT NULL,
  acceptDateTime datetime,
  acceptAdmin int(11),
  PRIMARY KEY (memberID),
  CONSTRAINT member_ibfk_1 FOREIGN KEY (acceptAdmin) REFERENCES Admin (adminID)
  ON DELETE RESTRICT ON UPDATE CASCADE
);

-- Table structure for Activity (superclass)
CREATE TABLE Activity (
  activityID int(11) NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  proposer int(11) NOT NULL,
  startDate date NOT NULL,
  endDate date NOT NULL,
  maxParticipant int(11) NOT NULL,
  format varchar(10) NOT NULL,
  description text NOT NULL,
  proposeDateTime datetime NOT NULL,
  acceptAdmin int(11),
  acceptDateTime datetime,
  applicationStatus varchar(20) ,
  PRIMARY KEY (activityID),
  CONSTRAINT activity_proposer_fk FOREIGN KEY (proposer) REFERENCES Member (memberID), -- to do set null if have time to do
  CONSTRAINT activity_acceptAdmin_fk FOREIGN KEY (acceptAdmin) REFERENCES Admin (adminID)
  ON DELETE RESTRICT ON UPDATE CASCADE
);

-- Subclass: Table structure for Project
CREATE TABLE Project (
  projectID int(11) NOT NULL,
  advisor varchar(255) DEFAULT NULL,
  PRIMARY KEY (projectID),
  CONSTRAINT project_fk FOREIGN KEY (projectID) REFERENCES Activity (activityID)
  ON DELETE CASCADE ON UPDATE CASCADE
);

-- Subclass: Table structure for Workshop
CREATE TABLE Workshop (
  workshopID int(11) NOT NULL,
  startTime varchar(255) NOT NULL,
  endTime varchar(255) NOT NULL,
  PRIMARY KEY (workshopID),
  CONSTRAINT workshop_fk FOREIGN KEY (workshopID) REFERENCES Activity (activityID)
  ON DELETE CASCADE ON UPDATE CASCADE
);

-- Table structure for ActivityRoles
CREATE TABLE ActivityRoles (
  activityID int(11) NOT NULL,
  activityRole varchar(50) NOT NULL,
  PRIMARY KEY (activityID, activityRole),
  CONSTRAINT activityroles_fk FOREIGN KEY (activityID) REFERENCES Activity (activityID)
  ON DELETE CASCADE ON UPDATE CASCADE
);

-- Table structure for Developer
CREATE TABLE Developer (
  developerID int(11) NOT NULL AUTO_INCREMENT,
  fName varchar(255) NOT NULL,
  lName varchar(255) NOT NULL,
  email varchar(320) NOT NULL,
  PRIMARY KEY (developerID)
);

-- Table structure for chatDevAd
CREATE TABLE chatDevAd (
  messageID int(11) NOT NULL AUTO_INCREMENT,
  adminID int(11) NOT NULL,
  developerID int(11) NOT NULL,
  sender varchar(30) NOT NULL,
  message text NOT NULL,
  timesent datetime NOT NULL,
  PRIMARY KEY (messageID),
  KEY adminID (adminID),
  KEY developerID (developerID),
  CONSTRAINT chatdevad_admin_fk FOREIGN KEY (adminID) REFERENCES Admin (adminID)
  ON DELETE RESTRICT ON UPDATE  CASCADE,
  CONSTRAINT chatdevad_dev_fk FOREIGN KEY (developerID) REFERENCES Developer (developerID)
  ON DELETE RESTRICT ON UPDATE  CASCADE
);

-- Table structure for Feedback
CREATE TABLE Feedback (
  feedbackID int(11) NOT NULL AUTO_INCREMENT,
  activityID int(11) NOT NULL,
  memberID int(11) NOT NULL,
  feedbackMessage text NOT NULL,
  feedbackDateTime datetime NOT NULL,
  PRIMARY KEY (feedbackID),
  KEY activityID (activityID),
  KEY memberID (memberID),
  CONSTRAINT feedback_activity_fk FOREIGN KEY (activityID) REFERENCES Activity (activityID), -- to do set null if have time to do ,
  CONSTRAINT feedback_member_fk FOREIGN KEY (memberID) REFERENCES Member (memberID)
  ON DELETE RESTRICT ON UPDATE RESTRICT
);

-- Table structure for associative entity ActivityRegistration
CREATE TABLE ActivityRegistration (
  activityID int(11) NOT NULL,
  memberID int(11) NOT NULL,
  role varchar(30) NOT NULL,
  expectation text NOT NULL,
  datetime datetime NOT NULL,
  PRIMARY KEY (memberID, activityID),
  KEY activityID (activityID),
  CONSTRAINT activityRegist_member_fk FOREIGN KEY (memberID) REFERENCES Member (memberID)
  ON DELETE CASCADE ON UPDATE CASCADE
  ,
  CONSTRAINT activityRegist_activity_fk FOREIGN KEY (activityID) REFERENCES Activity (activityID)
  ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE UNIQUE INDEX index_member_email
ON Member (email);

CREATE UNIQUE INDEX index_admin_email
ON Admin (email);

CREATE UNIQUE INDEX index_developer_email
ON Developer (email);

CREATE UNIQUE INDEX index_activity_title
ON Activity (title);
