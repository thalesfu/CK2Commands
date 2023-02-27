package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SibiCounty interface {
    feud.County
    BDadhar达哈尔() 	feud.Barony
    BMastanj马斯坦迦() 	feud.Barony
    BSaprenda娑薜荔陀() 	feud.Barony
    BSardoi萨豆伊() 	feud.Barony
    BSibi尸毗() 	feud.Barony
    BValikondapuram伐梨军荼补蓝() 	feud.Barony
    BWalishtan瓦利斯坦() 	feud.Barony
}

type 尸毗SibiCounty struct {
	feud.BaseCounty
	达哈尔Dadhar 	feud.Barony
	马斯坦迦Mastanj 	feud.Barony
	娑薜荔陀Saprenda 	feud.Barony
	萨豆伊Sardoi 	feud.Barony
	尸毗Sibi 	feud.Barony
	伐梨军荼补蓝Valikondapuram 	feud.Barony
	瓦利斯坦Walishtan 	feud.Barony
}

func (c *尸毗SibiCounty) BDadhar达哈尔() feud.Barony {
	return c.达哈尔Dadhar
}
    
func (c *尸毗SibiCounty) BMastanj马斯坦迦() feud.Barony {
	return c.马斯坦迦Mastanj
}
    
func (c *尸毗SibiCounty) BSaprenda娑薜荔陀() feud.Barony {
	return c.娑薜荔陀Saprenda
}
    
func (c *尸毗SibiCounty) BSardoi萨豆伊() feud.Barony {
	return c.萨豆伊Sardoi
}
    
func (c *尸毗SibiCounty) BSibi尸毗() feud.Barony {
	return c.尸毗Sibi
}
    
func (c *尸毗SibiCounty) BValikondapuram伐梨军荼补蓝() feud.Barony {
	return c.伐梨军荼补蓝Valikondapuram
}
    
func (c *尸毗SibiCounty) BWalishtan瓦利斯坦() feud.Barony {
	return c.瓦利斯坦Walishtan
}
    
var CSibi尸毗 SibiCounty = &尸毗SibiCounty{}

func init() {
	f := CSibi尸毗.(*尸毗SibiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1374",
		Title:     "sibi",
		TitleName: "尸毗",
		TitleCode: "c_sibi",
		Baronies:  map[string]feud.Barony{},
	}

	f.达哈尔Dadhar = BDadhar达哈尔
	f.达哈尔Dadhar.SetParent(f)
	
	f.马斯坦迦Mastanj = BMastanj马斯坦迦
	f.马斯坦迦Mastanj.SetParent(f)
	
	f.娑薜荔陀Saprenda = BSaprenda娑薜荔陀
	f.娑薜荔陀Saprenda.SetParent(f)
	
	f.萨豆伊Sardoi = BSardoi萨豆伊
	f.萨豆伊Sardoi.SetParent(f)
	
	f.尸毗Sibi = BSibi尸毗
	f.尸毗Sibi.SetParent(f)
	
	f.伐梨军荼补蓝Valikondapuram = BValikondapuram伐梨军荼补蓝
	f.伐梨军荼补蓝Valikondapuram.SetParent(f)
	
	f.瓦利斯坦Walishtan = BWalishtan瓦利斯坦
	f.瓦利斯坦Walishtan.SetParent(f)
	
}
