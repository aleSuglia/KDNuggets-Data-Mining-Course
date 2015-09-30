KDNuggets Data mining course downloader
=======================================

A new academic year starts and I always want to learn from the most famous authors or books.
In the data mining field the KDNuggets website (http://www.kdnuggets.com/) is an authority. It has a specific
section (data_mining_course) which let you to access to some useful slides about data mining.

Unfortunately the author has inserted the link separating the *url* from the *url path* of each single file.

I don't want to copy-paste things, so I've created this simple and useful Golang program that automatically 
generates in your current directory all the *ppt* files. This kind of script could be easily used to 
download all the other files which are available in the other site's sections. 


**Installation**

The script it is written in Golang (http://www.golang.org) and on the official website you can download
the installation binary for every operating system.

After that, you had to build the program:
```
go build kdnuggets_down.go
```

and run it:
```
go run kdnuggets_down.go
```
