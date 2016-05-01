# Real life Kubernetes

On this presentation we will give a short introduction on Kubernetes and
show the experience of learning and using Kubernetes on production for two very distinct systems.

The first one is a data acquisition system, involving running multiple instance of crawlers,
storage services, captcha breaking services, message brokers (like RabbitMQ) and database integration
outside the cluster.

The second one is a web application, involving network analysis using graphs with the ultimate goal
of fraud prevention. The application is strongly bounded with the microservices architecture and the
twelve factor app.

Graph and document databases, cache layers, a message broker and a distributed filesystem are some
of the technologies surrounding the application ecosystem.

Also making a great case of integration between AWS and k8s using tools such as autoscaling.

We will show how Kubernetes solved problems like:

* Load balancing
* Service discovery
* Port collisions
* Scheduling services on multiple nodes
* Handling outages
* QoS
* Deployment and rolling updates of services
* Deploy and update Kubernetes itself at AWS
* Autoscaling k8s nodes AWS
* Self-registering k8s nodes within AWS ELB's
* Debugging
* Monitoring


## Running the presentation

Install docker and run: `make`


## Fabio Favero Henkes profile:

Software craftsman, full stack developer. Focus on solving problems with simplicity.
Embrace new technologies and complex tasks. Twelve+ years of professional experience.
Bachelor at Computer Science. Postgraduate at Software Engineering.
Self-taught, always reading and improving. Prefer winter days.

* Powers: C, Java, Golang, Node.js, Javascript and Html5.

* Spells: Spring Framework and Spring Boot, Yeoman, Bower, Gulp, AngularJS.

* Weapons: MongoDB, Apache Solr, Neo4J, Postgresql, Memcached, Rabbitmq, NGINX, MacOS, Linux, FreeBSD, Docker, Maven, GoCD, Kubernetes, AWS, .. and also a decent cook and gamer.

Acquiring new powers in Scala, Apache Spark GraphX and OCaml.
