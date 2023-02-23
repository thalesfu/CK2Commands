package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HarerCounty interface {
	feud.County
	BAlemaya阿勒马亚() feud.Barony
	BBabile巴比莱() feud.Barony
	BDakkar达卡() feud.Barony
	BDiridawa德雷达瓦() feud.Barony
	BHarar哈勒尔() feud.Barony
	BKombolcha孔博勒查() feud.Barony
	BKurfachele库发彻拉() feud.Barony
}

type 哈勒尔HarerCounty struct {
	feud.BaseCounty
	阿勒马亚Alemaya    feud.Barony
	巴比莱Babile      feud.Barony
	达卡Dakkar       feud.Barony
	德雷达瓦Diridawa   feud.Barony
	哈勒尔Harar       feud.Barony
	孔博勒查Kombolcha  feud.Barony
	库发彻拉Kurfachele feud.Barony
}

func (c *哈勒尔HarerCounty) BAlemaya阿勒马亚() feud.Barony {
	return c.阿勒马亚Alemaya
}

func (c *哈勒尔HarerCounty) BBabile巴比莱() feud.Barony {
	return c.巴比莱Babile
}

func (c *哈勒尔HarerCounty) BDakkar达卡() feud.Barony {
	return c.达卡Dakkar
}

func (c *哈勒尔HarerCounty) BDiridawa德雷达瓦() feud.Barony {
	return c.德雷达瓦Diridawa
}

func (c *哈勒尔HarerCounty) BHarar哈勒尔() feud.Barony {
	return c.哈勒尔Harar
}

func (c *哈勒尔HarerCounty) BKombolcha孔博勒查() feud.Barony {
	return c.孔博勒查Kombolcha
}

func (c *哈勒尔HarerCounty) BKurfachele库发彻拉() feud.Barony {
	return c.库发彻拉Kurfachele
}

var CHarer哈勒尔 HarerCounty = &哈勒尔HarerCounty{}

func init() {
	f := CHarer哈勒尔.(*哈勒尔HarerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "873",
		Title:     "harer",
		TitleName: "哈勒尔",
		TitleCode: "c_harer",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿勒马亚Alemaya = BAlemaya阿勒马亚
	f.阿勒马亚Alemaya.SetParent(f)

	f.巴比莱Babile = BBabile巴比莱
	f.巴比莱Babile.SetParent(f)

	f.达卡Dakkar = BDakkar达卡
	f.达卡Dakkar.SetParent(f)

	f.德雷达瓦Diridawa = BDiridawa德雷达瓦
	f.德雷达瓦Diridawa.SetParent(f)

	f.哈勒尔Harar = BHarar哈勒尔
	f.哈勒尔Harar.SetParent(f)

	f.孔博勒查Kombolcha = BKombolcha孔博勒查
	f.孔博勒查Kombolcha.SetParent(f)

	f.库发彻拉Kurfachele = BKurfachele库发彻拉
	f.库发彻拉Kurfachele.SetParent(f)

}
