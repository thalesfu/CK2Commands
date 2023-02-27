package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RayapuraCounty interface {
    feud.County
    BNandgram难陀伽罗摩() 	feud.Barony
    BPutukipali富兜吉波梨() 	feud.Barony
    BRabluto罗卢豆() 	feud.Barony
    BRajim罗什伐() 	feud.Barony
    BRayapura罗耶补罗() 	feud.Barony
    BShivadurga湿婆突伽() 	feud.Barony
    BShivapura湿婆城() 	feud.Barony
}

type 罗耶补罗RayapuraCounty struct {
	feud.BaseCounty
	难陀伽罗摩Nandgram 	feud.Barony
	富兜吉波梨Putukipali 	feud.Barony
	罗卢豆Rabluto 	feud.Barony
	罗什伐Rajim 	feud.Barony
	罗耶补罗Rayapura 	feud.Barony
	湿婆突伽Shivadurga 	feud.Barony
	湿婆城Shivapura 	feud.Barony
}

func (c *罗耶补罗RayapuraCounty) BNandgram难陀伽罗摩() feud.Barony {
	return c.难陀伽罗摩Nandgram
}
    
func (c *罗耶补罗RayapuraCounty) BPutukipali富兜吉波梨() feud.Barony {
	return c.富兜吉波梨Putukipali
}
    
func (c *罗耶补罗RayapuraCounty) BRabluto罗卢豆() feud.Barony {
	return c.罗卢豆Rabluto
}
    
func (c *罗耶补罗RayapuraCounty) BRajim罗什伐() feud.Barony {
	return c.罗什伐Rajim
}
    
func (c *罗耶补罗RayapuraCounty) BRayapura罗耶补罗() feud.Barony {
	return c.罗耶补罗Rayapura
}
    
func (c *罗耶补罗RayapuraCounty) BShivadurga湿婆突伽() feud.Barony {
	return c.湿婆突伽Shivadurga
}
    
func (c *罗耶补罗RayapuraCounty) BShivapura湿婆城() feud.Barony {
	return c.湿婆城Shivapura
}
    
var CRayapura罗耶补罗 RayapuraCounty = &罗耶补罗RayapuraCounty{}

func init() {
	f := CRayapura罗耶补罗.(*罗耶补罗RayapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1160",
		Title:     "rayapura",
		TitleName: "罗耶补罗",
		TitleCode: "c_rayapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.难陀伽罗摩Nandgram = BNandgram难陀伽罗摩
	f.难陀伽罗摩Nandgram.SetParent(f)
	
	f.富兜吉波梨Putukipali = BPutukipali富兜吉波梨
	f.富兜吉波梨Putukipali.SetParent(f)
	
	f.罗卢豆Rabluto = BRabluto罗卢豆
	f.罗卢豆Rabluto.SetParent(f)
	
	f.罗什伐Rajim = BRajim罗什伐
	f.罗什伐Rajim.SetParent(f)
	
	f.罗耶补罗Rayapura = BRayapura罗耶补罗
	f.罗耶补罗Rayapura.SetParent(f)
	
	f.湿婆突伽Shivadurga = BShivadurga湿婆突伽
	f.湿婆突伽Shivadurga.SetParent(f)
	
	f.湿婆城Shivapura = BShivapura湿婆城
	f.湿婆城Shivapura.SetParent(f)
	
}
