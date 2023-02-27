package meath

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/meath/dublin"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/meath/kildare"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/meath/meath"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/meath/westmeath"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeathDuke interface {
    feud.Duke
    CDublin都柏林() 	dublin.DublinCounty
    CKildare基尔代尔() 	kildare.KildareCounty
    CMeath米斯() 	meath.MeathCounty
    CWestmeath韦斯特米斯() 	westmeath.WestmeathCounty
}

type 米斯MeathDuke struct {
	feud.BaseDuke
	都柏林Dublin 	dublin.DublinCounty
	基尔代尔Kildare 	kildare.KildareCounty
	米斯Meath 	meath.MeathCounty
	韦斯特米斯Westmeath 	westmeath.WestmeathCounty
}

func (d *米斯MeathDuke) CDublin都柏林() dublin.DublinCounty {
	return d.都柏林Dublin
}
    
func (d *米斯MeathDuke) CKildare基尔代尔() kildare.KildareCounty {
	return d.基尔代尔Kildare
}
    
func (d *米斯MeathDuke) CMeath米斯() meath.MeathCounty {
	return d.米斯Meath
}
    
func (d *米斯MeathDuke) CWestmeath韦斯特米斯() westmeath.WestmeathCounty {
	return d.韦斯特米斯Westmeath
}
    
var DMeath米斯 MeathDuke = &米斯MeathDuke{}

func init() {
	f := DMeath米斯.(*米斯MeathDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "meath",
		TitleName: "米斯",
		TitleCode: "d_meath",
		Counties:  map[string]feud.County{},
	}

	f.都柏林Dublin = dublin.CDublin都柏林
	f.都柏林Dublin.SetParent(f)
	
	f.基尔代尔Kildare = kildare.CKildare基尔代尔
	f.基尔代尔Kildare.SetParent(f)
	
	f.米斯Meath = meath.CMeath米斯
	f.米斯Meath.SetParent(f)
	
	f.韦斯特米斯Westmeath = westmeath.CWestmeath韦斯特米斯
	f.韦斯特米斯Westmeath.SetParent(f)
	
}
