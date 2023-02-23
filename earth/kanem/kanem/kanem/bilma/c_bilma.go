package bilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BilmaCounty interface {
	feud.County
	BBilma比尔马() feud.Barony
	BDanouala达努瓦拉() feud.Barony
	BGouam古阿姆() feud.Barony
	BKiroungou基伦古() feud.Barony
	BLibasso利巴索() feud.Barony
	BMateukous马特库斯() feud.Barony
	BTigue蒂盖() feud.Barony
}

type 比尔马BilmaCounty struct {
	feud.BaseCounty
	比尔马Bilma      feud.Barony
	达努瓦拉Danouala  feud.Barony
	古阿姆Gouam      feud.Barony
	基伦古Kiroungou  feud.Barony
	利巴索Libasso    feud.Barony
	马特库斯Mateukous feud.Barony
	蒂盖Tigue       feud.Barony
}

func (c *比尔马BilmaCounty) BBilma比尔马() feud.Barony {
	return c.比尔马Bilma
}

func (c *比尔马BilmaCounty) BDanouala达努瓦拉() feud.Barony {
	return c.达努瓦拉Danouala
}

func (c *比尔马BilmaCounty) BGouam古阿姆() feud.Barony {
	return c.古阿姆Gouam
}

func (c *比尔马BilmaCounty) BKiroungou基伦古() feud.Barony {
	return c.基伦古Kiroungou
}

func (c *比尔马BilmaCounty) BLibasso利巴索() feud.Barony {
	return c.利巴索Libasso
}

func (c *比尔马BilmaCounty) BMateukous马特库斯() feud.Barony {
	return c.马特库斯Mateukous
}

func (c *比尔马BilmaCounty) BTigue蒂盖() feud.Barony {
	return c.蒂盖Tigue
}

var CBilma比尔马 BilmaCounty = &比尔马BilmaCounty{}

func init() {
	f := CBilma比尔马.(*比尔马BilmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1735",
		Title:     "bilma",
		TitleName: "比尔马",
		TitleCode: "c_bilma",
		Baronies:  map[string]feud.Barony{},
	}

	f.比尔马Bilma = BBilma比尔马
	f.比尔马Bilma.SetParent(f)

	f.达努瓦拉Danouala = BDanouala达努瓦拉
	f.达努瓦拉Danouala.SetParent(f)

	f.古阿姆Gouam = BGouam古阿姆
	f.古阿姆Gouam.SetParent(f)

	f.基伦古Kiroungou = BKiroungou基伦古
	f.基伦古Kiroungou.SetParent(f)

	f.利巴索Libasso = BLibasso利巴索
	f.利巴索Libasso.SetParent(f)

	f.马特库斯Mateukous = BMateukous马特库斯
	f.马特库斯Mateukous.SetParent(f)

	f.蒂盖Tigue = BTigue蒂盖
	f.蒂盖Tigue.SetParent(f)

}
