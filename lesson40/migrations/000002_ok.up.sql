CREATE TABLE Students (
                          StudentID INT PRIMARY KEY,
                          FirstName VARCHAR(50),
                          LastName VARCHAR(50),
                          Age INT
);

-- Inserting data into 'Students'
INSERT INTO Students (StudentID, FirstName, LastName, Age) VALUES
                                                               (1, 'John', 'Doe', 20),
                                                               (2, 'Jane', 'Smith', 22);
