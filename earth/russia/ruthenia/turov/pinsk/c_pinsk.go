package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PinskCounty interface {
    feud.County
    BBiaroza别廖扎() 	feud.Barony
    BDrahichyn德罗吉钦() 	feud.Barony
    BHantsavichy甘采维奇() 	feud.Barony
    BIvanava伊万诺沃() 	feud.Barony
    BLuninets卢尼涅茨() 	feud.Barony
    BMotal莫托利() 	feud.Barony
    BPinsk平斯克() 	feud.Barony
}

type 平斯克PinskCounty struct {
	feud.BaseCounty
	别廖扎Biaroza 	feud.Barony
	德罗吉钦Drahichyn 	feud.Barony
	甘采维奇Hantsavichy 	feud.Barony
	伊万诺沃Ivanava 	feud.Barony
	卢尼涅茨Luninets 	feud.Barony
	莫托利Motal 	feud.Barony
	平斯克Pinsk 	feud.Barony
}

func (c *平斯克PinskCounty) BBiaroza别廖扎() feud.Barony {
	return c.别廖扎Biaroza
}
    
func (c *平斯克PinskCounty) BDrahichyn德罗吉钦() feud.Barony {
	return c.德罗吉钦Drahichyn
}
    
func (c *平斯克PinskCounty) BHantsavichy甘采维奇() feud.Barony {
	return c.甘采维奇Hantsavichy
}
    
func (c *平斯克PinskCounty) BIvanava伊万诺沃() feud.Barony {
	return c.伊万诺沃Ivanava
}
    
func (c *平斯克PinskCounty) BLuninets卢尼涅茨() feud.Barony {
	return c.卢尼涅茨Luninets
}
    
func (c *平斯克PinskCounty) BMotal莫托利() feud.Barony {
	return c.莫托利Motal
}
    
func (c *平斯克PinskCounty) BPinsk平斯克() feud.Barony {
	return c.平斯克Pinsk
}
    
var CPinsk平斯克 PinskCounty = &平斯克PinskCounty{}

func init() {
	f := CPinsk平斯克.(*平斯克PinskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "548",
		Title:     "pinsk",
		TitleName: "平斯克",
		TitleCode: "c_pinsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别廖扎Biaroza = BBiaroza别廖扎
	f.别廖扎Biaroza.SetParent(f)
	
	f.德罗吉钦Drahichyn = BDrahichyn德罗吉钦
	f.德罗吉钦Drahichyn.SetParent(f)
	
	f.甘采维奇Hantsavichy = BHantsavichy甘采维奇
	f.甘采维奇Hantsavichy.SetParent(f)
	
	f.伊万诺沃Ivanava = BIvanava伊万诺沃
	f.伊万诺沃Ivanava.SetParent(f)
	
	f.卢尼涅茨Luninets = BLuninets卢尼涅茨
	f.卢尼涅茨Luninets.SetParent(f)
	
	f.莫托利Motal = BMotal莫托利
	f.莫托利Motal.SetParent(f)
	
	f.平斯克Pinsk = BPinsk平斯克
	f.平斯克Pinsk.SetParent(f)
	
}
