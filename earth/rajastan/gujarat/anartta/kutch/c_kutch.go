package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KutchCounty interface {
    feud.County
    BAnjar阿迦尔() 	feud.Barony
    BBhuj普杰() 	feud.Barony
    BDhaneti陀内提() 	feud.Barony
    BHolavanhalli呼罗槃诃梨() 	feud.Barony
    BHora呼罗() 	feud.Barony
    BHukampura侯伽补罗() 	feud.Barony
    BKanthakota建多拘吒() 	feud.Barony
}

type 契吒KutchCounty struct {
	feud.BaseCounty
	阿迦尔Anjar 	feud.Barony
	普杰Bhuj 	feud.Barony
	陀内提Dhaneti 	feud.Barony
	呼罗槃诃梨Holavanhalli 	feud.Barony
	呼罗Hora 	feud.Barony
	侯伽补罗Hukampura 	feud.Barony
	建多拘吒Kanthakota 	feud.Barony
}

func (c *契吒KutchCounty) BAnjar阿迦尔() feud.Barony {
	return c.阿迦尔Anjar
}
    
func (c *契吒KutchCounty) BBhuj普杰() feud.Barony {
	return c.普杰Bhuj
}
    
func (c *契吒KutchCounty) BDhaneti陀内提() feud.Barony {
	return c.陀内提Dhaneti
}
    
func (c *契吒KutchCounty) BHolavanhalli呼罗槃诃梨() feud.Barony {
	return c.呼罗槃诃梨Holavanhalli
}
    
func (c *契吒KutchCounty) BHora呼罗() feud.Barony {
	return c.呼罗Hora
}
    
func (c *契吒KutchCounty) BHukampura侯伽补罗() feud.Barony {
	return c.侯伽补罗Hukampura
}
    
func (c *契吒KutchCounty) BKanthakota建多拘吒() feud.Barony {
	return c.建多拘吒Kanthakota
}
    
var CKutch契吒 KutchCounty = &契吒KutchCounty{}

func init() {
	f := CKutch契吒.(*契吒KutchCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1295",
		Title:     "kutch",
		TitleName: "契吒",
		TitleCode: "c_kutch",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿迦尔Anjar = BAnjar阿迦尔
	f.阿迦尔Anjar.SetParent(f)
	
	f.普杰Bhuj = BBhuj普杰
	f.普杰Bhuj.SetParent(f)
	
	f.陀内提Dhaneti = BDhaneti陀内提
	f.陀内提Dhaneti.SetParent(f)
	
	f.呼罗槃诃梨Holavanhalli = BHolavanhalli呼罗槃诃梨
	f.呼罗槃诃梨Holavanhalli.SetParent(f)
	
	f.呼罗Hora = BHora呼罗
	f.呼罗Hora.SetParent(f)
	
	f.侯伽补罗Hukampura = BHukampura侯伽补罗
	f.侯伽补罗Hukampura.SetParent(f)
	
	f.建多拘吒Kanthakota = BKanthakota建多拘吒
	f.建多拘吒Kanthakota.SetParent(f)
	
}
