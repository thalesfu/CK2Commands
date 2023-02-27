package idatarainadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IdatarainaduCounty interface {
    feud.County
    BGabburu加布罗() 	feud.Barony
    BHemagudda醯摩矩吒() 	feud.Barony
    BJaladurga阇罗突伽() 	feud.Barony
    BKammatadurga卡姆摩塔杜尔加() 	feud.Barony
    BManvi曼维() 	feud.Barony
    BMudgal勿伽罗() 	feud.Barony
    BRaichur罗耶注卢() 	feud.Barony
}

type 伊陀多来拿都IdatarainaduCounty struct {
	feud.BaseCounty
	加布罗Gabburu 	feud.Barony
	醯摩矩吒Hemagudda 	feud.Barony
	阇罗突伽Jaladurga 	feud.Barony
	卡姆摩塔杜尔加Kammatadurga 	feud.Barony
	曼维Manvi 	feud.Barony
	勿伽罗Mudgal 	feud.Barony
	罗耶注卢Raichur 	feud.Barony
}

func (c *伊陀多来拿都IdatarainaduCounty) BGabburu加布罗() feud.Barony {
	return c.加布罗Gabburu
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BHemagudda醯摩矩吒() feud.Barony {
	return c.醯摩矩吒Hemagudda
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BJaladurga阇罗突伽() feud.Barony {
	return c.阇罗突伽Jaladurga
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BKammatadurga卡姆摩塔杜尔加() feud.Barony {
	return c.卡姆摩塔杜尔加Kammatadurga
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BManvi曼维() feud.Barony {
	return c.曼维Manvi
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BMudgal勿伽罗() feud.Barony {
	return c.勿伽罗Mudgal
}
    
func (c *伊陀多来拿都IdatarainaduCounty) BRaichur罗耶注卢() feud.Barony {
	return c.罗耶注卢Raichur
}
    
var CIdatarainadu伊陀多来拿都 IdatarainaduCounty = &伊陀多来拿都IdatarainaduCounty{}

func init() {
	f := CIdatarainadu伊陀多来拿都.(*伊陀多来拿都IdatarainaduCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1203",
		Title:     "idatarainadu",
		TitleName: "伊陀多来拿都",
		TitleCode: "c_idatarainadu",
		Baronies:  map[string]feud.Barony{},
	}

	f.加布罗Gabburu = BGabburu加布罗
	f.加布罗Gabburu.SetParent(f)
	
	f.醯摩矩吒Hemagudda = BHemagudda醯摩矩吒
	f.醯摩矩吒Hemagudda.SetParent(f)
	
	f.阇罗突伽Jaladurga = BJaladurga阇罗突伽
	f.阇罗突伽Jaladurga.SetParent(f)
	
	f.卡姆摩塔杜尔加Kammatadurga = BKammatadurga卡姆摩塔杜尔加
	f.卡姆摩塔杜尔加Kammatadurga.SetParent(f)
	
	f.曼维Manvi = BManvi曼维
	f.曼维Manvi.SetParent(f)
	
	f.勿伽罗Mudgal = BMudgal勿伽罗
	f.勿伽罗Mudgal.SetParent(f)
	
	f.罗耶注卢Raichur = BRaichur罗耶注卢
	f.罗耶注卢Raichur.SetParent(f)
	
}
