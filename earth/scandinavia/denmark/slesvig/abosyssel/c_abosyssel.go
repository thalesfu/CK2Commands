package abosyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbosysselCounty interface {
	feud.County
	BAarhus奥胡斯() feud.Barony
	BHorsens霍尔森斯() feud.Barony
	BJelling耶灵() feud.Barony
	BLisbjerg利斯比约() feud.Barony
	BRanders兰讷斯() feud.Barony
	BSkanderborg斯坎讷堡() feud.Barony
	BVejle瓦埃勒() feud.Barony
}

type 奥胡斯AbosysselCounty struct {
	feud.BaseCounty
	奥胡斯Aarhus       feud.Barony
	霍尔森斯Horsens     feud.Barony
	耶灵Jelling       feud.Barony
	利斯比约Lisbjerg    feud.Barony
	兰讷斯Randers      feud.Barony
	斯坎讷堡Skanderborg feud.Barony
	瓦埃勒Vejle        feud.Barony
}

func (c *奥胡斯AbosysselCounty) BAarhus奥胡斯() feud.Barony {
	return c.奥胡斯Aarhus
}

func (c *奥胡斯AbosysselCounty) BHorsens霍尔森斯() feud.Barony {
	return c.霍尔森斯Horsens
}

func (c *奥胡斯AbosysselCounty) BJelling耶灵() feud.Barony {
	return c.耶灵Jelling
}

func (c *奥胡斯AbosysselCounty) BLisbjerg利斯比约() feud.Barony {
	return c.利斯比约Lisbjerg
}

func (c *奥胡斯AbosysselCounty) BRanders兰讷斯() feud.Barony {
	return c.兰讷斯Randers
}

func (c *奥胡斯AbosysselCounty) BSkanderborg斯坎讷堡() feud.Barony {
	return c.斯坎讷堡Skanderborg
}

func (c *奥胡斯AbosysselCounty) BVejle瓦埃勒() feud.Barony {
	return c.瓦埃勒Vejle
}

var CAbosyssel奥胡斯 AbosysselCounty = &奥胡斯AbosysselCounty{}

func init() {
	f := CAbosyssel奥胡斯.(*奥胡斯AbosysselCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1684",
		Title:     "abosyssel",
		TitleName: "奥胡斯",
		TitleCode: "c_abosyssel",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥胡斯Aarhus = BAarhus奥胡斯
	f.奥胡斯Aarhus.SetParent(f)

	f.霍尔森斯Horsens = BHorsens霍尔森斯
	f.霍尔森斯Horsens.SetParent(f)

	f.耶灵Jelling = BJelling耶灵
	f.耶灵Jelling.SetParent(f)

	f.利斯比约Lisbjerg = BLisbjerg利斯比约
	f.利斯比约Lisbjerg.SetParent(f)

	f.兰讷斯Randers = BRanders兰讷斯
	f.兰讷斯Randers.SetParent(f)

	f.斯坎讷堡Skanderborg = BSkanderborg斯坎讷堡
	f.斯坎讷堡Skanderborg.SetParent(f)

	f.瓦埃勒Vejle = BVejle瓦埃勒
	f.瓦埃勒Vejle.SetParent(f)

}
