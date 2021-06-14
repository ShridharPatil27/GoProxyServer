# GoProxyServer
A simple reverse proxy server in Go with features like request and response headers addition, changing response body contents before serving and caching the response data for specific time and serving from cache. 

Directory **OriginalServer** contains a simple http server which serves two routes with url as mentioned below:-
1. http://localhost:5000/about.html
![image](https://user-images.githubusercontent.com/85868251/121869320-8db40e80-cd1f-11eb-9193-b60c4c6fdc76.png)

2. http://localhost:5000/admin/contact.html
![image](https://user-images.githubusercontent.com/85868251/121869506-be944380-cd1f-11eb-92b7-82d0022af6f3.png)

Directory **GoProxyServer** contains the proxy server with proxies the **OriginalServer**.
While do so we can do following things:-
1. Add request headers before calling the **OriginalServer** (You can see in console log of **OriginalServer**)
![image](https://user-images.githubusercontent.com/85868251/121870805-08c9f480-cd21-11eb-9747-b70d362f2ca6.png)
2. Add response header before serving the response from **GoProxyServer** (You can see in console log of **GoProxyServer**)
![image](https://user-images.githubusercontent.com/85868251/121870926-28611d00-cd21-11eb-8ac2-1b97e03ebc67.png)
3. Change contents of HTML when servred from **GoProxyServer**. For example:- "mug" -> "bottle" and "pen" -> "pencil"
4. Cache on successful response in **GoProxyServer** from **OriginalServer** for a specific time and if present in cache return from cache instead of hitting again to **OriginalServer** (You can see in console log of **GoProxyServer** if the request is served from cache)
![image](https://user-images.githubusercontent.com/85868251/121871033-475faf00-cd21-11eb-809e-41353e456222.png)


Output:-
1. http://localhost:3000/about.html
![image](https://user-images.githubusercontent.com/85868251/121870623-da4c1980-cd20-11eb-93ba-e141395c20fc.png)

2. http://localhost:3000/admin/contact.html
![image](https://user-images.githubusercontent.com/85868251/121870682-e8019f00-cd20-11eb-8ae0-aa2bfcb61d90.png)
