# hacktoberfest
Check your Hacktoberfest progress

```
  curl http://localhost:5000/api/<GITHUB_USER_NAME>
```

```
  {
    "count": "",
    "pull_requests": [
      {
        "pull_request_number": "",
        "title": "",
        "user": "",
        "state": "",
        "repo_name": "",
        "url": "",
        "created_at": ""
      }
    ]
  }
```

# Frontend
A web GUI frontend is served from the static directory. 

To set this up you have to first install dependencies with:
```npm install```

Then you configure where your API endpoint is in the `static/.env` file.  

After that you build the Vue application with the following command
```npm run build```
