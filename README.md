This is demo service for Otus architect course. It implements simple CRUD user REST API.

Before running with helm at first time update helm dependencies:

```
make helm-update-dependencies
```

Then use:

````
make run
````

Ensure everything is running:

```
$ kubectl get all
NAME                                 READY   STATUS    RESTARTS   AGE
pod/otusdemo-685f9cbd7b-g5mrr        1/1     Running   2          65s
pod/otusdemo-685f9cbd7b-nsttc        1/1     Running   2          65s
pod/otusdemo-685f9cbd7b-schxk        1/1     Running   2          65s
pod/otusdemo-postgresql-0            1/1     Running   0          65s

NAME                                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/otusdemo                       NodePort    10.111.58.236   <none>        9000:31119/TCP   65s
service/otusdemo-postgresql            ClusterIP   10.100.24.81    <none>        5432/TCP         65s
service/otusdemo-postgresql-headless   ClusterIP   None            <none>        5432/TCP         65s

NAME                            READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/otusdemo        3/3     3            3           65s

NAME                                       DESIRED   CURRENT   READY   AGE
replicaset.apps/otusdemo-685f9cbd7b        3         3         3       65s

NAME                                   READY   AGE
statefulset.apps/otusdemo-postgresql   1/1     65s
```

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