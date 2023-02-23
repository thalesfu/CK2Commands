package damot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamotCounty interface {
	feud.County
	BAwasa阿瓦萨() feud.Barony
	BFasha法沙() feud.Barony
	BMalbarde马尔巴德() feud.Barony
	BMuger穆格尔() feud.Barony
	BTeltele特尔特勒() feud.Barony
	BZima济马() feud.Barony
}

type 达莫特DamotCounty struct {
	feud.BaseCounty
	阿瓦萨Awasa     feud.Barony
	法沙Fasha      feud.Barony
	马尔巴德Malbarde feud.Barony
	穆格尔Muger     feud.Barony
	特尔特勒Teltele  feud.Barony
	济马Zima       feud.Barony
}

func (c *达莫特DamotCounty) BAwasa阿瓦萨() feud.Barony {
	return c.阿瓦萨Awasa
}

func (c *达莫特DamotCounty) BFasha法沙() feud.Barony {
	return c.法沙Fasha
}

func (c *达莫特DamotCounty) BMalbarde马尔巴德() feud.Barony {
	return c.马尔巴德Malbarde
}

func (c *达莫特DamotCounty) BMuger穆格尔() feud.Barony {
	return c.穆格尔Muger
}

func (c *达莫特DamotCounty) BTeltele特尔特勒() feud.Barony {
	return c.特尔特勒Teltele
}

func (c *达莫特DamotCounty) BZima济马() feud.Barony {
	return c.济马Zima
}

var CDamot达莫特 DamotCounty = &达莫特DamotCounty{}

func init() {
	f := CDamot达莫特.(*达莫特DamotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1428",
		Title:     "damot",
		TitleName: "达莫特",
		TitleCode: "c_damot",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦萨Awasa = BAwasa阿瓦萨
	f.阿瓦萨Awasa.SetParent(f)

	f.法沙Fasha = BFasha法沙
	f.法沙Fasha.SetParent(f)

	f.马尔巴德Malbarde = BMalbarde马尔巴德
	f.马尔巴德Malbarde.SetParent(f)

	f.穆格尔Muger = BMuger穆格尔
	f.穆格尔Muger.SetParent(f)

	f.特尔特勒Teltele = BTeltele特尔特勒
	f.特尔特勒Teltele.SetParent(f)

	f.济马Zima = BZima济马
	f.济马Zima.SetParent(f)

}
