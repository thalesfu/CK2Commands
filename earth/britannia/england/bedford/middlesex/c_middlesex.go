package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MiddlesexCounty interface {
	feud.County
	BChelsea切尔西() feud.Barony
	BFulham富勒姆() feud.Barony
	BHarrow哈罗() feud.Barony
	BLondon伦敦() feud.Barony
	BStaines斯泰恩斯() feud.Barony
	BTottenham托特纳姆() feud.Barony
	BWestminster威斯敏斯特() feud.Barony
}

type 米德尔塞克斯MiddlesexCounty struct {
	feud.BaseCounty
	切尔西Chelsea       feud.Barony
	富勒姆Fulham        feud.Barony
	哈罗Harrow         feud.Barony
	伦敦London         feud.Barony
	斯泰恩斯Staines      feud.Barony
	托特纳姆Tottenham    feud.Barony
	威斯敏斯特Westminster feud.Barony
}

func (c *米德尔塞克斯MiddlesexCounty) BChelsea切尔西() feud.Barony {
	return c.切尔西Chelsea
}

func (c *米德尔塞克斯MiddlesexCounty) BFulham富勒姆() feud.Barony {
	return c.富勒姆Fulham
}

func (c *米德尔塞克斯MiddlesexCounty) BHarrow哈罗() feud.Barony {
	return c.哈罗Harrow
}

func (c *米德尔塞克斯MiddlesexCounty) BLondon伦敦() feud.Barony {
	return c.伦敦London
}

func (c *米德尔塞克斯MiddlesexCounty) BStaines斯泰恩斯() feud.Barony {
	return c.斯泰恩斯Staines
}

func (c *米德尔塞克斯MiddlesexCounty) BTottenham托特纳姆() feud.Barony {
	return c.托特纳姆Tottenham
}

func (c *米德尔塞克斯MiddlesexCounty) BWestminster威斯敏斯特() feud.Barony {
	return c.威斯敏斯特Westminster
}

var CMiddlesex米德尔塞克斯 MiddlesexCounty = &米德尔塞克斯MiddlesexCounty{}

func init() {
	f := CMiddlesex米德尔塞克斯.(*米德尔塞克斯MiddlesexCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "32",
		Title:     "middlesex",
		TitleName: "米德尔塞克斯",
		TitleCode: "c_middlesex",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔西Chelsea = BChelsea切尔西
	f.切尔西Chelsea.SetParent(f)

	f.富勒姆Fulham = BFulham富勒姆
	f.富勒姆Fulham.SetParent(f)

	f.哈罗Harrow = BHarrow哈罗
	f.哈罗Harrow.SetParent(f)

	f.伦敦London = BLondon伦敦
	f.伦敦London.SetParent(f)

	f.斯泰恩斯Staines = BStaines斯泰恩斯
	f.斯泰恩斯Staines.SetParent(f)

	f.托特纳姆Tottenham = BTottenham托特纳姆
	f.托特纳姆Tottenham.SetParent(f)

	f.威斯敏斯特Westminster = BWestminster威斯敏斯特
	f.威斯敏斯特Westminster.SetParent(f)

}
