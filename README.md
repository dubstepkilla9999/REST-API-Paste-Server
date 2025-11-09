# REST-API-Paste-Server
This simple REST web service stores any text and makes it available via a new randomized link (similar to some existing "paste" websites). It uses gin to do so.

Tested in a local network, using curl.

**Endpoints:**
*   **GET /storage/:id**: Retrieve a previously stored text by its ID.
*   **POST /storage**: Create a new entry in the storage. Returns a JSON containing the text and its new randomly generated ID. You will then be able to retrieve this text by its ID.
*   post request body:
  ```json
    {
      "content": "This is a test string."
    }
