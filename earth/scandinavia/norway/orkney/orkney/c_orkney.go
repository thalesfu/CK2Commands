package orkney

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrkneyCounty interface {
	feud.County
	BBirsay伯赛() feud.Barony
	BEgilsay埃吉尔赛() feud.Barony
	BKirkwall柯克沃尔() feud.Barony
	BOrphir奥费尔() feud.Barony
	BRonaldsay罗纳德赛() feud.Barony
	BSanday桑迪() feud.Barony
	BWestray韦斯特雷() feud.Barony
	BWyre怀尔() feud.Barony
}

type 奥克尼OrkneyCounty struct {
	feud.BaseCounty
	伯赛Birsay      feud.Barony
	埃吉尔赛Egilsay   feud.Barony
	柯克沃尔Kirkwall  feud.Barony
	奥费尔Orphir     feud.Barony
	罗纳德赛Ronaldsay feud.Barony
	桑迪Sanday      feud.Barony
	韦斯特雷Westray   feud.Barony
	怀尔Wyre        feud.Barony
}

func (c *奥克尼OrkneyCounty) BBirsay伯赛() feud.Barony {
	return c.伯赛Birsay
}

func (c *奥克尼OrkneyCounty) BEgilsay埃吉尔赛() feud.Barony {
	return c.埃吉尔赛Egilsay
}

func (c *奥克尼OrkneyCounty) BKirkwall柯克沃尔() feud.Barony {
	return c.柯克沃尔Kirkwall
}

func (c *奥克尼OrkneyCounty) BOrphir奥费尔() feud.Barony {
	return c.奥费尔Orphir
}

func (c *奥克尼OrkneyCounty) BRonaldsay罗纳德赛() feud.Barony {
	return c.罗纳德赛Ronaldsay
}

func (c *奥克尼OrkneyCounty) BSanday桑迪() feud.Barony {
	return c.桑迪Sanday
}

func (c *奥克尼OrkneyCounty) BWestray韦斯特雷() feud.Barony {
	return c.韦斯特雷Westray
}

func (c *奥克尼OrkneyCounty) BWyre怀尔() feud.Barony {
	return c.怀尔Wyre
}

var COrkney奥克尼 OrkneyCounty = &奥克尼OrkneyCounty{}

func init() {
	f := COrkney奥克尼.(*奥克尼OrkneyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "36",
		Title:     "orkney",
		TitleName: "奥克尼",
		TitleCode: "c_orkney",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯赛Birsay = BBirsay伯赛
	f.伯赛Birsay.SetParent(f)

	f.埃吉尔赛Egilsay = BEgilsay埃吉尔赛
	f.埃吉尔赛Egilsay.SetParent(f)

	f.柯克沃尔Kirkwall = BKirkwall柯克沃尔
	f.柯克沃尔Kirkwall.SetParent(f)

	f.奥费尔Orphir = BOrphir奥费尔
	f.奥费尔Orphir.SetParent(f)

	f.罗纳德赛Ronaldsay = BRonaldsay罗纳德赛
	f.罗纳德赛Ronaldsay.SetParent(f)

	f.桑迪Sanday = BSanday桑迪
	f.桑迪Sanday.SetParent(f)

	f.韦斯特雷Westray = BWestray韦斯特雷
	f.韦斯特雷Westray.SetParent(f)

	f.怀尔Wyre = BWyre怀尔
	f.怀尔Wyre.SetParent(f)

}
