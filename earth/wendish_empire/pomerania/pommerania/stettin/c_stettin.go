package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StettinCounty interface {
    feud.County
    BCedene采代内() 	feud.Barony
    BKammin卡明() 	feud.Barony
    BKolbatz科尔巴茨() 	feud.Barony
    BPyritz皮里茨() 	feud.Barony
    BSoldin索尔丁() 	feud.Barony
    BStargard施塔加德() 	feud.Barony
    BStettin斯德丁() 	feud.Barony
    BWollin沃林() 	feud.Barony
}

type 斯德丁StettinCounty struct {
	feud.BaseCounty
	采代内Cedene 	feud.Barony
	卡明Kammin 	feud.Barony
	科尔巴茨Kolbatz 	feud.Barony
	皮里茨Pyritz 	feud.Barony
	索尔丁Soldin 	feud.Barony
	施塔加德Stargard 	feud.Barony
	斯德丁Stettin 	feud.Barony
	沃林Wollin 	feud.Barony
}

func (c *斯德丁StettinCounty) BCedene采代内() feud.Barony {
	return c.采代内Cedene
}
    
func (c *斯德丁StettinCounty) BKammin卡明() feud.Barony {
	return c.卡明Kammin
}
    
func (c *斯德丁StettinCounty) BKolbatz科尔巴茨() feud.Barony {
	return c.科尔巴茨Kolbatz
}
    
func (c *斯德丁StettinCounty) BPyritz皮里茨() feud.Barony {
	return c.皮里茨Pyritz
}
    
func (c *斯德丁StettinCounty) BSoldin索尔丁() feud.Barony {
	return c.索尔丁Soldin
}
    
func (c *斯德丁StettinCounty) BStargard施塔加德() feud.Barony {
	return c.施塔加德Stargard
}
    
func (c *斯德丁StettinCounty) BStettin斯德丁() feud.Barony {
	return c.斯德丁Stettin
}
    
func (c *斯德丁StettinCounty) BWollin沃林() feud.Barony {
	return c.沃林Wollin
}
    
var CStettin斯德丁 StettinCounty = &斯德丁StettinCounty{}

func init() {
	f := CStettin斯德丁.(*斯德丁StettinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "366",
		Title:     "stettin",
		TitleName: "斯德丁",
		TitleCode: "c_stettin",
		Baronies:  map[string]feud.Barony{},
	}

	f.采代内Cedene = BCedene采代内
	f.采代内Cedene.SetParent(f)
	
	f.卡明Kammin = BKammin卡明
	f.卡明Kammin.SetParent(f)
	
	f.科尔巴茨Kolbatz = BKolbatz科尔巴茨
	f.科尔巴茨Kolbatz.SetParent(f)
	
	f.皮里茨Pyritz = BPyritz皮里茨
	f.皮里茨Pyritz.SetParent(f)
	
	f.索尔丁Soldin = BSoldin索尔丁
	f.索尔丁Soldin.SetParent(f)
	
	f.施塔加德Stargard = BStargard施塔加德
	f.施塔加德Stargard.SetParent(f)
	
	f.斯德丁Stettin = BStettin斯德丁
	f.斯德丁Stettin.SetParent(f)
	
	f.沃林Wollin = BWollin沃林
	f.沃林Wollin.SetParent(f)
	
}
