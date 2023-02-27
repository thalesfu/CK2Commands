package novosil

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novosil/novosil"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novosil/serpukhov"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novosil/vorotynsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NovosilDuke interface {
    feud.Duke
    CNovosil诺沃西利() 	novosil.NovosilCounty
    CSerpukhov谢尔普霍夫() 	serpukhov.SerpukhovCounty
    CVorotynsk沃罗滕斯克() 	vorotynsk.VorotynskCounty
}

type 诺沃西利NovosilDuke struct {
	feud.BaseDuke
	诺沃西利Novosil 	novosil.NovosilCounty
	谢尔普霍夫Serpukhov 	serpukhov.SerpukhovCounty
	沃罗滕斯克Vorotynsk 	vorotynsk.VorotynskCounty
}

func (d *诺沃西利NovosilDuke) CNovosil诺沃西利() novosil.NovosilCounty {
	return d.诺沃西利Novosil
}
    
func (d *诺沃西利NovosilDuke) CSerpukhov谢尔普霍夫() serpukhov.SerpukhovCounty {
	return d.谢尔普霍夫Serpukhov
}
    
func (d *诺沃西利NovosilDuke) CVorotynsk沃罗滕斯克() vorotynsk.VorotynskCounty {
	return d.沃罗滕斯克Vorotynsk
}
    
var DNovosil诺沃西利 NovosilDuke = &诺沃西利NovosilDuke{}

func init() {
	f := DNovosil诺沃西利.(*诺沃西利NovosilDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "novosil",
		TitleName: "诺沃西利",
		TitleCode: "d_novosil",
		Counties:  map[string]feud.County{},
	}

	f.诺沃西利Novosil = novosil.CNovosil诺沃西利
	f.诺沃西利Novosil.SetParent(f)
	
	f.谢尔普霍夫Serpukhov = serpukhov.CSerpukhov谢尔普霍夫
	f.谢尔普霍夫Serpukhov.SetParent(f)
	
	f.沃罗滕斯克Vorotynsk = vorotynsk.CVorotynsk沃罗滕斯克
	f.沃罗滕斯克Vorotynsk.SetParent(f)
	
}
