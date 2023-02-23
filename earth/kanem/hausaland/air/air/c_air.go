package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AirCounty interface {
	feud.County
	BAgades阿加德斯() feud.Barony
	BGabasour加巴苏尔() feud.Barony
	BKoufe库费() feud.Barony
	BNgargue恩加尔盖() feud.Barony
	BTakedda塔凯达() feud.Barony
	BTchalta查尔塔() feud.Barony
	BToumboula图姆布拉() feud.Barony
}

type 艾尔AirCounty struct {
	feud.BaseCounty
	阿加德斯Agades    feud.Barony
	加巴苏尔Gabasour  feud.Barony
	库费Koufe       feud.Barony
	恩加尔盖Ngargue   feud.Barony
	塔凯达Takedda    feud.Barony
	查尔塔Tchalta    feud.Barony
	图姆布拉Toumboula feud.Barony
}

func (c *艾尔AirCounty) BAgades阿加德斯() feud.Barony {
	return c.阿加德斯Agades
}

func (c *艾尔AirCounty) BGabasour加巴苏尔() feud.Barony {
	return c.加巴苏尔Gabasour
}

func (c *艾尔AirCounty) BKoufe库费() feud.Barony {
	return c.库费Koufe
}

func (c *艾尔AirCounty) BNgargue恩加尔盖() feud.Barony {
	return c.恩加尔盖Ngargue
}

func (c *艾尔AirCounty) BTakedda塔凯达() feud.Barony {
	return c.塔凯达Takedda
}

func (c *艾尔AirCounty) BTchalta查尔塔() feud.Barony {
	return c.查尔塔Tchalta
}

func (c *艾尔AirCounty) BToumboula图姆布拉() feud.Barony {
	return c.图姆布拉Toumboula
}

var CAir艾尔 AirCounty = &艾尔AirCounty{}

func init() {
	f := CAir艾尔.(*艾尔AirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1738",
		Title:     "air",
		TitleName: "艾尔",
		TitleCode: "c_air",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加德斯Agades = BAgades阿加德斯
	f.阿加德斯Agades.SetParent(f)

	f.加巴苏尔Gabasour = BGabasour加巴苏尔
	f.加巴苏尔Gabasour.SetParent(f)

	f.库费Koufe = BKoufe库费
	f.库费Koufe.SetParent(f)

	f.恩加尔盖Ngargue = BNgargue恩加尔盖
	f.恩加尔盖Ngargue.SetParent(f)

	f.塔凯达Takedda = BTakedda塔凯达
	f.塔凯达Takedda.SetParent(f)

	f.查尔塔Tchalta = BTchalta查尔塔
	f.查尔塔Tchalta.SetParent(f)

	f.图姆布拉Toumboula = BToumboula图姆布拉
	f.图姆布拉Toumboula.SetParent(f)

}
