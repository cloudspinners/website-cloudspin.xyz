
# What

This is an example of a website project using the cloudspin [website-stack](https://github.com/kief/website-stack) infrastructure code to manage hosting infrastructure for static websites.


# How to work with this site

## Prerequisites:

1. Docker. I use colima to install it on my Mac. In order to locally preview the site content, you need to enable exposing ports. Can do this by starting colima with the right parameters (`colima start --network-address`).
2. [Dojo](https://github.com/kudulab/dojo) (I install it on my Mac with homebrew)

Run dojo to download and run an instance of the kiefm/site-spin container, which has Jekyll preinstalled. After running the command, you should have a prompt open in the container instance. You can run jekyll commands to work with the site content. There's also a `go` script which wraps commonly used commands.

## Environment variables

In order to upload the content to a site instance, the following environment variables should be set:

|| WEBSITE_S3_BUCKET | The name of the S3 bucket that hosts the site contents |
|| WEBSITE_AWS_PROFILE | The AWS profile with permission to upload to the bucket |


## Commands

From the dojo prompt, you can run the script `./go` with the following commands:

	./go preview   Build and serve the website on http://localhost:4000, including unpublished posts
	./go plan      Build and serve the website on http://localhost:4000, only including published posts
	./go build     Build the website content in site/_site
	./go publish   Build the website content and upload to the website bucket

