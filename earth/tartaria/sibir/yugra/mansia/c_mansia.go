package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MansiaCounty interface {
	feud.County
	BMamontovo马蒙托沃() feud.Barony
	BMansiysk曼西斯克() feud.Barony
	BNefteyugansk涅夫捷尤甘斯克() feud.Barony
	BPoykovskiy波伊科夫斯基() feud.Barony
	BPytyakh佩季_亚赫() feud.Barony
	BSamza萨姆扎() feud.Barony
	BYabin亚宾() feud.Barony
	BYalbak亚尔巴克() feud.Barony
}

type 曼西MansiaCounty struct {
	feud.BaseCounty
	马蒙托沃Mamontovo       feud.Barony
	曼西斯克Mansiysk        feud.Barony
	涅夫捷尤甘斯克Nefteyugansk feud.Barony
	波伊科夫斯基Poykovskiy    feud.Barony
	佩季_亚赫Pytyakh        feud.Barony
	萨姆扎Samza            feud.Barony
	亚宾Yabin             feud.Barony
	亚尔巴克Yalbak          feud.Barony
}

func (c *曼西MansiaCounty) BMamontovo马蒙托沃() feud.Barony {
	return c.马蒙托沃Mamontovo
}

func (c *曼西MansiaCounty) BMansiysk曼西斯克() feud.Barony {
	return c.曼西斯克Mansiysk
}

func (c *曼西MansiaCounty) BNefteyugansk涅夫捷尤甘斯克() feud.Barony {
	return c.涅夫捷尤甘斯克Nefteyugansk
}

func (c *曼西MansiaCounty) BPoykovskiy波伊科夫斯基() feud.Barony {
	return c.波伊科夫斯基Poykovskiy
}

func (c *曼西MansiaCounty) BPytyakh佩季_亚赫() feud.Barony {
	return c.佩季_亚赫Pytyakh
}

func (c *曼西MansiaCounty) BSamza萨姆扎() feud.Barony {
	return c.萨姆扎Samza
}

func (c *曼西MansiaCounty) BYabin亚宾() feud.Barony {
	return c.亚宾Yabin
}

func (c *曼西MansiaCounty) BYalbak亚尔巴克() feud.Barony {
	return c.亚尔巴克Yalbak
}

var CMansia曼西 MansiaCounty = &曼西MansiaCounty{}

func init() {
	f := CMansia曼西.(*曼西MansiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "892",
		Title:     "mansia",
		TitleName: "曼西",
		TitleCode: "c_mansia",
		Baronies:  map[string]feud.Barony{},
	}

	f.马蒙托沃Mamontovo = BMamontovo马蒙托沃
	f.马蒙托沃Mamontovo.SetParent(f)

	f.曼西斯克Mansiysk = BMansiysk曼西斯克
	f.曼西斯克Mansiysk.SetParent(f)

	f.涅夫捷尤甘斯克Nefteyugansk = BNefteyugansk涅夫捷尤甘斯克
	f.涅夫捷尤甘斯克Nefteyugansk.SetParent(f)

	f.波伊科夫斯基Poykovskiy = BPoykovskiy波伊科夫斯基
	f.波伊科夫斯基Poykovskiy.SetParent(f)

	f.佩季_亚赫Pytyakh = BPytyakh佩季_亚赫
	f.佩季_亚赫Pytyakh.SetParent(f)

	f.萨姆扎Samza = BSamza萨姆扎
	f.萨姆扎Samza.SetParent(f)

	f.亚宾Yabin = BYabin亚宾
	f.亚宾Yabin.SetParent(f)

	f.亚尔巴克Yalbak = BYalbak亚尔巴克
	f.亚尔巴克Yalbak.SetParent(f)

}
