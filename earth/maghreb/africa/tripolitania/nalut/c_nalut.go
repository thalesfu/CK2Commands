package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NalutCounty interface {
	feud.County
	BFursata富尔塞塔() feud.Barony
	BNalut纳卢特() feud.Barony
	BTamaghzah塔马格扎() feud.Barony
}

type 纳卢特NalutCounty struct {
	feud.BaseCounty
	富尔塞塔Fursata   feud.Barony
	纳卢特Nalut      feud.Barony
	塔马格扎Tamaghzah feud.Barony
}

func (c *纳卢特NalutCounty) BFursata富尔塞塔() feud.Barony {
	return c.富尔塞塔Fursata
}

func (c *纳卢特NalutCounty) BNalut纳卢特() feud.Barony {
	return c.纳卢特Nalut
}

func (c *纳卢特NalutCounty) BTamaghzah塔马格扎() feud.Barony {
	return c.塔马格扎Tamaghzah
}

var CNalut纳卢特 NalutCounty = &纳卢特NalutCounty{}

func init() {
	f := CNalut纳卢特.(*纳卢特NalutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1724",
		Title:     "nalut",
		TitleName: "纳卢特",
		TitleCode: "c_nalut",
		Baronies:  map[string]feud.Barony{},
	}

	f.富尔塞塔Fursata = BFursata富尔塞塔
	f.富尔塞塔Fursata.SetParent(f)

	f.纳卢特Nalut = BNalut纳卢特
	f.纳卢特Nalut.SetParent(f)

	f.塔马格扎Tamaghzah = BTamaghzah塔马格扎
	f.塔马格扎Tamaghzah.SetParent(f)

}
