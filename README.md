## OpenShift Simple Demo App

A minimal Golang web server to demonstrate application build quickly

To try it simply run the following command:
```
$ oc new-project hello-openshift-demo
$ oc new-app https://github.com/giannisalinetti/hello-openshift --strategy=docker
```

Try to inspect the build logs and demonstrate that a pure Docker build has been triggered.
```
$ oc logs builds/hello-openshift-1 -f
```

### Maintainer
Gianni Salinetti gsalinet@redhat.com


