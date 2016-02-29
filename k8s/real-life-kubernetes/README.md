# Real life Kubernetes

On this presentation we will give a short introduction on Kubernetes and 
show the experience of learning and using Kubernetes on production for two very distinct systems.

The first one is a data acquisition system, involving running multiple instance of crawlers,
storage services, captcha breaking services, message brokers (like RabbitMQ) and database integration 
outside the cluster.

The second one is: FABIO TODO

We will show how Kubernetes solved problems like:

* Load balancing
* Service discovery
* Port collisions
* Scheduling services on multiple nodes
* Handling outages
* QoS
* Deployment and rolling updates of services
* Deploy and update Kubernetes itself at AWS
* Debugging
* Monitoring
