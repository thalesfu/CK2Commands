package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TroyesCounty interface {
	feud.County
	BBrienne布列讷() feud.Barony
	BChacenay沙瑟奈() feud.Barony
	BChaumont肖蒙() feud.Barony
	BClairvaux克莱尔沃() feud.Barony
	BLangres朗格勒() feud.Barony
	BRosnay罗奈() feud.Barony
	BTroyes特鲁瓦() feud.Barony
}

type 特鲁瓦TroyesCounty struct {
	feud.BaseCounty
	布列讷Brienne    feud.Barony
	沙瑟奈Chacenay   feud.Barony
	肖蒙Chaumont    feud.Barony
	克莱尔沃Clairvaux feud.Barony
	朗格勒Langres    feud.Barony
	罗奈Rosnay      feud.Barony
	特鲁瓦Troyes     feud.Barony
}

func (c *特鲁瓦TroyesCounty) BBrienne布列讷() feud.Barony {
	return c.布列讷Brienne
}

func (c *特鲁瓦TroyesCounty) BChacenay沙瑟奈() feud.Barony {
	return c.沙瑟奈Chacenay
}

func (c *特鲁瓦TroyesCounty) BChaumont肖蒙() feud.Barony {
	return c.肖蒙Chaumont
}

func (c *特鲁瓦TroyesCounty) BClairvaux克莱尔沃() feud.Barony {
	return c.克莱尔沃Clairvaux
}

func (c *特鲁瓦TroyesCounty) BLangres朗格勒() feud.Barony {
	return c.朗格勒Langres
}

func (c *特鲁瓦TroyesCounty) BRosnay罗奈() feud.Barony {
	return c.罗奈Rosnay
}

func (c *特鲁瓦TroyesCounty) BTroyes特鲁瓦() feud.Barony {
	return c.特鲁瓦Troyes
}

var CTroyes特鲁瓦 TroyesCounty = &特鲁瓦TroyesCounty{}

func init() {
	f := CTroyes特鲁瓦.(*特鲁瓦TroyesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "130",
		Title:     "troyes",
		TitleName: "特鲁瓦",
		TitleCode: "c_troyes",
		Baronies:  map[string]feud.Barony{},
	}

	f.布列讷Brienne = BBrienne布列讷
	f.布列讷Brienne.SetParent(f)

	f.沙瑟奈Chacenay = BChacenay沙瑟奈
	f.沙瑟奈Chacenay.SetParent(f)

	f.肖蒙Chaumont = BChaumont肖蒙
	f.肖蒙Chaumont.SetParent(f)

	f.克莱尔沃Clairvaux = BClairvaux克莱尔沃
	f.克莱尔沃Clairvaux.SetParent(f)

	f.朗格勒Langres = BLangres朗格勒
	f.朗格勒Langres.SetParent(f)

	f.罗奈Rosnay = BRosnay罗奈
	f.罗奈Rosnay.SetParent(f)

	f.特鲁瓦Troyes = BTroyes特鲁瓦
	f.特鲁瓦Troyes.SetParent(f)

}
