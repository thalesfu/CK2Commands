package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IstriaCounty interface {
	feud.County
	BDuino杜伊诺() feud.Barony
	BFiume菲乌梅() feud.Barony
	BKarstberg卡尔斯特堡() feud.Barony
	BLovrana洛夫拉纳() feud.Barony
	BMitterburg米特堡() feud.Barony
	BPula普拉() feud.Barony
	BWolauska沃劳斯卡() feud.Barony
}

type 伊斯特里亚IstriaCounty struct {
	feud.BaseCounty
	杜伊诺Duino       feud.Barony
	菲乌梅Fiume       feud.Barony
	卡尔斯特堡Karstberg feud.Barony
	洛夫拉纳Lovrana    feud.Barony
	米特堡Mitterburg  feud.Barony
	普拉Pula         feud.Barony
	沃劳斯卡Wolauska   feud.Barony
}

func (c *伊斯特里亚IstriaCounty) BDuino杜伊诺() feud.Barony {
	return c.杜伊诺Duino
}

func (c *伊斯特里亚IstriaCounty) BFiume菲乌梅() feud.Barony {
	return c.菲乌梅Fiume
}

func (c *伊斯特里亚IstriaCounty) BKarstberg卡尔斯特堡() feud.Barony {
	return c.卡尔斯特堡Karstberg
}

func (c *伊斯特里亚IstriaCounty) BLovrana洛夫拉纳() feud.Barony {
	return c.洛夫拉纳Lovrana
}

func (c *伊斯特里亚IstriaCounty) BMitterburg米特堡() feud.Barony {
	return c.米特堡Mitterburg
}

func (c *伊斯特里亚IstriaCounty) BPula普拉() feud.Barony {
	return c.普拉Pula
}

func (c *伊斯特里亚IstriaCounty) BWolauska沃劳斯卡() feud.Barony {
	return c.沃劳斯卡Wolauska
}

var CIstria伊斯特里亚 IstriaCounty = &伊斯特里亚IstriaCounty{}

func init() {
	f := CIstria伊斯特里亚.(*伊斯特里亚IstriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "458",
		Title:     "istria",
		TitleName: "伊斯特里亚",
		TitleCode: "c_istria",
		Baronies:  map[string]feud.Barony{},
	}

	f.杜伊诺Duino = BDuino杜伊诺
	f.杜伊诺Duino.SetParent(f)

	f.菲乌梅Fiume = BFiume菲乌梅
	f.菲乌梅Fiume.SetParent(f)

	f.卡尔斯特堡Karstberg = BKarstberg卡尔斯特堡
	f.卡尔斯特堡Karstberg.SetParent(f)

	f.洛夫拉纳Lovrana = BLovrana洛夫拉纳
	f.洛夫拉纳Lovrana.SetParent(f)

	f.米特堡Mitterburg = BMitterburg米特堡
	f.米特堡Mitterburg.SetParent(f)

	f.普拉Pula = BPula普拉
	f.普拉Pula.SetParent(f)

	f.沃劳斯卡Wolauska = BWolauska沃劳斯卡
	f.沃劳斯卡Wolauska.SetParent(f)

}
