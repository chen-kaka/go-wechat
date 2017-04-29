source ~/.bash_profile

cd $GOPATH/src/go-wechat
go build

tokill=`ps -ef | grep go-wechat | grep -v grep | awk '{print $2}'`
kill -9 $tokill

./go-wechat &

echo "deploy completed."