package maragha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaraghaCounty interface {
    feud.County
    BBahaduran巴哈杜兰() 	feud.Barony
    BKileh_shin基莱欣() 	feud.Barony
    BKursara库尔萨拉() 	feud.Barony
    BMansurlu曼苏尔卢() 	feud.Barony
    BMaragheh马拉盖() 	feud.Barony
    BMiyaneh米亚内() 	feud.Barony
    BUjan乌詹() 	feud.Barony
}

type 蔑剌哈MaraghaCounty struct {
	feud.BaseCounty
	巴哈杜兰Bahaduran 	feud.Barony
	基莱欣Kileh_shin 	feud.Barony
	库尔萨拉Kursara 	feud.Barony
	曼苏尔卢Mansurlu 	feud.Barony
	马拉盖Maragheh 	feud.Barony
	米亚内Miyaneh 	feud.Barony
	乌詹Ujan 	feud.Barony
}

func (c *蔑剌哈MaraghaCounty) BBahaduran巴哈杜兰() feud.Barony {
	return c.巴哈杜兰Bahaduran
}
    
func (c *蔑剌哈MaraghaCounty) BKileh_shin基莱欣() feud.Barony {
	return c.基莱欣Kileh_shin
}
    
func (c *蔑剌哈MaraghaCounty) BKursara库尔萨拉() feud.Barony {
	return c.库尔萨拉Kursara
}
    
func (c *蔑剌哈MaraghaCounty) BMansurlu曼苏尔卢() feud.Barony {
	return c.曼苏尔卢Mansurlu
}
    
func (c *蔑剌哈MaraghaCounty) BMaragheh马拉盖() feud.Barony {
	return c.马拉盖Maragheh
}
    
func (c *蔑剌哈MaraghaCounty) BMiyaneh米亚内() feud.Barony {
	return c.米亚内Miyaneh
}
    
func (c *蔑剌哈MaraghaCounty) BUjan乌詹() feud.Barony {
	return c.乌詹Ujan
}
    
var CMaragha蔑剌哈 MaraghaCounty = &蔑剌哈MaraghaCounty{}

func init() {
	f := CMaragha蔑剌哈.(*蔑剌哈MaraghaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1537",
		Title:     "maragha",
		TitleName: "蔑剌哈",
		TitleCode: "c_maragha",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴哈杜兰Bahaduran = BBahaduran巴哈杜兰
	f.巴哈杜兰Bahaduran.SetParent(f)
	
	f.基莱欣Kileh_shin = BKileh_shin基莱欣
	f.基莱欣Kileh_shin.SetParent(f)
	
	f.库尔萨拉Kursara = BKursara库尔萨拉
	f.库尔萨拉Kursara.SetParent(f)
	
	f.曼苏尔卢Mansurlu = BMansurlu曼苏尔卢
	f.曼苏尔卢Mansurlu.SetParent(f)
	
	f.马拉盖Maragheh = BMaragheh马拉盖
	f.马拉盖Maragheh.SetParent(f)
	
	f.米亚内Miyaneh = BMiyaneh米亚内
	f.米亚内Miyaneh.SetParent(f)
	
	f.乌詹Ujan = BUjan乌詹
	f.乌詹Ujan.SetParent(f)
	
}
