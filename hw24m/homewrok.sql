create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_name varchar
);

create table author
(
    id   int primary key,
    name varchar not null
);

INSERT INTO book (id, name, page, author_name) VALUES
(1, 'To Kill a Mockingbird', 281, 'Harper Lee'),
(2, '1984', 328, 'George Orwell'),
(3, 'The Great Gatsby', 180, 'F. Scott Fitzgerald'),
(4, 'Pride and Prejudice', 279, 'Jane Austen'),
(5, 'The Catcher in the Rye', 214, 'J.D. Salinger'),
(6, 'Moby-Dick', 635, 'Herman Melville'),
(7, 'War and Peace', 1225, 'Leo Tolstoy'),
(8, 'The Odyssey', 541, 'Homer'),
(9, 'Crime and Punishment', 671, 'Fyodor Dostoevsky'),
(10, 'The Brothers Karamazov', 824, 'Fyodor Dostoevsky'),
(11, 'Brave New World', 268, 'Aldous Huxley'),
(12, 'Ulysses', 730, 'James Joyce'),
(13, 'The Iliad', 704, 'Homer'),
(14, 'Jane Eyre', 507, 'Charlotte Brontë'),
(15, 'The Divine Comedy', 798, 'Dante Alighieri');


INSERT INTO author (id, name) VALUES
(1, 'Harper Lee'),
(2, 'George Orwell'),
(3, 'F. Scott Fitzgerald'),
(4, 'Jane Austen'),
(5, 'J.D. Salinger'),
(6, 'Herman Melville'),
(7, 'Leo Tolstoy');


select * from book;

select * from author;

select * from book join author on author.name = book.author_name;

select * from book left join author on author.name = book.author_name;

select * from book right join author on author.name = book.author_name;

select book.id as book_id, book.name as book_name, book.page, book.author_name, author.id as author_id, author.name as author_namefrom book cross join author;


select b.id, b.name, page, a.name from author as a
    full join book as b on b.author_name =a.name
where b.author_name is null or a.name is null;