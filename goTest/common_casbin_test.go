package goTest

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"testing"
)

var modelString = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act
`

var scvString = `
p, alice, data1, read
p, bob, data2, write
p, data_group_admin, data_group, write

g, alice, data_group_admin
g2, data1, data_group
g2, data2, data_group
`

func TestCasBin(t *testing.T) {
	// 从Go代码初始化模型
	// rbac_with_resource_roles_model.conf
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act, type")
	m.AddDef("p", "p", "sub, obj, act, type")
	//_, _表示角色继承关系的前项和后项，即前项继承后项角色的权限。 一般来讲，如果您需要进行角色和用户的绑定，直接使用g 即可
	m.AddDef("g", "g", "_, _") // 是RBAC角色继承关系的定义
	//目前不知道g2如何使用 测试后发现直接的g 就可以满足需求
	//m.AddDef("g", "g2", "_, _") // 是RBAC角色继承关系的定义
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	//m.AddDef("m", "m", "r.type == p.type && g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act || r.sub == 'superAccess' ")
	m.AddDef("m", "m", "r.type == p.type && g(r.sub, p.sub) && r.act == p.act || r.sub == 'superAccess' ")

	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a, _ := xormadapter.NewAdapter("mysql", "root:MYSQL@tcp(192.168.0.100:3306)/go_cloud?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	e, _ := casbin.NewEnforcer(m, a)

	// Load the policy from DB.
	//e.LoadPolicy()

	// Modify the policy.
	//e.AddPolicy("data_group", "data1", "read")                  //这个用户有这个资源的权限
	//e.AddPolicy("bob", "data2", "write")                   //这个用户有这个资源的权限
	//e.AddPolicy("data_group_admin", "data_group", "write") //这个用户有这个资源的权限
	//
	//e.AddGroupingPolicy("alice", "data_group_admin") //这个用户有这个角色
	//e.AddGroupingPolicy("data1", "data_group")       //这个资源属于某个角色
	//e.AddGroupingPolicy("data2", "data_group")       //这个资源属于某个角色
	//e.RemovePolicy(...)

	// Check the permission.
	//EnforceEx 返还生效的数组 Enforce只验证是否生效
	fmt.Println(e.EnforceEx("alice", "data_group", "write", "123")) //true
	fmt.Println(e.Enforce("data1", "data1", "read", ""))            //true
	fmt.Println(e.Enforce("bob", "data2", "write", ""))             //true
	fmt.Println(e.Enforce("data1", "data_group", "write"))          //false
	//从数据库重新加载   上面一行打赏断点    修改alice2 权限  从数据库重新加载然后可以使得下面权限得到false
	_ = e.LoadPolicy()
	fmt.Println(e.EnforceEx("alice2", "data_group", "write", "123")) //true

	fmt.Println(e.GetAllActions())  //有效的资源的请求模式
	fmt.Println(e.GetAllRoles())    //所有的角色
	fmt.Println(e.GetAllDomains())  //？？？
	fmt.Println(e.GetAllObjects())  //有效的资源路径
	fmt.Println(e.GetAllSubjects()) //有效的资源对应的角色
	fmt.Println(e.GetAllNamedActions("p"))
	fmt.Println(e.GetPermissionsForUser("alice2"))         //查询直属权限
	fmt.Println(e.GetImplicitPermissionsForUser("alice2")) //查询所有权限包括间接
	fmt.Println(e.GetRolesForUser("alice2"))               //查询直属角色
	fmt.Println(e.GetImplicitRolesForUser("alice2"))       //查询所有角色包括间接

	// Save the policy back to DB.
	//e.SavePolicy()

	//fmt.Println(m.ToText())
}
