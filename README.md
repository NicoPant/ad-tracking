# ad-tracking

This project implements a microservices-based ad tracking system using **Go**, **gRPC**, **MongoDB**, and **Docker**. **Dragonfly** is used for caching but is to be implemented later.

It includes two main services:

- **Ad Service (`ad`)**: Manages advertisements with title, description and url. It also serves ad to the tracker service
- **Tracker Service (`tracker`)**: Keep a count of a specific ad that was served.


##  Prerequisites
Make sure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)

Run the following command from the root of the project:

```bash
make install
```
This command will:

- Build Docker images for ad and tracker.

- Start all services using Docker Compose.

- Load environment variables from each service's .env file.

- Set up MongoDB and expose gRPC ports.

## Usage

ad is available at **localhost:8000** ((inside Docker network as adservice:8000)

tracker is available at **localhost:9000** ((inside Docker network as trackerservice:9000)

It is to be tested with **Postman** or any other gRPC client. If you do use Postman, make sure to use the **`ad.proto`** file.

Here are the different requests that you can make:

#### CreateAd

It creates an Ad

Message: 
```json
{
    "title": "title of the test",
    "description": "test of the endpoint",
    "url": "test.com"
}
```

Response (the id is a generated)

```json
{
    "id": "a97eb838-01d9-40f3-910a-b61d7f5bed98"
}
```

#### GetById

Get and Ad by its unique generated Id

Message:

```json
{
    "id": "a97eb838-01d9-40f3-910a-b61d7f5bed98"
}
```

Response: 

```json
{
    "ad": {
        "id": "a97eb838-01d9-40f3-910a-b61d7f5bed98",
        "title": "title of the test",
        "description": "test of the endpoint",
        "url": "test.com"
    }
}
```

#### ServeAd

Increment the counter on the tracker linked to a specific ad and returns the url of the latter.

Message

```json
{
    "adId": "a97eb838-01d9-40f3-910a-b61d7f5bed98"
}
```

Response

```json
{
    "url": "test.com" 
}
```


### Clean  Up

Once you want to stop the project: 

```bash
make down
```

This will stop the running containers a remove them.


Happy Tracking ! 
