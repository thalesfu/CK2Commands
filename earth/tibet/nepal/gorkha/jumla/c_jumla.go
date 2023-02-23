package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JumlaCounty interface {
	feud.County
	BChainpur钱普尔() feud.Barony
	BDunai多瑙() feud.Barony
	BGamgadhi甘迦提() feud.Barony
	BJumla久姆拉() feud.Barony
	BManma曼尼() feud.Barony
	BMartadi马尔塔蒂() feud.Barony
	BSimikot斯弥拘吒() feud.Barony
}

type 久姆拉JumlaCounty struct {
	feud.BaseCounty
	钱普尔Chainpur feud.Barony
	多瑙Dunai     feud.Barony
	甘迦提Gamgadhi feud.Barony
	久姆拉Jumla    feud.Barony
	曼尼Manma     feud.Barony
	马尔塔蒂Martadi feud.Barony
	斯弥拘吒Simikot feud.Barony
}

func (c *久姆拉JumlaCounty) BChainpur钱普尔() feud.Barony {
	return c.钱普尔Chainpur
}

func (c *久姆拉JumlaCounty) BDunai多瑙() feud.Barony {
	return c.多瑙Dunai
}

func (c *久姆拉JumlaCounty) BGamgadhi甘迦提() feud.Barony {
	return c.甘迦提Gamgadhi
}

func (c *久姆拉JumlaCounty) BJumla久姆拉() feud.Barony {
	return c.久姆拉Jumla
}

func (c *久姆拉JumlaCounty) BManma曼尼() feud.Barony {
	return c.曼尼Manma
}

func (c *久姆拉JumlaCounty) BMartadi马尔塔蒂() feud.Barony {
	return c.马尔塔蒂Martadi
}

func (c *久姆拉JumlaCounty) BSimikot斯弥拘吒() feud.Barony {
	return c.斯弥拘吒Simikot
}

var CJumla久姆拉 JumlaCounty = &久姆拉JumlaCounty{}

func init() {
	f := CJumla久姆拉.(*久姆拉JumlaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1472",
		Title:     "jumla",
		TitleName: "久姆拉",
		TitleCode: "c_jumla",
		Baronies:  map[string]feud.Barony{},
	}

	f.钱普尔Chainpur = BChainpur钱普尔
	f.钱普尔Chainpur.SetParent(f)

	f.多瑙Dunai = BDunai多瑙
	f.多瑙Dunai.SetParent(f)

	f.甘迦提Gamgadhi = BGamgadhi甘迦提
	f.甘迦提Gamgadhi.SetParent(f)

	f.久姆拉Jumla = BJumla久姆拉
	f.久姆拉Jumla.SetParent(f)

	f.曼尼Manma = BManma曼尼
	f.曼尼Manma.SetParent(f)

	f.马尔塔蒂Martadi = BMartadi马尔塔蒂
	f.马尔塔蒂Martadi.SetParent(f)

	f.斯弥拘吒Simikot = BSimikot斯弥拘吒
	f.斯弥拘吒Simikot.SetParent(f)

}
