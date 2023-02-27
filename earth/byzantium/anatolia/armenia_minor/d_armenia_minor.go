package armenia_minor

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/armenia_minor/adana"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/armenia_minor/lykandos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/armenia_minor/seleukeia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/armenia_minor/tarsos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/armenia_minor/teluch"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Armenia_minorDuke interface {
    feud.Duke
    CAdana阿达纳() 	adana.AdanaCounty
    CLykandos吕坎多斯() 	lykandos.LykandosCounty
    CSeleukeia塞琉西亚() 	seleukeia.SeleukeiaCounty
    CTarsos塔索斯() 	tarsos.TarsosCounty
    CTeluch特卢克() 	teluch.TeluchCounty
}

type 奇里乞亚Armenia_minorDuke struct {
	feud.BaseDuke
	阿达纳Adana 	adana.AdanaCounty
	吕坎多斯Lykandos 	lykandos.LykandosCounty
	塞琉西亚Seleukeia 	seleukeia.SeleukeiaCounty
	塔索斯Tarsos 	tarsos.TarsosCounty
	特卢克Teluch 	teluch.TeluchCounty
}

func (d *奇里乞亚Armenia_minorDuke) CAdana阿达纳() adana.AdanaCounty {
	return d.阿达纳Adana
}
    
func (d *奇里乞亚Armenia_minorDuke) CLykandos吕坎多斯() lykandos.LykandosCounty {
	return d.吕坎多斯Lykandos
}
    
func (d *奇里乞亚Armenia_minorDuke) CSeleukeia塞琉西亚() seleukeia.SeleukeiaCounty {
	return d.塞琉西亚Seleukeia
}
    
func (d *奇里乞亚Armenia_minorDuke) CTarsos塔索斯() tarsos.TarsosCounty {
	return d.塔索斯Tarsos
}
    
func (d *奇里乞亚Armenia_minorDuke) CTeluch特卢克() teluch.TeluchCounty {
	return d.特卢克Teluch
}
    
var DArmenia_minor奇里乞亚 Armenia_minorDuke = &奇里乞亚Armenia_minorDuke{}

func init() {
	f := DArmenia_minor奇里乞亚.(*奇里乞亚Armenia_minorDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "armenia_minor",
		TitleName: "奇里乞亚",
		TitleCode: "d_armenia_minor",
		Counties:  map[string]feud.County{},
	}

	f.阿达纳Adana = adana.CAdana阿达纳
	f.阿达纳Adana.SetParent(f)
	
	f.吕坎多斯Lykandos = lykandos.CLykandos吕坎多斯
	f.吕坎多斯Lykandos.SetParent(f)
	
	f.塞琉西亚Seleukeia = seleukeia.CSeleukeia塞琉西亚
	f.塞琉西亚Seleukeia.SetParent(f)
	
	f.塔索斯Tarsos = tarsos.CTarsos塔索斯
	f.塔索斯Tarsos.SetParent(f)
	
	f.特卢克Teluch = teluch.CTeluch特卢克
	f.特卢克Teluch.SetParent(f)
	
}
