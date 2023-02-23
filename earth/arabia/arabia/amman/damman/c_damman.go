package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DammanCounty interface {
	feud.County
	BAbuhadriya阿布哈德里耶() feud.Barony
	BAlaba艾巴() feud.Barony
	BAlhinnah欣纳赫() feud.Barony
	BAvan阿瓦恩() feud.Barony
	BDamman达曼() feud.Barony
	BJubail朱拜勒() feud.Barony
	BNajmah奈季迈() feud.Barony
	BQatif盖提夫() feud.Barony
}

type 盖提夫DammanCounty struct {
	feud.BaseCounty
	阿布哈德里耶Abuhadriya feud.Barony
	艾巴Alaba          feud.Barony
	欣纳赫Alhinnah      feud.Barony
	阿瓦恩Avan          feud.Barony
	达曼Damman         feud.Barony
	朱拜勒Jubail        feud.Barony
	奈季迈Najmah        feud.Barony
	盖提夫Qatif         feud.Barony
}

func (c *盖提夫DammanCounty) BAbuhadriya阿布哈德里耶() feud.Barony {
	return c.阿布哈德里耶Abuhadriya
}

func (c *盖提夫DammanCounty) BAlaba艾巴() feud.Barony {
	return c.艾巴Alaba
}

func (c *盖提夫DammanCounty) BAlhinnah欣纳赫() feud.Barony {
	return c.欣纳赫Alhinnah
}

func (c *盖提夫DammanCounty) BAvan阿瓦恩() feud.Barony {
	return c.阿瓦恩Avan
}

func (c *盖提夫DammanCounty) BDamman达曼() feud.Barony {
	return c.达曼Damman
}

func (c *盖提夫DammanCounty) BJubail朱拜勒() feud.Barony {
	return c.朱拜勒Jubail
}

func (c *盖提夫DammanCounty) BNajmah奈季迈() feud.Barony {
	return c.奈季迈Najmah
}

func (c *盖提夫DammanCounty) BQatif盖提夫() feud.Barony {
	return c.盖提夫Qatif
}

var CDamman盖提夫 DammanCounty = &盖提夫DammanCounty{}

func init() {
	f := CDamman盖提夫.(*盖提夫DammanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "651",
		Title:     "damman",
		TitleName: "盖提夫",
		TitleCode: "c_damman",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布哈德里耶Abuhadriya = BAbuhadriya阿布哈德里耶
	f.阿布哈德里耶Abuhadriya.SetParent(f)

	f.艾巴Alaba = BAlaba艾巴
	f.艾巴Alaba.SetParent(f)

	f.欣纳赫Alhinnah = BAlhinnah欣纳赫
	f.欣纳赫Alhinnah.SetParent(f)

	f.阿瓦恩Avan = BAvan阿瓦恩
	f.阿瓦恩Avan.SetParent(f)

	f.达曼Damman = BDamman达曼
	f.达曼Damman.SetParent(f)

	f.朱拜勒Jubail = BJubail朱拜勒
	f.朱拜勒Jubail.SetParent(f)

	f.奈季迈Najmah = BNajmah奈季迈
	f.奈季迈Najmah.SetParent(f)

	f.盖提夫Qatif = BQatif盖提夫
	f.盖提夫Qatif.SetParent(f)

}
