insert into casbin_rule (
    p_type, v0, v1, v2
) values 
-- Admin permissions
('p', 'admin', '/restaurant/create', 'POST'),  
('p', 'admin', '/restaurant/getall', 'GET'),  
('p', 'admin', '/restaurant/*', 'GET'),  
('p', 'admin', '/restaurant/update', 'PUT'),  
('p', 'admin', '/restaurant/*', 'DELETE'),  
('p', 'admin', '/users/create', 'POST'),  
('p', 'admin', '/users/getall', 'GET'),  
('p', 'admin', '/users/*', 'GET'),  
('p', 'admin', '/users/update', 'PUT'),  
('p', 'admin', '/users/*', 'DELETE'),  

-- User permissions
('p', 'user', '/restaurant/getall', 'GET'),
('p', 'user', '/restaurant/*', 'GET'),
('p', 'user', '/users/getall', 'GET'),
('p', 'user', '/users/*', 'GET');