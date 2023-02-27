package pechora

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/pechora/koshma"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/pechora/tobysh"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/pechora/tsilma"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PechoraDuke interface {
    feud.Duke
    CKoshma科什马() 	koshma.KoshmaCounty
    CTobysh托贝什() 	tobysh.TobyshCounty
    CTsilma齐利马() 	tsilma.TsilmaCounty
}

type 伯朝拉PechoraDuke struct {
	feud.BaseDuke
	科什马Koshma 	koshma.KoshmaCounty
	托贝什Tobysh 	tobysh.TobyshCounty
	齐利马Tsilma 	tsilma.TsilmaCounty
}

func (d *伯朝拉PechoraDuke) CKoshma科什马() koshma.KoshmaCounty {
	return d.科什马Koshma
}
    
func (d *伯朝拉PechoraDuke) CTobysh托贝什() tobysh.TobyshCounty {
	return d.托贝什Tobysh
}
    
func (d *伯朝拉PechoraDuke) CTsilma齐利马() tsilma.TsilmaCounty {
	return d.齐利马Tsilma
}
    
var DPechora伯朝拉 PechoraDuke = &伯朝拉PechoraDuke{}

func init() {
	f := DPechora伯朝拉.(*伯朝拉PechoraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pechora",
		TitleName: "伯朝拉",
		TitleCode: "d_pechora",
		Counties:  map[string]feud.County{},
	}

	f.科什马Koshma = koshma.CKoshma科什马
	f.科什马Koshma.SetParent(f)
	
	f.托贝什Tobysh = tobysh.CTobysh托贝什
	f.托贝什Tobysh.SetParent(f)
	
	f.齐利马Tsilma = tsilma.CTsilma齐利马
	f.齐利马Tsilma.SetParent(f)
	
}
