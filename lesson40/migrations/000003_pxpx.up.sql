CREATE TABLE Courses (
                         CourseID INT PRIMARY KEY,
                         CourseName VARCHAR(100),
                         Credits INT
);

-- Inserting data into 'Courses'
INSERT INTO Courses (CourseID, CourseName, Credits) VALUES
                                                        (101, 'Mathematics', 3),
                                                        (102, 'Physics', 4);
