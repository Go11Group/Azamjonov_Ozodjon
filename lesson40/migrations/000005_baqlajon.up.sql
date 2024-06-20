CREATE TABLE Professors (
                            ProfessorID INT PRIMARY KEY,
                            FirstName VARCHAR(50),
                            LastName VARCHAR(50),
                            Department VARCHAR(100)
);

-- Inserting data into 'Professors'
INSERT INTO Professors (ProfessorID, FirstName, LastName, Department) VALUES
                                                                          (1, 'Alice', 'Johnson', 'Mathematics'),
                                                                          (2, 'Bob', 'Williams', 'Physics');