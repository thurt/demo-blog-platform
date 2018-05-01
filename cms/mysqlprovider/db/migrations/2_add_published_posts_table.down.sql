-- note: this is a destructive change since the published_posts rows are not used at all in this migration.
-- this means that the posts rows will become the published version. 
-- this could be bad for the user if they have unpublished content in their posts that they dont want revealed yet.
-- also, since the published_posts rows are deleted, it means that the user cannot retreive the version of the post at the last published version (only the currenty unpublished version will remain and that version will become the published version)

-- add published column to posts table
ALTER TABLE `posts`
ADD COLUMN `published` BOOLEAN NOT NULL DEFAULT 0;

-- change published to true for all rows in the posts table that are found in the published_posts table
UPDATE `posts` 
INNER JOIN `published_posts` ON posts.id=published_posts.id
SET posts.published=TRUE;

-- remove published_posts table
DROP TABLE `published_posts`;
