# Witbot

Ask simple questions like:

- What is the weather like in New York
- Weather today Sydney

The API will forward the requests to the [Wit.ai API](https://wit.ai) and take action based on the response coming from Wit.ai
 
## Makefile

* ``make  all``  -> (dep, test, build)
* ``make  dep``  -> dep ensure
* ``make  build``  -> builds to /bin/
* ``make  test``  -> runs tests
* ``make  clean``  -> removes /bin/
* ``make  run`` -> runs project
* ``make  build-linux`` -> cross compilation