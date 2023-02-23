package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RatanpurCounty interface {
	feud.County
	BPali波尼() feud.Barony
	BRatanpur剌怛那补罗() feud.Barony
	BSaravpur娑罗婆城() feud.Barony
	BSavarinarayana娑伐利那罗衍那() feud.Barony
	BShegaon势伽罗摩() feud.Barony
	BShirwal尸尔瓦尔() feud.Barony
	BYellandu翳兰头() feud.Barony
}

type 剌怛那补罗RatanpurCounty struct {
	feud.BaseCounty
	波尼Pali                feud.Barony
	剌怛那补罗Ratanpur         feud.Barony
	娑罗婆城Saravpur          feud.Barony
	娑伐利那罗衍那Savarinarayana feud.Barony
	势伽罗摩Shegaon           feud.Barony
	尸尔瓦尔Shirwal           feud.Barony
	翳兰头Yellandu           feud.Barony
}

func (c *剌怛那补罗RatanpurCounty) BPali波尼() feud.Barony {
	return c.波尼Pali
}

func (c *剌怛那补罗RatanpurCounty) BRatanpur剌怛那补罗() feud.Barony {
	return c.剌怛那补罗Ratanpur
}

func (c *剌怛那补罗RatanpurCounty) BSaravpur娑罗婆城() feud.Barony {
	return c.娑罗婆城Saravpur
}

func (c *剌怛那补罗RatanpurCounty) BSavarinarayana娑伐利那罗衍那() feud.Barony {
	return c.娑伐利那罗衍那Savarinarayana
}

func (c *剌怛那补罗RatanpurCounty) BShegaon势伽罗摩() feud.Barony {
	return c.势伽罗摩Shegaon
}

func (c *剌怛那补罗RatanpurCounty) BShirwal尸尔瓦尔() feud.Barony {
	return c.尸尔瓦尔Shirwal
}

func (c *剌怛那补罗RatanpurCounty) BYellandu翳兰头() feud.Barony {
	return c.翳兰头Yellandu
}

var CRatanpur剌怛那补罗 RatanpurCounty = &剌怛那补罗RatanpurCounty{}

func init() {
	f := CRatanpur剌怛那补罗.(*剌怛那补罗RatanpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1272",
		Title:     "ratanpur",
		TitleName: "剌怛那补罗",
		TitleCode: "c_ratanpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.波尼Pali = BPali波尼
	f.波尼Pali.SetParent(f)

	f.剌怛那补罗Ratanpur = BRatanpur剌怛那补罗
	f.剌怛那补罗Ratanpur.SetParent(f)

	f.娑罗婆城Saravpur = BSaravpur娑罗婆城
	f.娑罗婆城Saravpur.SetParent(f)

	f.娑伐利那罗衍那Savarinarayana = BSavarinarayana娑伐利那罗衍那
	f.娑伐利那罗衍那Savarinarayana.SetParent(f)

	f.势伽罗摩Shegaon = BShegaon势伽罗摩
	f.势伽罗摩Shegaon.SetParent(f)

	f.尸尔瓦尔Shirwal = BShirwal尸尔瓦尔
	f.尸尔瓦尔Shirwal.SetParent(f)

	f.翳兰头Yellandu = BYellandu翳兰头
	f.翳兰头Yellandu.SetParent(f)

}
