package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GuriaCounty interface {
    feud.County
    BBatum巴统() 	feud.Barony
    BBukistsikhe布基斯齐赫() 	feud.Barony
    BGonio戈尼奥() 	feud.Barony
    BKhula胡拉() 	feud.Barony
    BOzurgeti奥祖尔盖蒂() 	feud.Barony
    BSkhalta斯哈尔塔() 	feud.Barony
    BTsikhisdziri齐希斯济里() 	feud.Barony
    BUdabno乌达比诺() 	feud.Barony
}

type 古里亚GuriaCounty struct {
	feud.BaseCounty
	巴统Batum 	feud.Barony
	布基斯齐赫Bukistsikhe 	feud.Barony
	戈尼奥Gonio 	feud.Barony
	胡拉Khula 	feud.Barony
	奥祖尔盖蒂Ozurgeti 	feud.Barony
	斯哈尔塔Skhalta 	feud.Barony
	齐希斯济里Tsikhisdziri 	feud.Barony
	乌达比诺Udabno 	feud.Barony
}

func (c *古里亚GuriaCounty) BBatum巴统() feud.Barony {
	return c.巴统Batum
}
    
func (c *古里亚GuriaCounty) BBukistsikhe布基斯齐赫() feud.Barony {
	return c.布基斯齐赫Bukistsikhe
}
    
func (c *古里亚GuriaCounty) BGonio戈尼奥() feud.Barony {
	return c.戈尼奥Gonio
}
    
func (c *古里亚GuriaCounty) BKhula胡拉() feud.Barony {
	return c.胡拉Khula
}
    
func (c *古里亚GuriaCounty) BOzurgeti奥祖尔盖蒂() feud.Barony {
	return c.奥祖尔盖蒂Ozurgeti
}
    
func (c *古里亚GuriaCounty) BSkhalta斯哈尔塔() feud.Barony {
	return c.斯哈尔塔Skhalta
}
    
func (c *古里亚GuriaCounty) BTsikhisdziri齐希斯济里() feud.Barony {
	return c.齐希斯济里Tsikhisdziri
}
    
func (c *古里亚GuriaCounty) BUdabno乌达比诺() feud.Barony {
	return c.乌达比诺Udabno
}
    
var CGuria古里亚 GuriaCounty = &古里亚GuriaCounty{}

func init() {
	f := CGuria古里亚.(*古里亚GuriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "677",
		Title:     "guria",
		TitleName: "古里亚",
		TitleCode: "c_guria",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴统Batum = BBatum巴统
	f.巴统Batum.SetParent(f)
	
	f.布基斯齐赫Bukistsikhe = BBukistsikhe布基斯齐赫
	f.布基斯齐赫Bukistsikhe.SetParent(f)
	
	f.戈尼奥Gonio = BGonio戈尼奥
	f.戈尼奥Gonio.SetParent(f)
	
	f.胡拉Khula = BKhula胡拉
	f.胡拉Khula.SetParent(f)
	
	f.奥祖尔盖蒂Ozurgeti = BOzurgeti奥祖尔盖蒂
	f.奥祖尔盖蒂Ozurgeti.SetParent(f)
	
	f.斯哈尔塔Skhalta = BSkhalta斯哈尔塔
	f.斯哈尔塔Skhalta.SetParent(f)
	
	f.齐希斯济里Tsikhisdziri = BTsikhisdziri齐希斯济里
	f.齐希斯济里Tsikhisdziri.SetParent(f)
	
	f.乌达比诺Udabno = BUdabno乌达比诺
	f.乌达比诺Udabno.SetParent(f)
	
}
