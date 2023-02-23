package byzantium

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/amalfi"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/aydin"
	"github.com/thalesfu/CK2Commands/earth/byzantium/bulgaria"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium"
	"github.com/thalesfu/CK2Commands/earth/byzantium/candar"
	"github.com/thalesfu/CK2Commands/earth/byzantium/croatia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/cyprus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus"
	"github.com/thalesfu/CK2Commands/earth/byzantium/eretnid"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/germiyan"
	"github.com/thalesfu/CK2Commands/earth/byzantium/karaman"
	"github.com/thalesfu/CK2Commands/earth/byzantium/mentese"
	"github.com/thalesfu/CK2Commands/earth/byzantium/ottoman"
	"github.com/thalesfu/CK2Commands/earth/byzantium/rum"
	"github.com/thalesfu/CK2Commands/earth/byzantium/saruhan"
	"github.com/thalesfu/CK2Commands/earth/byzantium/serbia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily"
	"github.com/thalesfu/CK2Commands/earth/byzantium/tekke"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trinacria"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ByzantiumEmpire interface {
	feud.Empire
	KAmalfi阿马尔菲() amalfi.AmalfiKingdom
	KAnatolia安纳托利亚() anatolia.AnatoliaKingdom
	KArmenia亚美尼亚() armenia.ArmeniaKingdom
	KAydin艾登() aydin.AydinKingdom
	KBulgaria保加利亚() bulgaria.BulgariaKingdom
	KByzantium希腊() byzantium.ByzantiumKingdom
	KCandar钱达尔() candar.CandarKingdom
	KCroatia克罗地亚() croatia.CroatiaKingdom
	KCyprus塞浦路斯() cyprus.CyprusKingdom
	KEpirus伊庇鲁斯() epirus.EpirusKingdom
	KEretnid埃雷特纳() eretnid.EretnidKingdom
	KGeorgia格鲁吉亚() georgia.GeorgiaKingdom
	KGermiyan格尔米扬() germiyan.GermiyanKingdom
	KKaraman卡拉曼() karaman.KaramanKingdom
	KMentese门泰谢() mentese.MenteseKingdom
	KOttoman奥斯曼() ottoman.OttomanKingdom
	KRum罗姆() rum.RumKingdom
	KSaruhan萨鲁汗() saruhan.SaruhanKingdom
	KSerbia塞尔维亚() serbia.SerbiaKingdom
	KSicily西西里() sicily.SicilyKingdom
	KTekke泰凯() tekke.TekkeKingdom
	KThrace色雷斯() thrace.ThraceKingdom
	KTrebizond特拉比松() trebizond.TrebizondKingdom
	KTrinacria特里纳克里亚() trinacria.TrinacriaKingdom
}

type 拜占庭帝国ByzantiumEmpire struct {
	feud.BaseEmpire
	阿马尔菲Amalfi      amalfi.AmalfiKingdom
	安纳托利亚Anatolia   anatolia.AnatoliaKingdom
	亚美尼亚Armenia     armenia.ArmeniaKingdom
	艾登Aydin         aydin.AydinKingdom
	保加利亚Bulgaria    bulgaria.BulgariaKingdom
	希腊Byzantium     byzantium.ByzantiumKingdom
	钱达尔Candar       candar.CandarKingdom
	克罗地亚Croatia     croatia.CroatiaKingdom
	塞浦路斯Cyprus      cyprus.CyprusKingdom
	伊庇鲁斯Epirus      epirus.EpirusKingdom
	埃雷特纳Eretnid     eretnid.EretnidKingdom
	格鲁吉亚Georgia     georgia.GeorgiaKingdom
	格尔米扬Germiyan    germiyan.GermiyanKingdom
	卡拉曼Karaman      karaman.KaramanKingdom
	门泰谢Mentese      mentese.MenteseKingdom
	奥斯曼Ottoman      ottoman.OttomanKingdom
	罗姆Rum           rum.RumKingdom
	萨鲁汗Saruhan      saruhan.SaruhanKingdom
	塞尔维亚Serbia      serbia.SerbiaKingdom
	西西里Sicily       sicily.SicilyKingdom
	泰凯Tekke         tekke.TekkeKingdom
	色雷斯Thrace       thrace.ThraceKingdom
	特拉比松Trebizond   trebizond.TrebizondKingdom
	特里纳克里亚Trinacria trinacria.TrinacriaKingdom
}

