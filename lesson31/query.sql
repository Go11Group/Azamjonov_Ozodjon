create table product
(
    id       uuid primary key,
    name     varchar,
    category varchar,
    cost     int
);

SELECT
    tablename,
    indexname,
    indexdef
FROM
    pg_indexes
WHERE
    schemaname = 'public'
ORDER BY
    tablename,
    indexname;
-- single index
create  index  name_index on product(name);
create  index  name_index_hash  on  product using  hash(surname);

drop  index  name_index;
drop  index  name_index_hash;

create  unique index  name_index on product(cost);
create  index  name_index on product(category,name);

explain (analyse )
select * from product where id='a5dde9e9-1a7e-4f23-bb9e-9203260a2918';

explain (analyse )
select  * from product where category="O'Kon";

explain (analyse )
select  * from product where name='Tatyana';