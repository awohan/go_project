1.install govendor<br/>
> brew install govendor

2.set GOPATH<br/>
> export GOPATH=~/gopath

3.goto $GOPATH<br/>
> cd $GOPATH

4.download<br/>
> git clone github.com/BastionPay/bas-auth-srv.git src/BastionPay/bas-auth-srv

5.goto source
> cd src/BastionPay/bas-auth-srv

6.switch your branch<br/>
> git checkout master/dev

7.sync dependency
> govendor sync

8.make build<br/>
> make

use build.sh to build<br/>
> export GOPATH=~/gopath

build master<br/>
> ./build.sh bas-auth-srv master

build dev<br/>
> ./build.sh bas-auth-srv dev

run<br/>
> ./bas-auth-srv -log_path=conf/log.xml -conf_path=conf