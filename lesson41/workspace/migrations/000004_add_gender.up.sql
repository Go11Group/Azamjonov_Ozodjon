create type gender as enum('Male', 'Female');

alter table examples add gender gender;