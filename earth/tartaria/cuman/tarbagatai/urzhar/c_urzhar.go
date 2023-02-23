package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UrzharCounty interface {
	feud.County
	BAkzhar阿克札尔() feud.Barony
	BAyagoz阿亚古兹() feud.Barony
	BAyzhar艾扎尔() feud.Barony
	BManzhar曼扎尔() feud.Barony
	BTaskesken塔斯克斯肯() feud.Barony
	BUrzhar乌尔札尔() feud.Barony
	BUshbik乌什比克() feud.Barony
}

type 阿亚古兹UrzharCounty struct {
	feud.BaseCounty
	阿克札尔Akzhar     feud.Barony
	阿亚古兹Ayagoz     feud.Barony
	艾扎尔Ayzhar      feud.Barony
	曼扎尔Manzhar     feud.Barony
	塔斯克斯肯Taskesken feud.Barony
	乌尔札尔Urzhar     feud.Barony
	乌什比克Ushbik     feud.Barony
}

func (c *阿亚古兹UrzharCounty) BAkzhar阿克札尔() feud.Barony {
	return c.阿克札尔Akzhar
}

func (c *阿亚古兹UrzharCounty) BAyagoz阿亚古兹() feud.Barony {
	return c.阿亚古兹Ayagoz
}

func (c *阿亚古兹UrzharCounty) BAyzhar艾扎尔() feud.Barony {
	return c.艾扎尔Ayzhar
}

func (c *阿亚古兹UrzharCounty) BManzhar曼扎尔() feud.Barony {
	return c.曼扎尔Manzhar
}

func (c *阿亚古兹UrzharCounty) BTaskesken塔斯克斯肯() feud.Barony {
	return c.塔斯克斯肯Taskesken
}

func (c *阿亚古兹UrzharCounty) BUrzhar乌尔札尔() feud.Barony {
	return c.乌尔札尔Urzhar
}

func (c *阿亚古兹UrzharCounty) BUshbik乌什比克() feud.Barony {
	return c.乌什比克Ushbik
}

var CUrzhar阿亚古兹 UrzharCounty = &阿亚古兹UrzharCounty{}

func init() {
	f := CUrzhar阿亚古兹.(*阿亚古兹UrzharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1427",
		Title:     "urzhar",
		TitleName: "阿亚古兹",
		TitleCode: "c_urzhar",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克札尔Akzhar = BAkzhar阿克札尔
	f.阿克札尔Akzhar.SetParent(f)

	f.阿亚古兹Ayagoz = BAyagoz阿亚古兹
	f.阿亚古兹Ayagoz.SetParent(f)

	f.艾扎尔Ayzhar = BAyzhar艾扎尔
	f.艾扎尔Ayzhar.SetParent(f)

	f.曼扎尔Manzhar = BManzhar曼扎尔
	f.曼扎尔Manzhar.SetParent(f)

	f.塔斯克斯肯Taskesken = BTaskesken塔斯克斯肯
	f.塔斯克斯肯Taskesken.SetParent(f)

	f.乌尔札尔Urzhar = BUrzhar乌尔札尔
	f.乌尔札尔Urzhar.SetParent(f)

	f.乌什比克Ushbik = BUshbik乌什比克
	f.乌什比克Ushbik.SetParent(f)

}
