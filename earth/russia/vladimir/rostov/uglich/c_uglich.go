package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UglichCounty interface {
	feud.County
	BBreytovo布列伊托沃() feud.Barony
	BKoprino科普里诺() feud.Barony
	BMologa莫洛加() feud.Barony
	BUglich乌格利奇() feud.Barony
	BUstyuzhna乌斯秋日纳() feud.Barony
}

type 乌格利奇UglichCounty struct {
	feud.BaseCounty
	布列伊托沃Breytovo  feud.Barony
	科普里诺Koprino    feud.Barony
	莫洛加Mologa      feud.Barony
	乌格利奇Uglich     feud.Barony
	乌斯秋日纳Ustyuzhna feud.Barony
}

func (c *乌格利奇UglichCounty) BBreytovo布列伊托沃() feud.Barony {
	return c.布列伊托沃Breytovo
}

func (c *乌格利奇UglichCounty) BKoprino科普里诺() feud.Barony {
	return c.科普里诺Koprino
}

func (c *乌格利奇UglichCounty) BMologa莫洛加() feud.Barony {
	return c.莫洛加Mologa
}

func (c *乌格利奇UglichCounty) BUglich乌格利奇() feud.Barony {
	return c.乌格利奇Uglich
}

func (c *乌格利奇UglichCounty) BUstyuzhna乌斯秋日纳() feud.Barony {
	return c.乌斯秋日纳Ustyuzhna
}

var CUglich乌格利奇 UglichCounty = &乌格利奇UglichCounty{}

func init() {
	f := CUglich乌格利奇.(*乌格利奇UglichCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "571",
		Title:     "uglich",
		TitleName: "乌格利奇",
		TitleCode: "c_uglich",
		Baronies:  map[string]feud.Barony{},
	}

	f.布列伊托沃Breytovo = BBreytovo布列伊托沃
	f.布列伊托沃Breytovo.SetParent(f)

	f.科普里诺Koprino = BKoprino科普里诺
	f.科普里诺Koprino.SetParent(f)

	f.莫洛加Mologa = BMologa莫洛加
	f.莫洛加Mologa.SetParent(f)

	f.乌格利奇Uglich = BUglich乌格利奇
	f.乌格利奇Uglich.SetParent(f)

	f.乌斯秋日纳Ustyuzhna = BUstyuzhna乌斯秋日纳
	f.乌斯秋日纳Ustyuzhna.SetParent(f)

}
