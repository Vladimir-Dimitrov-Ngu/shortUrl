# ShortUrl
URL Shortener is a server application that transforms long web links into shorter, more manageable ones.

## Description
URL Shortener allows users to create short URLs for long web links. This can be useful for shortening lengthy links in messages, text messages, or on social media, making them more convenient to share.

## Installation
1. Clone the repository:
```bash
git clone https://github.com/your_username/url-shortener.git
```
2. Navigate to the project directory:
```bash
cd url-shortener
```
3. Build the Go application:
```go
go build .
```
4. Run the application
```go
go run main.go
```

## Usage
### Creating a Short Link
Send a POST request to /api/urls with a JSON body containing the long URL:
```http request
POST /api/urls HTTP/1.1
Content-Type: application/json

{
    "url": "https://example.com/very-long-url-that-we-want-to-shorten"
}

```
Receive a response with the short URL:

```json
{
    "key": "abcde"
}

```

## Redirecting with Short Link
Open a browser and enter the URL in the format http://your-server/{key}, where {key} is the key obtained after creating the short link. For example:

```
http://your-server/abcde
```

## API
### POST /api/urls
Creates a short link.

Request Body Parameters

| Parameter    | Type   | Description|
|------------|--------| --- |
| Url	| String | The long URL to be shortened.|


### GET /{key}
Redirects to the long URL associated with the short key.

Path Parameters

| Parameter | Type   | Description|
|-----------|--------| --- |
| key	      | String | The short key created when URL is added.|
