package maine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaineCounty interface {
	feud.County
	BBeaumont博蒙() feud.Barony
	BCraon克朗() feud.Barony
	BEvron埃夫龙() feud.Barony
	BLaval拉瓦勒() feud.Barony
	BMayenne马耶讷() feud.Barony
	BMontenay蒙特奈() feud.Barony
	BSable萨布莱() feud.Barony
}

type 曼恩MaineCounty struct {
	feud.BaseCounty
	博蒙Beaumont  feud.Barony
	克朗Craon     feud.Barony
	埃夫龙Evron    feud.Barony
	拉瓦勒Laval    feud.Barony
	马耶讷Mayenne  feud.Barony
	蒙特奈Montenay feud.Barony
	萨布莱Sable    feud.Barony
}

func (c *曼恩MaineCounty) BBeaumont博蒙() feud.Barony {
	return c.博蒙Beaumont
}

func (c *曼恩MaineCounty) BCraon克朗() feud.Barony {
	return c.克朗Craon
}

func (c *曼恩MaineCounty) BEvron埃夫龙() feud.Barony {
	return c.埃夫龙Evron
}

func (c *曼恩MaineCounty) BLaval拉瓦勒() feud.Barony {
	return c.拉瓦勒Laval
}

func (c *曼恩MaineCounty) BMayenne马耶讷() feud.Barony {
	return c.马耶讷Mayenne
}

func (c *曼恩MaineCounty) BMontenay蒙特奈() feud.Barony {
	return c.蒙特奈Montenay
}

func (c *曼恩MaineCounty) BSable萨布莱() feud.Barony {
	return c.萨布莱Sable
}

var CMaine曼恩 MaineCounty = &曼恩MaineCounty{}

func init() {
	f := CMaine曼恩.(*曼恩MaineCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "108",
		Title:     "maine",
		TitleName: "曼恩",
		TitleCode: "c_maine",
		Baronies:  map[string]feud.Barony{},
	}

	f.博蒙Beaumont = BBeaumont博蒙
	f.博蒙Beaumont.SetParent(f)

	f.克朗Craon = BCraon克朗
	f.克朗Craon.SetParent(f)

	f.埃夫龙Evron = BEvron埃夫龙
	f.埃夫龙Evron.SetParent(f)

	f.拉瓦勒Laval = BLaval拉瓦勒
	f.拉瓦勒Laval.SetParent(f)

	f.马耶讷Mayenne = BMayenne马耶讷
	f.马耶讷Mayenne.SetParent(f)

	f.蒙特奈Montenay = BMontenay蒙特奈
	f.蒙特奈Montenay.SetParent(f)

	f.萨布莱Sable = BSable萨布莱
	f.萨布莱Sable.SetParent(f)

}
