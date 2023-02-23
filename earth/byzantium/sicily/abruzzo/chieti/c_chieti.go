package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChietiCounty interface {
	feud.County
	BAtessa阿泰萨() feud.Barony
	BChieti基耶蒂() feud.Barony
	BCrecchio克雷基奥() feud.Barony
	BLanciano兰恰诺() feud.Barony
	BTermoli泰尔莫利() feud.Barony
	BVasto瓦斯托() feud.Barony
}

type 基耶蒂ChietiCounty struct {
	feud.BaseCounty
	阿泰萨Atessa    feud.Barony
	基耶蒂Chieti    feud.Barony
	克雷基奥Crecchio feud.Barony
	兰恰诺Lanciano  feud.Barony
	泰尔莫利Termoli  feud.Barony
	瓦斯托Vasto     feud.Barony
}

func (c *基耶蒂ChietiCounty) BAtessa阿泰萨() feud.Barony {
	return c.阿泰萨Atessa
}

func (c *基耶蒂ChietiCounty) BChieti基耶蒂() feud.Barony {
	return c.基耶蒂Chieti
}

func (c *基耶蒂ChietiCounty) BCrecchio克雷基奥() feud.Barony {
	return c.克雷基奥Crecchio
}

func (c *基耶蒂ChietiCounty) BLanciano兰恰诺() feud.Barony {
	return c.兰恰诺Lanciano
}

func (c *基耶蒂ChietiCounty) BTermoli泰尔莫利() feud.Barony {
	return c.泰尔莫利Termoli
}

func (c *基耶蒂ChietiCounty) BVasto瓦斯托() feud.Barony {
	return c.瓦斯托Vasto
}

var CChieti基耶蒂 ChietiCounty = &基耶蒂ChietiCounty{}

func init() {
	f := CChieti基耶蒂.(*基耶蒂ChietiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1707",
		Title:     "chieti",
		TitleName: "基耶蒂",
		TitleCode: "c_chieti",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿泰萨Atessa = BAtessa阿泰萨
	f.阿泰萨Atessa.SetParent(f)

	f.基耶蒂Chieti = BChieti基耶蒂
	f.基耶蒂Chieti.SetParent(f)

	f.克雷基奥Crecchio = BCrecchio克雷基奥
	f.克雷基奥Crecchio.SetParent(f)

	f.兰恰诺Lanciano = BLanciano兰恰诺
	f.兰恰诺Lanciano.SetParent(f)

	f.泰尔莫利Termoli = BTermoli泰尔莫利
	f.泰尔莫利Termoli.SetParent(f)

	f.瓦斯托Vasto = BVasto瓦斯托
	f.瓦斯托Vasto.SetParent(f)

}
