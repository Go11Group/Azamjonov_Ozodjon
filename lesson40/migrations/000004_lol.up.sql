CREATE TABLE Enrollments (
                             EnrollmentID INT PRIMARY KEY,
                             StudentID INT,
                             CourseID INT,
                             Semester VARCHAR(20),
                             FOREIGN KEY (StudentID) REFERENCES Students(StudentID),
                             FOREIGN KEY (CourseID) REFERENCES Courses(CourseID)
);

-- Inserting data into 'Enrollments'
INSERT INTO Enrollments (EnrollmentID, StudentID, CourseID, Semester) VALUES
                                                                          (1, 1, 101, 'Fall 2024'),
                                                                          (2, 2, 102, 'Spring 2024');
