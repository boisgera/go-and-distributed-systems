Go and Distributed Systems
================================================================================

Go
--------------------------------------------------------------------------------

**TODO:**

  - go tour reference

  - selected examples (to make the following HTTP Section understandable
    first and foremost)

HTTP
--------------------------------------------------------------------------------

Steps:

  1. Make a very simple "Hello world" server

  2. Connect to it via curl, Python (requests) and Go

  3. Return "Hello world" HTML instead of plain text

  4. Return "Hello world" + IP adress + date

  5. Do it in JSON

  6. Routing & Query Parameters

     (Mention FastAPI alternative)

  7. mDNS ?

Q: How can I introduce concurrency fast here ? Implicitly with a slow request
and show that the HTTP Server IS concurrent. But explicitly? Start a 
long-running process that collects a log of the connections and give the list
back? So that would be a 4.5 ; yes.

Moar
--------------------------------------------------------------------------------

Concepts and Go libs

  - mDNS

  - mqtt

  - sockets / websockets

  - webRTC


