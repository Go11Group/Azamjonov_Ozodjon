DROP TABLE IF EXISTS users cascade ;

CREATE TABLE users (
                       user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(100) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       birthday DATE NOT NULL,
                       password VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at BIGINT DEFAULT 0
);
DROP TABLE IF EXISTS courses cascade ;

CREATE TABLE courses (
                         course_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         title VARCHAR(100) NOT NULL,
                         description TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at BIGINT DEFAULT 0
);
DROP TABLE IF EXISTS lessons cascade ;

CREATE TABLE lessons (
                         lesson_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         course_id UUID REFERENCES courses(course_id) ON DELETE CASCADE,
                         title VARCHAR(100) NOT NULL,
                         content TEXT,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at BIGINT DEFAULT 0
);






DROP TABLE IF EXISTS enrollments cascade ;

CREATE TABLE enrollments (
                             enrollment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
                             course_id UUID REFERENCES courses(course_id) ON DELETE CASCADE,
                             enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             deleted_at BIGINT DEFAULT 0
);



-- Insert 5 rows into the users table
INSERT INTO users (user_id, name, email, birthday, password, created_at, updated_at) VALUES
                                                                                                     (uuid_generate_v4(), 'Alice', 'alice@example.com', '1990-01-01 00:00:00', 'password123', NOW(), NOW()),
                                                                                                     (uuid_generate_v4(), 'Bob', 'bob@example.com', '1985-02-02 00:00:00', 'password456', NOW(), NOW()),
                                                                                                     (uuid_generate_v4(), 'Charlie', 'charlie@example.com', '1992-03-03 00:00:00', 'password789', NOW(), NOW()),
                                                                                                     (uuid_generate_v4(), 'David', 'david@example.com', '1988-04-04 00:00:00', 'password101', NOW(), NOW()),
                                                                                                     (uuid_generate_v4(), 'Eve', 'eve@example.com', '1995-05-05 00:00:00', 'password102', NOW(), NOW());

-- Insert 5 rows into the courses table
INSERT INTO courses (course_id, title, description, created_at, updated_at) VALUES
                                                                                            (uuid_generate_v4(), 'Course 1', 'Description for Course 1', NOW(), NOW()),
                                                                                            (uuid_generate_v4(), 'Course 2', 'Description for Course 2', NOW(), NOW()),
                                                                                            (uuid_generate_v4(), 'Course 3', 'Description for Course 3', NOW(), NOW()),
                                                                                            (uuid_generate_v4(), 'Course 4', 'Description for Course 4', NOW(), NOW()),
                                                                                            (uuid_generate_v4(), 'Course 5', 'Description for Course 5', NOW(), NOW());
-- update Users set
--     deleted_at = date_part('epoch', current_timestamp)::INT
-- where user_id = $1 and deleted_at = 0
