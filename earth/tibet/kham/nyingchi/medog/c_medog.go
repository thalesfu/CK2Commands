package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedogCounty interface {
    feud.County
    BBaibung背崩() 	feud.Barony
    BDeshing德兴() 	feud.Barony
    BGandain甘登() 	feud.Barony
    BGedang格当() 	feud.Barony
    BMedog墨脱() 	feud.Barony
    BPhomshen旁辛() 	feud.Barony
    BTamu达木() 	feud.Barony
}

type 墨脱MedogCounty struct {
	feud.BaseCounty
	背崩Baibung 	feud.Barony
	德兴Deshing 	feud.Barony
	甘登Gandain 	feud.Barony
	格当Gedang 	feud.Barony
	墨脱Medog 	feud.Barony
	旁辛Phomshen 	feud.Barony
	达木Tamu 	feud.Barony
}

func (c *墨脱MedogCounty) BBaibung背崩() feud.Barony {
	return c.背崩Baibung
}
    
func (c *墨脱MedogCounty) BDeshing德兴() feud.Barony {
	return c.德兴Deshing
}
    
func (c *墨脱MedogCounty) BGandain甘登() feud.Barony {
	return c.甘登Gandain
}
    
func (c *墨脱MedogCounty) BGedang格当() feud.Barony {
	return c.格当Gedang
}
    
func (c *墨脱MedogCounty) BMedog墨脱() feud.Barony {
	return c.墨脱Medog
}
    
func (c *墨脱MedogCounty) BPhomshen旁辛() feud.Barony {
	return c.旁辛Phomshen
}
    
func (c *墨脱MedogCounty) BTamu达木() feud.Barony {
	return c.达木Tamu
}
    
var CMedog墨脱 MedogCounty = &墨脱MedogCounty{}

func init() {
	f := CMedog墨脱.(*墨脱MedogCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1555",
		Title:     "medog",
		TitleName: "墨脱",
		TitleCode: "c_medog",
		Baronies:  map[string]feud.Barony{},
	}

	f.背崩Baibung = BBaibung背崩
	f.背崩Baibung.SetParent(f)
	
	f.德兴Deshing = BDeshing德兴
	f.德兴Deshing.SetParent(f)
	
	f.甘登Gandain = BGandain甘登
	f.甘登Gandain.SetParent(f)
	
	f.格当Gedang = BGedang格当
	f.格当Gedang.SetParent(f)
	
	f.墨脱Medog = BMedog墨脱
	f.墨脱Medog.SetParent(f)
	
	f.旁辛Phomshen = BPhomshen旁辛
	f.旁辛Phomshen.SetParent(f)
	
	f.达木Tamu = BTamu达木
	f.达木Tamu.SetParent(f)
	
}
