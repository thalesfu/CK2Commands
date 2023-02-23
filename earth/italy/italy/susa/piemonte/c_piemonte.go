package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PiemonteCounty interface {
	feud.County
	BAuriate奥里亚泰() feud.Barony
	BCrevacuore克雷瓦科雷() feud.Barony
	BFerrero费雷罗() feud.Barony
	BMessarano梅萨拉诺() feud.Barony
	BNovara诺瓦拉() feud.Barony
	BSettimo塞蒂莫() feud.Barony
	BTorino都灵() feud.Barony
}

type 皮埃蒙特PiemonteCounty struct {
	feud.BaseCounty
	奥里亚泰Auriate     feud.Barony
	克雷瓦科雷Crevacuore feud.Barony
	费雷罗Ferrero      feud.Barony
	梅萨拉诺Messarano   feud.Barony
	诺瓦拉Novara       feud.Barony
	塞蒂莫Settimo      feud.Barony
	都灵Torino        feud.Barony
}

func (c *皮埃蒙特PiemonteCounty) BAuriate奥里亚泰() feud.Barony {
	return c.奥里亚泰Auriate
}

func (c *皮埃蒙特PiemonteCounty) BCrevacuore克雷瓦科雷() feud.Barony {
	return c.克雷瓦科雷Crevacuore
}

func (c *皮埃蒙特PiemonteCounty) BFerrero费雷罗() feud.Barony {
	return c.费雷罗Ferrero
}

func (c *皮埃蒙特PiemonteCounty) BMessarano梅萨拉诺() feud.Barony {
	return c.梅萨拉诺Messarano
}

func (c *皮埃蒙特PiemonteCounty) BNovara诺瓦拉() feud.Barony {
	return c.诺瓦拉Novara
}

func (c *皮埃蒙特PiemonteCounty) BSettimo塞蒂莫() feud.Barony {
	return c.塞蒂莫Settimo
}

func (c *皮埃蒙特PiemonteCounty) BTorino都灵() feud.Barony {
	return c.都灵Torino
}

var CPiemonte皮埃蒙特 PiemonteCounty = &皮埃蒙特PiemonteCounty{}

func init() {
	f := CPiemonte皮埃蒙特.(*皮埃蒙特PiemonteCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "236",
		Title:     "piemonte",
		TitleName: "皮埃蒙特",
		TitleCode: "c_piemonte",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥里亚泰Auriate = BAuriate奥里亚泰
	f.奥里亚泰Auriate.SetParent(f)

	f.克雷瓦科雷Crevacuore = BCrevacuore克雷瓦科雷
	f.克雷瓦科雷Crevacuore.SetParent(f)

	f.费雷罗Ferrero = BFerrero费雷罗
	f.费雷罗Ferrero.SetParent(f)

	f.梅萨拉诺Messarano = BMessarano梅萨拉诺
	f.梅萨拉诺Messarano.SetParent(f)

	f.诺瓦拉Novara = BNovara诺瓦拉
	f.诺瓦拉Novara.SetParent(f)

	f.塞蒂莫Settimo = BSettimo塞蒂莫
	f.塞蒂莫Settimo.SetParent(f)

	f.都灵Torino = BTorino都灵
	f.都灵Torino.SetParent(f)

}
