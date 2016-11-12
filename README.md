# test stackdriver trace

```sh
$ go install github.com/tmc/teststackdrivertrace
$ teststackdrivertrace -projectID ${GCLOUD_PROJECT_ID} -serviceAccountKey ${PATH_TO_SERVICEACCOUNT_JSON} &
$ curl -s --header "X-Cloud-Trace-Context: 205445aa7843bc8bf206b12000100000/0;o=1" localhost:9000
```

