package songhay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SonghayCounty interface {
	feud.County
	BAnsongo昂松戈() feud.Barony
	BBoumoundo布蒙多() feud.Barony
	BDomossela多莫塞拉() feud.Barony
	BFalakoro法拉科罗() feud.Barony
	BKukiya库基亚() feud.Barony
	BTiende蒂安德() feud.Barony
	BTillaberi蒂拉贝里() feud.Barony
}

type 桑海SonghayCounty struct {
	feud.BaseCounty
	昂松戈Ansongo    feud.Barony
	布蒙多Boumoundo  feud.Barony
	多莫塞拉Domossela feud.Barony
	法拉科罗Falakoro  feud.Barony
	库基亚Kukiya     feud.Barony
	蒂安德Tiende     feud.Barony
	蒂拉贝里Tillaberi feud.Barony
}

func (c *桑海SonghayCounty) BAnsongo昂松戈() feud.Barony {
	return c.昂松戈Ansongo
}

func (c *桑海SonghayCounty) BBoumoundo布蒙多() feud.Barony {
	return c.布蒙多Boumoundo
}

func (c *桑海SonghayCounty) BDomossela多莫塞拉() feud.Barony {
	return c.多莫塞拉Domossela
}

func (c *桑海SonghayCounty) BFalakoro法拉科罗() feud.Barony {
	return c.法拉科罗Falakoro
}

func (c *桑海SonghayCounty) BKukiya库基亚() feud.Barony {
	return c.库基亚Kukiya
}

func (c *桑海SonghayCounty) BTiende蒂安德() feud.Barony {
	return c.蒂安德Tiende
}

func (c *桑海SonghayCounty) BTillaberi蒂拉贝里() feud.Barony {
	return c.蒂拉贝里Tillaberi
}

var CSonghay桑海 SonghayCounty = &桑海SonghayCounty{}

func init() {
	f := CSonghay桑海.(*桑海SonghayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1754",
		Title:     "songhay",
		TitleName: "桑海",
		TitleCode: "c_songhay",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂松戈Ansongo = BAnsongo昂松戈
	f.昂松戈Ansongo.SetParent(f)

	f.布蒙多Boumoundo = BBoumoundo布蒙多
	f.布蒙多Boumoundo.SetParent(f)

	f.多莫塞拉Domossela = BDomossela多莫塞拉
	f.多莫塞拉Domossela.SetParent(f)

	f.法拉科罗Falakoro = BFalakoro法拉科罗
	f.法拉科罗Falakoro.SetParent(f)

	f.库基亚Kukiya = BKukiya库基亚
	f.库基亚Kukiya.SetParent(f)

	f.蒂安德Tiende = BTiende蒂安德
	f.蒂安德Tiende.SetParent(f)

	f.蒂拉贝里Tillaberi = BTillaberi蒂拉贝里
	f.蒂拉贝里Tillaberi.SetParent(f)

}
