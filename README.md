# Benchmarking Web Servers

The purpose of this repository is to provide a benchmarking suite across a number of different web server implementation in a number
of languages. Each language/framework is detailed more below.

## General strategy
Each server is run on a remote server (not locally) to reduce the effect of processes already running on the machines. [K6](https://k6.io/)
was chosen as the testing tool, simply due to prior experience.

A SQLite database will be running on the machine to service the testing. Between each run the DB will be wiped and re-seeded.

The following tests are run against each implementation:
  * Simple healthcheck request.
    * This does not require any complex logic on the server. A simple string is returned rather than JSON to reduce the conputational
      effects of marshalling.
  * Another healthcheck style endpoint, returning JSON
    * Slightly more involved, as marshalling to JSON is required.
  * A GET request to retrieve data from a SQLite database
    * This will just be a simple
  * A POST request to add data to the database
    * This will be passed as a JSON body which will need to be decoded before insertion.

## Implementations

### Golang w/ Echo web server
The [Echo framework](https://echo.labstack.com/) is a well known and tested framework for creating webservers in Golang. Personally, I have
the most experience with it.

### Rust w/ Rocket
The [Rocket framework](https://rocket.rs/) is a well known framework. This is the language that I have the least knowledge in though, so
there are likely efficiency/performance gains that could be made by someone with more experience.

### Javascript w/ Express
The [Express framework](https://expressjs.com/) is pretty much old faithful to a lot of people these days. Including this more out of
curiosity than any expectation of high performance.

### NestJS
[NestJS](https://nestjs.com/) is a framework used across the whole organisation, and is the presumed default for new services at this point. This
implementation is that primary point of data for comparing against.

## Extensions
The following are possible extensions for this work if we wanted to explore this further.
  * Making more complex requests, perhaps through more complex data query requirements in the server.
  * Reaching out to external services from the web server (eg: AWS SSM)
