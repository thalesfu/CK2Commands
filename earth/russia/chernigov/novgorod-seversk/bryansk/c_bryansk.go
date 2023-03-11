package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BryanskCounty interface {
    feud.County
    BBezichi别日奇() 	feud.Barony
    BBryansk布良斯克() 	feud.Barony
    BKokino科基诺() 	feud.Barony
    BPochep波切普() 	feud.Barony
    BSuponevo苏波涅沃() 	feud.Barony
    BTrubetsk特鲁别茨克() 	feud.Barony
    BVygonichi维戈尼奇() 	feud.Barony
}

type 布良斯克BryanskCounty struct {
	feud.BaseCounty
	别日奇Bezichi 	feud.Barony
	布良斯克Bryansk 	feud.Barony
	科基诺Kokino 	feud.Barony
	波切普Pochep 	feud.Barony
	苏波涅沃Suponevo 	feud.Barony
	特鲁别茨克Trubetsk 	feud.Barony
	维戈尼奇Vygonichi 	feud.Barony
}

func (c *布良斯克BryanskCounty) BBezichi别日奇() feud.Barony {
	return c.别日奇Bezichi
}
    
func (c *布良斯克BryanskCounty) BBryansk布良斯克() feud.Barony {
	return c.布良斯克Bryansk
}
    
func (c *布良斯克BryanskCounty) BKokino科基诺() feud.Barony {
	return c.科基诺Kokino
}
    
func (c *布良斯克BryanskCounty) BPochep波切普() feud.Barony {
	return c.波切普Pochep
}
    
func (c *布良斯克BryanskCounty) BSuponevo苏波涅沃() feud.Barony {
	return c.苏波涅沃Suponevo
}
    
func (c *布良斯克BryanskCounty) BTrubetsk特鲁别茨克() feud.Barony {
	return c.特鲁别茨克Trubetsk
}
    
func (c *布良斯克BryanskCounty) BVygonichi维戈尼奇() feud.Barony {
	return c.维戈尼奇Vygonichi
}
    
var CBryansk布良斯克 BryanskCounty = &布良斯克BryanskCounty{}

func init() {
	f := CBryansk布良斯克.(*布良斯克BryanskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "576",
		Title:     "bryansk",
		TitleName: "布良斯克",
		TitleCode: "c_bryansk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别日奇Bezichi = BBezichi别日奇
	f.别日奇Bezichi.SetParent(f)
	
	f.布良斯克Bryansk = BBryansk布良斯克
	f.布良斯克Bryansk.SetParent(f)
	
	f.科基诺Kokino = BKokino科基诺
	f.科基诺Kokino.SetParent(f)
	
	f.波切普Pochep = BPochep波切普
	f.波切普Pochep.SetParent(f)
	
	f.苏波涅沃Suponevo = BSuponevo苏波涅沃
	f.苏波涅沃Suponevo.SetParent(f)
	
	f.特鲁别茨克Trubetsk = BTrubetsk特鲁别茨克
	f.特鲁别茨克Trubetsk.SetParent(f)
	
	f.维戈尼奇Vygonichi = BVygonichi维戈尼奇
	f.维戈尼奇Vygonichi.SetParent(f)
	
}
