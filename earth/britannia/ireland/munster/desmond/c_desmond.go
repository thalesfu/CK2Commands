package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DesmondCounty interface {
	feud.County
	BBlarney布拉尼() feud.Barony
	BClone克朗() feud.Barony
	BCloyne克洛因() feud.Barony
	BCork科克() feud.Barony
	BDunasead顿那瑟德() feud.Barony
	BFermoy弗莫伊() feud.Barony
	BKilbrittain基尔不列颠() feud.Barony
	BYoughal约尔() feud.Barony
}

type 德斯蒙德DesmondCounty struct {
	feud.BaseCounty
	布拉尼Blarney       feud.Barony
	克朗Clone          feud.Barony
	克洛因Cloyne        feud.Barony
	科克Cork           feud.Barony
	顿那瑟德Dunasead     feud.Barony
	弗莫伊Fermoy        feud.Barony
	基尔不列颠Kilbrittain feud.Barony
	约尔Youghal        feud.Barony
}

func (c *德斯蒙德DesmondCounty) BBlarney布拉尼() feud.Barony {
	return c.布拉尼Blarney
}

func (c *德斯蒙德DesmondCounty) BClone克朗() feud.Barony {
	return c.克朗Clone
}

func (c *德斯蒙德DesmondCounty) BCloyne克洛因() feud.Barony {
	return c.克洛因Cloyne
}

func (c *德斯蒙德DesmondCounty) BCork科克() feud.Barony {
	return c.科克Cork
}

func (c *德斯蒙德DesmondCounty) BDunasead顿那瑟德() feud.Barony {
	return c.顿那瑟德Dunasead
}

func (c *德斯蒙德DesmondCounty) BFermoy弗莫伊() feud.Barony {
	return c.弗莫伊Fermoy
}

func (c *德斯蒙德DesmondCounty) BKilbrittain基尔不列颠() feud.Barony {
	return c.基尔不列颠Kilbrittain
}

func (c *德斯蒙德DesmondCounty) BYoughal约尔() feud.Barony {
	return c.约尔Youghal
}

var CDesmond德斯蒙德 DesmondCounty = &德斯蒙德DesmondCounty{}

func init() {
	f := CDesmond德斯蒙德.(*德斯蒙德DesmondCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "14",
		Title:     "desmond",
		TitleName: "德斯蒙德",
		TitleCode: "c_desmond",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉尼Blarney = BBlarney布拉尼
	f.布拉尼Blarney.SetParent(f)

	f.克朗Clone = BClone克朗
	f.克朗Clone.SetParent(f)

	f.克洛因Cloyne = BCloyne克洛因
	f.克洛因Cloyne.SetParent(f)

	f.科克Cork = BCork科克
	f.科克Cork.SetParent(f)

	f.顿那瑟德Dunasead = BDunasead顿那瑟德
	f.顿那瑟德Dunasead.SetParent(f)

	f.弗莫伊Fermoy = BFermoy弗莫伊
	f.弗莫伊Fermoy.SetParent(f)

	f.基尔不列颠Kilbrittain = BKilbrittain基尔不列颠
	f.基尔不列颠Kilbrittain.SetParent(f)

	f.约尔Youghal = BYoughal约尔
	f.约尔Youghal.SetParent(f)

}
