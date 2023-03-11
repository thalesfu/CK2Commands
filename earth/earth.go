package earth

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia"
	"github.com/thalesfu/CK2Commands/earth/arabia"
	"github.com/thalesfu/CK2Commands/earth/bengal"
	"github.com/thalesfu/CK2Commands/earth/britannia"
	"github.com/thalesfu/CK2Commands/earth/bulgaria"
	"github.com/thalesfu/CK2Commands/earth/byzantium"
	"github.com/thalesfu/CK2Commands/earth/carpathia"
	"github.com/thalesfu/CK2Commands/earth/chagatai"
	"github.com/thalesfu/CK2Commands/earth/china_west_governor"
	"github.com/thalesfu/CK2Commands/earth/cordoba"
	"github.com/thalesfu/CK2Commands/earth/deccan"
	"github.com/thalesfu/CK2Commands/earth/france"
	"github.com/thalesfu/CK2Commands/earth/germany"
	"github.com/thalesfu/CK2Commands/earth/golden_horde"
	"github.com/thalesfu/CK2Commands/earth/hre"
	"github.com/thalesfu/CK2Commands/earth/idel_ural"
	"github.com/thalesfu/CK2Commands/earth/il_khanate"
	"github.com/thalesfu/CK2Commands/earth/india"
	"github.com/thalesfu/CK2Commands/earth/italy"
	"github.com/thalesfu/CK2Commands/earth/kanem"
	"github.com/thalesfu/CK2Commands/earth/latin_empire"
	"github.com/thalesfu/CK2Commands/earth/lechczechrus"
	"github.com/thalesfu/CK2Commands/earth/maghreb"
	"github.com/thalesfu/CK2Commands/earth/mali"
	"github.com/thalesfu/CK2Commands/earth/mexikha"
	"github.com/thalesfu/CK2Commands/earth/mongol_empire"
	"github.com/thalesfu/CK2Commands/earth/nicaea"
	"github.com/thalesfu/CK2Commands/earth/outremer"
	"github.com/thalesfu/CK2Commands/earth/persia"
	"github.com/thalesfu/CK2Commands/earth/pirates"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe"
	"github.com/thalesfu/CK2Commands/earth/rajastan"
	"github.com/thalesfu/CK2Commands/earth/rebels"
	"github.com/thalesfu/CK2Commands/earth/roman_empire"
	"github.com/thalesfu/CK2Commands/earth/russia"
	"github.com/thalesfu/CK2Commands/earth/russian_empire"
	"github.com/thalesfu/CK2Commands/earth/scandinavia"
	"github.com/thalesfu/CK2Commands/earth/seljuk_turks"
	"github.com/thalesfu/CK2Commands/earth/shiite"
	"github.com/thalesfu/CK2Commands/earth/spain"
	"github.com/thalesfu/CK2Commands/earth/sunni"
	"github.com/thalesfu/CK2Commands/earth/tartaria"
	"github.com/thalesfu/CK2Commands/earth/tibet"
	"github.com/thalesfu/CK2Commands/earth/timurids"
	"github.com/thalesfu/CK2Commands/earth/turkestan"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire"
)
var (
    Abyssinia阿比西尼亚 = abyssinia.EAbyssinia阿比西尼亚
    Arabia阿拉伯帝国 = arabia.EArabia阿拉伯帝国
    Bengal旁伽罗帝国 = bengal.EBengal旁伽罗帝国
    Britannia不列颠尼亚 = britannia.EBritannia不列颠尼亚
    Bulgaria保加利亚 = bulgaria.EBulgaria保加利亚
    Byzantium拜占庭帝国 = byzantium.EByzantium拜占庭帝国
    Carpathia喀尔巴阡 = carpathia.ECarpathia喀尔巴阡
    Chagatai察合台 = chagatai.EChagatai察合台
    China_west_governor西域都护府 = china_west_governor.EChina_west_governor西域都护府
    Cordoba科尔多瓦 = cordoba.ECordoba科尔多瓦
    Deccan德干帝国 = deccan.EDeccan德干帝国
    France法兰克 = france.EFrance法兰克
    Germany日耳曼尼亚 = germany.EGermany日耳曼尼亚
    Golden_horde金帐汗国 = golden_horde.EGolden_horde金帐汗国
    Hre神圣罗马帝国 = hre.EHre神圣罗马帝国
    Idel_ural伏尔加_乌拉尔 = idel_ural.EIdel_ural伏尔加_乌拉尔
    Il_khanate伊儿汗国 = il_khanate.EIl_khanate伊儿汗国
    India印度 = india.EIndia印度
    Italy意大利亚 = italy.EItaly意大利亚
    Kanem加涅姆_博尔努 = kanem.EKanem加涅姆_博尔努
    Latin_empire拉丁帝国 = latin_empire.ELatin_empire拉丁帝国
    Lechczechrus斯拉夫联盟 = lechczechrus.ELechczechrus斯拉夫联盟
    Maghreb马格里布 = maghreb.EMaghreb马格里布
    Mali马里 = mali.EMali马里
    Mexikha阿兹特克帝国 = mexikha.EMexikha阿兹特克帝国
    Mongol_empire大蒙古国 = mongol_empire.EMongol_empire大蒙古国
    Nicaea尼西亚帝国 = nicaea.ENicaea尼西亚帝国
    Outremer海外 = outremer.EOutremer海外
    Persia波斯帝国 = persia.EPersia波斯帝国
    Pirates海盗 = pirates.EPirates海盗
    Pontic_steppe黑海草原 = pontic_steppe.EPontic_steppe黑海草原
    Rajastan罗阇萨傥那 = rajastan.ERajastan罗阇萨傥那
    Rebels叛军 = rebels.ERebels叛军
    Roman_empire罗马帝国 = roman_empire.ERoman_empire罗马帝国
    Russia罗斯帝国 = russia.ERussia罗斯帝国
    Russian_empire俄罗斯 = russian_empire.ERussian_empire俄罗斯
    Scandinavia斯堪的纳维亚 = scandinavia.EScandinavia斯堪的纳维亚
    Seljuk_turks塞尔柱帝国 = seljuk_turks.ESeljuk_turks塞尔柱帝国
    Shiite什叶派哈里发国 = shiite.EShiite什叶派哈里发国
    Spain西班牙 = spain.ESpain西班牙
    Sunni逊尼派哈里发国 = sunni.ESunni逊尼派哈里发国
    Tartaria鞑靼 = tartaria.ETartaria鞑靼
    Tibet吐蕃 = tibet.ETibet吐蕃
    Timurids帖木儿 = timurids.ETimurids帖木儿
    Turkestan图兰 = turkestan.ETurkestan图兰
    Wendish_empire文德帝国 = wendish_empire.EWendish_empire文德帝国
)
