package auth

import (
	"fmt"
	"net"
	"net/http"

	"github.com/holywolfchan/yuncang/utils/logs"

	"github.com/holywolfchan/yuncang/model"
	"github.com/holywolfchan/yuncang/service"

	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var enforcer *casbin.Enforcer

func init() {

	a := xormadapter.NewAdapter("mysql", "root:***@tcp(localhost:3306)/yuncang?charset=utf8&parseTime=True&loc=Local", true)
	enforcer = casbin.NewEnforcer("./conf/casbin_model.conf", a)
	enforcer.AddFunction("IPFilter", IPFilterFunc)
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
func GetStr(v interface{}) string {
	if v != nil {

		return fmt.Sprint(v)
	}
	return ""
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ok, err := CheckPermission(c); ok {
			c.Next()

		} else {
			if err == nil {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
					"errmsg": "权限不足，拒绝访问",
				})
				return
			}
			logs.Infof("请求授权失败:%v", err)
			c.AbortWithStatusJSON(http.StatusUnavailableForLegalReasons, gin.H{
				"errmsg": err,
			})
			return
		}

	}
}

//逻辑是以r.dom::r.sub的形式构成role1去跟p中的r.dom::p.sub的role2比较
//若相等，就返回true；
//否则进入下一步——进入g中的role-role模型查找，g中的role-role模型为[g.dom::g.r.sub<g.dom::g.p.sub]
//若role1和role2在g中都存在，则进入下一步查找角色继承关系，否则返回false；
//若role1 不属于 role2,则返回false；
//在g = 1, 2, 3 中代表 域3 中 角色1 归属 角色2，继承角色2的权限
func CheckPermission(c *gin.Context) (bool, error) {
	var userinfo model.UserFullInfo
	var err error
	var sub string
	var dom string
	var path string
	var act string
	var ip string
	if userinfo, err = service.GetContextUser(c); err == nil {
		sub = GetStr(userinfo.User.RoleId)
		dom = GetStr(userinfo.User.DomainId)
		path = c.Request.URL.Path
		act = c.Request.Method
		ip = c.ClientIP()
		return enforcer.Enforce(sub, dom, path, act, ip), nil
	}
	return false, err
}

func IPFilter(ip1 string, ip2 string) bool {
	if ip2 == "0.0.0.0/0" {
		return true
	}
	objIP1 := net.ParseIP(ip1)
	if objIP1 == nil {
		panic("invalid argument: ip1 in IPMatch() function is not an IP address.")
	}

	_, cidr, err := net.ParseCIDR(ip2)
	if err != nil {
		objIP2 := net.ParseIP(ip2)
		if objIP2 == nil {
			panic("invalid argument: ip2 in IPMatch() function is neither an IP address nor a CIDR.")
		}

		return objIP1.Equal(objIP2)
	}

	return cidr.Contains(objIP1)
}

func IPFilterFunc(args ...interface{}) (interface{}, error) {
	ip1 := args[0].(string)
	ip2 := args[1].(string)

	return (bool)(IPFilter(ip1, ip2)), nil
}
