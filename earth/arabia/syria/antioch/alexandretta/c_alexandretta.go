package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlexandrettaCounty interface {
    feud.County
    BAlexandretta亚历山大勒塔() 	feud.Barony
    BDerbasak达尔伯萨克() 	feud.Barony
    BLarochederissole拉罗什德里索勒() 	feud.Barony
    BMyriandros迈里安德鲁斯() 	feud.Barony
    BPortbonnel博内勒港() 	feud.Barony
    BPortella罗特拉() 	feud.Barony
    BRhosus罗苏斯() 	feud.Barony
    BSarvantikar萨凡提卡() 	feud.Barony
}

type 亚历山大勒塔AlexandrettaCounty struct {
	feud.BaseCounty
	亚历山大勒塔Alexandretta 	feud.Barony
	达尔伯萨克Derbasak 	feud.Barony
	拉罗什德里索勒Larochederissole 	feud.Barony
	迈里安德鲁斯Myriandros 	feud.Barony
	博内勒港Portbonnel 	feud.Barony
	罗特拉Portella 	feud.Barony
	罗苏斯Rhosus 	feud.Barony
	萨凡提卡Sarvantikar 	feud.Barony
}

func (c *亚历山大勒塔AlexandrettaCounty) BAlexandretta亚历山大勒塔() feud.Barony {
	return c.亚历山大勒塔Alexandretta
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BDerbasak达尔伯萨克() feud.Barony {
	return c.达尔伯萨克Derbasak
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BLarochederissole拉罗什德里索勒() feud.Barony {
	return c.拉罗什德里索勒Larochederissole
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BMyriandros迈里安德鲁斯() feud.Barony {
	return c.迈里安德鲁斯Myriandros
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BPortbonnel博内勒港() feud.Barony {
	return c.博内勒港Portbonnel
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BPortella罗特拉() feud.Barony {
	return c.罗特拉Portella
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BRhosus罗苏斯() feud.Barony {
	return c.罗苏斯Rhosus
}
    
func (c *亚历山大勒塔AlexandrettaCounty) BSarvantikar萨凡提卡() feud.Barony {
	return c.萨凡提卡Sarvantikar
}
    
var CAlexandretta亚历山大勒塔 AlexandrettaCounty = &亚历山大勒塔AlexandrettaCounty{}

func init() {
	f := CAlexandretta亚历山大勒塔.(*亚历山大勒塔AlexandrettaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "763",
		Title:     "alexandretta",
		TitleName: "亚历山大勒塔",
		TitleCode: "c_alexandretta",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚历山大勒塔Alexandretta = BAlexandretta亚历山大勒塔
	f.亚历山大勒塔Alexandretta.SetParent(f)
	
	f.达尔伯萨克Derbasak = BDerbasak达尔伯萨克
	f.达尔伯萨克Derbasak.SetParent(f)
	
	f.拉罗什德里索勒Larochederissole = BLarochederissole拉罗什德里索勒
	f.拉罗什德里索勒Larochederissole.SetParent(f)
	
	f.迈里安德鲁斯Myriandros = BMyriandros迈里安德鲁斯
	f.迈里安德鲁斯Myriandros.SetParent(f)
	
	f.博内勒港Portbonnel = BPortbonnel博内勒港
	f.博内勒港Portbonnel.SetParent(f)
	
	f.罗特拉Portella = BPortella罗特拉
	f.罗特拉Portella.SetParent(f)
	
	f.罗苏斯Rhosus = BRhosus罗苏斯
	f.罗苏斯Rhosus.SetParent(f)
	
	f.萨凡提卡Sarvantikar = BSarvantikar萨凡提卡
	f.萨凡提卡Sarvantikar.SetParent(f)
	
}
