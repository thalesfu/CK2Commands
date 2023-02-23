package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlmatyCounty interface {
	feud.County
	BKapchagai卡普恰盖() feud.Barony
	BKargal卡尔加尔() feud.Barony
	BSaryesik萨雷耶西克() feud.Barony
}

type 阿拉木图AlmatyCounty struct {
	feud.BaseCounty
	卡普恰盖Kapchagai feud.Barony
	卡尔加尔Kargal    feud.Barony
	萨雷耶西克Saryesik feud.Barony
}

func (c *阿拉木图AlmatyCounty) BKapchagai卡普恰盖() feud.Barony {
	return c.卡普恰盖Kapchagai
}

func (c *阿拉木图AlmatyCounty) BKargal卡尔加尔() feud.Barony {
	return c.卡尔加尔Kargal
}

func (c *阿拉木图AlmatyCounty) BSaryesik萨雷耶西克() feud.Barony {
	return c.萨雷耶西克Saryesik
}

var CAlmaty阿拉木图 AlmatyCounty = &阿拉木图AlmatyCounty{}

func init() {
	f := CAlmaty阿拉木图.(*阿拉木图AlmatyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1797",
		Title:     "almaty",
		TitleName: "阿拉木图",
		TitleCode: "c_almaty",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡普恰盖Kapchagai = BKapchagai卡普恰盖
	f.卡普恰盖Kapchagai.SetParent(f)

	f.卡尔加尔Kargal = BKargal卡尔加尔
	f.卡尔加尔Kargal.SetParent(f)

	f.萨雷耶西克Saryesik = BSaryesik萨雷耶西克
	f.萨雷耶西克Saryesik.SetParent(f)

}
