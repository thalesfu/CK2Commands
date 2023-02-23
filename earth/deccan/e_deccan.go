package deccan

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra"
	"github.com/thalesfu/CK2Commands/earth/deccan/karnata"
	"github.com/thalesfu/CK2Commands/earth/deccan/lanka"
	"github.com/thalesfu/CK2Commands/earth/deccan/maharastra"
	"github.com/thalesfu/CK2Commands/earth/deccan/tamilakam"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeccanEmpire interface {
	feud.Empire
	KAndhra案达罗() andhra.AndhraKingdom
	KKarnata羯罗拏吒() karnata.KarnataKingdom
	KLanka楞迦() lanka.LankaKingdom
	KMaharastra摩诃剌侘() maharastra.MaharastraKingdom
	KTamilakam昙密罗迦摩() tamilakam.TamilakamKingdom
	KTelingana得楞伽那() telingana.TelinganaKingdom
}

type 德干帝国DeccanEmpire struct {
	feud.BaseEmpire
	案达罗Andhra      andhra.AndhraKingdom
	羯罗拏吒Karnata    karnata.KarnataKingdom
	楞迦Lanka        lanka.LankaKingdom
	摩诃剌侘Maharastra maharastra.MaharastraKingdom
	昙密罗迦摩Tamilakam tamilakam.TamilakamKingdom
	得楞伽那Telingana  telingana.TelinganaKingdom
}

func (e *德干帝国DeccanEmpire) KAndhra案达罗() andhra.AndhraKingdom {
	return e.案达罗Andhra
}

func (e *德干帝国DeccanEmpire) KKarnata羯罗拏吒() karnata.KarnataKingdom {
	return e.羯罗拏吒Karnata
}

func (e *德干帝国DeccanEmpire) KLanka楞迦() lanka.LankaKingdom {
	return e.楞迦Lanka
}

func (e *德干帝国DeccanEmpire) KMaharastra摩诃剌侘() maharastra.MaharastraKingdom {
	return e.摩诃剌侘Maharastra
}

func (e *德干帝国DeccanEmpire) KTamilakam昙密罗迦摩() tamilakam.TamilakamKingdom {
	return e.昙密罗迦摩Tamilakam
}

func (e *德干帝国DeccanEmpire) KTelingana得楞伽那() telingana.TelinganaKingdom {
	return e.得楞伽那Telingana
}

var EDeccan德干帝国 DeccanEmpire = &德干帝国DeccanEmpire{}

func init() {
	f := EDeccan德干帝国.(*德干帝国DeccanEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "deccan",
		TitleName: "德干帝国",
		TitleCode: "e_deccan",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.案达罗Andhra = andhra.KAndhra案达罗
	f.案达罗Andhra.SetParent(f)
	f.羯罗拏吒Karnata = karnata.KKarnata羯罗拏吒
	f.羯罗拏吒Karnata.SetParent(f)
	f.楞迦Lanka = lanka.KLanka楞迦
	f.楞迦Lanka.SetParent(f)
	f.摩诃剌侘Maharastra = maharastra.KMaharastra摩诃剌侘
	f.摩诃剌侘Maharastra.SetParent(f)
	f.昙密罗迦摩Tamilakam = tamilakam.KTamilakam昙密罗迦摩
	f.昙密罗迦摩Tamilakam.SetParent(f)
	f.得楞伽那Telingana = telingana.KTelingana得楞伽那
	f.得楞伽那Telingana.SetParent(f)
}
