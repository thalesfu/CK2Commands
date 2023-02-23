package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MesopotamiaCounty interface {
	feud.County
	BArghana阿格哈纳() feud.Barony
	BHani哈尼() feud.Barony
	BJermuk吉尔穆克() feud.Barony
	BLice利杰() feud.Barony
	BMartyropolis马图罗波利斯() feud.Barony
	BMayafaraqin玛亚法拉琴() feud.Barony
	BTigranakert迪亚巴克尔() feud.Barony
	BTzimisca兹米沙() feud.Barony
}

type 赫利亚特MesopotamiaCounty struct {
	feud.BaseCounty
	阿格哈纳Arghana        feud.Barony
	哈尼Hani             feud.Barony
	吉尔穆克Jermuk         feud.Barony
	利杰Lice             feud.Barony
	马图罗波利斯Martyropolis feud.Barony
	玛亚法拉琴Mayafaraqin   feud.Barony
	迪亚巴克尔Tigranakert   feud.Barony
	兹米沙Tzimisca        feud.Barony
}

func (c *赫利亚特MesopotamiaCounty) BArghana阿格哈纳() feud.Barony {
	return c.阿格哈纳Arghana
}

func (c *赫利亚特MesopotamiaCounty) BHani哈尼() feud.Barony {
	return c.哈尼Hani
}

func (c *赫利亚特MesopotamiaCounty) BJermuk吉尔穆克() feud.Barony {
	return c.吉尔穆克Jermuk
}

func (c *赫利亚特MesopotamiaCounty) BLice利杰() feud.Barony {
	return c.利杰Lice
}

func (c *赫利亚特MesopotamiaCounty) BMartyropolis马图罗波利斯() feud.Barony {
	return c.马图罗波利斯Martyropolis
}

func (c *赫利亚特MesopotamiaCounty) BMayafaraqin玛亚法拉琴() feud.Barony {
	return c.玛亚法拉琴Mayafaraqin
}

func (c *赫利亚特MesopotamiaCounty) BTigranakert迪亚巴克尔() feud.Barony {
	return c.迪亚巴克尔Tigranakert
}

func (c *赫利亚特MesopotamiaCounty) BTzimisca兹米沙() feud.Barony {
	return c.兹米沙Tzimisca
}

var CMesopotamia赫利亚特 MesopotamiaCounty = &赫利亚特MesopotamiaCounty{}

func init() {
	f := CMesopotamia赫利亚特.(*赫利亚特MesopotamiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "702",
		Title:     "mesopotamia",
		TitleName: "赫利亚特",
		TitleCode: "c_mesopotamia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格哈纳Arghana = BArghana阿格哈纳
	f.阿格哈纳Arghana.SetParent(f)

	f.哈尼Hani = BHani哈尼
	f.哈尼Hani.SetParent(f)

	f.吉尔穆克Jermuk = BJermuk吉尔穆克
	f.吉尔穆克Jermuk.SetParent(f)

	f.利杰Lice = BLice利杰
	f.利杰Lice.SetParent(f)

	f.马图罗波利斯Martyropolis = BMartyropolis马图罗波利斯
	f.马图罗波利斯Martyropolis.SetParent(f)

	f.玛亚法拉琴Mayafaraqin = BMayafaraqin玛亚法拉琴
	f.玛亚法拉琴Mayafaraqin.SetParent(f)

	f.迪亚巴克尔Tigranakert = BTigranakert迪亚巴克尔
	f.迪亚巴克尔Tigranakert.SetParent(f)

	f.兹米沙Tzimisca = BTzimisca兹米沙
	f.兹米沙Tzimisca.SetParent(f)

}
