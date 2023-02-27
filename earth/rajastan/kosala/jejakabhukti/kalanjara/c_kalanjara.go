package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalanjaraCounty interface {
    feud.County
    BBhatta跋多() 	feud.Barony
    BDevendranagar提云陀罗城() 	feud.Barony
    BJayapura阇耶城() 	feud.Barony
    BKalinjar羯陵阇罗() 	feud.Barony
    BKhajuraho渴树罗城() 	feud.Barony
    BLongching长青() 	feud.Barony
    BPanna般那() 	feud.Barony
}

type 迦兰阇罗KalanjaraCounty struct {
	feud.BaseCounty
	跋多Bhatta 	feud.Barony
	提云陀罗城Devendranagar 	feud.Barony
	阇耶城Jayapura 	feud.Barony
	羯陵阇罗Kalinjar 	feud.Barony
	渴树罗城Khajuraho 	feud.Barony
	长青Longching 	feud.Barony
	般那Panna 	feud.Barony
}

func (c *迦兰阇罗KalanjaraCounty) BBhatta跋多() feud.Barony {
	return c.跋多Bhatta
}
    
func (c *迦兰阇罗KalanjaraCounty) BDevendranagar提云陀罗城() feud.Barony {
	return c.提云陀罗城Devendranagar
}
    
func (c *迦兰阇罗KalanjaraCounty) BJayapura阇耶城() feud.Barony {
	return c.阇耶城Jayapura
}
    
func (c *迦兰阇罗KalanjaraCounty) BKalinjar羯陵阇罗() feud.Barony {
	return c.羯陵阇罗Kalinjar
}
    
func (c *迦兰阇罗KalanjaraCounty) BKhajuraho渴树罗城() feud.Barony {
	return c.渴树罗城Khajuraho
}
    
func (c *迦兰阇罗KalanjaraCounty) BLongching长青() feud.Barony {
	return c.长青Longching
}
    
func (c *迦兰阇罗KalanjaraCounty) BPanna般那() feud.Barony {
	return c.般那Panna
}
    
var CKalanjara迦兰阇罗 KalanjaraCounty = &迦兰阇罗KalanjaraCounty{}

func init() {
	f := CKalanjara迦兰阇罗.(*迦兰阇罗KalanjaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1171",
		Title:     "kalanjara",
		TitleName: "迦兰阇罗",
		TitleCode: "c_kalanjara",
		Baronies:  map[string]feud.Barony{},
	}

	f.跋多Bhatta = BBhatta跋多
	f.跋多Bhatta.SetParent(f)
	
	f.提云陀罗城Devendranagar = BDevendranagar提云陀罗城
	f.提云陀罗城Devendranagar.SetParent(f)
	
	f.阇耶城Jayapura = BJayapura阇耶城
	f.阇耶城Jayapura.SetParent(f)
	
	f.羯陵阇罗Kalinjar = BKalinjar羯陵阇罗
	f.羯陵阇罗Kalinjar.SetParent(f)
	
	f.渴树罗城Khajuraho = BKhajuraho渴树罗城
	f.渴树罗城Khajuraho.SetParent(f)
	
	f.长青Longching = BLongching长青
	f.长青Longching.SetParent(f)
	
	f.般那Panna = BPanna般那
	f.般那Panna.SetParent(f)
	
}
