docker pull mysql

docker run -p 3306:3306 --name mysql -v $PWD/conf/my.cnf:/etc/mysql/my.cnf -v $PWD/logs:/logs -v $PWD/data:/mysql_data -e MYSQL_ROOT_PASSWORD=123456 -d mysql


 -p 3306:3306：将容器的3306端口映射到主机的3306端口
 -v $PWD/conf/my.cnf:/etc/mysql/my.cnf：将主机当前目录下的conf/my.cnf挂载到容器的/etc/mysql/my.cnf
 -v $PWD/logs:/logs：将主机当前目录下的logs目录挂载到容器的/logs
 -v $PWD/data:/mysql_data：将主机当前目录下的data目录挂载到容器的/mysql_data
 -e MYSQL_ROOT_PASSWORD=123456：初始化root用户的密码
