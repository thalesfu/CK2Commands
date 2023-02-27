package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DalarnaCounty interface {
    feud.County
    BBorganas博尔加内斯() 	feud.Barony
    BHedemora海德穆拉() 	feud.Barony
    BHusby许斯比() 	feud.Barony
    BIdre伊德勒() 	feud.Barony
    BKopparberg科帕尔贝里() 	feud.Barony
    BMora穆拉() 	feud.Barony
    BSater塞特() 	feud.Barony
    BTuna吞拿() 	feud.Barony
}

type 耶恩贝拉兰DalarnaCounty struct {
	feud.BaseCounty
	博尔加内斯Borganas 	feud.Barony
	海德穆拉Hedemora 	feud.Barony
	许斯比Husby 	feud.Barony
	伊德勒Idre 	feud.Barony
	科帕尔贝里Kopparberg 	feud.Barony
	穆拉Mora 	feud.Barony
	塞特Sater 	feud.Barony
	吞拿Tuna 	feud.Barony
}

func (c *耶恩贝拉兰DalarnaCounty) BBorganas博尔加内斯() feud.Barony {
	return c.博尔加内斯Borganas
}
    
func (c *耶恩贝拉兰DalarnaCounty) BHedemora海德穆拉() feud.Barony {
	return c.海德穆拉Hedemora
}
    
func (c *耶恩贝拉兰DalarnaCounty) BHusby许斯比() feud.Barony {
	return c.许斯比Husby
}
    
func (c *耶恩贝拉兰DalarnaCounty) BIdre伊德勒() feud.Barony {
	return c.伊德勒Idre
}
    
func (c *耶恩贝拉兰DalarnaCounty) BKopparberg科帕尔贝里() feud.Barony {
	return c.科帕尔贝里Kopparberg
}
    
func (c *耶恩贝拉兰DalarnaCounty) BMora穆拉() feud.Barony {
	return c.穆拉Mora
}
    
func (c *耶恩贝拉兰DalarnaCounty) BSater塞特() feud.Barony {
	return c.塞特Sater
}
    
func (c *耶恩贝拉兰DalarnaCounty) BTuna吞拿() feud.Barony {
	return c.吞拿Tuna
}
    
var CDalarna耶恩贝拉兰 DalarnaCounty = &耶恩贝拉兰DalarnaCounty{}

func init() {
	f := CDalarna耶恩贝拉兰.(*耶恩贝拉兰DalarnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "287",
		Title:     "dalarna",
		TitleName: "耶恩贝拉兰",
		TitleCode: "c_dalarna",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔加内斯Borganas = BBorganas博尔加内斯
	f.博尔加内斯Borganas.SetParent(f)
	
	f.海德穆拉Hedemora = BHedemora海德穆拉
	f.海德穆拉Hedemora.SetParent(f)
	
	f.许斯比Husby = BHusby许斯比
	f.许斯比Husby.SetParent(f)
	
	f.伊德勒Idre = BIdre伊德勒
	f.伊德勒Idre.SetParent(f)
	
	f.科帕尔贝里Kopparberg = BKopparberg科帕尔贝里
	f.科帕尔贝里Kopparberg.SetParent(f)
	
	f.穆拉Mora = BMora穆拉
	f.穆拉Mora.SetParent(f)
	
	f.塞特Sater = BSater塞特
	f.塞特Sater.SetParent(f)
	
	f.吞拿Tuna = BTuna吞拿
	f.吞拿Tuna.SetParent(f)
	
}
