package italy

import (
	"github.com/thalesfu/CK2Commands/earth/italy/genoa"
	"github.com/thalesfu/CK2Commands/earth/italy/italy"
	"github.com/thalesfu/CK2Commands/earth/italy/naples"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy"
	"github.com/thalesfu/CK2Commands/earth/italy/pisa"
	"github.com/thalesfu/CK2Commands/earth/italy/romagna"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia"
	"github.com/thalesfu/CK2Commands/earth/italy/venice"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ItalyEmpire interface {
	feud.Empire
	KGenoa热那亚() genoa.GenoaKingdom
	KItaly意大利() italy.ItalyKingdom
	KNaples那不勒斯() naples.NaplesKingdom
	KPapacy教廷属邦() papacy.PapacyKingdom
	KPisa比萨() pisa.PisaKingdom
	KRomagna罗马涅() romagna.RomagnaKingdom
	KSardinia撒丁与科西嘉() sardinia.SardiniaKingdom
	KVenice威尼斯() venice.VeniceKingdom
}

type 意大利亚ItalyEmpire struct {
	feud.BaseEmpire
	热那亚Genoa       genoa.GenoaKingdom
	意大利Italy       italy.ItalyKingdom
	那不勒斯Naples     naples.NaplesKingdom
	教廷属邦Papacy     papacy.PapacyKingdom
	比萨Pisa         pisa.PisaKingdom
	罗马涅Romagna     romagna.RomagnaKingdom
	撒丁与科西嘉Sardinia sardinia.SardiniaKingdom
	威尼斯Venice      venice.VeniceKingdom
}

func (e *意大利亚ItalyEmpire) KGenoa热那亚() genoa.GenoaKingdom {
	return e.热那亚Genoa
}

func (e *意大利亚ItalyEmpire) KItaly意大利() italy.ItalyKingdom {
	return e.意大利Italy
}

func (e *意大利亚ItalyEmpire) KNaples那不勒斯() naples.NaplesKingdom {
	return e.那不勒斯Naples
}

func (e *意大利亚ItalyEmpire) KPapacy教廷属邦() papacy.PapacyKingdom {
	return e.教廷属邦Papacy
}

func (e *意大利亚ItalyEmpire) KPisa比萨() pisa.PisaKingdom {
	return e.比萨Pisa
}

func (e *意大利亚ItalyEmpire) KRomagna罗马涅() romagna.RomagnaKingdom {
	return e.罗马涅Romagna
}

func (e *意大利亚ItalyEmpire) KSardinia撒丁与科西嘉() sardinia.SardiniaKingdom {
	return e.撒丁与科西嘉Sardinia
}

func (e *意大利亚ItalyEmpire) KVenice威尼斯() venice.VeniceKingdom {
	return e.威尼斯Venice
}

var EItaly意大利亚 ItalyEmpire = &意大利亚ItalyEmpire{}

func init() {
	f := EItaly意大利亚.(*意大利亚ItalyEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "italy",
		TitleName: "意大利亚",
		TitleCode: "e_italy",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.热那亚Genoa = genoa.KGenoa热那亚
	f.热那亚Genoa.SetParent(f)
	f.意大利Italy = italy.KItaly意大利
	f.意大利Italy.SetParent(f)
	f.那不勒斯Naples = naples.KNaples那不勒斯
	f.那不勒斯Naples.SetParent(f)
	f.教廷属邦Papacy = papacy.KPapacy教廷属邦
	f.教廷属邦Papacy.SetParent(f)
	f.比萨Pisa = pisa.KPisa比萨
	f.比萨Pisa.SetParent(f)
	f.罗马涅Romagna = romagna.KRomagna罗马涅
	f.罗马涅Romagna.SetParent(f)
	f.撒丁与科西嘉Sardinia = sardinia.KSardinia撒丁与科西嘉
	f.撒丁与科西嘉Sardinia.SetParent(f)
	f.威尼斯Venice = venice.KVenice威尼斯
	f.威尼斯Venice.SetParent(f)
}
