package upper_lorraine

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/upper_lorraine/bar"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/upper_lorraine/lorraine"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/upper_lorraine/saintois"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/upper_lorraine/verdun"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Upper_lorraineDuke interface {
    feud.Duke
    CBar巴鲁瓦() 	bar.BarCounty
    CLorraine洛林() 	lorraine.LorraineCounty
    CSaintois桑图瓦() 	saintois.SaintoisCounty
    CVerdun凡尔登() 	verdun.VerdunCounty
}

type 上洛林Upper_lorraineDuke struct {
	feud.BaseDuke
	巴鲁瓦Bar 	bar.BarCounty
	洛林Lorraine 	lorraine.LorraineCounty
	桑图瓦Saintois 	saintois.SaintoisCounty
	凡尔登Verdun 	verdun.VerdunCounty
}

func (d *上洛林Upper_lorraineDuke) CBar巴鲁瓦() bar.BarCounty {
	return d.巴鲁瓦Bar
}
    
func (d *上洛林Upper_lorraineDuke) CLorraine洛林() lorraine.LorraineCounty {
	return d.洛林Lorraine
}
    
func (d *上洛林Upper_lorraineDuke) CSaintois桑图瓦() saintois.SaintoisCounty {
	return d.桑图瓦Saintois
}
    
func (d *上洛林Upper_lorraineDuke) CVerdun凡尔登() verdun.VerdunCounty {
	return d.凡尔登Verdun
}
    
var DUpper_lorraine上洛林 Upper_lorraineDuke = &上洛林Upper_lorraineDuke{}

func init() {
	f := DUpper_lorraine上洛林.(*上洛林Upper_lorraineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "upper_lorraine",
		TitleName: "上洛林",
		TitleCode: "d_upper_lorraine",
		Counties:  map[string]feud.County{},
	}

	f.巴鲁瓦Bar = bar.CBar巴鲁瓦
	f.巴鲁瓦Bar.SetParent(f)
	
	f.洛林Lorraine = lorraine.CLorraine洛林
	f.洛林Lorraine.SetParent(f)
	
	f.桑图瓦Saintois = saintois.CSaintois桑图瓦
	f.桑图瓦Saintois.SetParent(f)
	
	f.凡尔登Verdun = verdun.CVerdun凡尔登
	f.凡尔登Verdun.SetParent(f)
	
}
