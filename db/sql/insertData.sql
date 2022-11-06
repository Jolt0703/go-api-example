insert into articles (title, contents, username, nice, created_at)
values('1st article', 'first blog post', 'John Smith', 2, now());

insert into articles (title, contents, username, nice)
values ('2nd article', 'Second blog post', 'John Doe', 4);

insert into comments (article_id, message, created_at)
values (1, '1st comment', now());

insert into comments (article_id, message)
values (1, '2nd comment');