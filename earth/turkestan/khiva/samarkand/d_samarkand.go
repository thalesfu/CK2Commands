package samarkand

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/samarkand/bukhara"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/samarkand/nakhshab"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/samarkand/oshrusana"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/samarkand/samarkand"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamarkandDuke interface {
    feud.Duke
    CBukhara布哈拉() 	bukhara.BukharaCounty
    CNakhshab那色波() 	nakhshab.NakhshabCounty
    COshrusana奥什鲁萨纳() 	oshrusana.OshrusanaCounty
    CSamarkand撒马尔罕() 	samarkand.SamarkandCounty
}

type 撒马尔罕SamarkandDuke struct {
	feud.BaseDuke
	布哈拉Bukhara 	bukhara.BukharaCounty
	那色波Nakhshab 	nakhshab.NakhshabCounty
	奥什鲁萨纳Oshrusana 	oshrusana.OshrusanaCounty
	撒马尔罕Samarkand 	samarkand.SamarkandCounty
}

func (d *撒马尔罕SamarkandDuke) CBukhara布哈拉() bukhara.BukharaCounty {
	return d.布哈拉Bukhara
}
    
func (d *撒马尔罕SamarkandDuke) CNakhshab那色波() nakhshab.NakhshabCounty {
	return d.那色波Nakhshab
}
    
func (d *撒马尔罕SamarkandDuke) COshrusana奥什鲁萨纳() oshrusana.OshrusanaCounty {
	return d.奥什鲁萨纳Oshrusana
}
    
func (d *撒马尔罕SamarkandDuke) CSamarkand撒马尔罕() samarkand.SamarkandCounty {
	return d.撒马尔罕Samarkand
}
    
var DSamarkand撒马尔罕 SamarkandDuke = &撒马尔罕SamarkandDuke{}

func init() {
	f := DSamarkand撒马尔罕.(*撒马尔罕SamarkandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "samarkand",
		TitleName: "撒马尔罕",
		TitleCode: "d_samarkand",
		Counties:  map[string]feud.County{},
	}

	f.布哈拉Bukhara = bukhara.CBukhara布哈拉
	f.布哈拉Bukhara.SetParent(f)
	
	f.那色波Nakhshab = nakhshab.CNakhshab那色波
	f.那色波Nakhshab.SetParent(f)
	
	f.奥什鲁萨纳Oshrusana = oshrusana.COshrusana奥什鲁萨纳
	f.奥什鲁萨纳Oshrusana.SetParent(f)
	
	f.撒马尔罕Samarkand = samarkand.CSamarkand撒马尔罕
	f.撒马尔罕Samarkand.SetParent(f)
	
}
