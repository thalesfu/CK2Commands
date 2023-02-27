package kiranapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KiranapuraCounty interface {
    feud.County
    BBalaghat婆罗伽吒() 	feud.Barony
    BDorkang多尔康() 	feud.Barony
    BDurgi补耆() 	feud.Barony
    BGangrar乾罗() 	feud.Barony
    BKankauli乾乔梨() 	feud.Barony
    BKiranapura日光城() 	feud.Barony
    BLanji兰基() 	feud.Barony
}

type 日光城KiranapuraCounty struct {
	feud.BaseCounty
	婆罗伽吒Balaghat 	feud.Barony
	多尔康Dorkang 	feud.Barony
	补耆Durgi 	feud.Barony
	乾罗Gangrar 	feud.Barony
	乾乔梨Kankauli 	feud.Barony
	日光城Kiranapura 	feud.Barony
	兰基Lanji 	feud.Barony
}

func (c *日光城KiranapuraCounty) BBalaghat婆罗伽吒() feud.Barony {
	return c.婆罗伽吒Balaghat
}
    
func (c *日光城KiranapuraCounty) BDorkang多尔康() feud.Barony {
	return c.多尔康Dorkang
}
    
func (c *日光城KiranapuraCounty) BDurgi补耆() feud.Barony {
	return c.补耆Durgi
}
    
func (c *日光城KiranapuraCounty) BGangrar乾罗() feud.Barony {
	return c.乾罗Gangrar
}
    
func (c *日光城KiranapuraCounty) BKankauli乾乔梨() feud.Barony {
	return c.乾乔梨Kankauli
}
    
func (c *日光城KiranapuraCounty) BKiranapura日光城() feud.Barony {
	return c.日光城Kiranapura
}
    
func (c *日光城KiranapuraCounty) BLanji兰基() feud.Barony {
	return c.兰基Lanji
}
    
var CKiranapura日光城 KiranapuraCounty = &日光城KiranapuraCounty{}

func init() {
	f := CKiranapura日光城.(*日光城KiranapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1270",
		Title:     "kiranapura",
		TitleName: "日光城",
		TitleCode: "c_kiranapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗伽吒Balaghat = BBalaghat婆罗伽吒
	f.婆罗伽吒Balaghat.SetParent(f)
	
	f.多尔康Dorkang = BDorkang多尔康
	f.多尔康Dorkang.SetParent(f)
	
	f.补耆Durgi = BDurgi补耆
	f.补耆Durgi.SetParent(f)
	
	f.乾罗Gangrar = BGangrar乾罗
	f.乾罗Gangrar.SetParent(f)
	
	f.乾乔梨Kankauli = BKankauli乾乔梨
	f.乾乔梨Kankauli.SetParent(f)
	
	f.日光城Kiranapura = BKiranapura日光城
	f.日光城Kiranapura.SetParent(f)
	
	f.兰基Lanji = BLanji兰基
	f.兰基Lanji.SetParent(f)
	
}
