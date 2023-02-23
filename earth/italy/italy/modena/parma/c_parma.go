package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ParmaCounty interface {
	feud.County
	BFidenza菲登扎() feud.Barony
	BFontanellato丰塔内拉托() feud.Barony
	BFornovo福尔诺沃() feud.Barony
	BGuastalla瓜斯塔拉() feud.Barony
	BLaspezia拉斯佩齐亚() feud.Barony
	BMassa马萨() feud.Barony
	BNoceto诺切托() feud.Barony
	BParma帕尔马() feud.Barony
}

type 帕尔马ParmaCounty struct {
	feud.BaseCounty
	菲登扎Fidenza        feud.Barony
	丰塔内拉托Fontanellato feud.Barony
	福尔诺沃Fornovo       feud.Barony
	瓜斯塔拉Guastalla     feud.Barony
	拉斯佩齐亚Laspezia     feud.Barony
	马萨Massa           feud.Barony
	诺切托Noceto         feud.Barony
	帕尔马Parma          feud.Barony
}

func (c *帕尔马ParmaCounty) BFidenza菲登扎() feud.Barony {
	return c.菲登扎Fidenza
}

func (c *帕尔马ParmaCounty) BFontanellato丰塔内拉托() feud.Barony {
	return c.丰塔内拉托Fontanellato
}

func (c *帕尔马ParmaCounty) BFornovo福尔诺沃() feud.Barony {
	return c.福尔诺沃Fornovo
}

func (c *帕尔马ParmaCounty) BGuastalla瓜斯塔拉() feud.Barony {
	return c.瓜斯塔拉Guastalla
}

func (c *帕尔马ParmaCounty) BLaspezia拉斯佩齐亚() feud.Barony {
	return c.拉斯佩齐亚Laspezia
}

func (c *帕尔马ParmaCounty) BMassa马萨() feud.Barony {
	return c.马萨Massa
}

func (c *帕尔马ParmaCounty) BNoceto诺切托() feud.Barony {
	return c.诺切托Noceto
}

func (c *帕尔马ParmaCounty) BParma帕尔马() feud.Barony {
	return c.帕尔马Parma
}

var CParma帕尔马 ParmaCounty = &帕尔马ParmaCounty{}

func init() {
	f := CParma帕尔马.(*帕尔马ParmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "321",
		Title:     "parma",
		TitleName: "帕尔马",
		TitleCode: "c_parma",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲登扎Fidenza = BFidenza菲登扎
	f.菲登扎Fidenza.SetParent(f)

	f.丰塔内拉托Fontanellato = BFontanellato丰塔内拉托
	f.丰塔内拉托Fontanellato.SetParent(f)

	f.福尔诺沃Fornovo = BFornovo福尔诺沃
	f.福尔诺沃Fornovo.SetParent(f)

	f.瓜斯塔拉Guastalla = BGuastalla瓜斯塔拉
	f.瓜斯塔拉Guastalla.SetParent(f)

	f.拉斯佩齐亚Laspezia = BLaspezia拉斯佩齐亚
	f.拉斯佩齐亚Laspezia.SetParent(f)

	f.马萨Massa = BMassa马萨
	f.马萨Massa.SetParent(f)

	f.诺切托Noceto = BNoceto诺切托
	f.诺切托Noceto.SetParent(f)

	f.帕尔马Parma = BParma帕尔马
	f.帕尔马Parma.SetParent(f)

}
