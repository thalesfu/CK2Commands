package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeneventoCounty interface {
	feud.County
	BAscoli阿斯科利() feud.Barony
	BAvellino阿韦利诺() feud.Barony
	BBenevento贝内文托() feud.Barony
	BConza孔扎() feud.Barony
	BFrigento弗里真托() feud.Barony
	BMontemarono蒙特马罗诺() feud.Barony
	BSanangelo圣安杰洛() feud.Barony
	BTrevico特雷维科() feud.Barony
}

type 贝内文托BeneventoCounty struct {
	feud.BaseCounty
	阿斯科利Ascoli       feud.Barony
	阿韦利诺Avellino     feud.Barony
	贝内文托Benevento    feud.Barony
	孔扎Conza          feud.Barony
	弗里真托Frigento     feud.Barony
	蒙特马罗诺Montemarono feud.Barony
	圣安杰洛Sanangelo    feud.Barony
	特雷维科Trevico      feud.Barony
}

func (c *贝内文托BeneventoCounty) BAscoli阿斯科利() feud.Barony {
	return c.阿斯科利Ascoli
}

func (c *贝内文托BeneventoCounty) BAvellino阿韦利诺() feud.Barony {
	return c.阿韦利诺Avellino
}

func (c *贝内文托BeneventoCounty) BBenevento贝内文托() feud.Barony {
	return c.贝内文托Benevento
}

func (c *贝内文托BeneventoCounty) BConza孔扎() feud.Barony {
	return c.孔扎Conza
}

func (c *贝内文托BeneventoCounty) BFrigento弗里真托() feud.Barony {
	return c.弗里真托Frigento
}

func (c *贝内文托BeneventoCounty) BMontemarono蒙特马罗诺() feud.Barony {
	return c.蒙特马罗诺Montemarono
}

func (c *贝内文托BeneventoCounty) BSanangelo圣安杰洛() feud.Barony {
	return c.圣安杰洛Sanangelo
}

func (c *贝内文托BeneventoCounty) BTrevico特雷维科() feud.Barony {
	return c.特雷维科Trevico
}

var CBenevento贝内文托 BeneventoCounty = &贝内文托BeneventoCounty{}

func init() {
	f := CBenevento贝内文托.(*贝内文托BeneventoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "335",
		Title:     "benevento",
		TitleName: "贝内文托",
		TitleCode: "c_benevento",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯科利Ascoli = BAscoli阿斯科利
	f.阿斯科利Ascoli.SetParent(f)

	f.阿韦利诺Avellino = BAvellino阿韦利诺
	f.阿韦利诺Avellino.SetParent(f)

	f.贝内文托Benevento = BBenevento贝内文托
	f.贝内文托Benevento.SetParent(f)

	f.孔扎Conza = BConza孔扎
	f.孔扎Conza.SetParent(f)

	f.弗里真托Frigento = BFrigento弗里真托
	f.弗里真托Frigento.SetParent(f)

	f.蒙特马罗诺Montemarono = BMontemarono蒙特马罗诺
	f.蒙特马罗诺Montemarono.SetParent(f)

	f.圣安杰洛Sanangelo = BSanangelo圣安杰洛
	f.圣安杰洛Sanangelo.SetParent(f)

	f.特雷维科Trevico = BTrevico特雷维科
	f.特雷维科Trevico.SetParent(f)

}
