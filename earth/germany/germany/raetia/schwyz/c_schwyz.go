package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SchwyzCounty interface {
	feud.County
	BAltdorf阿尔特多夫() feud.Barony
	BDisentis迪森蒂斯() feud.Barony
	BEngelberg恩格尔贝格() feud.Barony
	BGlarus格拉鲁斯() feud.Barony
	BKussnacht屈斯纳赫特() feud.Barony
	BSchanis谢尼斯() feud.Barony
	BSchwyz施维茨() feud.Barony
}

type 施维茨SchwyzCounty struct {
	feud.BaseCounty
	阿尔特多夫Altdorf   feud.Barony
	迪森蒂斯Disentis   feud.Barony
	恩格尔贝格Engelberg feud.Barony
	格拉鲁斯Glarus     feud.Barony
	屈斯纳赫特Kussnacht feud.Barony
	谢尼斯Schanis     feud.Barony
	施维茨Schwyz      feud.Barony
}

func (c *施维茨SchwyzCounty) BAltdorf阿尔特多夫() feud.Barony {
	return c.阿尔特多夫Altdorf
}

func (c *施维茨SchwyzCounty) BDisentis迪森蒂斯() feud.Barony {
	return c.迪森蒂斯Disentis
}

func (c *施维茨SchwyzCounty) BEngelberg恩格尔贝格() feud.Barony {
	return c.恩格尔贝格Engelberg
}

func (c *施维茨SchwyzCounty) BGlarus格拉鲁斯() feud.Barony {
	return c.格拉鲁斯Glarus
}

func (c *施维茨SchwyzCounty) BKussnacht屈斯纳赫特() feud.Barony {
	return c.屈斯纳赫特Kussnacht
}

func (c *施维茨SchwyzCounty) BSchanis谢尼斯() feud.Barony {
	return c.谢尼斯Schanis
}

func (c *施维茨SchwyzCounty) BSchwyz施维茨() feud.Barony {
	return c.施维茨Schwyz
}

var CSchwyz施维茨 SchwyzCounty = &施维茨SchwyzCounty{}

func init() {
	f := CSchwyz施维茨.(*施维茨SchwyzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "246",
		Title:     "schwyz",
		TitleName: "施维茨",
		TitleCode: "c_schwyz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔特多夫Altdorf = BAltdorf阿尔特多夫
	f.阿尔特多夫Altdorf.SetParent(f)

	f.迪森蒂斯Disentis = BDisentis迪森蒂斯
	f.迪森蒂斯Disentis.SetParent(f)

	f.恩格尔贝格Engelberg = BEngelberg恩格尔贝格
	f.恩格尔贝格Engelberg.SetParent(f)

	f.格拉鲁斯Glarus = BGlarus格拉鲁斯
	f.格拉鲁斯Glarus.SetParent(f)

	f.屈斯纳赫特Kussnacht = BKussnacht屈斯纳赫特
	f.屈斯纳赫特Kussnacht.SetParent(f)

	f.谢尼斯Schanis = BSchanis谢尼斯
	f.谢尼斯Schanis.SetParent(f)

	f.施维茨Schwyz = BSchwyz施维茨
	f.施维茨Schwyz.SetParent(f)

}
