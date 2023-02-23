package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FoggiaCounty interface {
	feud.County
	BBovino博维诺() feud.Barony
	BFoggia福贾() feud.Barony
	BLucera卢切拉() feud.Barony
	BSalapla萨拉普拉() feud.Barony
	BSiponto西蓬托() feud.Barony
	BTroia特罗亚() feud.Barony
	BVieste维耶斯泰() feud.Barony
}

type 福贾FoggiaCounty struct {
	feud.BaseCounty
	博维诺Bovino   feud.Barony
	福贾Foggia    feud.Barony
	卢切拉Lucera   feud.Barony
	萨拉普拉Salapla feud.Barony
	西蓬托Siponto  feud.Barony
	特罗亚Troia    feud.Barony
	维耶斯泰Vieste  feud.Barony
}

func (c *福贾FoggiaCounty) BBovino博维诺() feud.Barony {
	return c.博维诺Bovino
}

func (c *福贾FoggiaCounty) BFoggia福贾() feud.Barony {
	return c.福贾Foggia
}

func (c *福贾FoggiaCounty) BLucera卢切拉() feud.Barony {
	return c.卢切拉Lucera
}

func (c *福贾FoggiaCounty) BSalapla萨拉普拉() feud.Barony {
	return c.萨拉普拉Salapla
}

func (c *福贾FoggiaCounty) BSiponto西蓬托() feud.Barony {
	return c.西蓬托Siponto
}

func (c *福贾FoggiaCounty) BTroia特罗亚() feud.Barony {
	return c.特罗亚Troia
}

func (c *福贾FoggiaCounty) BVieste维耶斯泰() feud.Barony {
	return c.维耶斯泰Vieste
}

var CFoggia福贾 FoggiaCounty = &福贾FoggiaCounty{}

func init() {
	f := CFoggia福贾.(*福贾FoggiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "348",
		Title:     "foggia",
		TitleName: "福贾",
		TitleCode: "c_foggia",
		Baronies:  map[string]feud.Barony{},
	}

	f.博维诺Bovino = BBovino博维诺
	f.博维诺Bovino.SetParent(f)

	f.福贾Foggia = BFoggia福贾
	f.福贾Foggia.SetParent(f)

	f.卢切拉Lucera = BLucera卢切拉
	f.卢切拉Lucera.SetParent(f)

	f.萨拉普拉Salapla = BSalapla萨拉普拉
	f.萨拉普拉Salapla.SetParent(f)

	f.西蓬托Siponto = BSiponto西蓬托
	f.西蓬托Siponto.SetParent(f)

	f.特罗亚Troia = BTroia特罗亚
	f.特罗亚Troia.SetParent(f)

	f.维耶斯泰Vieste = BVieste维耶斯泰
	f.维耶斯泰Vieste.SetParent(f)

}
