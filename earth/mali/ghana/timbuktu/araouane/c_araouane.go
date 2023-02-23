package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AraouaneCounty interface {
	feud.County
	BAbalessa阿巴来萨() feud.Barony
	BAguelhok阿格勒霍克() feud.Barony
	BAraouane阿拉万() feud.Barony
	BGuezzam盖赞() feud.Barony
	BTamanrasset塔曼拉塞特() feud.Barony
	BTazrouk塔兹鲁克() feud.Barony
	BTessalit泰萨利特() feud.Barony
	BTinzaouten廷扎乌登() feud.Barony
}

type 阿拉万AraouaneCounty struct {
	feud.BaseCounty
	阿巴来萨Abalessa     feud.Barony
	阿格勒霍克Aguelhok    feud.Barony
	阿拉万Araouane      feud.Barony
	盖赞Guezzam        feud.Barony
	塔曼拉塞特Tamanrasset feud.Barony
	塔兹鲁克Tazrouk      feud.Barony
	泰萨利特Tessalit     feud.Barony
	廷扎乌登Tinzaouten   feud.Barony
}

func (c *阿拉万AraouaneCounty) BAbalessa阿巴来萨() feud.Barony {
	return c.阿巴来萨Abalessa
}

func (c *阿拉万AraouaneCounty) BAguelhok阿格勒霍克() feud.Barony {
	return c.阿格勒霍克Aguelhok
}

func (c *阿拉万AraouaneCounty) BAraouane阿拉万() feud.Barony {
	return c.阿拉万Araouane
}

func (c *阿拉万AraouaneCounty) BGuezzam盖赞() feud.Barony {
	return c.盖赞Guezzam
}

func (c *阿拉万AraouaneCounty) BTamanrasset塔曼拉塞特() feud.Barony {
	return c.塔曼拉塞特Tamanrasset
}

func (c *阿拉万AraouaneCounty) BTazrouk塔兹鲁克() feud.Barony {
	return c.塔兹鲁克Tazrouk
}

func (c *阿拉万AraouaneCounty) BTessalit泰萨利特() feud.Barony {
	return c.泰萨利特Tessalit
}

func (c *阿拉万AraouaneCounty) BTinzaouten廷扎乌登() feud.Barony {
	return c.廷扎乌登Tinzaouten
}

var CAraouane阿拉万 AraouaneCounty = &阿拉万AraouaneCounty{}

func init() {
	f := CAraouane阿拉万.(*阿拉万AraouaneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "917",
		Title:     "araouane",
		TitleName: "阿拉万",
		TitleCode: "c_araouane",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴来萨Abalessa = BAbalessa阿巴来萨
	f.阿巴来萨Abalessa.SetParent(f)

	f.阿格勒霍克Aguelhok = BAguelhok阿格勒霍克
	f.阿格勒霍克Aguelhok.SetParent(f)

	f.阿拉万Araouane = BAraouane阿拉万
	f.阿拉万Araouane.SetParent(f)

	f.盖赞Guezzam = BGuezzam盖赞
	f.盖赞Guezzam.SetParent(f)

	f.塔曼拉塞特Tamanrasset = BTamanrasset塔曼拉塞特
	f.塔曼拉塞特Tamanrasset.SetParent(f)

	f.塔兹鲁克Tazrouk = BTazrouk塔兹鲁克
	f.塔兹鲁克Tazrouk.SetParent(f)

	f.泰萨利特Tessalit = BTessalit泰萨利特
	f.泰萨利特Tessalit.SetParent(f)

	f.廷扎乌登Tinzaouten = BTinzaouten廷扎乌登
	f.廷扎乌登Tinzaouten.SetParent(f)

}
