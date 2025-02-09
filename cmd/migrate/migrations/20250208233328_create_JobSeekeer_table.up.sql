CREATE TABLE IF NOT EXISTS JobSeeker (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    firstName VARCHAR(100) NOT NULL,
    lastName VARCHAR(100) NOT NULL,
    profileSummary TEXT,
    skills JSON,  
    experience INT,
    education VARCHAR(255),
    userID INT UNSIGNED,
    FOREIGN KEY (userID) REFERENCES User(id) ON DELETE CASCADE
)