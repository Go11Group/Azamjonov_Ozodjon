CREATE TABLE users (
                       user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       birthday DATE,
                       password VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP
);


CREATE TABLE Courses (
                         course_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         title VARCHAR(100) NOT NULL,
                         description TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP
);

CREATE TABLE Lessons (
                         lesson_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         course_id UUID REFERENCES Courses(course_id) ON DELETE CASCADE,
                         title VARCHAR(100) NOT NULL,
                         content TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP
);

CREATE TABLE Enrollments (
                             enrollment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             user_id UUID REFERENCES Users(user_id) ON DELETE CASCADE,
                             course_id UUID REFERENCES Courses(course_id) ON DELETE CASCADE,
                             enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP
);


-- Insert 5 rows into the users table
INSERT INTO users (user_id, name, email, birthday, password, created_at, updated_at, deleted_at) VALUES
                                                                                                     (uuid_generate_v4(), 'Alice', 'alice@example.com', '1990-01-01 00:00:00', 'password123', NOW(), NOW(), NULL),
                                                                                                     (uuid_generate_v4(), 'Bob', 'bob@example.com', '1985-02-02 00:00:00', 'password456', NOW(), NOW(), NULL),
                                                                                                     (uuid_generate_v4(), 'Charlie', 'charlie@example.com', '1992-03-03 00:00:00', 'password789', NOW(), NOW(), NULL),
                                                                                                     (uuid_generate_v4(), 'David', 'david@example.com', '1988-04-04 00:00:00', 'password101', NOW(), NOW(), NULL),
                                                                                                     (uuid_generate_v4(), 'Eve', 'eve@example.com', '1995-05-05 00:00:00', 'password102', NOW(), NOW(), NULL);

-- Insert 5 rows into the courses table
INSERT INTO courses (course_id, title, description, created_at, updated_at, deleted_at) VALUES
                                                                                            (uuid_generate_v4(), 'Course 1', 'Description for Course 1', NOW(), NOW(), NULL),
                                                                                            (uuid_generate_v4(), 'Course 2', 'Description for Course 2', NOW(), NOW(), NULL),
                                                                                            (uuid_generate_v4(), 'Course 3', 'Description for Course 3', NOW(), NOW(), NULL),
                                                                                            (uuid_generate_v4(), 'Course 4', 'Description for Course 4', NOW(), NOW(), NULL),
                                                                                            (uuid_generate_v4(), 'Course 5', 'Description for Course 5', NOW(), NOW(), NULL);
-- update Users set
--     deleted_at = date_part('epoch', current_timestamp)::INT
-- where user_id = $1 and deleted_at = 0
