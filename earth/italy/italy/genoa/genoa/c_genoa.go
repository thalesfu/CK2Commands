package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GenoaCounty interface {
	feud.County
	BCarara卡拉拉() feud.Barony
	BChiavari基亚瓦里() feud.Barony
	BFosdinovo福斯迪诺沃() feud.Barony
	BGenoa热那亚() feud.Barony
	BLuna卢纳() feud.Barony
	BRapallo拉帕洛() feud.Barony
	BSiestri塞斯特里() feud.Barony
}

type 热那亚GenoaCounty struct {
	feud.BaseCounty
	卡拉拉Carara      feud.Barony
	基亚瓦里Chiavari   feud.Barony
	福斯迪诺沃Fosdinovo feud.Barony
	热那亚Genoa       feud.Barony
	卢纳Luna         feud.Barony
	拉帕洛Rapallo     feud.Barony
	塞斯特里Siestri    feud.Barony
}

func (c *热那亚GenoaCounty) BCarara卡拉拉() feud.Barony {
	return c.卡拉拉Carara
}

func (c *热那亚GenoaCounty) BChiavari基亚瓦里() feud.Barony {
	return c.基亚瓦里Chiavari
}

func (c *热那亚GenoaCounty) BFosdinovo福斯迪诺沃() feud.Barony {
	return c.福斯迪诺沃Fosdinovo
}

func (c *热那亚GenoaCounty) BGenoa热那亚() feud.Barony {
	return c.热那亚Genoa
}

func (c *热那亚GenoaCounty) BLuna卢纳() feud.Barony {
	return c.卢纳Luna
}

func (c *热那亚GenoaCounty) BRapallo拉帕洛() feud.Barony {
	return c.拉帕洛Rapallo
}

func (c *热那亚GenoaCounty) BSiestri塞斯特里() feud.Barony {
	return c.塞斯特里Siestri
}

var CGenoa热那亚 GenoaCounty = &热那亚GenoaCounty{}

func init() {
	f := CGenoa热那亚.(*热那亚GenoaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "233",
		Title:     "genoa",
		TitleName: "热那亚",
		TitleCode: "c_genoa",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉拉Carara = BCarara卡拉拉
	f.卡拉拉Carara.SetParent(f)

	f.基亚瓦里Chiavari = BChiavari基亚瓦里
	f.基亚瓦里Chiavari.SetParent(f)

	f.福斯迪诺沃Fosdinovo = BFosdinovo福斯迪诺沃
	f.福斯迪诺沃Fosdinovo.SetParent(f)

	f.热那亚Genoa = BGenoa热那亚
	f.热那亚Genoa.SetParent(f)

	f.卢纳Luna = BLuna卢纳
	f.卢纳Luna.SetParent(f)

	f.拉帕洛Rapallo = BRapallo拉帕洛
	f.拉帕洛Rapallo.SetParent(f)

	f.塞斯特里Siestri = BSiestri塞斯特里
	f.塞斯特里Siestri.SetParent(f)

}
