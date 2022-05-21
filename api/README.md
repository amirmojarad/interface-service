# API Endpoints

## Auth

## File

## Movie
**Group url**: /movies
### `Search Movies`
* **`/search`** :
    - POST
    ```
    json body:
    {
        "title": $title
    }
    ```
    **Summary**:

    if movies with given title are exists in database return them, else search and query movies with given title from imdb-api. 

## User

