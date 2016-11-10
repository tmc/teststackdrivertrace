# test stackdriver trace

```sh
$ go install github.com/tmc/teststackdrivertrace
$ teststackdrivertrace -projectID ${GCLOUD_PROJECT_ID} -serviceAccountKey ${PATH_TO_SERVICEACCOUNT_JSON} &
$ curl -s --header "X-Cloud-Trace-Context: 205445aa7843bc8bf206b120001000/0;o=1" localhost:9000
(*googleapi.Error)(0xc4203a2c80)(googleapi: Error 400: Request contains an invalid argument., badRequest)
```

