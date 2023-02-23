package forez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ForezCounty interface {
	feud.County
	BChalmazel沙尔马泽尔() feud.Barony
	BCharlieu沙尔略() feud.Barony
	BCouzan库藏() feud.Barony
	BFeurs弗尔() feud.Barony
	BMontbrison蒙布里松() feud.Barony
	BRoanne罗昂() feud.Barony
	BStetienne圣埃蒂安() feud.Barony
	BThiers梯也尔() feud.Barony
}

type 福雷ForezCounty struct {
	feud.BaseCounty
	沙尔马泽尔Chalmazel feud.Barony
	沙尔略Charlieu    feud.Barony
	库藏Couzan       feud.Barony
	弗尔Feurs        feud.Barony
	蒙布里松Montbrison feud.Barony
	罗昂Roanne       feud.Barony
	圣埃蒂安Stetienne  feud.Barony
	梯也尔Thiers      feud.Barony
}

func (c *福雷ForezCounty) BChalmazel沙尔马泽尔() feud.Barony {
	return c.沙尔马泽尔Chalmazel
}

func (c *福雷ForezCounty) BCharlieu沙尔略() feud.Barony {
	return c.沙尔略Charlieu
}

func (c *福雷ForezCounty) BCouzan库藏() feud.Barony {
	return c.库藏Couzan
}

func (c *福雷ForezCounty) BFeurs弗尔() feud.Barony {
	return c.弗尔Feurs
}

func (c *福雷ForezCounty) BMontbrison蒙布里松() feud.Barony {
	return c.蒙布里松Montbrison
}

func (c *福雷ForezCounty) BRoanne罗昂() feud.Barony {
	return c.罗昂Roanne
}

func (c *福雷ForezCounty) BStetienne圣埃蒂安() feud.Barony {
	return c.圣埃蒂安Stetienne
}

func (c *福雷ForezCounty) BThiers梯也尔() feud.Barony {
	return c.梯也尔Thiers
}

var CForez福雷 ForezCounty = &福雷ForezCounty{}

func init() {
	f := CForez福雷.(*福雷ForezCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "224",
		Title:     "forez",
		TitleName: "福雷",
		TitleCode: "c_forez",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙尔马泽尔Chalmazel = BChalmazel沙尔马泽尔
	f.沙尔马泽尔Chalmazel.SetParent(f)

	f.沙尔略Charlieu = BCharlieu沙尔略
	f.沙尔略Charlieu.SetParent(f)

	f.库藏Couzan = BCouzan库藏
	f.库藏Couzan.SetParent(f)

	f.弗尔Feurs = BFeurs弗尔
	f.弗尔Feurs.SetParent(f)

	f.蒙布里松Montbrison = BMontbrison蒙布里松
	f.蒙布里松Montbrison.SetParent(f)

	f.罗昂Roanne = BRoanne罗昂
	f.罗昂Roanne.SetParent(f)

	f.圣埃蒂安Stetienne = BStetienne圣埃蒂安
	f.圣埃蒂安Stetienne.SetParent(f)

	f.梯也尔Thiers = BThiers梯也尔
	f.梯也尔Thiers.SetParent(f)

}
