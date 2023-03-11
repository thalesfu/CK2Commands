package novgorod_seversk

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novgorod_seversk/bryansk"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novgorod_seversk/kromy"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novgorod_seversk/kursk"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/novgorod_seversk/rylsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Novgorod_severskDuke interface {
    feud.Duke
    CBryansk布良斯克() 	bryansk.BryanskCounty
    CKromy克罗梅() 	kromy.KromyCounty
    CKursk库尔斯克() 	kursk.KurskCounty
    CRylsk雷利斯克() 	rylsk.RylskCounty
}

type 布良斯克Novgorod_severskDuke struct {
	feud.BaseDuke
	布良斯克Bryansk 	bryansk.BryanskCounty
	克罗梅Kromy 	kromy.KromyCounty
	库尔斯克Kursk 	kursk.KurskCounty
	雷利斯克Rylsk 	rylsk.RylskCounty
}

func (d *布良斯克Novgorod_severskDuke) CBryansk布良斯克() bryansk.BryanskCounty {
	return d.布良斯克Bryansk
}
    
func (d *布良斯克Novgorod_severskDuke) CKromy克罗梅() kromy.KromyCounty {
	return d.克罗梅Kromy
}
    
func (d *布良斯克Novgorod_severskDuke) CKursk库尔斯克() kursk.KurskCounty {
	return d.库尔斯克Kursk
}
    
func (d *布良斯克Novgorod_severskDuke) CRylsk雷利斯克() rylsk.RylskCounty {
	return d.雷利斯克Rylsk
}
    
var DNovgorod_seversk布良斯克 Novgorod_severskDuke = &布良斯克Novgorod_severskDuke{}

func init() {
	f := DNovgorod_seversk布良斯克.(*布良斯克Novgorod_severskDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "novgorod_seversk",
		TitleName: "布良斯克",
		TitleCode: "d_novgorod-seversk",
		Counties:  map[string]feud.County{},
	}

	f.布良斯克Bryansk = bryansk.CBryansk布良斯克
	f.布良斯克Bryansk.SetParent(f)
	
	f.克罗梅Kromy = kromy.CKromy克罗梅
	f.克罗梅Kromy.SetParent(f)
	
	f.库尔斯克Kursk = kursk.CKursk库尔斯克
	f.库尔斯克Kursk.SetParent(f)
	
	f.雷利斯克Rylsk = rylsk.CRylsk雷利斯克
	f.雷利斯克Rylsk.SetParent(f)
	
}
