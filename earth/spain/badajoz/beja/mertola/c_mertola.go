package mertola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MertolaCounty interface {
    feud.County
    BBeja贝雅() 	feud.Barony
    BMertola梅尔图拉() 	feud.Barony
    BMonsaraz蒙萨拉什() 	feud.Barony
    BMoura莫拉() 	feud.Barony
    BMourao莫朗() 	feud.Barony
    BNoudal诺达尔() 	feud.Barony
    BPortel波特尔() 	feud.Barony
    BSerpa塞尔帕() 	feud.Barony
}

type 梅尔图拉MertolaCounty struct {
	feud.BaseCounty
	贝雅Beja 	feud.Barony
	梅尔图拉Mertola 	feud.Barony
	蒙萨拉什Monsaraz 	feud.Barony
	莫拉Moura 	feud.Barony
	莫朗Mourao 	feud.Barony
	诺达尔Noudal 	feud.Barony
	波特尔Portel 	feud.Barony
	塞尔帕Serpa 	feud.Barony
}

func (c *梅尔图拉MertolaCounty) BBeja贝雅() feud.Barony {
	return c.贝雅Beja
}
    
func (c *梅尔图拉MertolaCounty) BMertola梅尔图拉() feud.Barony {
	return c.梅尔图拉Mertola
}
    
func (c *梅尔图拉MertolaCounty) BMonsaraz蒙萨拉什() feud.Barony {
	return c.蒙萨拉什Monsaraz
}
    
func (c *梅尔图拉MertolaCounty) BMoura莫拉() feud.Barony {
	return c.莫拉Moura
}
    
func (c *梅尔图拉MertolaCounty) BMourao莫朗() feud.Barony {
	return c.莫朗Mourao
}
    
func (c *梅尔图拉MertolaCounty) BNoudal诺达尔() feud.Barony {
	return c.诺达尔Noudal
}
    
func (c *梅尔图拉MertolaCounty) BPortel波特尔() feud.Barony {
	return c.波特尔Portel
}
    
func (c *梅尔图拉MertolaCounty) BSerpa塞尔帕() feud.Barony {
	return c.塞尔帕Serpa
}
    
var CMertola梅尔图拉 MertolaCounty = &梅尔图拉MertolaCounty{}

func init() {
	f := CMertola梅尔图拉.(*梅尔图拉MertolaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "185",
		Title:     "mertola",
		TitleName: "梅尔图拉",
		TitleCode: "c_mertola",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝雅Beja = BBeja贝雅
	f.贝雅Beja.SetParent(f)
	
	f.梅尔图拉Mertola = BMertola梅尔图拉
	f.梅尔图拉Mertola.SetParent(f)
	
	f.蒙萨拉什Monsaraz = BMonsaraz蒙萨拉什
	f.蒙萨拉什Monsaraz.SetParent(f)
	
	f.莫拉Moura = BMoura莫拉
	f.莫拉Moura.SetParent(f)
	
	f.莫朗Mourao = BMourao莫朗
	f.莫朗Mourao.SetParent(f)
	
	f.诺达尔Noudal = BNoudal诺达尔
	f.诺达尔Noudal.SetParent(f)
	
	f.波特尔Portel = BPortel波特尔
	f.波特尔Portel.SetParent(f)
	
	f.塞尔帕Serpa = BSerpa塞尔帕
	f.塞尔帕Serpa.SetParent(f)
	
}
