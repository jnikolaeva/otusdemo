This is demo service for Otus architect course. It implements simple CRUD user REST API.

To run service with helm use command:

````
make run
````

To remove:

```
make remove
```

To run without helm use:

```
make k8s-run
```

To remove:

```
make k8s-remove
```

To run end-2-end tests using newman:

```
newman run ./api/api.postman_collection.json
```

After running service is available by url http://arch.homework/otusapp/.