package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DithmarschenCounty interface {
	feud.County
	BBusum比苏姆() feud.Barony
	BElmshorn埃尔姆斯霍恩() feud.Barony
	BHeide海德() feud.Barony
	BItzehoe伊策霍() feud.Barony
	BMeldorf梅尔多夫() feud.Barony
	BUlstrup乌尔斯楚普() feud.Barony
	BWacken瓦肯() feud.Barony
}

type 迪特马申DithmarschenCounty struct {
	feud.BaseCounty
	比苏姆Busum       feud.Barony
	埃尔姆斯霍恩Elmshorn feud.Barony
	海德Heide        feud.Barony
	伊策霍Itzehoe     feud.Barony
	梅尔多夫Meldorf    feud.Barony
	乌尔斯楚普Ulstrup   feud.Barony
	瓦肯Wacken       feud.Barony
}

func (c *迪特马申DithmarschenCounty) BBusum比苏姆() feud.Barony {
	return c.比苏姆Busum
}

func (c *迪特马申DithmarschenCounty) BElmshorn埃尔姆斯霍恩() feud.Barony {
	return c.埃尔姆斯霍恩Elmshorn
}

func (c *迪特马申DithmarschenCounty) BHeide海德() feud.Barony {
	return c.海德Heide
}

func (c *迪特马申DithmarschenCounty) BItzehoe伊策霍() feud.Barony {
	return c.伊策霍Itzehoe
}

func (c *迪特马申DithmarschenCounty) BMeldorf梅尔多夫() feud.Barony {
	return c.梅尔多夫Meldorf
}

func (c *迪特马申DithmarschenCounty) BUlstrup乌尔斯楚普() feud.Barony {
	return c.乌尔斯楚普Ulstrup
}

func (c *迪特马申DithmarschenCounty) BWacken瓦肯() feud.Barony {
	return c.瓦肯Wacken
}

var CDithmarschen迪特马申 DithmarschenCounty = &迪特马申DithmarschenCounty{}

func init() {
	f := CDithmarschen迪特马申.(*迪特马申DithmarschenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1688",
		Title:     "dithmarschen",
		TitleName: "迪特马申",
		TitleCode: "c_dithmarschen",
		Baronies:  map[string]feud.Barony{},
	}

	f.比苏姆Busum = BBusum比苏姆
	f.比苏姆Busum.SetParent(f)

	f.埃尔姆斯霍恩Elmshorn = BElmshorn埃尔姆斯霍恩
	f.埃尔姆斯霍恩Elmshorn.SetParent(f)

	f.海德Heide = BHeide海德
	f.海德Heide.SetParent(f)

	f.伊策霍Itzehoe = BItzehoe伊策霍
	f.伊策霍Itzehoe.SetParent(f)

	f.梅尔多夫Meldorf = BMeldorf梅尔多夫
	f.梅尔多夫Meldorf.SetParent(f)

	f.乌尔斯楚普Ulstrup = BUlstrup乌尔斯楚普
	f.乌尔斯楚普Ulstrup.SetParent(f)

	f.瓦肯Wacken = BWacken瓦肯
	f.瓦肯Wacken.SetParent(f)

}
