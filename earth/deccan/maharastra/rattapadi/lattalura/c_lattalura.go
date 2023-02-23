package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LattaluraCounty interface {
	feud.County
	BBangi槃耆() feud.Barony
	BBarshi巴尔斯希() feud.Barony
	BBir比尔() feud.Barony
	BDarur达鲁尔() feud.Barony
	BLattalura罗多楼罗() feud.Barony
	BTagara塔加拉() feud.Barony
	BUdgir优陀耆厘() feud.Barony
}

type 罗多楼罗LattaluraCounty struct {
	feud.BaseCounty
	槃耆Bangi       feud.Barony
	巴尔斯希Barshi    feud.Barony
	比尔Bir         feud.Barony
	达鲁尔Darur      feud.Barony
	罗多楼罗Lattalura feud.Barony
	塔加拉Tagara     feud.Barony
	优陀耆厘Udgir     feud.Barony
}

func (c *罗多楼罗LattaluraCounty) BBangi槃耆() feud.Barony {
	return c.槃耆Bangi
}

func (c *罗多楼罗LattaluraCounty) BBarshi巴尔斯希() feud.Barony {
	return c.巴尔斯希Barshi
}

func (c *罗多楼罗LattaluraCounty) BBir比尔() feud.Barony {
	return c.比尔Bir
}

func (c *罗多楼罗LattaluraCounty) BDarur达鲁尔() feud.Barony {
	return c.达鲁尔Darur
}

func (c *罗多楼罗LattaluraCounty) BLattalura罗多楼罗() feud.Barony {
	return c.罗多楼罗Lattalura
}

func (c *罗多楼罗LattaluraCounty) BTagara塔加拉() feud.Barony {
	return c.塔加拉Tagara
}

func (c *罗多楼罗LattaluraCounty) BUdgir优陀耆厘() feud.Barony {
	return c.优陀耆厘Udgir
}

var CLattalura罗多楼罗 LattaluraCounty = &罗多楼罗LattaluraCounty{}

func init() {
	f := CLattalura罗多楼罗.(*罗多楼罗LattaluraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1212",
		Title:     "lattalura",
		TitleName: "罗多楼罗",
		TitleCode: "c_lattalura",
		Baronies:  map[string]feud.Barony{},
	}

	f.槃耆Bangi = BBangi槃耆
	f.槃耆Bangi.SetParent(f)

	f.巴尔斯希Barshi = BBarshi巴尔斯希
	f.巴尔斯希Barshi.SetParent(f)

	f.比尔Bir = BBir比尔
	f.比尔Bir.SetParent(f)

	f.达鲁尔Darur = BDarur达鲁尔
	f.达鲁尔Darur.SetParent(f)

	f.罗多楼罗Lattalura = BLattalura罗多楼罗
	f.罗多楼罗Lattalura.SetParent(f)

	f.塔加拉Tagara = BTagara塔加拉
	f.塔加拉Tagara.SetParent(f)

	f.优陀耆厘Udgir = BUdgir优陀耆厘
	f.优陀耆厘Udgir.SetParent(f)

}
