package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BacsCounty interface {
	feud.County
	BApatin阿帕廷() feud.Barony
	BBacs巴奇() feud.Barony
	BBacsalmas巴乔尔马什() feud.Barony
	BBaja鲍尧() feud.Barony
	BPancsova蓬乔沃() feud.Barony
	BPardany帕尔达尼() feud.Barony
	BSzintarev津陶雷夫() feud.Barony
	BZombor宗博尔() feud.Barony
}

type 巴奇BacsCounty struct {
	feud.BaseCounty
	阿帕廷Apatin      feud.Barony
	巴奇Bacs         feud.Barony
	巴乔尔马什Bacsalmas feud.Barony
	鲍尧Baja         feud.Barony
	蓬乔沃Pancsova    feud.Barony
	帕尔达尼Pardany    feud.Barony
	津陶雷夫Szintarev  feud.Barony
	宗博尔Zombor      feud.Barony
}

func (c *巴奇BacsCounty) BApatin阿帕廷() feud.Barony {
	return c.阿帕廷Apatin
}

func (c *巴奇BacsCounty) BBacs巴奇() feud.Barony {
	return c.巴奇Bacs
}

func (c *巴奇BacsCounty) BBacsalmas巴乔尔马什() feud.Barony {
	return c.巴乔尔马什Bacsalmas
}

func (c *巴奇BacsCounty) BBaja鲍尧() feud.Barony {
	return c.鲍尧Baja
}

func (c *巴奇BacsCounty) BPancsova蓬乔沃() feud.Barony {
	return c.蓬乔沃Pancsova
}

func (c *巴奇BacsCounty) BPardany帕尔达尼() feud.Barony {
	return c.帕尔达尼Pardany
}

func (c *巴奇BacsCounty) BSzintarev津陶雷夫() feud.Barony {
	return c.津陶雷夫Szintarev
}

func (c *巴奇BacsCounty) BZombor宗博尔() feud.Barony {
	return c.宗博尔Zombor
}

var CBacs巴奇 BacsCounty = &巴奇BacsCounty{}

func init() {
	f := CBacs巴奇.(*巴奇BacsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "518",
		Title:     "bacs",
		TitleName: "巴奇",
		TitleCode: "c_bacs",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿帕廷Apatin = BApatin阿帕廷
	f.阿帕廷Apatin.SetParent(f)

	f.巴奇Bacs = BBacs巴奇
	f.巴奇Bacs.SetParent(f)

	f.巴乔尔马什Bacsalmas = BBacsalmas巴乔尔马什
	f.巴乔尔马什Bacsalmas.SetParent(f)

	f.鲍尧Baja = BBaja鲍尧
	f.鲍尧Baja.SetParent(f)

	f.蓬乔沃Pancsova = BPancsova蓬乔沃
	f.蓬乔沃Pancsova.SetParent(f)

	f.帕尔达尼Pardany = BPardany帕尔达尼
	f.帕尔达尼Pardany.SetParent(f)

	f.津陶雷夫Szintarev = BSzintarev津陶雷夫
	f.津陶雷夫Szintarev.SetParent(f)

	f.宗博尔Zombor = BZombor宗博尔
	f.宗博尔Zombor.SetParent(f)

}
