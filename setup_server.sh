#启动redis服务
redis-server  ./conf/redis.conf

#启动trackerd
fdfs_trackerd ./conf/tracker.conf restart

#启动storaged
fdfs_storaged $GOPATH/src/images/conf/storage.conf restart
