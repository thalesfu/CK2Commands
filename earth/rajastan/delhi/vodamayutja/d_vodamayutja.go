package vodamayutja

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/vodamayutja/katehar"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/vodamayutja/sambhal"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/vodamayutja/vodamayutja"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VodamayutjaDuke interface {
    feud.Duke
    CKatehar揭谛诃罗() 	katehar.KateharCounty
    CSambhal三婆罗() 	sambhal.SambhalCounty
    CVodamayutja菩陀摩瑜多() 	vodamayutja.VodamayutjaCounty
}

type 菩陀摩瑜多VodamayutjaDuke struct {
	feud.BaseDuke
	揭谛诃罗Katehar 	katehar.KateharCounty
	三婆罗Sambhal 	sambhal.SambhalCounty
	菩陀摩瑜多Vodamayutja 	vodamayutja.VodamayutjaCounty
}

func (d *菩陀摩瑜多VodamayutjaDuke) CKatehar揭谛诃罗() katehar.KateharCounty {
	return d.揭谛诃罗Katehar
}
    
func (d *菩陀摩瑜多VodamayutjaDuke) CSambhal三婆罗() sambhal.SambhalCounty {
	return d.三婆罗Sambhal
}
    
func (d *菩陀摩瑜多VodamayutjaDuke) CVodamayutja菩陀摩瑜多() vodamayutja.VodamayutjaCounty {
	return d.菩陀摩瑜多Vodamayutja
}
    
var DVodamayutja菩陀摩瑜多 VodamayutjaDuke = &菩陀摩瑜多VodamayutjaDuke{}

func init() {
	f := DVodamayutja菩陀摩瑜多.(*菩陀摩瑜多VodamayutjaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vodamayutja",
		TitleName: "菩陀摩瑜多",
		TitleCode: "d_vodamayutja",
		Counties:  map[string]feud.County{},
	}

	f.揭谛诃罗Katehar = katehar.CKatehar揭谛诃罗
	f.揭谛诃罗Katehar.SetParent(f)
	
	f.三婆罗Sambhal = sambhal.CSambhal三婆罗
	f.三婆罗Sambhal.SetParent(f)
	
	f.菩陀摩瑜多Vodamayutja = vodamayutja.CVodamayutja菩陀摩瑜多
	f.菩陀摩瑜多Vodamayutja.SetParent(f)
	
}
