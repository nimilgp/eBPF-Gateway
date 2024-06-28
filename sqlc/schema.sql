CREATE TABLE Accounts (
    userName VARCHAR(64) PRIMARY KEY,
    passwdHash VARCHAR(128) NOT NULL,
    powerLevel INTEGER NOT NULL,
    firstName VARCHAR(32) NOT NULL,
    lastName VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL,
    valid BOOLEAN DEFAULT True 
);

CREATE TABLE BearerTokens (
    tokenString CHAR(128) PRIMARY KEY,
    validTill TIMESTAMP NOT NULL,
    userName VARCHAR(64) NOT NULL,
    valid BOOLEAN DEFAULT True,
    FOREIGN KEY(userName) REFERENCES Accounts(userName)
); 