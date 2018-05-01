-- create the published_posts table using posts schema
CREATE TABLE `published_posts` like `posts`;

-- copy published posts into published_posts table
INSERT INTO `published_posts` 
SELECT * FROM `posts`
WHERE published=TRUE;

-- delete the original copy of the rows
DELETE FROM `posts`
WHERE published=TRUE;

-- remove the published column from posts and published_posts
ALTER TABLE `posts`
DROP COLUMN `published`;

ALTER TABLE `published_posts`
DROP COLUMN `published`;
