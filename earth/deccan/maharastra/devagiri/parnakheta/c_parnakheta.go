package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ParnakhetaCounty interface {
	feud.County
	BAkola阿拘罗() feud.Barony
	BAmbadapura菴婆陀补罗() feud.Barony
	BBetora吠都罗() feud.Barony
	BBuldhana菩罗荼那() feud.Barony
	BKaranja迦兰阇() feud.Barony
	BMahkar摩诃羯罗() feud.Barony
	BParnakheta半挐契吒() feud.Barony
}

type 半挐契吒ParnakhetaCounty struct {
	feud.BaseCounty
	阿拘罗Akola        feud.Barony
	菴婆陀补罗Ambadapura feud.Barony
	吠都罗Betora       feud.Barony
	菩罗荼那Buldhana    feud.Barony
	迦兰阇Karanja      feud.Barony
	摩诃羯罗Mahkar      feud.Barony
	半挐契吒Parnakheta  feud.Barony
}

func (c *半挐契吒ParnakhetaCounty) BAkola阿拘罗() feud.Barony {
	return c.阿拘罗Akola
}

func (c *半挐契吒ParnakhetaCounty) BAmbadapura菴婆陀补罗() feud.Barony {
	return c.菴婆陀补罗Ambadapura
}

func (c *半挐契吒ParnakhetaCounty) BBetora吠都罗() feud.Barony {
	return c.吠都罗Betora
}

func (c *半挐契吒ParnakhetaCounty) BBuldhana菩罗荼那() feud.Barony {
	return c.菩罗荼那Buldhana
}

func (c *半挐契吒ParnakhetaCounty) BKaranja迦兰阇() feud.Barony {
	return c.迦兰阇Karanja
}

func (c *半挐契吒ParnakhetaCounty) BMahkar摩诃羯罗() feud.Barony {
	return c.摩诃羯罗Mahkar
}

func (c *半挐契吒ParnakhetaCounty) BParnakheta半挐契吒() feud.Barony {
	return c.半挐契吒Parnakheta
}

var CParnakheta半挐契吒 ParnakhetaCounty = &半挐契吒ParnakhetaCounty{}

func init() {
	f := CParnakheta半挐契吒.(*半挐契吒ParnakhetaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1261",
		Title:     "parnakheta",
		TitleName: "半挐契吒",
		TitleCode: "c_parnakheta",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拘罗Akola = BAkola阿拘罗
	f.阿拘罗Akola.SetParent(f)

	f.菴婆陀补罗Ambadapura = BAmbadapura菴婆陀补罗
	f.菴婆陀补罗Ambadapura.SetParent(f)

	f.吠都罗Betora = BBetora吠都罗
	f.吠都罗Betora.SetParent(f)

	f.菩罗荼那Buldhana = BBuldhana菩罗荼那
	f.菩罗荼那Buldhana.SetParent(f)

	f.迦兰阇Karanja = BKaranja迦兰阇
	f.迦兰阇Karanja.SetParent(f)

	f.摩诃羯罗Mahkar = BMahkar摩诃羯罗
	f.摩诃羯罗Mahkar.SetParent(f)

	f.半挐契吒Parnakheta = BParnakheta半挐契吒
	f.半挐契吒Parnakheta.SetParent(f)

}
