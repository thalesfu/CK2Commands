package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BordeauxCounty interface {
	feud.County
	BBlaye布莱() feud.Barony
	BBordeaux波尔多() feud.Barony
	BBourg布尔() feud.Barony
	BCastillon卡斯蒂永() feud.Barony
	BLareole拉雷奥尔() feud.Barony
	BLasauve拉索夫() feud.Barony
	BLibourne利布尔讷() feud.Barony
	BStemilion圣埃米利翁() feud.Barony
}

type 波尔多BordeauxCounty struct {
	feud.BaseCounty
	布莱Blaye        feud.Barony
	波尔多Bordeaux    feud.Barony
	布尔Bourg        feud.Barony
	卡斯蒂永Castillon  feud.Barony
	拉雷奥尔Lareole    feud.Barony
	拉索夫Lasauve     feud.Barony
	利布尔讷Libourne   feud.Barony
	圣埃米利翁Stemilion feud.Barony
}

func (c *波尔多BordeauxCounty) BBlaye布莱() feud.Barony {
	return c.布莱Blaye
}

func (c *波尔多BordeauxCounty) BBordeaux波尔多() feud.Barony {
	return c.波尔多Bordeaux
}

func (c *波尔多BordeauxCounty) BBourg布尔() feud.Barony {
	return c.布尔Bourg
}

func (c *波尔多BordeauxCounty) BCastillon卡斯蒂永() feud.Barony {
	return c.卡斯蒂永Castillon
}

func (c *波尔多BordeauxCounty) BLareole拉雷奥尔() feud.Barony {
	return c.拉雷奥尔Lareole
}

func (c *波尔多BordeauxCounty) BLasauve拉索夫() feud.Barony {
	return c.拉索夫Lasauve
}

func (c *波尔多BordeauxCounty) BLibourne利布尔讷() feud.Barony {
	return c.利布尔讷Libourne
}

func (c *波尔多BordeauxCounty) BStemilion圣埃米利翁() feud.Barony {
	return c.圣埃米利翁Stemilion
}

var CBordeaux波尔多 BordeauxCounty = &波尔多BordeauxCounty{}

func init() {
	f := CBordeaux波尔多.(*波尔多BordeauxCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "149",
		Title:     "bordeaux",
		TitleName: "波尔多",
		TitleCode: "c_bordeaux",
		Baronies:  map[string]feud.Barony{},
	}

	f.布莱Blaye = BBlaye布莱
	f.布莱Blaye.SetParent(f)

	f.波尔多Bordeaux = BBordeaux波尔多
	f.波尔多Bordeaux.SetParent(f)

	f.布尔Bourg = BBourg布尔
	f.布尔Bourg.SetParent(f)

	f.卡斯蒂永Castillon = BCastillon卡斯蒂永
	f.卡斯蒂永Castillon.SetParent(f)

	f.拉雷奥尔Lareole = BLareole拉雷奥尔
	f.拉雷奥尔Lareole.SetParent(f)

	f.拉索夫Lasauve = BLasauve拉索夫
	f.拉索夫Lasauve.SetParent(f)

	f.利布尔讷Libourne = BLibourne利布尔讷
	f.利布尔讷Libourne.SetParent(f)

	f.圣埃米利翁Stemilion = BStemilion圣埃米利翁
	f.圣埃米利翁Stemilion.SetParent(f)

}
