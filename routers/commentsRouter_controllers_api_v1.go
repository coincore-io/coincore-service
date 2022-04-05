package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:ConfigController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:ConfigController"],
        beego.ControllerComments{
            Method: "GetConfigList",
            Router: `/get_config_list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:MarketController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:MarketController"],
        beego.ControllerComments{
            Method: "GetMarketPrice",
            Router: `/get_market_price`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NewsController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NewsController"],
        beego.ControllerComments{
            Method: "GetNewsList",
            Router: `/get_news_list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NewsController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NewsController"],
        beego.ControllerComments{
            Method: "GetNewsDetail",
            Router: `/news_detail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"],
        beego.ControllerComments{
            Method: "AddNoteBook",
            Router: `/add_note_book`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"],
        beego.ControllerComments{
            Method: "DelNoteBook",
            Router: `/del_note_book`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"],
        beego.ControllerComments{
            Method: "GetNoteBook",
            Router: `/get_note_book`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:NodeBookController"],
        beego.ControllerComments{
            Method: "UpdNoteBook",
            Router: `/upd_note_book`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"],
        beego.ControllerComments{
            Method: "DeleteWalletToken",
            Router: `/delete_wallet_token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"],
        beego.ControllerComments{
            Method: "HotTokenList",
            Router: `/hot_token_list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:TokenController"],
        beego.ControllerComments{
            Method: "SourchAddToken",
            Router: `/sourch_add_token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:VersionController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:VersionController"],
        beego.ControllerComments{
            Method: "GetVersionInfo",
            Router: `/version_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "BatchSubmitWallet",
            Router: `/batch_submit_wallet`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "DeleteWallet",
            Router: `/delete_wallet`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "GetAddressBalance",
            Router: `/get_address_balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "GetSignTxInfo",
            Router: `/get_sign_tx_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "GetTxByAddress",
            Router: `/get_tx_by_address`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "GetWalletAsset",
            Router: `/get_wallet_asset`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "GetWalletBalance",
            Router: `/get_wallet_balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "SendTx",
            Router: `/send_tx`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"] = append(beego.GlobalControllerRouter["coinwallet/controllers/api_v1:WalletController"],
        beego.ControllerComments{
            Method: "SubmitWalletInfo",
            Router: `/submit_wallet_info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
