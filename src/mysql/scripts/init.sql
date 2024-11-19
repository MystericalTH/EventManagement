CREATE DATABASE IF NOT EXISTS ClubManagement;
USE ClubManagement;
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%';
-- Table structure for table `Admin`
CREATE TABLE `Admin` (
  `adminID` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(320) NOT NULL,
  PRIMARY KEY (`adminID`)
);

-- Table structure for table `Activity`
CREATE TABLE `Activity` (
  `activityID` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `proposer` int(11) NOT NULL,
  `startDate` date NOT NULL,
  `endDate` date NOT NULL,
  `maxNumber` int(11),
  `format` varchar(10),
  `description` text NOT NULL,
  `proposeDateTime` datetime NOT NULL,
  `acceptAdmin` int(11),
  `acceptDateTime` datetime,
  `applicationStatus` varchar(20),
  PRIMARY KEY (`activityID`)
);

-- Table structure for table `Developer`
CREATE TABLE `Developer` (
  `developerID` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(320) NOT NULL,
  PRIMARY KEY (`developerID`)
);

-- Table structure for table `Member`
CREATE TABLE `Member` (
  `memberID` int(11) NOT NULL AUTO_INCREMENT,
  `fName` varchar(255) NOT NULL,
  `lName` varchar(255) NOT NULL,
  `email` varchar(320) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `githubUrl` varchar(320) NOT NULL,
  `interest` text NOT NULL,
  `reason` text NOT NULL,
  `acceptDateTime` datetime,
  `acceptAdmin` int(11),
  PRIMARY KEY (`memberID`),
  CONSTRAINT `member_ibfk_1` FOREIGN KEY (`acceptAdmin`) REFERENCES `Admin` (`adminID`)
);

-- Table structure for table `Project`
CREATE TABLE `Project` (
  `projectID` int(11) NOT NULL AUTO_INCREMENT,
  `advisor` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`projectID`),
  CONSTRAINT `project_ibfk_1` FOREIGN KEY (`projectID`) REFERENCES `Activity` (`activityID`)
);

-- Table structure for table `Workshop`
CREATE TABLE `Workshop` (
  `workshopID` int(11) NOT NULL AUTO_INCREMENT,
  `startTime` time NOT NULL,
  `endTime` time NOT NULL,
  PRIMARY KEY (`workshopID`)
);

-- Table structure for table `ActivityRoles`
CREATE TABLE `ActivityRoles` (
  `activityID` int(11) NOT NULL,
  `activityRole` varchar(50) NOT NULL,
  PRIMARY KEY (`activityID`, `activityRole`),
  CONSTRAINT `activityroles_ibfk_1` FOREIGN KEY (`activityID`) REFERENCES `Activity` (`activityID`)
);

-- Table structure for table `chatDevAd`
CREATE TABLE `chatDevAd` (
  `messageID` int(11) NOT NULL AUTO_INCREMENT,
  `adminID` int(11) NOT NULL,
  `developerID` int(11) NOT NULL,
  `message` text NOT NULL,
  `datetime` datetime NOT NULL,
  PRIMARY KEY (`messageID`),
  KEY `adminID` (`adminID`),
  KEY `developerID` (`developerID`),
  CONSTRAINT `chatdevad_ibfk_1` FOREIGN KEY (`adminID`) REFERENCES `Admin` (`adminID`),
  CONSTRAINT `chatdevad_ibfk_2` FOREIGN KEY (`developerID`) REFERENCES `Developer` (`developerID`)
);

-- Table structure for table `Feedback`
CREATE TABLE `Feedback` (
  `feedbackID` int(11) NOT NULL AUTO_INCREMENT,
  `activityID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `feedbackMessage` text NOT NULL,
  `feedbackDateTime` datetime NOT NULL,
  PRIMARY KEY (`feedbackID`),
  KEY `activityID` (`activityID`),
  KEY `memberID` (`memberID`),
  CONSTRAINT `feedback_ibfk_1` FOREIGN KEY (`activityID`) REFERENCES `Activity` (`activityID`),
  CONSTRAINT `feedback_ibfk_2` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`)
);

-- Table structure for table `pjRegist`
CREATE TABLE `pjRegist` (
  `projectID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `role` varchar(30) NOT NULL,
  `expectation` text NOT NULL,
  `datetime` datetime NOT NULL,
  PRIMARY KEY (`memberID`, `projectID`),
  KEY `projectID` (`projectID`),
  CONSTRAINT `pjregist_ibfk_1` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`),
  CONSTRAINT `pjregist_ibfk_2` FOREIGN KEY (`projectID`) REFERENCES `Project` (`projectID`)
);

-- Table structure for table `wsRegist`
CREATE TABLE `wsRegist` (
  `workshopID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `role` varchar(30) NOT NULL,
  `expectation` text NOT NULL,
  `datetime` datetime NOT NULL,
  PRIMARY KEY (`memberID`, `workshopID`),
  KEY `workshopID` (`workshopID`),
  CONSTRAINT `wsregist_ibfk_1` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`),
  CONSTRAINT `wsregist_ibfk_2` FOREIGN KEY (`workshopID`) REFERENCES `Workshop` (`workshopID`)
);

