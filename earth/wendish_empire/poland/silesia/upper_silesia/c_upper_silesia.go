package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Upper_silesiaCounty interface {
    feud.County
    BBoleslawiec博莱斯瓦维茨() 	feud.Barony
    BBrzeg布热格() 	feud.Barony
    BHenrykow亨利库夫() 	feud.Barony
    BLegnica莱格尼察() 	feud.Barony
    BNiemcza涅姆恰() 	feud.Barony
    BOlesnica奥莱希尼察() 	feud.Barony
    BWroclaw弗罗茨瓦夫() 	feud.Barony
}

type 弗罗茨瓦夫Upper_silesiaCounty struct {
	feud.BaseCounty
	博莱斯瓦维茨Boleslawiec 	feud.Barony
	布热格Brzeg 	feud.Barony
	亨利库夫Henrykow 	feud.Barony
	莱格尼察Legnica 	feud.Barony
	涅姆恰Niemcza 	feud.Barony
	奥莱希尼察Olesnica 	feud.Barony
	弗罗茨瓦夫Wroclaw 	feud.Barony
}

func (c *弗罗茨瓦夫Upper_silesiaCounty) BBoleslawiec博莱斯瓦维茨() feud.Barony {
	return c.博莱斯瓦维茨Boleslawiec
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BBrzeg布热格() feud.Barony {
	return c.布热格Brzeg
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BHenrykow亨利库夫() feud.Barony {
	return c.亨利库夫Henrykow
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BLegnica莱格尼察() feud.Barony {
	return c.莱格尼察Legnica
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BNiemcza涅姆恰() feud.Barony {
	return c.涅姆恰Niemcza
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BOlesnica奥莱希尼察() feud.Barony {
	return c.奥莱希尼察Olesnica
}
    
func (c *弗罗茨瓦夫Upper_silesiaCounty) BWroclaw弗罗茨瓦夫() feud.Barony {
	return c.弗罗茨瓦夫Wroclaw
}
    
var CUpper_silesia弗罗茨瓦夫 Upper_silesiaCounty = &弗罗茨瓦夫Upper_silesiaCounty{}

func init() {
	f := CUpper_silesia弗罗茨瓦夫.(*弗罗茨瓦夫Upper_silesiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "435",
		Title:     "upper_silesia",
		TitleName: "弗罗茨瓦夫",
		TitleCode: "c_upper_silesia",
		Baronies:  map[string]feud.Barony{},
	}

	f.博莱斯瓦维茨Boleslawiec = BBoleslawiec博莱斯瓦维茨
	f.博莱斯瓦维茨Boleslawiec.SetParent(f)
	
	f.布热格Brzeg = BBrzeg布热格
	f.布热格Brzeg.SetParent(f)
	
	f.亨利库夫Henrykow = BHenrykow亨利库夫
	f.亨利库夫Henrykow.SetParent(f)
	
	f.莱格尼察Legnica = BLegnica莱格尼察
	f.莱格尼察Legnica.SetParent(f)
	
	f.涅姆恰Niemcza = BNiemcza涅姆恰
	f.涅姆恰Niemcza.SetParent(f)
	
	f.奥莱希尼察Olesnica = BOlesnica奥莱希尼察
	f.奥莱希尼察Olesnica.SetParent(f)
	
	f.弗罗茨瓦夫Wroclaw = BWroclaw弗罗茨瓦夫
	f.弗罗茨瓦夫Wroclaw.SetParent(f)
	
}
