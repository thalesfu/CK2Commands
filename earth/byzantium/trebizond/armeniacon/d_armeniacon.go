package armeniacon

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/armeniacon/amisos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/armeniacon/sinope"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArmeniaconDuke interface {
    feud.Duke
    CAmisos阿米索斯() 	amisos.AmisosCounty
    CSinope锡诺皮() 	sinope.SinopeCounty
}

type 亚美尼亚坎ArmeniaconDuke struct {
	feud.BaseDuke
	阿米索斯Amisos 	amisos.AmisosCounty
	锡诺皮Sinope 	sinope.SinopeCounty
}

func (d *亚美尼亚坎ArmeniaconDuke) CAmisos阿米索斯() amisos.AmisosCounty {
	return d.阿米索斯Amisos
}
    
func (d *亚美尼亚坎ArmeniaconDuke) CSinope锡诺皮() sinope.SinopeCounty {
	return d.锡诺皮Sinope
}
    
var DArmeniacon亚美尼亚坎 ArmeniaconDuke = &亚美尼亚坎ArmeniaconDuke{}

func init() {
	f := DArmeniacon亚美尼亚坎.(*亚美尼亚坎ArmeniaconDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "armeniacon",
		TitleName: "亚美尼亚坎",
		TitleCode: "d_armeniacon",
		Counties:  map[string]feud.County{},
	}

	f.阿米索斯Amisos = amisos.CAmisos阿米索斯
	f.阿米索斯Amisos.SetParent(f)
	
	f.锡诺皮Sinope = sinope.CSinope锡诺皮
	f.锡诺皮Sinope.SetParent(f)
	
}
