# Witbot

Ask simple questions like:

- What is the weather like in New York
- Weather today Sydney

The API will forward the requests to the [Wit.ai API](https://wit.ai) and take action based on the response coming from Wit.ai
 
## Makefile

``make ...``
* ``all``  -> (dep, test, build)
* ``dep``  -> dep ensure
* ``build``  -> builds to /bin/
* ``test``  -> runs tests
* ``clean``  -> removes /bin/
* ``run`` -> runs project
* ``build-linux`` -> cross compilation