package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SinopeCounty interface {
	feud.County
	BAndrapa安德拉帕() feud.Barony
	BCarusus卡洛索斯() feud.Barony
	BSinope锡诺皮() feud.Barony
	BTalaura塔劳拉() feud.Barony
	BThemiscyra忒弥斯库拉() feud.Barony
}

type 锡诺皮SinopeCounty struct {
	feud.BaseCounty
	安德拉帕Andrapa     feud.Barony
	卡洛索斯Carusus     feud.Barony
	锡诺皮Sinope       feud.Barony
	塔劳拉Talaura      feud.Barony
	忒弥斯库拉Themiscyra feud.Barony
}

func (c *锡诺皮SinopeCounty) BAndrapa安德拉帕() feud.Barony {
	return c.安德拉帕Andrapa
}

func (c *锡诺皮SinopeCounty) BCarusus卡洛索斯() feud.Barony {
	return c.卡洛索斯Carusus
}

func (c *锡诺皮SinopeCounty) BSinope锡诺皮() feud.Barony {
	return c.锡诺皮Sinope
}

func (c *锡诺皮SinopeCounty) BTalaura塔劳拉() feud.Barony {
	return c.塔劳拉Talaura
}

func (c *锡诺皮SinopeCounty) BThemiscyra忒弥斯库拉() feud.Barony {
	return c.忒弥斯库拉Themiscyra
}

var CSinope锡诺皮 SinopeCounty = &锡诺皮SinopeCounty{}

func init() {
	f := CSinope锡诺皮.(*锡诺皮SinopeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "739",
		Title:     "sinope",
		TitleName: "锡诺皮",
		TitleCode: "c_sinope",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德拉帕Andrapa = BAndrapa安德拉帕
	f.安德拉帕Andrapa.SetParent(f)

	f.卡洛索斯Carusus = BCarusus卡洛索斯
	f.卡洛索斯Carusus.SetParent(f)

	f.锡诺皮Sinope = BSinope锡诺皮
	f.锡诺皮Sinope.SetParent(f)

	f.塔劳拉Talaura = BTalaura塔劳拉
	f.塔劳拉Talaura.SetParent(f)

	f.忒弥斯库拉Themiscyra = BThemiscyra忒弥斯库拉
	f.忒弥斯库拉Themiscyra.SetParent(f)

}
