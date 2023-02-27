package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KangraCounty interface {
    feud.County
    BBharmour婆罗吸摩补罗() 	feud.Barony
    BChamba瞻波() 	feud.Barony
    BKahlur贾罗尔() 	feud.Barony
    BKangra建伽罗() 	feud.Barony
    BKullu屈露() 	feud.Barony
    BNurpur努尔布尔() 	feud.Barony
    BShimla西姆拉() 	feud.Barony
}

type 建伽罗KangraCounty struct {
	feud.BaseCounty
	婆罗吸摩补罗Bharmour 	feud.Barony
	瞻波Chamba 	feud.Barony
	贾罗尔Kahlur 	feud.Barony
	建伽罗Kangra 	feud.Barony
	屈露Kullu 	feud.Barony
	努尔布尔Nurpur 	feud.Barony
	西姆拉Shimla 	feud.Barony
}

func (c *建伽罗KangraCounty) BBharmour婆罗吸摩补罗() feud.Barony {
	return c.婆罗吸摩补罗Bharmour
}
    
func (c *建伽罗KangraCounty) BChamba瞻波() feud.Barony {
	return c.瞻波Chamba
}
    
func (c *建伽罗KangraCounty) BKahlur贾罗尔() feud.Barony {
	return c.贾罗尔Kahlur
}
    
func (c *建伽罗KangraCounty) BKangra建伽罗() feud.Barony {
	return c.建伽罗Kangra
}
    
func (c *建伽罗KangraCounty) BKullu屈露() feud.Barony {
	return c.屈露Kullu
}
    
func (c *建伽罗KangraCounty) BNurpur努尔布尔() feud.Barony {
	return c.努尔布尔Nurpur
}
    
func (c *建伽罗KangraCounty) BShimla西姆拉() feud.Barony {
	return c.西姆拉Shimla
}
    
var CKangra建伽罗 KangraCounty = &建伽罗KangraCounty{}

func init() {
	f := CKangra建伽罗.(*建伽罗KangraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1469",
		Title:     "kangra",
		TitleName: "建伽罗",
		TitleCode: "c_kangra",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗吸摩补罗Bharmour = BBharmour婆罗吸摩补罗
	f.婆罗吸摩补罗Bharmour.SetParent(f)
	
	f.瞻波Chamba = BChamba瞻波
	f.瞻波Chamba.SetParent(f)
	
	f.贾罗尔Kahlur = BKahlur贾罗尔
	f.贾罗尔Kahlur.SetParent(f)
	
	f.建伽罗Kangra = BKangra建伽罗
	f.建伽罗Kangra.SetParent(f)
	
	f.屈露Kullu = BKullu屈露
	f.屈露Kullu.SetParent(f)
	
	f.努尔布尔Nurpur = BNurpur努尔布尔
	f.努尔布尔Nurpur.SetParent(f)
	
	f.西姆拉Shimla = BShimla西姆拉
	f.西姆拉Shimla.SetParent(f)
	
}
