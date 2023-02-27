package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DelgerkhangaiCounty interface {
    feud.County
    BDelgerkhangai德勒格尔杭爱() 	feud.Barony
    BDelgertsogt德勒格尔朝格特() 	feud.Barony
    BKhuld_delgerkhangai呼勒德() 	feud.Barony
    BLuus卢斯() 	feud.Barony
    BSaikhan赛罕() 	feud.Barony
    BSangiindalai桑根达来() 	feud.Barony
    BTsant昌特() 	feud.Barony
}

type 德勒格尔杭爱DelgerkhangaiCounty struct {
	feud.BaseCounty
	德勒格尔杭爱Delgerkhangai 	feud.Barony
	德勒格尔朝格特Delgertsogt 	feud.Barony
	呼勒德Khuld_delgerkhangai 	feud.Barony
	卢斯Luus 	feud.Barony
	赛罕Saikhan 	feud.Barony
	桑根达来Sangiindalai 	feud.Barony
	昌特Tsant 	feud.Barony
}

func (c *德勒格尔杭爱DelgerkhangaiCounty) BDelgerkhangai德勒格尔杭爱() feud.Barony {
	return c.德勒格尔杭爱Delgerkhangai
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BDelgertsogt德勒格尔朝格特() feud.Barony {
	return c.德勒格尔朝格特Delgertsogt
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BKhuld_delgerkhangai呼勒德() feud.Barony {
	return c.呼勒德Khuld_delgerkhangai
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BLuus卢斯() feud.Barony {
	return c.卢斯Luus
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BSaikhan赛罕() feud.Barony {
	return c.赛罕Saikhan
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BSangiindalai桑根达来() feud.Barony {
	return c.桑根达来Sangiindalai
}
    
func (c *德勒格尔杭爱DelgerkhangaiCounty) BTsant昌特() feud.Barony {
	return c.昌特Tsant
}
    
var CDelgerkhangai德勒格尔杭爱 DelgerkhangaiCounty = &德勒格尔杭爱DelgerkhangaiCounty{}

func init() {
	f := CDelgerkhangai德勒格尔杭爱.(*德勒格尔杭爱DelgerkhangaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1918",
		Title:     "delgerkhangai",
		TitleName: "德勒格尔杭爱",
		TitleCode: "c_delgerkhangai",
		Baronies:  map[string]feud.Barony{},
	}

	f.德勒格尔杭爱Delgerkhangai = BDelgerkhangai德勒格尔杭爱
	f.德勒格尔杭爱Delgerkhangai.SetParent(f)
	
	f.德勒格尔朝格特Delgertsogt = BDelgertsogt德勒格尔朝格特
	f.德勒格尔朝格特Delgertsogt.SetParent(f)
	
	f.呼勒德Khuld_delgerkhangai = BKhuld_delgerkhangai呼勒德
	f.呼勒德Khuld_delgerkhangai.SetParent(f)
	
	f.卢斯Luus = BLuus卢斯
	f.卢斯Luus.SetParent(f)
	
	f.赛罕Saikhan = BSaikhan赛罕
	f.赛罕Saikhan.SetParent(f)
	
	f.桑根达来Sangiindalai = BSangiindalai桑根达来
	f.桑根达来Sangiindalai.SetParent(f)
	
	f.昌特Tsant = BTsant昌特
	f.昌特Tsant.SetParent(f)
	
}
