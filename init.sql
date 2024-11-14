DROP DATABASE IF EXISTS ClubManagement;
CREATE DATABASE IF NOT EXISTS ClubManagement;
USE ClubManagement;

--
-- Table structure for table `Activity`
--

CREATE TABLE `Activity` (
  `activityID` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `proposer` int(11) NOT NULL,
  `startDate` date NOT NULL,
  `endDate` date NOT NULL,
  `maxNumber` int(11) DEFAULT NULL,
  `format` varchar(10) DEFAULT NULL,
  `description` text NOT NULL,
  `proposeDateTime` datetime NOT NULL,
  `acceptAdmin` int(11) NOT NULL,
  `acceptDateTime` datetime NOT NULL,
  `applicationStatus` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `ActivityRoles`
--

CREATE TABLE `ActivityRoles` (
  `activityID` int(11) NOT NULL,
  `activityRole` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Admin`
--

CREATE TABLE `Admin` (
  `adminID` int(11) NOT NULL,
  `email` varchar(320) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `chatDevAd`
--

CREATE TABLE `chatDevAd` (
  `messageID` int(11) NOT NULL,
  `adminID` int(11) NOT NULL,
  `developerID` int(11) NOT NULL,
  `message` text NOT NULL,
  `datetime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Developer`
--

CREATE TABLE `Developer` (
  `developerID` int(11) NOT NULL,
  `email` varchar(320) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Feedback`
--

CREATE TABLE `Feedback` (
  `feedbackID` int(11) NOT NULL,
  `activityID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `feedbackMessage` text NOT NULL,
  `feedbackDateTime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Member`
--

CREATE TABLE `Member` (
  `memberID` int(11) NOT NULL,
  `fName` varchar(255) NOT NULL,
  `lName` varchar(255) NOT NULL,
  `email` varchar(320) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `githubUrl` varchar(320) DEFAULT NULL,
  `interest` text NOT NULL,
  `reason` text NOT NULL,
  `acceptDateTime` datetime DEFAULT NULL,
  `acceptAdmin` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `pjRegist`
--

CREATE TABLE `pjRegist` (
  `projectID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `role` varchar(30) NOT NULL,
  `expectation` text NOT NULL,
  `datetime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Project`
--

CREATE TABLE `Project` (
  `projectID` int(11) NOT NULL,
  `advisor` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `Workshop`
--

CREATE TABLE `Workshop` (
  `workshopID` int(11) NOT NULL,
  `startTime` time NOT NULL,
  `endTime` time NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `wsRegist`
--

CREATE TABLE `wsRegist` (
  `workshopID` int(11) NOT NULL,
  `memberID` int(11) NOT NULL,
  `role` varchar(30) NOT NULL,
  `expectation` text NOT NULL,
  `datetime` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Activity`
--
ALTER TABLE `Activity`
  ADD PRIMARY KEY (`activityID`),
  ADD KEY `acceptAdmin` (`acceptAdmin`),
  ADD KEY `proposer` (`proposer`);

--
-- Indexes for table `ActivityRoles`
--
ALTER TABLE `ActivityRoles`
  ADD PRIMARY KEY (`activityID`,`activityRole`);

--
-- Indexes for table `Admin`
--
ALTER TABLE `Admin`
  ADD PRIMARY KEY (`adminID`);

--
-- Indexes for table `chatDevAd`
--
ALTER TABLE `chatDevAd`
  ADD PRIMARY KEY (`messageID`),
  ADD KEY `adminID` (`adminID`),
  ADD KEY `developerID` (`developerID`);

--
-- Indexes for table `Developer`
--
ALTER TABLE `Developer`
  ADD PRIMARY KEY (`developerID`);

--
-- Indexes for table `Feedback`
--
ALTER TABLE `Feedback`
  ADD PRIMARY KEY (`feedbackID`),
  ADD KEY `activityID` (`activityID`),
  ADD KEY `memberID` (`memberID`);

--
-- Indexes for table `Member`
--
ALTER TABLE `Member`
  ADD PRIMARY KEY (`memberID`),
  ADD KEY `acceptAdmin` (`acceptAdmin`);

--
-- Indexes for table `pjRegist`
--
ALTER TABLE `pjRegist`
  ADD PRIMARY KEY (`memberID`,`projectID`),
  ADD KEY `projectID` (`projectID`);

--
-- Indexes for table `Project`
--
ALTER TABLE `Project`
  ADD PRIMARY KEY (`projectID`);

--
-- Indexes for table `Workshop`
--
ALTER TABLE `Workshop`
  ADD PRIMARY KEY (`workshopID`);

--
-- Indexes for table `wsRegist`
--
ALTER TABLE `wsRegist`
  ADD PRIMARY KEY (`memberID`,`workshopID`),
  ADD KEY `workshopID` (`workshopID`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Activity`
--
ALTER TABLE `Activity`
  ADD CONSTRAINT `activity_ibfk_1` FOREIGN KEY (`acceptAdmin`) REFERENCES `Admin` (`adminID`),
  ADD CONSTRAINT `activity_ibfk_2` FOREIGN KEY (`proposer`) REFERENCES `Member` (`memberID`);

--
-- Constraints for table `ActivityRoles`
--
ALTER TABLE `ActivityRoles`
  ADD CONSTRAINT `activityroles_ibfk_1` FOREIGN KEY (`activityID`) REFERENCES `Activity` (`activityID`);

--
-- Constraints for table `chatDevAd`
--
ALTER TABLE `chatDevAd`
  ADD CONSTRAINT `chatdevad_ibfk_1` FOREIGN KEY (`adminID`) REFERENCES `Admin` (`adminID`),
  ADD CONSTRAINT `chatdevad_ibfk_2` FOREIGN KEY (`developerID`) REFERENCES `Developer` (`developerID`);

--
-- Constraints for table `Feedback`
--
ALTER TABLE `Feedback`
  ADD CONSTRAINT `feedback_ibfk_1` FOREIGN KEY (`activityID`) REFERENCES `Activity` (`activityID`),
  ADD CONSTRAINT `feedback_ibfk_2` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`);

--
-- Constraints for table `Member`
--
ALTER TABLE `Member`
  ADD CONSTRAINT `member_ibfk_1` FOREIGN KEY (`acceptAdmin`) REFERENCES `Admin` (`adminID`);

--
-- Constraints for table `pjRegist`
--
ALTER TABLE `pjRegist`
  ADD CONSTRAINT `pjregist_ibfk_1` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`),
  ADD CONSTRAINT `pjregist_ibfk_2` FOREIGN KEY (`projectID`) REFERENCES `Project` (`projectID`);

--
-- Constraints for table `Project`
--
ALTER TABLE `Project`
  ADD CONSTRAINT `project_ibfk_1` FOREIGN KEY (`projectID`) REFERENCES `Activity` (`activityID`);

--
-- Constraints for table `Workshop`
--
ALTER TABLE `Workshop`
  ADD CONSTRAINT `workshop_ibfk_1` FOREIGN KEY (`workshopID`) REFERENCES `Activity` (`activityID`);

--
-- Constraints for table `wsRegist`
--
ALTER TABLE `wsRegist`
  ADD CONSTRAINT `wsregist_ibfk_1` FOREIGN KEY (`memberID`) REFERENCES `Member` (`memberID`),
  ADD CONSTRAINT `wsregist_ibfk_2` FOREIGN KEY (`workshopID`) REFERENCES `Workshop` (`workshopID`);