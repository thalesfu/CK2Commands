package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SiwistanCounty interface {
    feud.County
    BBahera婆醯罗() 	feud.Barony
    BDahlilah陀梨罗那() 	feud.Barony
    BGuja古奢() 	feud.Barony
    BKamharia甘诃梨() 	feud.Barony
    BMir_rukan弥尔楼乾() 	feud.Barony
    BSiwistan室毗湿檀那() 	feud.Barony
    BVicalapura毗遮罗补罗() 	feud.Barony
}

type 室毗湿檀那SiwistanCounty struct {
	feud.BaseCounty
	婆醯罗Bahera 	feud.Barony
	陀梨罗那Dahlilah 	feud.Barony
	古奢Guja 	feud.Barony
	甘诃梨Kamharia 	feud.Barony
	弥尔楼乾Mir_rukan 	feud.Barony
	室毗湿檀那Siwistan 	feud.Barony
	毗遮罗补罗Vicalapura 	feud.Barony
}

func (c *室毗湿檀那SiwistanCounty) BBahera婆醯罗() feud.Barony {
	return c.婆醯罗Bahera
}
    
func (c *室毗湿檀那SiwistanCounty) BDahlilah陀梨罗那() feud.Barony {
	return c.陀梨罗那Dahlilah
}
    
func (c *室毗湿檀那SiwistanCounty) BGuja古奢() feud.Barony {
	return c.古奢Guja
}
    
func (c *室毗湿檀那SiwistanCounty) BKamharia甘诃梨() feud.Barony {
	return c.甘诃梨Kamharia
}
    
func (c *室毗湿檀那SiwistanCounty) BMir_rukan弥尔楼乾() feud.Barony {
	return c.弥尔楼乾Mir_rukan
}
    
func (c *室毗湿檀那SiwistanCounty) BSiwistan室毗湿檀那() feud.Barony {
	return c.室毗湿檀那Siwistan
}
    
func (c *室毗湿檀那SiwistanCounty) BVicalapura毗遮罗补罗() feud.Barony {
	return c.毗遮罗补罗Vicalapura
}
    
var CSiwistan室毗湿檀那 SiwistanCounty = &室毗湿檀那SiwistanCounty{}

func init() {
	f := CSiwistan室毗湿檀那.(*室毗湿檀那SiwistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1333",
		Title:     "siwistan",
		TitleName: "室毗湿檀那",
		TitleCode: "c_siwistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆醯罗Bahera = BBahera婆醯罗
	f.婆醯罗Bahera.SetParent(f)
	
	f.陀梨罗那Dahlilah = BDahlilah陀梨罗那
	f.陀梨罗那Dahlilah.SetParent(f)
	
	f.古奢Guja = BGuja古奢
	f.古奢Guja.SetParent(f)
	
	f.甘诃梨Kamharia = BKamharia甘诃梨
	f.甘诃梨Kamharia.SetParent(f)
	
	f.弥尔楼乾Mir_rukan = BMir_rukan弥尔楼乾
	f.弥尔楼乾Mir_rukan.SetParent(f)
	
	f.室毗湿檀那Siwistan = BSiwistan室毗湿檀那
	f.室毗湿檀那Siwistan.SetParent(f)
	
	f.毗遮罗补罗Vicalapura = BVicalapura毗遮罗补罗
	f.毗遮罗补罗Vicalapura.SetParent(f)
	
}
