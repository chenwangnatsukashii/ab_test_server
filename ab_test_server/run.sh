package ab_test_server

# 开启nacos
sh /Users/cl60006/Downloads/nacos/bin/startup.sh -m standalone
# 关闭nacos
sh /Users/cl60006/Downloads/nacos/bin/shutdown.sh
# nacos后台管理界面
http://127.0.0.1:8848/nacos/#/login


# 安装brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
# 卸载brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/uninstall.sh)"


# 1.安装brew失败报错：error: Not a valid ref: refs/remotes/origin/master，需要卸载然后重新安装

# 2.在安装kafka的时候遇到No available formula with the name "*"
# 解决：
rm -rf /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core
brew update


# 安装kafka
brew install kafka
# 修改配置文件
vim /usr/local/etc/kafka/server.properties
# listeners=PLAINTEXT://localhost:9092
# 启动zk和kafka
brew services start zookeeper
brew services start kafka

# 创建topic
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test1
# 查看topic
kafka-topics --list --bootstrap-server localhost:9092
# 消费指定topic中的消息
sh kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test1 --from-beginning


# 添加环境变量
vim ./.bash_profile

# 查看当前目录
pwd

export PATH="${PATH}:/usr/local/Cellar/kafka/3.3.1_1/bin"

# 查看端口被哪个进程监听
sudo lsof -i:8082




