# GoProxyServer
A simple reverse proxy server in Go with features like request and response headers addition, changing response body contents before serving and caching the response data for specific time and serving from cache. 

Directory **OriginalServer** contains a simple http server which serves two routes with url as mentioned below:-
1. http://localhost:5000/about.html

![image](https://user-images.githubusercontent.com/85868251/121872601-16807980-cd23-11eb-9325-a7082a41d401.png)

2. http://localhost:5000/admin/contact.html

![image](https://user-images.githubusercontent.com/85868251/121872559-08325d80-cd23-11eb-97e5-efa77637175e.png)

Directory **GoProxyServer** contains the proxy server with proxies the **OriginalServer**.
While doing so we can perform following things:-
1. Add request headers before calling the **OriginalServer** (You can see in console log of **OriginalServer**)
![image](https://user-images.githubusercontent.com/85868251/121870805-08c9f480-cd21-11eb-9747-b70d362f2ca6.png)
2. Add response headers before serving the response from **GoProxyServer** (You can see in console log of **GoProxyServer**)
![image](https://user-images.githubusercontent.com/85868251/121870926-28611d00-cd21-11eb-8ac2-1b97e03ebc67.png)
3. Change contents of HTML when served from **GoProxyServer**. For example:- "mug" -> "bottle" and "pen" -> "pencil"
4. Cache on successful response to **GoProxyServer** from **OriginalServer** for a specific time and if present in cache return from cache instead of hitting again to **OriginalServer** (You can see in console log of **GoProxyServer** if the request is served from cache)
![image](https://user-images.githubusercontent.com/85868251/121871033-475faf00-cd21-11eb-809e-41353e456222.png)


Output:-
1. http://localhost:3000/about.html

![image](https://user-images.githubusercontent.com/85868251/121872696-34e67500-cd23-11eb-9b44-8fe18a248ed5.png)

2. http://localhost:3000/admin/contact.html

![image](https://user-images.githubusercontent.com/85868251/121872656-27c98600-cd23-11eb-8a46-57b1b8a4cd83.png)
