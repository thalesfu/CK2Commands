package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PittenCounty interface {
    feud.County
    BEdelz埃德尔茨() 	feud.Barony
    BFischau菲绍() 	feud.Barony
    BGloggnitz格洛格尼茨() 	feud.Barony
    BNeunkirchen诺因基兴() 	feud.Barony
    BPitten皮滕() 	feud.Barony
    BTernitz泰尔尼茨() 	feud.Barony
    BWr_neustadt维也纳新城() 	feud.Barony
}

type 皮滕PittenCounty struct {
	feud.BaseCounty
	埃德尔茨Edelz 	feud.Barony
	菲绍Fischau 	feud.Barony
	格洛格尼茨Gloggnitz 	feud.Barony
	诺因基兴Neunkirchen 	feud.Barony
	皮滕Pitten 	feud.Barony
	泰尔尼茨Ternitz 	feud.Barony
	维也纳新城Wr_neustadt 	feud.Barony
}

func (c *皮滕PittenCounty) BEdelz埃德尔茨() feud.Barony {
	return c.埃德尔茨Edelz
}
    
func (c *皮滕PittenCounty) BFischau菲绍() feud.Barony {
	return c.菲绍Fischau
}
    
func (c *皮滕PittenCounty) BGloggnitz格洛格尼茨() feud.Barony {
	return c.格洛格尼茨Gloggnitz
}
    
func (c *皮滕PittenCounty) BNeunkirchen诺因基兴() feud.Barony {
	return c.诺因基兴Neunkirchen
}
    
func (c *皮滕PittenCounty) BPitten皮滕() feud.Barony {
	return c.皮滕Pitten
}
    
func (c *皮滕PittenCounty) BTernitz泰尔尼茨() feud.Barony {
	return c.泰尔尼茨Ternitz
}
    
func (c *皮滕PittenCounty) BWr_neustadt维也纳新城() feud.Barony {
	return c.维也纳新城Wr_neustadt
}
    
var CPitten皮滕 PittenCounty = &皮滕PittenCounty{}

func init() {
	f := CPitten皮滕.(*皮滕PittenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1694",
		Title:     "pitten",
		TitleName: "皮滕",
		TitleCode: "c_pitten",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃德尔茨Edelz = BEdelz埃德尔茨
	f.埃德尔茨Edelz.SetParent(f)
	
	f.菲绍Fischau = BFischau菲绍
	f.菲绍Fischau.SetParent(f)
	
	f.格洛格尼茨Gloggnitz = BGloggnitz格洛格尼茨
	f.格洛格尼茨Gloggnitz.SetParent(f)
	
	f.诺因基兴Neunkirchen = BNeunkirchen诺因基兴
	f.诺因基兴Neunkirchen.SetParent(f)
	
	f.皮滕Pitten = BPitten皮滕
	f.皮滕Pitten.SetParent(f)
	
	f.泰尔尼茨Ternitz = BTernitz泰尔尼茨
	f.泰尔尼茨Ternitz.SetParent(f)
	
	f.维也纳新城Wr_neustadt = BWr_neustadt维也纳新城
	f.维也纳新城Wr_neustadt.SetParent(f)
	
}
