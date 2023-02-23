package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IdjilCounty interface {
	feud.County
	BFderick弗德里克() feud.Barony
	BKrossa克罗萨() feud.Barony
	BMosa莫萨() feud.Barony
	BRouessa鲁韦撒() feud.Barony
	BTazadit塔扎尔迪特() feud.Barony
	BZouerat祖韦拉特() feud.Barony
}

type 比尔乌姆格兰IdjilCounty struct {
	feud.BaseCounty
	弗德里克Fderick  feud.Barony
	克罗萨Krossa    feud.Barony
	莫萨Mosa       feud.Barony
	鲁韦撒Rouessa   feud.Barony
	塔扎尔迪特Tazadit feud.Barony
	祖韦拉特Zouerat  feud.Barony
}

func (c *比尔乌姆格兰IdjilCounty) BFderick弗德里克() feud.Barony {
	return c.弗德里克Fderick
}

func (c *比尔乌姆格兰IdjilCounty) BKrossa克罗萨() feud.Barony {
	return c.克罗萨Krossa
}

func (c *比尔乌姆格兰IdjilCounty) BMosa莫萨() feud.Barony {
	return c.莫萨Mosa
}

func (c *比尔乌姆格兰IdjilCounty) BRouessa鲁韦撒() feud.Barony {
	return c.鲁韦撒Rouessa
}

func (c *比尔乌姆格兰IdjilCounty) BTazadit塔扎尔迪特() feud.Barony {
	return c.塔扎尔迪特Tazadit
}

func (c *比尔乌姆格兰IdjilCounty) BZouerat祖韦拉特() feud.Barony {
	return c.祖韦拉特Zouerat
}

var CIdjil比尔乌姆格兰 IdjilCounty = &比尔乌姆格兰IdjilCounty{}

func init() {
	f := CIdjil比尔乌姆格兰.(*比尔乌姆格兰IdjilCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "921",
		Title:     "idjil",
		TitleName: "比尔乌姆格兰",
		TitleCode: "c_idjil",
		Baronies:  map[string]feud.Barony{},
	}

	f.弗德里克Fderick = BFderick弗德里克
	f.弗德里克Fderick.SetParent(f)

	f.克罗萨Krossa = BKrossa克罗萨
	f.克罗萨Krossa.SetParent(f)

	f.莫萨Mosa = BMosa莫萨
	f.莫萨Mosa.SetParent(f)

	f.鲁韦撒Rouessa = BRouessa鲁韦撒
	f.鲁韦撒Rouessa.SetParent(f)

	f.塔扎尔迪特Tazadit = BTazadit塔扎尔迪特
	f.塔扎尔迪特Tazadit.SetParent(f)

	f.祖韦拉特Zouerat = BZouerat祖韦拉特
	f.祖韦拉特Zouerat.SetParent(f)

}
