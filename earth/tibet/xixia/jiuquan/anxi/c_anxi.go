package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnxiCounty interface {
	feud.County
	BAnxi安西() feud.Barony
	BGuazhou瓜州() feud.Barony
	BRenhou仁厚() feud.Barony
	BSuoyang锁阳() feud.Barony
	BXincheng新城() feud.Barony
	BYuanquan渊泉() feud.Barony
	BYulin榆林() feud.Barony
}

type 安西AnxiCounty struct {
	feud.BaseCounty
	安西Anxi     feud.Barony
	瓜州Guazhou  feud.Barony
	仁厚Renhou   feud.Barony
	锁阳Suoyang  feud.Barony
	新城Xincheng feud.Barony
	渊泉Yuanquan feud.Barony
	榆林Yulin    feud.Barony
}

func (c *安西AnxiCounty) BAnxi安西() feud.Barony {
	return c.安西Anxi
}

func (c *安西AnxiCounty) BGuazhou瓜州() feud.Barony {
	return c.瓜州Guazhou
}

func (c *安西AnxiCounty) BRenhou仁厚() feud.Barony {
	return c.仁厚Renhou
}

func (c *安西AnxiCounty) BSuoyang锁阳() feud.Barony {
	return c.锁阳Suoyang
}

func (c *安西AnxiCounty) BXincheng新城() feud.Barony {
	return c.新城Xincheng
}

func (c *安西AnxiCounty) BYuanquan渊泉() feud.Barony {
	return c.渊泉Yuanquan
}

func (c *安西AnxiCounty) BYulin榆林() feud.Barony {
	return c.榆林Yulin
}

var CAnxi安西 AnxiCounty = &安西AnxiCounty{}

func init() {
	f := CAnxi安西.(*安西AnxiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1408",
		Title:     "anxi",
		TitleName: "安西",
		TitleCode: "c_anxi",
		Baronies:  map[string]feud.Barony{},
	}

	f.安西Anxi = BAnxi安西
	f.安西Anxi.SetParent(f)

	f.瓜州Guazhou = BGuazhou瓜州
	f.瓜州Guazhou.SetParent(f)

	f.仁厚Renhou = BRenhou仁厚
	f.仁厚Renhou.SetParent(f)

	f.锁阳Suoyang = BSuoyang锁阳
	f.锁阳Suoyang.SetParent(f)

	f.新城Xincheng = BXincheng新城
	f.新城Xincheng.SetParent(f)

	f.渊泉Yuanquan = BYuanquan渊泉
	f.渊泉Yuanquan.SetParent(f)

	f.榆林Yulin = BYulin榆林
	f.榆林Yulin.SetParent(f)

}
