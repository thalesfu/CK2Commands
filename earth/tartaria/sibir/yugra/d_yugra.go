package yugra

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/yugra/khantia"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/yugra/mansia"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/yugra/yamalia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YugraDuke interface {
    feud.Duke
    CKhantia汉特() 	khantia.KhantiaCounty
    CMansia曼西() 	mansia.MansiaCounty
    CYamalia亚马利亚() 	yamalia.YamaliaCounty
}

type 尤格拉YugraDuke struct {
	feud.BaseDuke
	汉特Khantia 	khantia.KhantiaCounty
	曼西Mansia 	mansia.MansiaCounty
	亚马利亚Yamalia 	yamalia.YamaliaCounty
}

func (d *尤格拉YugraDuke) CKhantia汉特() khantia.KhantiaCounty {
	return d.汉特Khantia
}
    
func (d *尤格拉YugraDuke) CMansia曼西() mansia.MansiaCounty {
	return d.曼西Mansia
}
    
func (d *尤格拉YugraDuke) CYamalia亚马利亚() yamalia.YamaliaCounty {
	return d.亚马利亚Yamalia
}
    
var DYugra尤格拉 YugraDuke = &尤格拉YugraDuke{}

func init() {
	f := DYugra尤格拉.(*尤格拉YugraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "yugra",
		TitleName: "尤格拉",
		TitleCode: "d_yugra",
		Counties:  map[string]feud.County{},
	}

	f.汉特Khantia = khantia.CKhantia汉特
	f.汉特Khantia.SetParent(f)
	
	f.曼西Mansia = mansia.CMansia曼西
	f.曼西Mansia.SetParent(f)
	
	f.亚马利亚Yamalia = yamalia.CYamalia亚马利亚
	f.亚马利亚Yamalia.SetParent(f)
	
}
