package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QuzdarCounty interface {
	feud.County
	BBodein博登() feud.Barony
	BQuzdar古兹达尔() feud.Barony
	BSumandal苏曼多() feud.Barony
	BSupa苏波() feud.Barony
	BSuratgarh苏剌多姞利呬() feud.Barony
	BTalya多俩() feud.Barony
	BTekkumuri嫡鸠牟利() feud.Barony
}

type 古兹达尔QuzdarCounty struct {
	feud.BaseCounty
	博登Bodein        feud.Barony
	古兹达尔Quzdar      feud.Barony
	苏曼多Sumandal     feud.Barony
	苏波Supa          feud.Barony
	苏剌多姞利呬Suratgarh feud.Barony
	多俩Talya         feud.Barony
	嫡鸠牟利Tekkumuri   feud.Barony
}

func (c *古兹达尔QuzdarCounty) BBodein博登() feud.Barony {
	return c.博登Bodein
}

func (c *古兹达尔QuzdarCounty) BQuzdar古兹达尔() feud.Barony {
	return c.古兹达尔Quzdar
}

func (c *古兹达尔QuzdarCounty) BSumandal苏曼多() feud.Barony {
	return c.苏曼多Sumandal
}

func (c *古兹达尔QuzdarCounty) BSupa苏波() feud.Barony {
	return c.苏波Supa
}

func (c *古兹达尔QuzdarCounty) BSuratgarh苏剌多姞利呬() feud.Barony {
	return c.苏剌多姞利呬Suratgarh
}

func (c *古兹达尔QuzdarCounty) BTalya多俩() feud.Barony {
	return c.多俩Talya
}

func (c *古兹达尔QuzdarCounty) BTekkumuri嫡鸠牟利() feud.Barony {
	return c.嫡鸠牟利Tekkumuri
}

var CQuzdar古兹达尔 QuzdarCounty = &古兹达尔QuzdarCounty{}

func init() {
	f := CQuzdar古兹达尔.(*古兹达尔QuzdarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1371",
		Title:     "quzdar",
		TitleName: "古兹达尔",
		TitleCode: "c_quzdar",
		Baronies:  map[string]feud.Barony{},
	}

	f.博登Bodein = BBodein博登
	f.博登Bodein.SetParent(f)

	f.古兹达尔Quzdar = BQuzdar古兹达尔
	f.古兹达尔Quzdar.SetParent(f)

	f.苏曼多Sumandal = BSumandal苏曼多
	f.苏曼多Sumandal.SetParent(f)

	f.苏波Supa = BSupa苏波
	f.苏波Supa.SetParent(f)

	f.苏剌多姞利呬Suratgarh = BSuratgarh苏剌多姞利呬
	f.苏剌多姞利呬Suratgarh.SetParent(f)

	f.多俩Talya = BTalya多俩
	f.多俩Talya.SetParent(f)

	f.嫡鸠牟利Tekkumuri = BTekkumuri嫡鸠牟利
	f.嫡鸠牟利Tekkumuri.SetParent(f)

}
