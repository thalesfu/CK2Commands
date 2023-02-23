package tregor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TregorCounty interface {
	feud.County
	BGuingamp甘冈() feud.Barony
	BLannion朗尼永() feud.Barony
	BLocquirec洛基雷克() feud.Barony
	BPedernec佩代尔内克() feud.Barony
	BPlougonven普卢贡旺() feud.Barony
	BPlourivo普卢里沃() feud.Barony
	BTreguier特雷吉耶() feud.Barony
}

type 特雷戈尔TregorCounty struct {
	feud.BaseCounty
	甘冈Guingamp     feud.Barony
	朗尼永Lannion     feud.Barony
	洛基雷克Locquirec  feud.Barony
	佩代尔内克Pedernec  feud.Barony
	普卢贡旺Plougonven feud.Barony
	普卢里沃Plourivo   feud.Barony
	特雷吉耶Treguier   feud.Barony
}

func (c *特雷戈尔TregorCounty) BGuingamp甘冈() feud.Barony {
	return c.甘冈Guingamp
}

func (c *特雷戈尔TregorCounty) BLannion朗尼永() feud.Barony {
	return c.朗尼永Lannion
}

func (c *特雷戈尔TregorCounty) BLocquirec洛基雷克() feud.Barony {
	return c.洛基雷克Locquirec
}

func (c *特雷戈尔TregorCounty) BPedernec佩代尔内克() feud.Barony {
	return c.佩代尔内克Pedernec
}

func (c *特雷戈尔TregorCounty) BPlougonven普卢贡旺() feud.Barony {
	return c.普卢贡旺Plougonven
}

func (c *特雷戈尔TregorCounty) BPlourivo普卢里沃() feud.Barony {
	return c.普卢里沃Plourivo
}

func (c *特雷戈尔TregorCounty) BTreguier特雷吉耶() feud.Barony {
	return c.特雷吉耶Treguier
}

var CTregor特雷戈尔 TregorCounty = &特雷戈尔TregorCounty{}

func init() {
	f := CTregor特雷戈尔.(*特雷戈尔TregorCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1936",
		Title:     "tregor",
		TitleName: "特雷戈尔",
		TitleCode: "c_tregor",
		Baronies:  map[string]feud.Barony{},
	}

	f.甘冈Guingamp = BGuingamp甘冈
	f.甘冈Guingamp.SetParent(f)

	f.朗尼永Lannion = BLannion朗尼永
	f.朗尼永Lannion.SetParent(f)

	f.洛基雷克Locquirec = BLocquirec洛基雷克
	f.洛基雷克Locquirec.SetParent(f)

	f.佩代尔内克Pedernec = BPedernec佩代尔内克
	f.佩代尔内克Pedernec.SetParent(f)

	f.普卢贡旺Plougonven = BPlougonven普卢贡旺
	f.普卢贡旺Plougonven.SetParent(f)

	f.普卢里沃Plourivo = BPlourivo普卢里沃
	f.普卢里沃Plourivo.SetParent(f)

	f.特雷吉耶Treguier = BTreguier特雷吉耶
	f.特雷吉耶Treguier.SetParent(f)

}
