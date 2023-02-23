package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DulanCounty interface {
	feud.County
	BDulan都兰() feud.Barony
	BGouli沟里() feud.Barony
	BReshui热水() feud.Barony
	BXiangjia香加() feud.Barony
	BXiangride香日德() feud.Barony
	BXiariha夏日哈() feud.Barony
	BZongjia宗加() feud.Barony
}

type 都兰DulanCounty struct {
	feud.BaseCounty
	都兰Dulan      feud.Barony
	沟里Gouli      feud.Barony
	热水Reshui     feud.Barony
	香加Xiangjia   feud.Barony
	香日德Xiangride feud.Barony
	夏日哈Xiariha   feud.Barony
	宗加Zongjia    feud.Barony
}

func (c *都兰DulanCounty) BDulan都兰() feud.Barony {
	return c.都兰Dulan
}

func (c *都兰DulanCounty) BGouli沟里() feud.Barony {
	return c.沟里Gouli
}

func (c *都兰DulanCounty) BReshui热水() feud.Barony {
	return c.热水Reshui
}

func (c *都兰DulanCounty) BXiangjia香加() feud.Barony {
	return c.香加Xiangjia
}

func (c *都兰DulanCounty) BXiangride香日德() feud.Barony {
	return c.香日德Xiangride
}

func (c *都兰DulanCounty) BXiariha夏日哈() feud.Barony {
	return c.夏日哈Xiariha
}

func (c *都兰DulanCounty) BZongjia宗加() feud.Barony {
	return c.宗加Zongjia
}

var CDulan都兰 DulanCounty = &都兰DulanCounty{}

func init() {
	f := CDulan都兰.(*都兰DulanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1578",
		Title:     "dulan",
		TitleName: "都兰",
		TitleCode: "c_dulan",
		Baronies:  map[string]feud.Barony{},
	}

	f.都兰Dulan = BDulan都兰
	f.都兰Dulan.SetParent(f)

	f.沟里Gouli = BGouli沟里
	f.沟里Gouli.SetParent(f)

	f.热水Reshui = BReshui热水
	f.热水Reshui.SetParent(f)

	f.香加Xiangjia = BXiangjia香加
	f.香加Xiangjia.SetParent(f)

	f.香日德Xiangride = BXiangride香日德
	f.香日德Xiangride.SetParent(f)

	f.夏日哈Xiariha = BXiariha夏日哈
	f.夏日哈Xiariha.SetParent(f)

	f.宗加Zongjia = BZongjia宗加
	f.宗加Zongjia.SetParent(f)

}
