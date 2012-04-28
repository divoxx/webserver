Project Information
===================

For now, this library is just an abstraction for running a webserver for an [Go][go] web application. It provides the following:

* A structure for defining environment specific configurations
* Logging of server behavior
* Simple interface for running an web application
* Handy method for running the the web app accepting parameters from CLI

The long tearm goal is to provide a standard interface for frameworks and webapp, as well as the following:

* Serves static files, which is handy for small apps that won't have Apache or Nginx in front of it or for development
* Logging of requests information

[go]: http://golang.org

-------------

Copyright (C) 2012 Rodrigo Kochenburger <divoxx@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
