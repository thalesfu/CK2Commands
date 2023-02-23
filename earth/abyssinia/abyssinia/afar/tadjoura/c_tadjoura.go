package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TadjouraCounty interface {
	feud.County
	BBerenice贝勒尼基() feud.Barony
	BDikhil迪基勒() feud.Barony
	BDjibouti吉布提() feud.Barony
	BGalafi加拉费() feud.Barony
	BHolhol霍霍() feud.Barony
	BObock奥博克() feud.Barony
	BRanda兰达() feud.Barony
	BTadjoura塔朱拉() feud.Barony
}

type 塔朱拉TadjouraCounty struct {
	feud.BaseCounty
	贝勒尼基Berenice feud.Barony
	迪基勒Dikhil    feud.Barony
	吉布提Djibouti  feud.Barony
	加拉费Galafi    feud.Barony
	霍霍Holhol     feud.Barony
	奥博克Obock     feud.Barony
	兰达Randa      feud.Barony
	塔朱拉Tadjoura  feud.Barony
}

func (c *塔朱拉TadjouraCounty) BBerenice贝勒尼基() feud.Barony {
	return c.贝勒尼基Berenice
}

func (c *塔朱拉TadjouraCounty) BDikhil迪基勒() feud.Barony {
	return c.迪基勒Dikhil
}

func (c *塔朱拉TadjouraCounty) BDjibouti吉布提() feud.Barony {
	return c.吉布提Djibouti
}

func (c *塔朱拉TadjouraCounty) BGalafi加拉费() feud.Barony {
	return c.加拉费Galafi
}

func (c *塔朱拉TadjouraCounty) BHolhol霍霍() feud.Barony {
	return c.霍霍Holhol
}

func (c *塔朱拉TadjouraCounty) BObock奥博克() feud.Barony {
	return c.奥博克Obock
}

func (c *塔朱拉TadjouraCounty) BRanda兰达() feud.Barony {
	return c.兰达Randa
}

func (c *塔朱拉TadjouraCounty) BTadjoura塔朱拉() feud.Barony {
	return c.塔朱拉Tadjoura
}

var CTadjoura塔朱拉 TadjouraCounty = &塔朱拉TadjouraCounty{}

func init() {
	f := CTadjoura塔朱拉.(*塔朱拉TadjouraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "874",
		Title:     "tadjoura",
		TitleName: "塔朱拉",
		TitleCode: "c_tadjoura",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝勒尼基Berenice = BBerenice贝勒尼基
	f.贝勒尼基Berenice.SetParent(f)

	f.迪基勒Dikhil = BDikhil迪基勒
	f.迪基勒Dikhil.SetParent(f)

	f.吉布提Djibouti = BDjibouti吉布提
	f.吉布提Djibouti.SetParent(f)

	f.加拉费Galafi = BGalafi加拉费
	f.加拉费Galafi.SetParent(f)

	f.霍霍Holhol = BHolhol霍霍
	f.霍霍Holhol.SetParent(f)

	f.奥博克Obock = BObock奥博克
	f.奥博克Obock.SetParent(f)

	f.兰达Randa = BRanda兰达
	f.兰达Randa.SetParent(f)

	f.塔朱拉Tadjoura = BTadjoura塔朱拉
	f.塔朱拉Tadjoura.SetParent(f)

}
