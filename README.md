# Label Printer

An API server to create a tag image for a text and print it through the NIIMBOT D110 label printer.

## Requirements

- Python `>=3.10`
- Golang
- Linux (macOS doesn't support Bluetooth sockets, not tested on Windows)

## Getting started

- In the `niimprint` directory, run 
    
    ```
    pip3 install -r requirements.txt
    ```
- Run golang server
    
    ```
    go run .
    ```

- POST request on `http://localhost:8769/print` with payload in the following JSON format

    ```json
    {
        "text" : "MYLABEL",
        "qr_text" : "https://www.example.com/MYLABEL"
    }
    ```

## cURL

```shell
curl -d '{"text": "MYLABEL", "qr_text": "https://www.example.com/MYLABEL"}' \
     -H "Content-Type: application/json" \
     -X POST \
     http://localhost:8769/print
```

## Authors

* [Faizan Khalid](https://github.com/IamFaizanKhalid)
