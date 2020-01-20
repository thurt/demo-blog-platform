# Demo Blog Platform

## Setup Instructions

Prerequisites: `docker` and `docker-compose` must be installed on your computer.

```bash
git clone <clone-url>
cd demo-blog-platform
docker-compose up
```

- UI/HTTP API Domain: `localhost:8080`
- API Specification: `localhost:8080/api/_swagger`
- Mailhog: `localhost:8025`
- gRPC API Domain: `localhost:10000`

Once `docker-compose up` has finished starting up the services, then you can open your web-browser and navigate to `localhost:8080`, which will display UI with a setup page where you will create the admin account for the blog.

After you create an admin username and password, you will be redirected to the login screen where you can re-enter the username and password to login to your blog.

> NOTE: this setup page will only display while the admin account does not exist. As admin, you also have the option to delete the admin account which will reset the blog back to this setup page.

## Walkthrough

When logged in as admin, you have access to the Post Editor which will display as a link in the top-middle of the UI, "Go to Post Editor".

### Creating and managing posts

Within the Post Editor, you can create a new post by clicking the "Create" button from the left-hand menu. When you create a new post, it will initially have no title, no text content, and will not yet be published.

As you make changes to the title and text content, the post details will be autosaved after a couple seconds of debounce time. You can also publish or unpublish any post by toggling the "Published" select box. Be aware, however, that you should wait a couple seconds after making any changes to your post in order to ensure that those changes are saved before navigating away from the post editor.

Any post which has been published will appear on your blog homepage, which you can go to by clicking the header in the top-left of the screen.

The homepage provides an index list of all of your published posts, showing their titles and organized by their date of creation. Clicking one of the titles will take you to the published post as it will look to public users.

So a published post means that anyone can see it when they visit your hosting domain (they do not have to be logged in). Although, remember that you can always unpublish any post that has previously been published and it will disappear immediately from the homepage because this blog is entirely dynamic.

### Post content

After clicking a post title to view the post, you may notice in the url bar that it contains an automatically generated url slug (based on its title). This is useful for bookmarking and sharing with others.

At the bottom of the post page, you may also add comments to this post as an admin user. This could be used to respond to other commentors of your post.

Adding comments to a post is one thing that public users cannot do. However, public users are free to create an account on your blog which gives them the capability to add comments to your posts.

In order to walkthrough creating a new user, first logout of your admin account by clicking the "Logout" link in the top-right of the UI. Once you are logged out, you will see a new link option in the top-right called "Create User".

### Creating a general user

Clicking "Create User" will start you on a 2-step process to create a new user account. The first step involves entering your email, desired username, and password. You are welcome to enter a fake email address because no real emails will be sent across the internet for this demo. Instead, a service called mailhog is used to capture smtp requests sent from the server so you can view them locally without sending real spam.

Once you finish filling out the form for step 1, press the "Register" button. For step 2, you will be asked to enter the verification code which was just sent to the email address you provided in the previous step. Since we are using mailhog, you can get the verification code by opening a new tab in your browser and navigate to `localhost:8025`. This will take you to the mailhog UI where you should see the verification confirmation email.

Copy the verification code provided in the email and paste it back in the blog website for step 2. Press "Verify" to finish creating your account. You will be redirected back to the login screen where you can now enter your new username and password to login as that user.

Now that you are logged in, you may add comments at the bottom of any post page! You may also view your profile by clicking your username link in the top-right. Within this page, you can easily get back to any post that you previously commented on or delete your account.
