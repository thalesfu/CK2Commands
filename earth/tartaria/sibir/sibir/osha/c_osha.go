package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OshaCounty interface {
	feud.County
	BBalagany巴拉加内() feud.Barony
	BKargaly卡尔加雷() feud.Barony
	BKotochigi科托奇吉() feud.Barony
	BOsha奥沙() feud.Barony
	BSartam萨尔塔姆() feud.Barony
	BVikulovo维库洛沃() feud.Barony
	BZaborka扎博尔卡() feud.Barony
}

type 奥沙OshaCounty struct {
	feud.BaseCounty
	巴拉加内Balagany  feud.Barony
	卡尔加雷Kargaly   feud.Barony
	科托奇吉Kotochigi feud.Barony
	奥沙Osha        feud.Barony
	萨尔塔姆Sartam    feud.Barony
	维库洛沃Vikulovo  feud.Barony
	扎博尔卡Zaborka   feud.Barony
}

func (c *奥沙OshaCounty) BBalagany巴拉加内() feud.Barony {
	return c.巴拉加内Balagany
}

func (c *奥沙OshaCounty) BKargaly卡尔加雷() feud.Barony {
	return c.卡尔加雷Kargaly
}

func (c *奥沙OshaCounty) BKotochigi科托奇吉() feud.Barony {
	return c.科托奇吉Kotochigi
}

func (c *奥沙OshaCounty) BOsha奥沙() feud.Barony {
	return c.奥沙Osha
}

func (c *奥沙OshaCounty) BSartam萨尔塔姆() feud.Barony {
	return c.萨尔塔姆Sartam
}

func (c *奥沙OshaCounty) BVikulovo维库洛沃() feud.Barony {
	return c.维库洛沃Vikulovo
}

func (c *奥沙OshaCounty) BZaborka扎博尔卡() feud.Barony {
	return c.扎博尔卡Zaborka
}

var COsha奥沙 OshaCounty = &奥沙OshaCounty{}

func init() {
	f := COsha奥沙.(*奥沙OshaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1864",
		Title:     "osha",
		TitleName: "奥沙",
		TitleCode: "c_osha",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉加内Balagany = BBalagany巴拉加内
	f.巴拉加内Balagany.SetParent(f)

	f.卡尔加雷Kargaly = BKargaly卡尔加雷
	f.卡尔加雷Kargaly.SetParent(f)

	f.科托奇吉Kotochigi = BKotochigi科托奇吉
	f.科托奇吉Kotochigi.SetParent(f)

	f.奥沙Osha = BOsha奥沙
	f.奥沙Osha.SetParent(f)

	f.萨尔塔姆Sartam = BSartam萨尔塔姆
	f.萨尔塔姆Sartam.SetParent(f)

	f.维库洛沃Vikulovo = BVikulovo维库洛沃
	f.维库洛沃Vikulovo.SetParent(f)

	f.扎博尔卡Zaborka = BZaborka扎博尔卡
	f.扎博尔卡Zaborka.SetParent(f)

}
