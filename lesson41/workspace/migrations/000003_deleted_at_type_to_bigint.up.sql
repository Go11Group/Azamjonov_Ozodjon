alter table examples add deleted bigint;

update 
    examples
set
    deleted = extract(epoch from deleted_at);

alter table examples drop column deleted_at;
alter table examples rename column deleted to deleted_at;