func (e *拜占庭帝国ByzantiumEmpire) KAmalfi阿马尔菲() amalfi.AmalfiKingdom {
	return e.阿马尔菲Amalfi
}

func (e *拜占庭帝国ByzantiumEmpire) KAnatolia安纳托利亚() anatolia.AnatoliaKingdom {
	return e.安纳托利亚Anatolia
}

func (e *拜占庭帝国ByzantiumEmpire) KArmenia亚美尼亚() armenia.ArmeniaKingdom {
	return e.亚美尼亚Armenia
}

func (e *拜占庭帝国ByzantiumEmpire) KAydin艾登() aydin.AydinKingdom {
	return e.艾登Aydin
}

func (e *拜占庭帝国ByzantiumEmpire) KBulgaria保加利亚() bulgaria.BulgariaKingdom {
	return e.保加利亚Bulgaria
}

func (e *拜占庭帝国ByzantiumEmpire) KByzantium希腊() byzantium.ByzantiumKingdom {
	return e.希腊Byzantium
}

func (e *拜占庭帝国ByzantiumEmpire) KCandar钱达尔() candar.CandarKingdom {
	return e.钱达尔Candar
}

func (e *拜占庭帝国ByzantiumEmpire) KCroatia克罗地亚() croatia.CroatiaKingdom {
	return e.克罗地亚Croatia
}

func (e *拜占庭帝国ByzantiumEmpire) KCyprus塞浦路斯() cyprus.CyprusKingdom {
	return e.塞浦路斯Cyprus
}

func (e *拜占庭帝国ByzantiumEmpire) KEpirus伊庇鲁斯() epirus.EpirusKingdom {
	return e.伊庇鲁斯Epirus
}

func (e *拜占庭帝国ByzantiumEmpire) KEretnid埃雷特纳() eretnid.EretnidKingdom {
	return e.埃雷特纳Eretnid
}

func (e *拜占庭帝国ByzantiumEmpire) KGeorgia格鲁吉亚() georgia.GeorgiaKingdom {
	return e.格鲁吉亚Georgia
}

func (e *拜占庭帝国ByzantiumEmpire) KGermiyan格尔米扬() germiyan.GermiyanKingdom {
	return e.格尔米扬Germiyan
}

func (e *拜占庭帝国ByzantiumEmpire) KKaraman卡拉曼() karaman.KaramanKingdom {
	return e.卡拉曼Karaman
}

func (e *拜占庭帝国ByzantiumEmpire) KMentese门泰谢() mentese.MenteseKingdom {
	return e.门泰谢Mentese
}

func (e *拜占庭帝国ByzantiumEmpire) KOttoman奥斯曼() ottoman.OttomanKingdom {
	return e.奥斯曼Ottoman
}

func (e *拜占庭帝国ByzantiumEmpire) KRum罗姆() rum.RumKingdom {
	return e.罗姆Rum
}

func (e *拜占庭帝国ByzantiumEmpire) KSaruhan萨鲁汗() saruhan.SaruhanKingdom {
	return e.萨鲁汗Saruhan
}

func (e *拜占庭帝国ByzantiumEmpire) KSerbia塞尔维亚() serbia.SerbiaKingdom {
	return e.塞尔维亚Serbia
}

func (e *拜占庭帝国ByzantiumEmpire) KSicily西西里() sicily.SicilyKingdom {
	return e.西西里Sicily
}

func (e *拜占庭帝国ByzantiumEmpire) KTekke泰凯() tekke.TekkeKingdom {
	return e.泰凯Tekke
}

