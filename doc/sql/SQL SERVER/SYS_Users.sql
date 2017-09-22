CREATE TABLE SYS_Users (
	ID INT --ID
	,
	Code VARCHAR (32) --用户, 这里可以自定义用户的代码, 也可以取自员工表的员工工号
	,
	EmployeeID INT --对照员工表的ID, 如果为0表示不连到员工表
	,
	Name VARCHAR (20) --如果是员工表的用户，姓名取自员工表
	,
	FullName VARCHAR (64) --全名
	,
	Password VARCHAR (128) --密码, 密码以用户名作密钥加密，这样相同明文，密文也不同, 再作Hash操作，变成不可逆
	,
	Description VARCHAR (255) --说明
	,
	Privilege TINYINT --特权用户, 此类用户是系统内置帐号，不允许外部删除及修改用户名, 此字段不用于编辑
	,
	Authorized TINYINT --授权： 0=正常, 1=禁用
	,
	CreateDate DATETIME --创建日期
	,
	ExpiryDate DATETIME --有效日期, 帐户有效期
	,
	LoginHost VARCHAR (256) --当前操作员登录的机器名
	,
	LoginState TINYINT --登录状态  0：无 1：已登录  2：登出  4：签到   8：签出
	,
	WorkPass VARCHAR (128) --业务密码
)