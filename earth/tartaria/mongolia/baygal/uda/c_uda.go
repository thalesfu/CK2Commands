package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdaCounty interface {
	feud.County
	BBomnak博姆纳克() feud.Barony
	BChumikan丘米坎() feud.Barony
	BTorom托罗姆() feud.Barony
	BUda兀的() feud.Barony
	BUmlekan乌姆列坎() feud.Barony
	BZeya结雅() feud.Barony
	BZeysky结雅斯基() feud.Barony
}

type 兀的UdaCounty struct {
	feud.BaseCounty
	博姆纳克Bomnak  feud.Barony
	丘米坎Chumikan feud.Barony
	托罗姆Torom    feud.Barony
	兀的Uda       feud.Barony
	乌姆列坎Umlekan feud.Barony
	结雅Zeya      feud.Barony
	结雅斯基Zeysky  feud.Barony
}

func (c *兀的UdaCounty) BBomnak博姆纳克() feud.Barony {
	return c.博姆纳克Bomnak
}

func (c *兀的UdaCounty) BChumikan丘米坎() feud.Barony {
	return c.丘米坎Chumikan
}

func (c *兀的UdaCounty) BTorom托罗姆() feud.Barony {
	return c.托罗姆Torom
}

func (c *兀的UdaCounty) BUda兀的() feud.Barony {
	return c.兀的Uda
}

func (c *兀的UdaCounty) BUmlekan乌姆列坎() feud.Barony {
	return c.乌姆列坎Umlekan
}

func (c *兀的UdaCounty) BZeya结雅() feud.Barony {
	return c.结雅Zeya
}

func (c *兀的UdaCounty) BZeysky结雅斯基() feud.Barony {
	return c.结雅斯基Zeysky
}

var CUda兀的 UdaCounty = &兀的UdaCounty{}

func init() {
	f := CUda兀的.(*兀的UdaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1922",
		Title:     "uda",
		TitleName: "兀的",
		TitleCode: "c_uda",
		Baronies:  map[string]feud.Barony{},
	}

	f.博姆纳克Bomnak = BBomnak博姆纳克
	f.博姆纳克Bomnak.SetParent(f)

	f.丘米坎Chumikan = BChumikan丘米坎
	f.丘米坎Chumikan.SetParent(f)

	f.托罗姆Torom = BTorom托罗姆
	f.托罗姆Torom.SetParent(f)

	f.兀的Uda = BUda兀的
	f.兀的Uda.SetParent(f)

	f.乌姆列坎Umlekan = BUmlekan乌姆列坎
	f.乌姆列坎Umlekan.SetParent(f)

	f.结雅Zeya = BZeya结雅
	f.结雅Zeya.SetParent(f)

	f.结雅斯基Zeysky = BZeysky结雅斯基
	f.结雅斯基Zeysky.SetParent(f)

}
