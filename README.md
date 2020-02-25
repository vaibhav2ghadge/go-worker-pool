# Implemented app by using SOLID principal and domain driven design
https://medium.com/learnfazz/domain-driven-design-in-go-253155543bb1 
- We can easily add and remove any part of code easily
- Given csv file with hotel details we have to push into db and json file and it may happen that storing the output change like sqlite3 to mysql , json to xml so our code needs to be moduler and functional.
- We observe that processing and pushing data into db slower the app. to slove this problem we build **worker pool model**. we have fix number of worker we pass data to worker and worker push it to the db.
- [Worker Pool Implementation](https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923 "Worker Pool Implementation")


## Library Used
- go get github.com/mattn/go-sqlite3 used for sqlite3

##Build and Installation
1. make build #for build binary
2.  ./main -sqliteDBPath=/home/user/go/src/trivago/sqlite.db -inputFilePath=/home/user/go/src/trivago/delivery/hotels.csv -jsonFilePath=/home/user/go/src/trivago/utils/op.json -workers=10


#### Also added build for linux in directory builds if faild to build 
