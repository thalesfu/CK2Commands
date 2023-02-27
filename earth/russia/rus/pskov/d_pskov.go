package pskov

import (
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov/gdov"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov/izborsk"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov/pskov"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov/velikiye_luki"
	"github.com/thalesfu/CK2Commands/earth/russia/rus/pskov/vodi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PskovDuke interface {
    feud.Duke
    CGdov格多夫() 	gdov.GdovCounty
    CIzborsk伊兹博尔斯克() 	izborsk.IzborskCounty
    CPskov普斯科夫() 	pskov.PskovCounty
    CVelikiye_luki卢基() 	velikiye_luki.Velikiye_lukiCounty
    CVodi英格里亚() 	vodi.VodiCounty
}

type 普斯科夫PskovDuke struct {
	feud.BaseDuke
	格多夫Gdov 	gdov.GdovCounty
	伊兹博尔斯克Izborsk 	izborsk.IzborskCounty
	普斯科夫Pskov 	pskov.PskovCounty
	卢基Velikiye_luki 	velikiye_luki.Velikiye_lukiCounty
	英格里亚Vodi 	vodi.VodiCounty
}

func (d *普斯科夫PskovDuke) CGdov格多夫() gdov.GdovCounty {
	return d.格多夫Gdov
}
    
func (d *普斯科夫PskovDuke) CIzborsk伊兹博尔斯克() izborsk.IzborskCounty {
	return d.伊兹博尔斯克Izborsk
}
    
func (d *普斯科夫PskovDuke) CPskov普斯科夫() pskov.PskovCounty {
	return d.普斯科夫Pskov
}
    
func (d *普斯科夫PskovDuke) CVelikiye_luki卢基() velikiye_luki.Velikiye_lukiCounty {
	return d.卢基Velikiye_luki
}
    
func (d *普斯科夫PskovDuke) CVodi英格里亚() vodi.VodiCounty {
	return d.英格里亚Vodi
}
    
var DPskov普斯科夫 PskovDuke = &普斯科夫PskovDuke{}

func init() {
	f := DPskov普斯科夫.(*普斯科夫PskovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "pskov",
		TitleName: "普斯科夫",
		TitleCode: "d_pskov",
		Counties:  map[string]feud.County{},
	}

	f.格多夫Gdov = gdov.CGdov格多夫
	f.格多夫Gdov.SetParent(f)
	
	f.伊兹博尔斯克Izborsk = izborsk.CIzborsk伊兹博尔斯克
	f.伊兹博尔斯克Izborsk.SetParent(f)
	
	f.普斯科夫Pskov = pskov.CPskov普斯科夫
	f.普斯科夫Pskov.SetParent(f)
	
	f.卢基Velikiye_luki = velikiye_luki.CVelikiye_luki卢基
	f.卢基Velikiye_luki.SetParent(f)
	
	f.英格里亚Vodi = vodi.CVodi英格里亚
	f.英格里亚Vodi.SetParent(f)
	
}
