package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanyakubjaCounty interface {
    feud.County
    BBithor毗豆罗() 	feud.Barony
    BEtawah伊吒婆() 	feud.Barony
    BJai_chandra阇耶旃陀罗() 	feud.Barony
    BJajmau阇祇牟() 	feud.Barony
    BKampil剑毕离() 	feud.Barony
    BKanyakubja曲女城() 	feud.Barony
    BSwargdwari莎罗伽堕利() 	feud.Barony
}

type 曲女城KanyakubjaCounty struct {
	feud.BaseCounty
	毗豆罗Bithor 	feud.Barony
	伊吒婆Etawah 	feud.Barony
	阇耶旃陀罗Jai_chandra 	feud.Barony
	阇祇牟Jajmau 	feud.Barony
	剑毕离Kampil 	feud.Barony
	曲女城Kanyakubja 	feud.Barony
	莎罗伽堕利Swargdwari 	feud.Barony
}

func (c *曲女城KanyakubjaCounty) BBithor毗豆罗() feud.Barony {
	return c.毗豆罗Bithor
}
    
func (c *曲女城KanyakubjaCounty) BEtawah伊吒婆() feud.Barony {
	return c.伊吒婆Etawah
}
    
func (c *曲女城KanyakubjaCounty) BJai_chandra阇耶旃陀罗() feud.Barony {
	return c.阇耶旃陀罗Jai_chandra
}
    
func (c *曲女城KanyakubjaCounty) BJajmau阇祇牟() feud.Barony {
	return c.阇祇牟Jajmau
}
    
func (c *曲女城KanyakubjaCounty) BKampil剑毕离() feud.Barony {
	return c.剑毕离Kampil
}
    
func (c *曲女城KanyakubjaCounty) BKanyakubja曲女城() feud.Barony {
	return c.曲女城Kanyakubja
}
    
func (c *曲女城KanyakubjaCounty) BSwargdwari莎罗伽堕利() feud.Barony {
	return c.莎罗伽堕利Swargdwari
}
    
var CKanyakubja曲女城 KanyakubjaCounty = &曲女城KanyakubjaCounty{}

func init() {
	f := CKanyakubja曲女城.(*曲女城KanyakubjaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1356",
		Title:     "kanyakubja",
		TitleName: "曲女城",
		TitleCode: "c_kanyakubja",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗豆罗Bithor = BBithor毗豆罗
	f.毗豆罗Bithor.SetParent(f)
	
	f.伊吒婆Etawah = BEtawah伊吒婆
	f.伊吒婆Etawah.SetParent(f)
	
	f.阇耶旃陀罗Jai_chandra = BJai_chandra阇耶旃陀罗
	f.阇耶旃陀罗Jai_chandra.SetParent(f)
	
	f.阇祇牟Jajmau = BJajmau阇祇牟
	f.阇祇牟Jajmau.SetParent(f)
	
	f.剑毕离Kampil = BKampil剑毕离
	f.剑毕离Kampil.SetParent(f)
	
	f.曲女城Kanyakubja = BKanyakubja曲女城
	f.曲女城Kanyakubja.SetParent(f)
	
	f.莎罗伽堕利Swargdwari = BSwargdwari莎罗伽堕利
	f.莎罗伽堕利Swargdwari.SetParent(f)
	
}
