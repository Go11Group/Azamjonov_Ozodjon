-- Create User Table
CREATE TABLE Users (
                      id INT PRIMARY KEY,
                      name VARCHAR(100),
                      age INT
);

-- Create Problems Table
CREATE TABLE Problems (
                          id INT PRIMARY KEY,
                          description TEXT,
                          difficulty VARCHAR(50)
);

-- Create Solved_Problems Table
CREATE TABLE Solved_Problems (
                                 id SERIAL PRIMARY KEY,
                                 user_id INT,
                                 problem_id INT,
                                 solve_date DATE,
                                 FOREIGN KEY (user_id) REFERENCES Users(id),
                                 FOREIGN KEY (problem_id) REFERENCES Problems(id)
);

