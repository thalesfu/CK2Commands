package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CorfuCounty interface {
	feud.County
	BAcinganorum阿金加诺鲁姆() feud.Barony
	BCorfu科孚() feud.Barony
	BDrepane德瑞帕农() feud.Barony
	BKorkyra克尔基拉() feud.Barony
	BPhaeacia费阿刻斯() feud.Barony
	BScheria斯刻里亚() feud.Barony
	BTaphos塔褔斯() feud.Barony
}

type 克尔基拉CorfuCounty struct {
	feud.BaseCounty
	阿金加诺鲁姆Acinganorum feud.Barony
	科孚Corfu           feud.Barony
	德瑞帕农Drepane       feud.Barony
	克尔基拉Korkyra       feud.Barony
	费阿刻斯Phaeacia      feud.Barony
	斯刻里亚Scheria       feud.Barony
	塔褔斯Taphos         feud.Barony
}

func (c *克尔基拉CorfuCounty) BAcinganorum阿金加诺鲁姆() feud.Barony {
	return c.阿金加诺鲁姆Acinganorum
}

func (c *克尔基拉CorfuCounty) BCorfu科孚() feud.Barony {
	return c.科孚Corfu
}

func (c *克尔基拉CorfuCounty) BDrepane德瑞帕农() feud.Barony {
	return c.德瑞帕农Drepane
}

func (c *克尔基拉CorfuCounty) BKorkyra克尔基拉() feud.Barony {
	return c.克尔基拉Korkyra
}

func (c *克尔基拉CorfuCounty) BPhaeacia费阿刻斯() feud.Barony {
	return c.费阿刻斯Phaeacia
}

func (c *克尔基拉CorfuCounty) BScheria斯刻里亚() feud.Barony {
	return c.斯刻里亚Scheria
}

func (c *克尔基拉CorfuCounty) BTaphos塔褔斯() feud.Barony {
	return c.塔褔斯Taphos
}

var CCorfu克尔基拉 CorfuCounty = &克尔基拉CorfuCounty{}

func init() {
	f := CCorfu克尔基拉.(*克尔基拉CorfuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1890",
		Title:     "corfu",
		TitleName: "克尔基拉",
		TitleCode: "c_corfu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿金加诺鲁姆Acinganorum = BAcinganorum阿金加诺鲁姆
	f.阿金加诺鲁姆Acinganorum.SetParent(f)

	f.科孚Corfu = BCorfu科孚
	f.科孚Corfu.SetParent(f)

	f.德瑞帕农Drepane = BDrepane德瑞帕农
	f.德瑞帕农Drepane.SetParent(f)

	f.克尔基拉Korkyra = BKorkyra克尔基拉
	f.克尔基拉Korkyra.SetParent(f)

	f.费阿刻斯Phaeacia = BPhaeacia费阿刻斯
	f.费阿刻斯Phaeacia.SetParent(f)

	f.斯刻里亚Scheria = BScheria斯刻里亚
	f.斯刻里亚Scheria.SetParent(f)

	f.塔褔斯Taphos = BTaphos塔褔斯
	f.塔褔斯Taphos.SetParent(f)

}
