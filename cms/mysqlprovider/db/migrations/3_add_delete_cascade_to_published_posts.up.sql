-- add foreign key constraint. published_posts.id is also the primary key
-- this change causes the published_posts row to be deleted when corresponding posts row is deleted
ALTER TABLE published_posts 
ADD CONSTRAINT FK_id
FOREIGN KEY (id) REFERENCES posts(id) ON DELETE CASCADE;
