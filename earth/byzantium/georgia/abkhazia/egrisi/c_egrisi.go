package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EgrisiCounty interface {
	feud.County
	BBedia贝迪亚() feud.Barony
	BBuchlus布赫卢斯() feud.Barony
	BEgrisi埃格里西() feud.Barony
	BKhobi霍比() feud.Barony
	BMokvi莫克维() feud.Barony
	BNesos内索斯() feud.Barony
	BPhasis法锡斯() feud.Barony
}

type 埃格里西EgrisiCounty struct {
	feud.BaseCounty
	贝迪亚Bedia    feud.Barony
	布赫卢斯Buchlus feud.Barony
	埃格里西Egrisi  feud.Barony
	霍比Khobi     feud.Barony
	莫克维Mokvi    feud.Barony
	内索斯Nesos    feud.Barony
	法锡斯Phasis   feud.Barony
}

func (c *埃格里西EgrisiCounty) BBedia贝迪亚() feud.Barony {
	return c.贝迪亚Bedia
}

func (c *埃格里西EgrisiCounty) BBuchlus布赫卢斯() feud.Barony {
	return c.布赫卢斯Buchlus
}

func (c *埃格里西EgrisiCounty) BEgrisi埃格里西() feud.Barony {
	return c.埃格里西Egrisi
}

func (c *埃格里西EgrisiCounty) BKhobi霍比() feud.Barony {
	return c.霍比Khobi
}

func (c *埃格里西EgrisiCounty) BMokvi莫克维() feud.Barony {
	return c.莫克维Mokvi
}

func (c *埃格里西EgrisiCounty) BNesos内索斯() feud.Barony {
	return c.内索斯Nesos
}

func (c *埃格里西EgrisiCounty) BPhasis法锡斯() feud.Barony {
	return c.法锡斯Phasis
}

var CEgrisi埃格里西 EgrisiCounty = &埃格里西EgrisiCounty{}

func init() {
	f := CEgrisi埃格里西.(*埃格里西EgrisiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1806",
		Title:     "egrisi",
		TitleName: "埃格里西",
		TitleCode: "c_egrisi",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝迪亚Bedia = BBedia贝迪亚
	f.贝迪亚Bedia.SetParent(f)

	f.布赫卢斯Buchlus = BBuchlus布赫卢斯
	f.布赫卢斯Buchlus.SetParent(f)

	f.埃格里西Egrisi = BEgrisi埃格里西
	f.埃格里西Egrisi.SetParent(f)

	f.霍比Khobi = BKhobi霍比
	f.霍比Khobi.SetParent(f)

	f.莫克维Mokvi = BMokvi莫克维
	f.莫克维Mokvi.SetParent(f)

	f.内索斯Nesos = BNesos内索斯
	f.内索斯Nesos.SetParent(f)

	f.法锡斯Phasis = BPhasis法锡斯
	f.法锡斯Phasis.SetParent(f)

}
