package yatenga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YatengaCounty interface {
	feud.County
	BDirelabe迪雷拉贝() feud.Barony
	BDubare杜巴雷() feud.Barony
	BGiti吉提() feud.Barony
	BKotola科托拉() feud.Barony
	BWalguyo瓦勒古约() feud.Barony
	BYatenga亚滕加() feud.Barony
	BZondoma宗多马() feud.Barony
}

type 亚滕加YatengaCounty struct {
	feud.BaseCounty
	迪雷拉贝Direlabe feud.Barony
	杜巴雷Dubare    feud.Barony
	吉提Giti       feud.Barony
	科托拉Kotola    feud.Barony
	瓦勒古约Walguyo  feud.Barony
	亚滕加Yatenga   feud.Barony
	宗多马Zondoma   feud.Barony
}

func (c *亚滕加YatengaCounty) BDirelabe迪雷拉贝() feud.Barony {
	return c.迪雷拉贝Direlabe
}

func (c *亚滕加YatengaCounty) BDubare杜巴雷() feud.Barony {
	return c.杜巴雷Dubare
}

func (c *亚滕加YatengaCounty) BGiti吉提() feud.Barony {
	return c.吉提Giti
}

func (c *亚滕加YatengaCounty) BKotola科托拉() feud.Barony {
	return c.科托拉Kotola
}

func (c *亚滕加YatengaCounty) BWalguyo瓦勒古约() feud.Barony {
	return c.瓦勒古约Walguyo
}

func (c *亚滕加YatengaCounty) BYatenga亚滕加() feud.Barony {
	return c.亚滕加Yatenga
}

func (c *亚滕加YatengaCounty) BZondoma宗多马() feud.Barony {
	return c.宗多马Zondoma
}

var CYatenga亚滕加 YatengaCounty = &亚滕加YatengaCounty{}

func init() {
	f := CYatenga亚滕加.(*亚滕加YatengaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1757",
		Title:     "yatenga",
		TitleName: "亚滕加",
		TitleCode: "c_yatenga",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪雷拉贝Direlabe = BDirelabe迪雷拉贝
	f.迪雷拉贝Direlabe.SetParent(f)

	f.杜巴雷Dubare = BDubare杜巴雷
	f.杜巴雷Dubare.SetParent(f)

	f.吉提Giti = BGiti吉提
	f.吉提Giti.SetParent(f)

	f.科托拉Kotola = BKotola科托拉
	f.科托拉Kotola.SetParent(f)

	f.瓦勒古约Walguyo = BWalguyo瓦勒古约
	f.瓦勒古约Walguyo.SetParent(f)

	f.亚滕加Yatenga = BYatenga亚滕加
	f.亚滕加Yatenga.SetParent(f)

	f.宗多马Zondoma = BZondoma宗多马
	f.宗多马Zondoma.SetParent(f)

}
