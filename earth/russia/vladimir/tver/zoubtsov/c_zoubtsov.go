package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZoubtsovCounty interface {
	feud.County
	BBatakovo巴塔科沃() feud.Barony
	BPapino帕皮诺() feud.Barony
	BUlyankovo乌良科沃() feud.Barony
	BVerigino韦里吉诺() feud.Barony
	BVysokino维索基诺() feud.Barony
	BYesinka叶辛卡() feud.Barony
	BZoubtsov祖布佐夫() feud.Barony
}

type 祖布佐夫ZoubtsovCounty struct {
	feud.BaseCounty
	巴塔科沃Batakovo  feud.Barony
	帕皮诺Papino     feud.Barony
	乌良科沃Ulyankovo feud.Barony
	韦里吉诺Verigino  feud.Barony
	维索基诺Vysokino  feud.Barony
	叶辛卡Yesinka    feud.Barony
	祖布佐夫Zoubtsov  feud.Barony
}

func (c *祖布佐夫ZoubtsovCounty) BBatakovo巴塔科沃() feud.Barony {
	return c.巴塔科沃Batakovo
}

func (c *祖布佐夫ZoubtsovCounty) BPapino帕皮诺() feud.Barony {
	return c.帕皮诺Papino
}

func (c *祖布佐夫ZoubtsovCounty) BUlyankovo乌良科沃() feud.Barony {
	return c.乌良科沃Ulyankovo
}

func (c *祖布佐夫ZoubtsovCounty) BVerigino韦里吉诺() feud.Barony {
	return c.韦里吉诺Verigino
}

func (c *祖布佐夫ZoubtsovCounty) BVysokino维索基诺() feud.Barony {
	return c.维索基诺Vysokino
}

func (c *祖布佐夫ZoubtsovCounty) BYesinka叶辛卡() feud.Barony {
	return c.叶辛卡Yesinka
}

func (c *祖布佐夫ZoubtsovCounty) BZoubtsov祖布佐夫() feud.Barony {
	return c.祖布佐夫Zoubtsov
}

var CZoubtsov祖布佐夫 ZoubtsovCounty = &祖布佐夫ZoubtsovCounty{}

func init() {
	f := CZoubtsov祖布佐夫.(*祖布佐夫ZoubtsovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1672",
		Title:     "zoubtsov",
		TitleName: "祖布佐夫",
		TitleCode: "c_zoubtsov",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴塔科沃Batakovo = BBatakovo巴塔科沃
	f.巴塔科沃Batakovo.SetParent(f)

	f.帕皮诺Papino = BPapino帕皮诺
	f.帕皮诺Papino.SetParent(f)

	f.乌良科沃Ulyankovo = BUlyankovo乌良科沃
	f.乌良科沃Ulyankovo.SetParent(f)

	f.韦里吉诺Verigino = BVerigino韦里吉诺
	f.韦里吉诺Verigino.SetParent(f)

	f.维索基诺Vysokino = BVysokino维索基诺
	f.维索基诺Vysokino.SetParent(f)

	f.叶辛卡Yesinka = BYesinka叶辛卡
	f.叶辛卡Yesinka.SetParent(f)

	f.祖布佐夫Zoubtsov = BZoubtsov祖布佐夫
	f.祖布佐夫Zoubtsov.SetParent(f)

}
