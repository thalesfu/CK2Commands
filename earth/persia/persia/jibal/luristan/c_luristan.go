package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LuristanCounty interface {
	feud.County
	BAlashtar阿拉什塔尔() feud.Barony
	BAligoodarz阿里古达尔兹() feud.Barony
	BBorujerd博鲁杰尔德() feud.Barony
	BDezbar德巴尔() feud.Barony
	BDorood道鲁德() feud.Barony
	BKhorramabad霍拉马巴德() feud.Barony
	BKoohdasht库赫达什特() feud.Barony
	BPoledokhtar波尔多赫塔尔() feud.Barony
}

type 洛雷斯坦LuristanCounty struct {
	feud.BaseCounty
	阿拉什塔尔Alashtar     feud.Barony
	阿里古达尔兹Aligoodarz  feud.Barony
	博鲁杰尔德Borujerd     feud.Barony
	德巴尔Dezbar         feud.Barony
	道鲁德Dorood         feud.Barony
	霍拉马巴德Khorramabad  feud.Barony
	库赫达什特Koohdasht    feud.Barony
	波尔多赫塔尔Poledokhtar feud.Barony
}

func (c *洛雷斯坦LuristanCounty) BAlashtar阿拉什塔尔() feud.Barony {
	return c.阿拉什塔尔Alashtar
}

func (c *洛雷斯坦LuristanCounty) BAligoodarz阿里古达尔兹() feud.Barony {
	return c.阿里古达尔兹Aligoodarz
}

func (c *洛雷斯坦LuristanCounty) BBorujerd博鲁杰尔德() feud.Barony {
	return c.博鲁杰尔德Borujerd
}

func (c *洛雷斯坦LuristanCounty) BDezbar德巴尔() feud.Barony {
	return c.德巴尔Dezbar
}

func (c *洛雷斯坦LuristanCounty) BDorood道鲁德() feud.Barony {
	return c.道鲁德Dorood
}

func (c *洛雷斯坦LuristanCounty) BKhorramabad霍拉马巴德() feud.Barony {
	return c.霍拉马巴德Khorramabad
}

func (c *洛雷斯坦LuristanCounty) BKoohdasht库赫达什特() feud.Barony {
	return c.库赫达什特Koohdasht
}

func (c *洛雷斯坦LuristanCounty) BPoledokhtar波尔多赫塔尔() feud.Barony {
	return c.波尔多赫塔尔Poledokhtar
}

var CLuristan洛雷斯坦 LuristanCounty = &洛雷斯坦LuristanCounty{}

func init() {
	f := CLuristan洛雷斯坦.(*洛雷斯坦LuristanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "657",
		Title:     "luristan",
		TitleName: "洛雷斯坦",
		TitleCode: "c_luristan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉什塔尔Alashtar = BAlashtar阿拉什塔尔
	f.阿拉什塔尔Alashtar.SetParent(f)

	f.阿里古达尔兹Aligoodarz = BAligoodarz阿里古达尔兹
	f.阿里古达尔兹Aligoodarz.SetParent(f)

	f.博鲁杰尔德Borujerd = BBorujerd博鲁杰尔德
	f.博鲁杰尔德Borujerd.SetParent(f)

	f.德巴尔Dezbar = BDezbar德巴尔
	f.德巴尔Dezbar.SetParent(f)

	f.道鲁德Dorood = BDorood道鲁德
	f.道鲁德Dorood.SetParent(f)

	f.霍拉马巴德Khorramabad = BKhorramabad霍拉马巴德
	f.霍拉马巴德Khorramabad.SetParent(f)

	f.库赫达什特Koohdasht = BKoohdasht库赫达什特
	f.库赫达什特Koohdasht.SetParent(f)

	f.波尔多赫塔尔Poledokhtar = BPoledokhtar波尔多赫塔尔
	f.波尔多赫塔尔Poledokhtar.SetParent(f)

}
