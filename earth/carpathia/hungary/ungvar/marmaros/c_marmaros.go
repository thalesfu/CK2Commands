package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarmarosCounty interface {
    feud.County
    BAknasugatag奥克瑙舒高陶格() 	feud.Barony
    BFelsoviso费尔舍维绍() 	feud.Barony
    BHuszt胡斯特() 	feud.Barony
    BMaramarossziget马劳毛罗什西盖特() 	feud.Barony
    BNagybanya瑙吉巴尼奥() 	feud.Barony
    BNagykaroly瑙吉卡罗伊() 	feud.Barony
    BRaho劳霍() 	feud.Barony
    BTecso泰什克() 	feud.Barony
}

type 马尔毛罗什MarmarosCounty struct {
	feud.BaseCounty
	奥克瑙舒高陶格Aknasugatag 	feud.Barony
	费尔舍维绍Felsoviso 	feud.Barony
	胡斯特Huszt 	feud.Barony
	马劳毛罗什西盖特Maramarossziget 	feud.Barony
	瑙吉巴尼奥Nagybanya 	feud.Barony
	瑙吉卡罗伊Nagykaroly 	feud.Barony
	劳霍Raho 	feud.Barony
	泰什克Tecso 	feud.Barony
}

func (c *马尔毛罗什MarmarosCounty) BAknasugatag奥克瑙舒高陶格() feud.Barony {
	return c.奥克瑙舒高陶格Aknasugatag
}
    
func (c *马尔毛罗什MarmarosCounty) BFelsoviso费尔舍维绍() feud.Barony {
	return c.费尔舍维绍Felsoviso
}
    
func (c *马尔毛罗什MarmarosCounty) BHuszt胡斯特() feud.Barony {
	return c.胡斯特Huszt
}
    
func (c *马尔毛罗什MarmarosCounty) BMaramarossziget马劳毛罗什西盖特() feud.Barony {
	return c.马劳毛罗什西盖特Maramarossziget
}
    
func (c *马尔毛罗什MarmarosCounty) BNagybanya瑙吉巴尼奥() feud.Barony {
	return c.瑙吉巴尼奥Nagybanya
}
    
func (c *马尔毛罗什MarmarosCounty) BNagykaroly瑙吉卡罗伊() feud.Barony {
	return c.瑙吉卡罗伊Nagykaroly
}
    
func (c *马尔毛罗什MarmarosCounty) BRaho劳霍() feud.Barony {
	return c.劳霍Raho
}
    
func (c *马尔毛罗什MarmarosCounty) BTecso泰什克() feud.Barony {
	return c.泰什克Tecso
}
    
var CMarmaros马尔毛罗什 MarmarosCounty = &马尔毛罗什MarmarosCounty{}

func init() {
	f := CMarmaros马尔毛罗什.(*马尔毛罗什MarmarosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "539",
		Title:     "marmaros",
		TitleName: "马尔毛罗什",
		TitleCode: "c_marmaros",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥克瑙舒高陶格Aknasugatag = BAknasugatag奥克瑙舒高陶格
	f.奥克瑙舒高陶格Aknasugatag.SetParent(f)
	
	f.费尔舍维绍Felsoviso = BFelsoviso费尔舍维绍
	f.费尔舍维绍Felsoviso.SetParent(f)
	
	f.胡斯特Huszt = BHuszt胡斯特
	f.胡斯特Huszt.SetParent(f)
	
	f.马劳毛罗什西盖特Maramarossziget = BMaramarossziget马劳毛罗什西盖特
	f.马劳毛罗什西盖特Maramarossziget.SetParent(f)
	
	f.瑙吉巴尼奥Nagybanya = BNagybanya瑙吉巴尼奥
	f.瑙吉巴尼奥Nagybanya.SetParent(f)
	
	f.瑙吉卡罗伊Nagykaroly = BNagykaroly瑙吉卡罗伊
	f.瑙吉卡罗伊Nagykaroly.SetParent(f)
	
	f.劳霍Raho = BRaho劳霍
	f.劳霍Raho.SetParent(f)
	
	f.泰什克Tecso = BTecso泰什克
	f.泰什克Tecso.SetParent(f)
	
}
