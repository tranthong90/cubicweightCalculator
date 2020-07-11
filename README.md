# Cubic Weight Calculator

This is a application which will retrieve the products data from an API and calculate the average cubic weight for a give category

### Language

Go 1.14

### Prerequisites

* Docker:          <https://www.docker.com/>

### Run the app with docker-compose

The app allows user to set the Source URL and the category they want to get the average cubic weight for. However it has the default settings:
SourceURL : http://wp8m3he1wt.s3-website-ap-southeast-2.amazonaws.com/api/products/1
Category: "Air Conditioners"

To change those settings, you need to set up the environment variables in your local machine

For example,
To will set the TEST environment variable to "DO"

For Mac/ Linux
TEST=DO docker-compose up 

For Windows (powershell)
$env:TEST="do";docker-compose up

So to run the app with default settings, please run:

```bash
docker-compose up
```

To change the category, please run

```bash
CATEGORY={Your category} docker-compose build
CATEGORY={Your category} docker-compose up
```