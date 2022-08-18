# status-checker

Status checker for some common websites that exist online. The application continuously polls the status of these websites and exposes APIs to retrieve the information.

# APIs

## 1. /POST/websites

### Sample Input (/POST/websites)

    {
        "websites": ["https://www.youtube.com", "https://www.google.com", "https://www.facebook.com", "https:www.fakewebsite123.com"]
    }

### Sample Response (/POST/websites)

    List of websites updated to [https://www.youtube.com https://www.google.com https://www.facebook.com https:www.fakewebsite123.com] Successfully.

## 2. /GET/websites
    Query param (optional) ?name=www.facebook.com

### Sample Response (/GET/websites)

    {
        "https://www.facebook.com": "UP",
        "https://www.google.com": "UP",
        "https://www.youtube.com": "UP",
        "https:www.fakewebsite123.com": "DOWN"
    }

### Sample Response (/GET/websites?name=https://www.facebook.com)

    {
        "https://www.facebook.com": "UP"
    }