func (e *拜占庭帝国ByzantiumEmpire) KThrace色雷斯() thrace.ThraceKingdom {
	return e.色雷斯Thrace
}

func (e *拜占庭帝国ByzantiumEmpire) KTrebizond特拉比松() trebizond.TrebizondKingdom {
	return e.特拉比松Trebizond
}

func (e *拜占庭帝国ByzantiumEmpire) KTrinacria特里纳克里亚() trinacria.TrinacriaKingdom {
	return e.特里纳克里亚Trinacria
}

var EByzantium拜占庭帝国 ByzantiumEmpire = &拜占庭帝国ByzantiumEmpire{}

func init() {
	f := EByzantium拜占庭帝国.(*拜占庭帝国ByzantiumEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "byzantium",
		TitleName: "拜占庭帝国",
		TitleCode: "e_byzantium",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿马尔菲Amalfi = amalfi.KAmalfi阿马尔菲
	f.阿马尔菲Amalfi.SetParent(f)
	f.安纳托利亚Anatolia = anatolia.KAnatolia安纳托利亚
	f.安纳托利亚Anatolia.SetParent(f)
	f.亚美尼亚Armenia = armenia.KArmenia亚美尼亚
	f.亚美尼亚Armenia.SetParent(f)
	f.艾登Aydin = aydin.KAydin艾登
	f.艾登Aydin.SetParent(f)
	f.保加利亚Bulgaria = bulgaria.KBulgaria保加利亚
	f.保加利亚Bulgaria.SetParent(f)
	f.希腊Byzantium = byzantium.KByzantium希腊
	f.希腊Byzantium.SetParent(f)
	f.钱达尔Candar = candar.KCandar钱达尔
	f.钱达尔Candar.SetParent(f)
	f.克罗地亚Croatia = croatia.KCroatia克罗地亚
	f.克罗地亚Croatia.SetParent(f)
	f.塞浦路斯Cyprus = cyprus.KCyprus塞浦路斯
	f.塞浦路斯Cyprus.SetParent(f)
	f.伊庇鲁斯Epirus = epirus.KEpirus伊庇鲁斯
	f.伊庇鲁斯Epirus.SetParent(f)
	f.埃雷特纳Eretnid = eretnid.KEretnid埃雷特纳
	f.埃雷特纳Eretnid.SetParent(f)
	f.格鲁吉亚Georgia = georgia.KGeorgia格鲁吉亚
	f.格鲁吉亚Georgia.SetParent(f)
	f.格尔米扬Germiyan = germiyan.KGermiyan格尔米扬
	f.格尔米扬Germiyan.SetParent(f)
	f.卡拉曼Karaman = karaman.KKaraman卡拉曼
	f.卡拉曼Karaman.SetParent(f)
	f.门泰谢Mentese = mentese.KMentese门泰谢
	f.门泰谢Mentese.SetParent(f)
	f.奥斯曼Ottoman = ottoman.KOttoman奥斯曼
	f.奥斯曼Ottoman.SetParent(f)
	f.罗姆Rum = rum.KRum罗姆
	f.罗姆Rum.SetParent(f)
	f.萨鲁汗Saruhan = saruhan.KSaruhan萨鲁汗
	f.萨鲁汗Saruhan.SetParent(f)
	f.塞尔维亚Serbia = serbia.KSerbia塞尔维亚
	f.塞尔维亚Serbia.SetParent(f)
	f.西西里Sicily = sicily.KSicily西西里
	f.西西里Sicily.SetParent(f)
	f.泰凯Tekke = tekke.KTekke泰凯
	f.泰凯Tekke.SetParent(f)
	f.色雷斯Thrace = thrace.KThrace色雷斯
	f.色雷斯Thrace.SetParent(f)
	f.特拉比松Trebizond = trebizond.KTrebizond特拉比松
	f.特拉比松Trebizond.SetParent(f)
	f.特里纳克里亚Trinacria = trinacria.KTrinacria特里纳克里亚
	f.特里纳克里亚Trinacria.SetParent(f)
}
