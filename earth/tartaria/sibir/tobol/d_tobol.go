package tobol

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/tobol/miass"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/tobol/tavda"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/tobol/tobol"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/tobol/tyumen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TobolDuke interface {
    feud.Duke
    CMiass米阿斯() 	miass.MiassCounty
    CTavda塔夫达() 	tavda.TavdaCounty
    CTobol托博尔() 	tobol.TobolCounty
    CTyumen成吉_图拉() 	tyumen.TyumenCounty
}

type 托博尔TobolDuke struct {
	feud.BaseDuke
	米阿斯Miass 	miass.MiassCounty
	塔夫达Tavda 	tavda.TavdaCounty
	托博尔Tobol 	tobol.TobolCounty
	成吉_图拉Tyumen 	tyumen.TyumenCounty
}

func (d *托博尔TobolDuke) CMiass米阿斯() miass.MiassCounty {
	return d.米阿斯Miass
}
    
func (d *托博尔TobolDuke) CTavda塔夫达() tavda.TavdaCounty {
	return d.塔夫达Tavda
}
    
func (d *托博尔TobolDuke) CTobol托博尔() tobol.TobolCounty {
	return d.托博尔Tobol
}
    
func (d *托博尔TobolDuke) CTyumen成吉_图拉() tyumen.TyumenCounty {
	return d.成吉_图拉Tyumen
}
    
var DTobol托博尔 TobolDuke = &托博尔TobolDuke{}

func init() {
	f := DTobol托博尔.(*托博尔TobolDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tobol",
		TitleName: "托博尔",
		TitleCode: "d_tobol",
		Counties:  map[string]feud.County{},
	}

	f.米阿斯Miass = miass.CMiass米阿斯
	f.米阿斯Miass.SetParent(f)
	
	f.塔夫达Tavda = tavda.CTavda塔夫达
	f.塔夫达Tavda.SetParent(f)
	
	f.托博尔Tobol = tobol.CTobol托博尔
	f.托博尔Tobol.SetParent(f)
	
	f.成吉_图拉Tyumen = tyumen.CTyumen成吉_图拉
	f.成吉_图拉Tyumen.SetParent(f)
	
}
