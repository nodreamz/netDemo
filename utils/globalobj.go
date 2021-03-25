package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"zinx/ziface"
)

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
一些参数是可以通过zinx.json由用户进行配置
 */

type GlobalObj struct {
	/*
	Server
	 */
	TcpServer ziface.IServer 	//当前Zinx全局的Server对象
	Host string		//当前服务器主机监听的IP
	TcpPort int		//当前服务器主机监听的端口号
	Name string 	//当前服务器的名称

	/*
	Zinx
	 */

	Version string 	//当前Zinx的版本号
	MaxConn int		//当前服务器主机允许的最大链接数
	MaxPackageSize uint32	//当前服务器框架数据包的最大值

}

func GetPath(ePath string) string {
	return ePath+"conf/zinx.json"
}
/*
从zinx.json去加载用于自定义的参数
 */
func (g *GlobalObj) Reload() {
	ePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(ePath)
	data, err := ioutil.ReadFile(ePath+"/MyDemo/ZinxV0.4/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}


/*
定义一个全局的对外Globalobj
 */

var GlobalObject *GlobalObj


//提供一个init方法，初始化当前的GlobalObject
func init() {
	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name: "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 8999,
		Host: "0.0.0.0",
		MaxConn: 1000,
		MaxPackageSize: 4096,
	}

	//应该常识从conf/zinx.json
	GlobalObject.Reload()
}
