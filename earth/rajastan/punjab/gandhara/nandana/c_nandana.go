package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NandanaCounty interface {
	feud.County
	BHund候德() feud.Barony
	BKhewra凯沃拉() feud.Barony
	BKhushab库沙布() feud.Barony
	BNandana欢喜园() feud.Barony
	BSimhapura狮子城() feud.Barony
	BTulaja陀罗奢() feud.Barony
}

type 欢喜园NandanaCounty struct {
	feud.BaseCounty
	候德Hund       feud.Barony
	凯沃拉Khewra    feud.Barony
	库沙布Khushab   feud.Barony
	欢喜园Nandana   feud.Barony
	狮子城Simhapura feud.Barony
	陀罗奢Tulaja    feud.Barony
}

func (c *欢喜园NandanaCounty) BHund候德() feud.Barony {
	return c.候德Hund
}

func (c *欢喜园NandanaCounty) BKhewra凯沃拉() feud.Barony {
	return c.凯沃拉Khewra
}

func (c *欢喜园NandanaCounty) BKhushab库沙布() feud.Barony {
	return c.库沙布Khushab
}

func (c *欢喜园NandanaCounty) BNandana欢喜园() feud.Barony {
	return c.欢喜园Nandana
}

func (c *欢喜园NandanaCounty) BSimhapura狮子城() feud.Barony {
	return c.狮子城Simhapura
}

func (c *欢喜园NandanaCounty) BTulaja陀罗奢() feud.Barony {
	return c.陀罗奢Tulaja
}

var CNandana欢喜园 NandanaCounty = &欢喜园NandanaCounty{}

func init() {
	f := CNandana欢喜园.(*欢喜园NandanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1341",
		Title:     "nandana",
		TitleName: "欢喜园",
		TitleCode: "c_nandana",
		Baronies:  map[string]feud.Barony{},
	}

	f.候德Hund = BHund候德
	f.候德Hund.SetParent(f)

	f.凯沃拉Khewra = BKhewra凯沃拉
	f.凯沃拉Khewra.SetParent(f)

	f.库沙布Khushab = BKhushab库沙布
	f.库沙布Khushab.SetParent(f)

	f.欢喜园Nandana = BNandana欢喜园
	f.欢喜园Nandana.SetParent(f)

	f.狮子城Simhapura = BSimhapura狮子城
	f.狮子城Simhapura.SetParent(f)

	f.陀罗奢Tulaja = BTulaja陀罗奢
	f.陀罗奢Tulaja.SetParent(f)

}
