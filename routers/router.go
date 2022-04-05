package routers

import (
	controllers "coinwallet/controllers/admin"
	"coinwallet/controllers/api_v1"
	"coinwallet/middleware"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dchest/captcha"
	"net/http"
)

func init() {
	//授权登录中间件
	middleware.AuthMiddle()
	beego.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := beego.NewNamespace("/admin",
		//操作日志
		beego.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),
		//登录页
		beego.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
		//退出登录
		beego.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		beego.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		beego.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		beego.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		beego.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

		beego.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//菜单管理
		beego.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		beego.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		beego.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		beego.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		beego.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		beego.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		beego.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		beego.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		beego.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		beego.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-添加界面
		beego.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-添加
		beego.NSRouter("/admin_user/create", &controllers.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		beego.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		beego.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-启用
		beego.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-禁用
		beego.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		beego.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),

		//系统管理-角色管理
		beego.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-添加界面
		beego.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-添加
		beego.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		beego.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		beego.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		beego.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-启用角色
		beego.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-禁用角色
		beego.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		beego.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		beego.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),

		//币种管理-币种管理
		beego.NSRouter("/asset/index", &controllers.AssetController{}, "get:Index"),
		//币种管理-币种管理-添加界面
		beego.NSRouter("/asset/add", &controllers.AssetController{}, "get:Add"),
		//币种管理-币种管理-添加
		beego.NSRouter("/asset/create", &controllers.AssetController{}, "post:Create"),
		//币种管理-币种管理-修改界面
		beego.NSRouter("/asset/edit", &controllers.AssetController{}, "get:Edit"),
		//币种管理-币种管理-修改
		beego.NSRouter("/asset/update", &controllers.AssetController{}, "post:Update"),
		//币种管理-币种管理-删除
		beego.NSRouter("/asset/del", &controllers.AssetController{}, "post:Del"),

		//币种管理-市场币种管理
		beego.NSRouter("/market/asset/index", &controllers.MarketAssetController{}, "get:Index"),
		//币种管理-市场币种管理-添加界面
		beego.NSRouter("/market/asset/add", &controllers.MarketAssetController{}, "get:Add"),
		//币种管理-市场币种管理-添加
		beego.NSRouter("/market/asset/create", &controllers.MarketAssetController{}, "post:Create"),
		//币种管理-市场币种管理-修改界面
		beego.NSRouter("/market/asset/edit", &controllers.MarketAssetController{}, "get:Edit"),
		//币种管理-市场币种管理-修改
		beego.NSRouter("/market/asset/update", &controllers.MarketAssetController{}, "post:Update"),
		//币种管理-市场币种管理-删除
		beego.NSRouter("/market/asset/del", &controllers.MarketAssetController{}, "post:Del"),


		//公告管理-公告管理
		beego.NSRouter("/news/index", &controllers.NewsController{}, "get:Index"),
		//公告管理-公告管理-添加界面
		beego.NSRouter("/news/add", &controllers.NewsController{}, "get:Add"),
		//公告管理-公告管理-添加
		beego.NSRouter("/news/create", &controllers.NewsController{}, "post:Create"),
		//公告管理-公告管理-修改界面
		beego.NSRouter("/news/edit", &controllers.NewsController{}, "get:Edit"),
		//公告管理-公告管理-修改
		beego.NSRouter("/news/update", &controllers.NewsController{}, "post:Update"),
		//公告管理-公告管理-删除
		beego.NSRouter("/news/del", &controllers.NewsController{}, "post:Del"),

		//合约管理-合约管理
		beego.NSRouter("/token/config/index", &controllers.TokenConfigController{}, "get:Index"),
		//合约管理-合约管理-添加界面
		beego.NSRouter("/token/config/add", &controllers.TokenConfigController{}, "get:Add"),
		//合约管理-合约管理-添加
		beego.NSRouter("/token/config/create", &controllers.TokenConfigController{}, "post:Create"),
		//合约管理-合约管理-修改界面
		beego.NSRouter("/token/config/edit", &controllers.TokenConfigController{}, "get:Edit"),
		//合约管理-合约管理-修改
		beego.NSRouter("/token/config/update", &controllers.TokenConfigController{}, "post:Update"),
		//合约管理-合约管理-删除
		beego.NSRouter("/token/config/del", &controllers.TokenConfigController{}, "post:Del"),

		//地址管理-地址管理
		beego.NSRouter("/address/index", &controllers.AddressController{}, "get:Index"),
		//地址管理-地址管理-修改界面
		beego.NSRouter("/address/edit", &controllers.AddressController{}, "get:Edit"),
		//地址管理-地址管理-修改
		beego.NSRouter("/address/update", &controllers.AddressController{}, "post:Update"),
		//地址管理-地址管理-删除
		beego.NSRouter("/address/del", &controllers.AddressController{}, "post:Del"),

		//版本管理-版本管理
		beego.NSRouter("/version/index", &controllers.VersionController{}, "get:Index"),
		//版本管理-添加界面
		beego.NSRouter("/version/add", &controllers.VersionController{}, "get:Add"),
		//版本管理-创建
		beego.NSRouter("/version/create", &controllers.VersionController{}, "post:Create"),
		//版本管理-版本管理-修改界面
		beego.NSRouter("/version/edit", &controllers.VersionController{}, "get:Edit"),
		//版本管理-版本管理-修改
		beego.NSRouter("/version/update", &controllers.VersionController{}, "post:Update"),
		//版本管理-版本管理-删除
		beego.NSRouter("/version/del", &controllers.VersionController{}, "post:Del"),

		//链管理-链管理
		beego.NSRouter("/chain/index", &controllers.ChainController{}, "get:Index"),
		//链管理-添加界面
		beego.NSRouter("/chain/add", &controllers.ChainController{}, "get:Add"),
		//链管理-创建
		beego.NSRouter("/chain/create", &controllers.ChainController{}, "post:Create"),
		//链管理-修改界面
		beego.NSRouter("/chain/edit", &controllers.ChainController{}, "get:Edit"),
		//链管理-修改
		beego.NSRouter("/chain/update", &controllers.ChainController{}, "post:Update"),
		//链管理-删除
		beego.NSRouter("/chain/del", &controllers.ChainController{}, "post:Del"),

		//设备地址-设备管理
		beego.NSRouter("/notebook/index", &controllers.NoteBookAddrController{}, "get:Index"),
		//设备地址-添加界面
		beego.NSRouter("/notebook/add", &controllers.NoteBookAddrController{}, "get:Add"),
		//设备地址-创建
		beego.NSRouter("/notebook/create", &controllers.NoteBookAddrController{}, "post:Create"),
		//设备地址-修改界面
		beego.NSRouter("/notebook/edit", &controllers.NoteBookAddrController{}, "get:Edit"),
		//设备地址-修改
		beego.NSRouter("/notebook/update", &controllers.NoteBookAddrController{}, "post:Update"),
		//设备地址-删除
		beego.NSRouter("/notebook/del", &controllers.NoteBookAddrController{}, "post:Del"),

		//市场管理-市场管理
		beego.NSRouter("/market/index", &controllers.MarketController{}, "get:Index"),

		//燃油管理-燃油管理
		beego.NSRouter("/gas/index", &controllers.GasController{}, "get:Index"),

		//记录管理-记录管理
		beego.NSRouter("/record/index", &controllers.RecordController{}, "get:Index"),
	)
	beego.AddNamespace(admin)

	api_path := beego.NewNamespace("/v1",
		beego.NSNamespace("/news",
			beego.NSInclude(
				&api_v1.NewsController{},
			),
		),
		beego.NSNamespace("/wallet",
			beego.NSInclude(
				&api_v1.WalletController{},
			),
		),

		beego.NSNamespace("/token",
			beego.NSInclude(
				&api_v1.TokenController{},
			),
		),

		beego.NSNamespace("/market",
			beego.NSInclude(
				&api_v1.MarketController{},
			),
		),

		beego.NSNamespace("/version",
			beego.NSInclude(
				&api_v1.VersionController{},
			),
		),

		beego.NSNamespace("/config",
			beego.NSInclude(
				&api_v1.ConfigController{},
			),
		),

		beego.NSNamespace("/notebook",
			beego.NSInclude(
				&api_v1.NodeBookController{},
			),
		),
	)
	beego.AddNamespace(api_path)
}
