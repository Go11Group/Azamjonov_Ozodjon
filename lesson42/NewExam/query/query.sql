-- select  u.id, u.name,u.email,u.birthday,
--         u.password,u.created_at,u.updated_at,
--         u.deleted_at from users u  where
--     u.id in (select e.user_id from enrollments e inner join courses c on e.id=c.id )


-- select  u.id, u.name,u.email,
--         c.id,
--         c.title,
--         c.description,
--         c.created_at,
--         c.updated_at,
--         c.deleted_at,
--         u.password,
--         u.birthday,
--         u. created_at,
--         u.updated_at,
--         u.deleted_at
--    from users u  inner join  enrollments e  on u.id=e.user_id  inner join courses c on e.course_id=c.id where u.id=1


-- select  c.id,c.title, l.id,l.course_id,l.title,l.content,l.created_at,l.updated_at,l.deleted_at , c.description,c.created_at,c.updated_at,c.deleted_at from courses c inner join lessons l on c.id=l.course_id


-- select   c.id,c.title, e.id,e.user_id,e.enrollment_date,e.created_at,e.updated_at,e.deleted_at,c.description,c.created_at,c.updated_at,c.deleted_at c from courses c inner join  enrollments e on c.id=e.course_id;


--     keyinroq yoziladi
-- select  u.id,u.name,u.email,u.birthday, e.id,e.user_id,e.course_id,e.enrollment_date,e.created_at,e.updated_at,e.deleted_at ,u.password,u.created_at,u.updated_at,u.deleted_at from users u inner join  enrollments e on u.id=e.user_id where true

-- select *from  courses c  inner join  enrollments e on c.id=e.course_id where  c.created_at between  start_date and end_time;
-- select id, , course_id,title, content,created_at, updated_at, deleted_at
-- from lessons where true limit 0 offset 0;

-- select   c.id,c.title,c.description,c.created_at,c.updated_at,c.deleted_at,u.id,u.name,u.email,u.birthday ,c.created_at,c.updated_at,c.deleted_at  from courses c   inner join  enrollments e on c.id=e.course_id inner join users u on e.user_id=e.id

-- ALTER TABLE users
--     ALTER COLUMN deleted_at SET DATA TYPE bigint;

-- select  c.id ,c.title ,enrollments_count from  courses c inner join enrollments   e on c.id=e.course_id  where e
-- SELECT c.id, c.title, COUNT(*)
-- FROM courses c
--          INNER JOIN enrollments e ON c.id = e.course_id
-- WHERE e.enrollment_date BETWEEN 'start_time' AND 'end_time'
-- GROUP BY c.id, c.title;


-- lments.enrollment_date BETWEEN $2 AND $1 GROUP BY course_id, course_title ORDER BY enrollments_count DESC
-- select c.id, c.title, COUNT(e.id) AS enrollments_count from courses c join enrollments e on c.id = e.course_id where e.enrollment_date between time1 AND time2 group by  c.id, c.title order by  enrollments_count desc


-- select u.id,c.id,c.title,c.description from users  u  inner join  enrollments  e on u.id=e.user_id inner join courses c on c.id=e.course_id

select c.id,u.id,u.name,u.email  from courses c inner join  enrollments e on c.id=e.course_id inner join users u on e.user_id=u.id where c.id=
