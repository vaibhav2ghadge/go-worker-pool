implemented app by using SOLID principal and domain driven design(https://medium.com/learnfazz/domain-driven-design-in-go-253155543bb1),12factory so we can easily add and remove any part of code easily 
Bulded and tested on linux 
for fast execuation add worker pool(https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923) by using go routine we can process file input fastly but problem is here sqlite3 locking mechanism it hadnle one thread at time.

library used
go get github.com/mattn/go-sqlite3 used for sqlite3

build and installation
please put trivago/ folder inside go/src/
make build #for build binary
for run 
./main -sqliteDBPath=/home/user/go/src/trivago/sqlite.db -inputFilePath=/home/user/go/src/trivago/delivery/hotels.csv -jsonFilePath=/home/user/go/src/trivago/utils/op.json -workers=10
you can change env variable by your choice used 12factory

# aslo added build for linux in directory builds if faild to create directory