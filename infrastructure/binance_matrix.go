package infrastructure

type binanceConfig struct {
	tableName string
	dbIndex   int
}

var binanceMatrix = map[string]binanceConfig{
	"aavebnb": {
		tableName: "aave_bnb",
		dbIndex:   4,
	},
	"aavebtc": {
		tableName: "aave_btc",
		dbIndex:   2,
	},
	"aaveeth": {
		tableName: "aave_eth",
		dbIndex:   3,
	},
	"aaveusdt": {
		tableName: "aave_usdt",
		dbIndex:   1,
	},
	"adabnb": {
		tableName: "ada_bnb",
		dbIndex:   4,
	},
	"adabtc": {
		tableName: "ada_btc",
		dbIndex:   2,
	},
	"adaeth": {
		tableName: "ada_eth",
		dbIndex:   3,
	},
	"adausdt": {
		tableName: "ada_usdt",
		dbIndex:   1,
	},
	"algobnb": {
		tableName: "algo_bnb",
		dbIndex:   4,
	},
	"algobtc": {
		tableName: "algo_btc",
		dbIndex:   2,
	},
	"algousdt": {
		tableName: "algo_usdt",
		dbIndex:   1,
	},
	"atombnb": {
		tableName: "atom_bnb",
		dbIndex:   4,
	},
	"atombtc": {
		tableName: "atom_btc",
		dbIndex:   2,
	},
	"atomusdt": {
		tableName: "atom_usdt",
		dbIndex:   1,
	},
	"balbnb": {
		tableName: "bal_bnb",
		dbIndex:   4,
	},
	"balbtc": {
		tableName: "bal_btc",
		dbIndex:   2,
	},
	"balusdt": {
		tableName: "bal_usdt",
		dbIndex:   1,
	},
	"batbnb": {
		tableName: "bat_bnb",
		dbIndex:   4,
	},
	"batbtc": {
		tableName: "bat_btc",
		dbIndex:   2,
	},
	"bateth": {
		tableName: "bat_eth",
		dbIndex:   3,
	},
	"batusdt": {
		tableName: "bat_usdt",
		dbIndex:   1,
	},
	"bcdbtc": {
		tableName: "bcd_btc",
		dbIndex:   2,
	},
	"bcdeth": {
		tableName: "bcd_eth",
		dbIndex:   3,
	},
	"bchbnb": {
		tableName: "bch_bnb",
		dbIndex:   4,
	},
	"bchbtc": {
		tableName: "bch_btc",
		dbIndex:   2,
	},
	"bchusdt": {
		tableName: "bch_usdt",
		dbIndex:   1,
	},
	"bchabcbtc": {
		tableName: "bchabc_btc",
		dbIndex:   2,
	},
	"bchabcusdt": {
		tableName: "bchabc_usdt",
		dbIndex:   1,
	},
	"bchsvbtc": {
		tableName: "bchsv_btc",
		dbIndex:   2,
	},
	"bchsvusdt": {
		tableName: "bchsv_usdt",
		dbIndex:   1,
	},
	"bnbusdt": {
		tableName: "bnb_usdt",
		dbIndex:   1,
	},
	"btcusdt": {
		tableName: "btc_usdt",
		dbIndex:   1,
	},
	"btgbtc": {
		tableName: "btg_btc",
		dbIndex:   2,
	},
	"btgeth": {
		tableName: "btg_eth",
		dbIndex:   3,
	},
	"btsbnb": {
		tableName: "bts_bnb",
		dbIndex:   4,
	},
	"btsbtc": {
		tableName: "bts_btc",
		dbIndex:   2,
	},
	"btseth": {
		tableName: "bts_eth",
		dbIndex:   3,
	},
	"btsusdt": {
		tableName: "bts_usdt",
		dbIndex:   1,
	},
	"bttbnb": {
		tableName: "btt_bnb",
		dbIndex:   4,
	},
	"bttbtc": {
		tableName: "btt_btc",
		dbIndex:   2,
	},
	"bttusdt": {
		tableName: "btt_usdt",
		dbIndex:   1,
	},
	"celobtc": {
		tableName: "celo_btc",
		dbIndex:   2,
	},
	"celousdt": {
		tableName: "celo_usdt",
		dbIndex:   1,
	},
	"compbnb": {
		tableName: "comp_bnb",
		dbIndex:   4,
	},
	"compbtc": {
		tableName: "comp_btc",
		dbIndex:   2,
	},
	"compusdt": {
		tableName: "comp_usdt",
		dbIndex:   1,
	},
	"cotibnb": {
		tableName: "coti_bnb",
		dbIndex:   4,
	},
	"cotibtc": {
		tableName: "coti_btc",
		dbIndex:   2,
	},
	"cotiusdt": {
		tableName: "coti_usdt",
		dbIndex:   1,
	},
	"daibnb": {
		tableName: "dai_bnb",
		dbIndex:   4,
	},
	"daibtc": {
		tableName: "dai_btc",
		dbIndex:   2,
	},
	"daiusdt": {
		tableName: "dai_usdt",
		dbIndex:   1,
	},
	"dashbnb": {
		tableName: "dash_bnb",
		dbIndex:   4,
	},
	"dashbtc": {
		tableName: "dash_btc",
		dbIndex:   2,
	},
	"dasheth": {
		tableName: "dash_eth",
		dbIndex:   3,
	},
	"dashusdt": {
		tableName: "dash_usdt",
		dbIndex:   1,
	},
	"dgbbnb": {
		tableName: "dgb_bnb",
		dbIndex:   4,
	},
	"dgbbtc": {
		tableName: "dgb_btc",
		dbIndex:   2,
	},
	"dgbusdt": {
		tableName: "dgb_usdt",
		dbIndex:   1,
	},
	"dogebnb": {
		tableName: "doge_bnb",
		dbIndex:   4,
	},
	"dogebtc": {
		tableName: "doge_btc",
		dbIndex:   2,
	},
	"dogeusdt": {
		tableName: "doge_usdt",
		dbIndex:   1,
	},
	"dotbnb": {
		tableName: "dot_bnb",
		dbIndex:   4,
	},
	"dotbtc": {
		tableName: "dot_btc",
		dbIndex:   2,
	},
	"dotusdt": {
		tableName: "dot_usdt",
		dbIndex:   1,
	},
	"egldbnb": {
		tableName: "egld_bnb",
		dbIndex:   4,
	},
	"egldbtc": {
		tableName: "egld_btc",
		dbIndex:   2,
	},
	"egldusdt": {
		tableName: "egld_usdt",
		dbIndex:   1,
	},
	"eosbnb": {
		tableName: "eos_bnb",
		dbIndex:   4,
	},
	"eosbtc": {
		tableName: "eos_btc",
		dbIndex:   2,
	},
	"eoseth": {
		tableName: "eos_eth",
		dbIndex:   3,
	},
	"eosusdt": {
		tableName: "eos_usdt",
		dbIndex:   1,
	},
	"etcbnb": {
		tableName: "etc_bnb",
		dbIndex:   4,
	},
	"etcbtc": {
		tableName: "etc_btc",
		dbIndex:   2,
	},
	"etceth": {
		tableName: "etc_eth",
		dbIndex:   3,
	},
	"etcusdt": {
		tableName: "etc_usdt",
		dbIndex:   1,
	},
	"ethusdt": {
		tableName: "eth_usdt",
		dbIndex:   1,
	},
	"filbnb": {
		tableName: "fil_bnb",
		dbIndex:   4,
	},
	"filbtc": {
		tableName: "fil_btc",
		dbIndex:   2,
	},
	"filusdt": {
		tableName: "fil_usdt",
		dbIndex:   1,
	},
	"grtbtc": {
		tableName: "grt_btc",
		dbIndex:   2,
	},
	"grteth": {
		tableName: "grt_eth",
		dbIndex:   3,
	},
	"grtusdt": {
		tableName: "grt_usdt",
		dbIndex:   1,
	},
	"gxsbtc": {
		tableName: "gxs_btc",
		dbIndex:   2,
	},
	"gxseth": {
		tableName: "gxs_eth",
		dbIndex:   3,
	},
	"gxsusdt": {
		tableName: "gxs_usdt",
		dbIndex:   1,
	},
	"hivebnb": {
		tableName: "hive_bnb",
		dbIndex:   4,
	},
	"hivebtc": {
		tableName: "hive_btc",
		dbIndex:   2,
	},
	"hiveusdt": {
		tableName: "hive_usdt",
		dbIndex:   1,
	},
	"iostbnb": {
		tableName: "iost_bnb",
		dbIndex:   4,
	},
	"iostbtc": {
		tableName: "iost_btc",
		dbIndex:   2,
	},
	"iosteth": {
		tableName: "iost_eth",
		dbIndex:   3,
	},
	"iostusdt": {
		tableName: "iost_usdt",
		dbIndex:   1,
	},
	"jstbnb": {
		tableName: "jst_bnb",
		dbIndex:   4,
	},
	"jstbtc": {
		tableName: "jst_btc",
		dbIndex:   2,
	},
	"jstusdt": {
		tableName: "jst_usdt",
		dbIndex:   1,
	},
	"linkbtc": {
		tableName: "link_btc",
		dbIndex:   2,
	},
	"linketh": {
		tableName: "link_eth",
		dbIndex:   3,
	},
	"linkusdt": {
		tableName: "link_usdt",
		dbIndex:   1,
	},
	"ltcbnb": {
		tableName: "ltc_bnb",
		dbIndex:   4,
	},
	"ltcbtc": {
		tableName: "ltc_btc",
		dbIndex:   2,
	},
	"ltceth": {
		tableName: "ltc_eth",
		dbIndex:   3,
	},
	"ltcusdt": {
		tableName: "ltc_usdt",
		dbIndex:   1,
	},
	"manabtc": {
		tableName: "mana_btc",
		dbIndex:   2,
	},
	"manaeth": {
		tableName: "mana_eth",
		dbIndex:   3,
	},
	"manausdt": {
		tableName: "mana_usdt",
		dbIndex:   1,
	},
	"maticbnb": {
		tableName: "matic_bnb",
		dbIndex:   4,
	},
	"maticbtc": {
		tableName: "matic_btc",
		dbIndex:   2,
	},
	"maticusdt": {
		tableName: "matic_usdt",
		dbIndex:   1,
	},
	"mblbnb": {
		tableName: "mbl_bnb",
		dbIndex:   4,
	},
	"mblbtc": {
		tableName: "mbl_btc",
		dbIndex:   2,
	},
	"mblusdt": {
		tableName: "mbl_usdt",
		dbIndex:   1,
	},
	"mkrbnb": {
		tableName: "mkr_bnb",
		dbIndex:   4,
	},
	"mkrbtc": {
		tableName: "mkr_btc",
		dbIndex:   2,
	},
	"mkrusdt": {
		tableName: "mkr_usdt",
		dbIndex:   1,
	},
	"neobnb": {
		tableName: "neo_bnb",
		dbIndex:   4,
	},
	"neobtc": {
		tableName: "neo_btc",
		dbIndex:   2,
	},
	"neoeth": {
		tableName: "neo_eth",
		dbIndex:   3,
	},
	"neousdt": {
		tableName: "neo_usdt",
		dbIndex:   1,
	},
	"oceanbnb": {
		tableName: "ocean_bnb",
		dbIndex:   4,
	},
	"oceanbtc": {
		tableName: "ocean_btc",
		dbIndex:   2,
	},
	"oceanusdt": {
		tableName: "ocean_usdt",
		dbIndex:   1,
	},
	"omgbnb": {
		tableName: "omg_bnb",
		dbIndex:   4,
	},
	"omgbtc": {
		tableName: "omg_btc",
		dbIndex:   2,
	},
	"omgeth": {
		tableName: "omg_eth",
		dbIndex:   3,
	},
	"omgusdt": {
		tableName: "omg_usdt",
		dbIndex:   1,
	},
	"ontbnb": {
		tableName: "ont_bnb",
		dbIndex:   4,
	},
	"ontbtc": {
		tableName: "ont_btc",
		dbIndex:   2,
	},
	"onteth": {
		tableName: "ont_eth",
		dbIndex:   3,
	},
	"ontusdt": {
		tableName: "ont_usdt",
		dbIndex:   1,
	},
	"oxtbtc": {
		tableName: "oxt_btc",
		dbIndex:   2,
	},
	"oxtusdt": {
		tableName: "oxt_usdt",
		dbIndex:   1,
	},
	"paxgbnb": {
		tableName: "paxg_bnb",
		dbIndex:   4,
	},
	"paxgbtc": {
		tableName: "paxg_btc",
		dbIndex:   2,
	},
	"paxgusdt": {
		tableName: "paxg_usdt",
		dbIndex:   1,
	},
	"qtumbnb": {
		tableName: "qtum_bnb",
		dbIndex:   4,
	},
	"qtumbtc": {
		tableName: "qtum_btc",
		dbIndex:   2,
	},
	"qtumeth": {
		tableName: "qtum_eth",
		dbIndex:   3,
	},
	"qtumusdt": {
		tableName: "qtum_usdt",
		dbIndex:   1,
	},
	"renbnb": {
		tableName: "ren_bnb",
		dbIndex:   4,
	},
	"renbtc": {
		tableName: "ren_btc",
		dbIndex:   2,
	},
	"renusdt": {
		tableName: "ren_usdt",
		dbIndex:   1,
	},
	"repbtc": {
		tableName: "rep_btc",
		dbIndex:   2,
	},
	"repeth": {
		tableName: "rep_eth",
		dbIndex:   3,
	},
	"repusdt": {
		tableName: "rep_usdt",
		dbIndex:   1,
	},
	"rvnbnb": {
		tableName: "rvn_bnb",
		dbIndex:   4,
	},
	"rvnbtc": {
		tableName: "rvn_btc",
		dbIndex:   2,
	},
	"rvnusdt": {
		tableName: "rvn_usdt",
		dbIndex:   1,
	},
	"snxbnb": {
		tableName: "snx_bnb",
		dbIndex:   4,
	},
	"snxbtc": {
		tableName: "snx_btc",
		dbIndex:   2,
	},
	"snxusdt": {
		tableName: "snx_usdt",
		dbIndex:   1,
	},
	"sushibnb": {
		tableName: "sushi_bnb",
		dbIndex:   4,
	},
	"sushibtc": {
		tableName: "sushi_btc",
		dbIndex:   2,
	},
	"sushiusdt": {
		tableName: "sushi_usdt",
		dbIndex:   1,
	},
	"tfuelbnb": {
		tableName: "tfuel_bnb",
		dbIndex:   4,
	},
	"tfuelbtc": {
		tableName: "tfuel_btc",
		dbIndex:   2,
	},
	"tfuelusdt": {
		tableName: "tfuel_usdt",
		dbIndex:   1,
	},
	"thetabnb": {
		tableName: "theta_bnb",
		dbIndex:   4,
	},
	"thetabtc": {
		tableName: "theta_btc",
		dbIndex:   2,
	},
	"thetaeth": {
		tableName: "theta_eth",
		dbIndex:   3,
	},
	"thetausdt": {
		tableName: "theta_usdt",
		dbIndex:   1,
	},
	"trxbnb": {
		tableName: "trx_bnb",
		dbIndex:   4,
	},
	"trxbtc": {
		tableName: "trx_btc",
		dbIndex:   2,
	},
	"trxeth": {
		tableName: "trx_eth",
		dbIndex:   3,
	},
	"trxusdt": {
		tableName: "trx_usdt",
		dbIndex:   1,
	},
	"umabtc": {
		tableName: "uma_btc",
		dbIndex:   2,
	},
	"umausdt": {
		tableName: "uma_usdt",
		dbIndex:   1,
	},
	"unibnb": {
		tableName: "uni_bnb",
		dbIndex:   4,
	},
	"unibtc": {
		tableName: "uni_btc",
		dbIndex:   2,
	},
	"uniusdt": {
		tableName: "uni_usdt",
		dbIndex:   1,
	},
	"usdcbnb": {
		tableName: "usdc_bnb",
		dbIndex:   4,
	},
	"usdcusdt": {
		tableName: "usdc_usdt",
		dbIndex:   1,
	},
	"usdtidr": {
		tableName: "usdt_idr",
		dbIndex:   5,
	},
	"wavesbnb": {
		tableName: "waves_bnb",
		dbIndex:   4,
	},
	"wavesbtc": {
		tableName: "waves_btc",
		dbIndex:   2,
	},
	"waveseth": {
		tableName: "waves_eth",
		dbIndex:   3,
	},
	"wavesusdt": {
		tableName: "waves_usdt",
		dbIndex:   1,
	},
	"wbtcbtc": {
		tableName: "wbtc_btc",
		dbIndex:   2,
	},
	"wbtceth": {
		tableName: "wbtc_eth",
		dbIndex:   3,
	},
	"wnxmbnb": {
		tableName: "wnxm_bnb",
		dbIndex:   4,
	},
	"wnxmbtc": {
		tableName: "wnxm_btc",
		dbIndex:   2,
	},
	"wnxmusdt": {
		tableName: "wnxm_usdt",
		dbIndex:   1,
	},
	"xembtc": {
		tableName: "xem_btc",
		dbIndex:   2,
	},
	"xemeth": {
		tableName: "xem_eth",
		dbIndex:   3,
	},
	"xlmbnb": {
		tableName: "xlm_bnb",
		dbIndex:   4,
	},
	"xlmbtc": {
		tableName: "xlm_btc",
		dbIndex:   2,
	},
	"xlmeth": {
		tableName: "xlm_eth",
		dbIndex:   3,
	},
	"xlmusdt": {
		tableName: "xlm_usdt",
		dbIndex:   1,
	},
	"xrpbnb": {
		tableName: "xrp_bnb",
		dbIndex:   4,
	},
	"xrpbtc": {
		tableName: "xrp_btc",
		dbIndex:   2,
	},
	"xrpeth": {
		tableName: "xrp_eth",
		dbIndex:   3,
	},
	"xrpusdt": {
		tableName: "xrp_usdt",
		dbIndex:   1,
	},
	"xtzbnb": {
		tableName: "xtz_bnb",
		dbIndex:   4,
	},
	"xtzbtc": {
		tableName: "xtz_btc",
		dbIndex:   2,
	},
	"xtzusdt": {
		tableName: "xtz_usdt",
		dbIndex:   1,
	},
	"xzcbtc": {
		tableName: "xzc_btc",
		dbIndex:   2,
	},
	"xzceth": {
		tableName: "xzc_eth",
		dbIndex:   3,
	},
	"xzcusdt": {
		tableName: "xzc_usdt",
		dbIndex:   1,
	},
	"yfibnb": {
		tableName: "yfi_bnb",
		dbIndex:   4,
	},
	"yfibtc": {
		tableName: "yfi_btc",
		dbIndex:   2,
	},
	"yfiusdt": {
		tableName: "yfi_usdt",
		dbIndex:   1,
	},
	"yfiibnb": {
		tableName: "yfii_bnb",
		dbIndex:   4,
	},
	"yfiibtc": {
		tableName: "yfii_btc",
		dbIndex:   2,
	},
	"yfiiusdt": {
		tableName: "yfii_usdt",
		dbIndex:   1,
	},
	"zecbnb": {
		tableName: "zec_bnb",
		dbIndex:   4,
	},
	"zecbtc": {
		tableName: "zec_btc",
		dbIndex:   2,
	},
	"zeceth": {
		tableName: "zec_eth",
		dbIndex:   3,
	},
	"zecusdt": {
		tableName: "zec_usdt",
		dbIndex:   1,
	},
	"zilbnb": {
		tableName: "zil_bnb",
		dbIndex:   4,
	},
	"zilbtc": {
		tableName: "zil_btc",
		dbIndex:   2,
	},
	"zileth": {
		tableName: "zil_eth",
		dbIndex:   3,
	},
	"zilusdt": {
		tableName: "zil_usdt",
		dbIndex:   1,
	},
	"zrxbnb": {
		tableName: "zrx_bnb",
		dbIndex:   4,
	},
	"zrxbtc": {
		tableName: "zrx_btc",
		dbIndex:   2,
	},
	"zrxeth": {
		tableName: "zrx_eth",
		dbIndex:   3,
	},
	"zrxusdt": {
		tableName: "zrx_usdt",
		dbIndex:   1,
	},
}

func GetBinanceIndex(streamName string) (string, int) {
	if config, found := binanceMatrix[streamName]; found {
		return config.tableName, config.dbIndex
	}
	return "", -1
}
