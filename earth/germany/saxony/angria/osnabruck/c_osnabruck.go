package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OsnabruckCounty interface {
	feud.County
	BBentheim本特海姆() feud.Barony
	BLingen林根() feud.Barony
	BMeppen梅彭() feud.Barony
	BOsnabruck奥斯纳布吕克() feud.Barony
	BQuackenbruck夸肯布吕克() feud.Barony
	BTecklenburg泰克伦堡() feud.Barony
	BWildeshausen维尔德斯豪森() feud.Barony
}

type 奥斯纳布吕克OsnabruckCounty struct {
	feud.BaseCounty
	本特海姆Bentheim       feud.Barony
	林根Lingen           feud.Barony
	梅彭Meppen           feud.Barony
	奥斯纳布吕克Osnabruck    feud.Barony
	夸肯布吕克Quackenbruck  feud.Barony
	泰克伦堡Tecklenburg    feud.Barony
	维尔德斯豪森Wildeshausen feud.Barony
}

func (c *奥斯纳布吕克OsnabruckCounty) BBentheim本特海姆() feud.Barony {
	return c.本特海姆Bentheim
}

func (c *奥斯纳布吕克OsnabruckCounty) BLingen林根() feud.Barony {
	return c.林根Lingen
}

func (c *奥斯纳布吕克OsnabruckCounty) BMeppen梅彭() feud.Barony {
	return c.梅彭Meppen
}

func (c *奥斯纳布吕克OsnabruckCounty) BOsnabruck奥斯纳布吕克() feud.Barony {
	return c.奥斯纳布吕克Osnabruck
}

func (c *奥斯纳布吕克OsnabruckCounty) BQuackenbruck夸肯布吕克() feud.Barony {
	return c.夸肯布吕克Quackenbruck
}

func (c *奥斯纳布吕克OsnabruckCounty) BTecklenburg泰克伦堡() feud.Barony {
	return c.泰克伦堡Tecklenburg
}

func (c *奥斯纳布吕克OsnabruckCounty) BWildeshausen维尔德斯豪森() feud.Barony {
	return c.维尔德斯豪森Wildeshausen
}

var COsnabruck奥斯纳布吕克 OsnabruckCounty = &奥斯纳布吕克OsnabruckCounty{}

func init() {
	f := COsnabruck奥斯纳布吕克.(*奥斯纳布吕克OsnabruckCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "87",
		Title:     "osnabruck",
		TitleName: "奥斯纳布吕克",
		TitleCode: "c_osnabruck",
		Baronies:  map[string]feud.Barony{},
	}

	f.本特海姆Bentheim = BBentheim本特海姆
	f.本特海姆Bentheim.SetParent(f)

	f.林根Lingen = BLingen林根
	f.林根Lingen.SetParent(f)

	f.梅彭Meppen = BMeppen梅彭
	f.梅彭Meppen.SetParent(f)

	f.奥斯纳布吕克Osnabruck = BOsnabruck奥斯纳布吕克
	f.奥斯纳布吕克Osnabruck.SetParent(f)

	f.夸肯布吕克Quackenbruck = BQuackenbruck夸肯布吕克
	f.夸肯布吕克Quackenbruck.SetParent(f)

	f.泰克伦堡Tecklenburg = BTecklenburg泰克伦堡
	f.泰克伦堡Tecklenburg.SetParent(f)

	f.维尔德斯豪森Wildeshausen = BWildeshausen维尔德斯豪森
	f.维尔德斯豪森Wildeshausen.SetParent(f)

}
