package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FurstenbergCounty interface {
	feud.County
	BBaar巴尔() feud.Barony
	BDonaueschingen多瑙埃兴根() feud.Barony
	BFurstenberg菲尔斯滕贝格() feud.Barony
	BHaslach哈斯拉赫() feud.Barony
	BHirsau希尔绍() feud.Barony
	BVillingen菲林根() feud.Barony
	BWolfach沃尔法赫() feud.Barony
	BZollern索伦() feud.Barony
}

type 菲尔斯滕贝格FurstenbergCounty struct {
	feud.BaseCounty
	巴尔Baar              feud.Barony
	多瑙埃兴根Donaueschingen feud.Barony
	菲尔斯滕贝格Furstenberg   feud.Barony
	哈斯拉赫Haslach         feud.Barony
	希尔绍Hirsau           feud.Barony
	菲林根Villingen        feud.Barony
	沃尔法赫Wolfach         feud.Barony
	索伦Zollern           feud.Barony
}

func (c *菲尔斯滕贝格FurstenbergCounty) BBaar巴尔() feud.Barony {
	return c.巴尔Baar
}

func (c *菲尔斯滕贝格FurstenbergCounty) BDonaueschingen多瑙埃兴根() feud.Barony {
	return c.多瑙埃兴根Donaueschingen
}

func (c *菲尔斯滕贝格FurstenbergCounty) BFurstenberg菲尔斯滕贝格() feud.Barony {
	return c.菲尔斯滕贝格Furstenberg
}

func (c *菲尔斯滕贝格FurstenbergCounty) BHaslach哈斯拉赫() feud.Barony {
	return c.哈斯拉赫Haslach
}

func (c *菲尔斯滕贝格FurstenbergCounty) BHirsau希尔绍() feud.Barony {
	return c.希尔绍Hirsau
}

func (c *菲尔斯滕贝格FurstenbergCounty) BVillingen菲林根() feud.Barony {
	return c.菲林根Villingen
}

func (c *菲尔斯滕贝格FurstenbergCounty) BWolfach沃尔法赫() feud.Barony {
	return c.沃尔法赫Wolfach
}

func (c *菲尔斯滕贝格FurstenbergCounty) BZollern索伦() feud.Barony {
	return c.索伦Zollern
}

var CFurstenberg菲尔斯滕贝格 FurstenbergCounty = &菲尔斯滕贝格FurstenbergCounty{}

func init() {
	f := CFurstenberg菲尔斯滕贝格.(*菲尔斯滕贝格FurstenbergCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "251",
		Title:     "furstenberg",
		TitleName: "菲尔斯滕贝格",
		TitleCode: "c_furstenberg",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔Baar = BBaar巴尔
	f.巴尔Baar.SetParent(f)

	f.多瑙埃兴根Donaueschingen = BDonaueschingen多瑙埃兴根
	f.多瑙埃兴根Donaueschingen.SetParent(f)

	f.菲尔斯滕贝格Furstenberg = BFurstenberg菲尔斯滕贝格
	f.菲尔斯滕贝格Furstenberg.SetParent(f)

	f.哈斯拉赫Haslach = BHaslach哈斯拉赫
	f.哈斯拉赫Haslach.SetParent(f)

	f.希尔绍Hirsau = BHirsau希尔绍
	f.希尔绍Hirsau.SetParent(f)

	f.菲林根Villingen = BVillingen菲林根
	f.菲林根Villingen.SetParent(f)

	f.沃尔法赫Wolfach = BWolfach沃尔法赫
	f.沃尔法赫Wolfach.SetParent(f)

	f.索伦Zollern = BZollern索伦
	f.索伦Zollern.SetParent(f)

}
