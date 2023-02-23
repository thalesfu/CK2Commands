package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BadghisCounty interface {
	feud.County
	BBuzgan布兹甘() feud.Barony
	BKuhsim库哈斯姆() feud.Barony
	BKusuya库苏亚() feud.Barony
	BMalin马林() feud.Barony
	BPushang普尚格() feud.Barony
	BZurabad祖拉巴德() feud.Barony
}

type 巴德吉斯BadghisCounty struct {
	feud.BaseCounty
	布兹甘Buzgan   feud.Barony
	库哈斯姆Kuhsim  feud.Barony
	库苏亚Kusuya   feud.Barony
	马林Malin     feud.Barony
	普尚格Pushang  feud.Barony
	祖拉巴德Zurabad feud.Barony
}

func (c *巴德吉斯BadghisCounty) BBuzgan布兹甘() feud.Barony {
	return c.布兹甘Buzgan
}

func (c *巴德吉斯BadghisCounty) BKuhsim库哈斯姆() feud.Barony {
	return c.库哈斯姆Kuhsim
}

func (c *巴德吉斯BadghisCounty) BKusuya库苏亚() feud.Barony {
	return c.库苏亚Kusuya
}

func (c *巴德吉斯BadghisCounty) BMalin马林() feud.Barony {
	return c.马林Malin
}

func (c *巴德吉斯BadghisCounty) BPushang普尚格() feud.Barony {
	return c.普尚格Pushang
}

func (c *巴德吉斯BadghisCounty) BZurabad祖拉巴德() feud.Barony {
	return c.祖拉巴德Zurabad
}

var CBadghis巴德吉斯 BadghisCounty = &巴德吉斯BadghisCounty{}

func init() {
	f := CBadghis巴德吉斯.(*巴德吉斯BadghisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1538",
		Title:     "badghis",
		TitleName: "巴德吉斯",
		TitleCode: "c_badghis",
		Baronies:  map[string]feud.Barony{},
	}

	f.布兹甘Buzgan = BBuzgan布兹甘
	f.布兹甘Buzgan.SetParent(f)

	f.库哈斯姆Kuhsim = BKuhsim库哈斯姆
	f.库哈斯姆Kuhsim.SetParent(f)

	f.库苏亚Kusuya = BKusuya库苏亚
	f.库苏亚Kusuya.SetParent(f)

	f.马林Malin = BMalin马林
	f.马林Malin.SetParent(f)

	f.普尚格Pushang = BPushang普尚格
	f.普尚格Pushang.SetParent(f)

	f.祖拉巴德Zurabad = BZurabad祖拉巴德
	f.祖拉巴德Zurabad.SetParent(f)

}
