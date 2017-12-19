rm ./server >/dev/null 2>&1
CGO_ENABLED=0 go build -o server -a -installsuffix cgo -ldflags "-s" -ldflags "-X main.btime=`date -u +.%Y%m%d.%H%M%S`" . 
if [ ! $? -eq 0 ]; then
 echo -e "go build failed"
 exit 1
fi

IMGVERSION=v1

docker build -t dbenque/helloserver:$IMGVERSION .
if [ ! $? -eq 0 ]; then
 rm ./kchain >/dev/null 2>&1
 echo -e "docker go build failed"
 exit 1
fi
rm ./kchain >/dev/null 2>&1

docker tag dbenque/kchain:$IMGVERSION dbenque/kchain:latest

read -p "Do you want to push to docker hub (y/n)?  " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
    docker push dbenque/kchain:$IMGVERSION
    docker push dbenque/kchain:latest
fi
