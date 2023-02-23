package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SeletytenizCounty interface {
	feud.County
	BBabatay巴巴泰() feud.Barony
	BBereke别列克() feud.Barony
	BKoschi科斯希() feud.Barony
	BKoyandy科扬德() feud.Barony
	BSaryoba萨雷奥巴() feud.Barony
	BSeletyteniz谢列特捷尼兹() feud.Barony
}

type 谢列特捷尼兹SeletytenizCounty struct {
	feud.BaseCounty
	巴巴泰Babatay        feud.Barony
	别列克Bereke         feud.Barony
	科斯希Koschi         feud.Barony
	科扬德Koyandy        feud.Barony
	萨雷奥巴Saryoba       feud.Barony
	谢列特捷尼兹Seletyteniz feud.Barony
}

func (c *谢列特捷尼兹SeletytenizCounty) BBabatay巴巴泰() feud.Barony {
	return c.巴巴泰Babatay
}

func (c *谢列特捷尼兹SeletytenizCounty) BBereke别列克() feud.Barony {
	return c.别列克Bereke
}

func (c *谢列特捷尼兹SeletytenizCounty) BKoschi科斯希() feud.Barony {
	return c.科斯希Koschi
}

func (c *谢列特捷尼兹SeletytenizCounty) BKoyandy科扬德() feud.Barony {
	return c.科扬德Koyandy
}

func (c *谢列特捷尼兹SeletytenizCounty) BSaryoba萨雷奥巴() feud.Barony {
	return c.萨雷奥巴Saryoba
}

func (c *谢列特捷尼兹SeletytenizCounty) BSeletyteniz谢列特捷尼兹() feud.Barony {
	return c.谢列特捷尼兹Seletyteniz
}

var CSeletyteniz谢列特捷尼兹 SeletytenizCounty = &谢列特捷尼兹SeletytenizCounty{}

func init() {
	f := CSeletyteniz谢列特捷尼兹.(*谢列特捷尼兹SeletytenizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1878",
		Title:     "seletyteniz",
		TitleName: "谢列特捷尼兹",
		TitleCode: "c_seletyteniz",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴巴泰Babatay = BBabatay巴巴泰
	f.巴巴泰Babatay.SetParent(f)

	f.别列克Bereke = BBereke别列克
	f.别列克Bereke.SetParent(f)

	f.科斯希Koschi = BKoschi科斯希
	f.科斯希Koschi.SetParent(f)

	f.科扬德Koyandy = BKoyandy科扬德
	f.科扬德Koyandy.SetParent(f)

	f.萨雷奥巴Saryoba = BSaryoba萨雷奥巴
	f.萨雷奥巴Saryoba.SetParent(f)

	f.谢列特捷尼兹Seletyteniz = BSeletyteniz谢列特捷尼兹
	f.谢列特捷尼兹Seletyteniz.SetParent(f)

}
