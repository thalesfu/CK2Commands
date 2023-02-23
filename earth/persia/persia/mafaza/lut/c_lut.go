package lut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LutCounty interface {
	feud.County
	BAmanieh阿玛尼什() feud.Barony
	BAspak阿斯帕克() feud.Barony
	BBibaz比巴茨() feud.Barony
	BDayhouk代胡克() feud.Barony
	BEsfandiar埃斯凡迪亚尔() feud.Barony
	BKalateh卡拉泰() feud.Barony
	BMazanabad马赞阿巴德() feud.Barony
	BTabas塔巴斯() feud.Barony
}

type 玛法扎LutCounty struct {
	feud.BaseCounty
	阿玛尼什Amanieh     feud.Barony
	阿斯帕克Aspak       feud.Barony
	比巴茨Bibaz        feud.Barony
	代胡克Dayhouk      feud.Barony
	埃斯凡迪亚尔Esfandiar feud.Barony
	卡拉泰Kalateh      feud.Barony
	马赞阿巴德Mazanabad  feud.Barony
	塔巴斯Tabas        feud.Barony
}

func (c *玛法扎LutCounty) BAmanieh阿玛尼什() feud.Barony {
	return c.阿玛尼什Amanieh
}

func (c *玛法扎LutCounty) BAspak阿斯帕克() feud.Barony {
	return c.阿斯帕克Aspak
}

func (c *玛法扎LutCounty) BBibaz比巴茨() feud.Barony {
	return c.比巴茨Bibaz
}

func (c *玛法扎LutCounty) BDayhouk代胡克() feud.Barony {
	return c.代胡克Dayhouk
}

func (c *玛法扎LutCounty) BEsfandiar埃斯凡迪亚尔() feud.Barony {
	return c.埃斯凡迪亚尔Esfandiar
}

func (c *玛法扎LutCounty) BKalateh卡拉泰() feud.Barony {
	return c.卡拉泰Kalateh
}

func (c *玛法扎LutCounty) BMazanabad马赞阿巴德() feud.Barony {
	return c.马赞阿巴德Mazanabad
}

func (c *玛法扎LutCounty) BTabas塔巴斯() feud.Barony {
	return c.塔巴斯Tabas
}

var CLut玛法扎 LutCounty = &玛法扎LutCounty{}

func init() {
	f := CLut玛法扎.(*玛法扎LutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "636",
		Title:     "lut",
		TitleName: "玛法扎",
		TitleCode: "c_lut",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿玛尼什Amanieh = BAmanieh阿玛尼什
	f.阿玛尼什Amanieh.SetParent(f)

	f.阿斯帕克Aspak = BAspak阿斯帕克
	f.阿斯帕克Aspak.SetParent(f)

	f.比巴茨Bibaz = BBibaz比巴茨
	f.比巴茨Bibaz.SetParent(f)

	f.代胡克Dayhouk = BDayhouk代胡克
	f.代胡克Dayhouk.SetParent(f)

	f.埃斯凡迪亚尔Esfandiar = BEsfandiar埃斯凡迪亚尔
	f.埃斯凡迪亚尔Esfandiar.SetParent(f)

	f.卡拉泰Kalateh = BKalateh卡拉泰
	f.卡拉泰Kalateh.SetParent(f)

	f.马赞阿巴德Mazanabad = BMazanabad马赞阿巴德
	f.马赞阿巴德Mazanabad.SetParent(f)

	f.塔巴斯Tabas = BTabas塔巴斯
	f.塔巴斯Tabas.SetParent(f)

}
