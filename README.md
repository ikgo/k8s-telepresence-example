## Kubernetes Telepresence Example

[Telepresence][telepresence] is a tool for swapping out a local process for a
service running in a kubernetes cluster.  This is an example demonstrating some
common development tasks.

The user_service returns a list of users:

```json
[
    {
        "name": "Matt"
    },
    {
        "name": "Steve"
    },
    {
        "name": "Greg"
    }
]
```

The widget_service accesses the user_service through the kubernetes DNS entry
`http://user-service.default/` returns a list of users and wigets:

```json
{
    "users": [
        {
            "name": "Matt"
        },
        {
            "name": "Steve"
        },
        {
            "name": "Greg"
        }
    ],
    "widgets": [
        {
            "name": "Thingamajig"
        },
        {
            "name": "Thingamabob"
        },
        {
            "name": "Plumbus"
        }
    ]
}
```

## Usage Start a minikube cluster:

```bash
minikube start
```

build the docker files in the minikube environment:

```bash
eval $(minikube docker-env)
make build
```

Start the services:

```bash
make run
```

print the urls using `make print-services` and navigate to the widget service.
Observe there are only three users.

Now add a user by modifying `user_service/main.go`.  Build swap out the user
service using Telepresence:

```bash
make swap-user
```

Refresh the browser.  The widget service queries the user service running on
your local machine.

## How this works

Telepresence mounts your local filesystem into a kubernetes pod behind your
existing Kubernetes service.  All incoming and outgoing requests to the service
are proxied to your local machine, creating a bridge between the host and the
minikube cluster.

Additionally, Telepresence takes care of the DNS entries and environment
ariables a pod running on the cluster would normally see.  This means that your
services do not need special configuration.

[telepresence]: http://www.telepresence.io/