package alania

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/alania/alania"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/alania/kasogs"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/alania/yegorlyk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlaniaDuke interface {
    feud.Duke
    CAlania阿兰尼亚() 	alania.AlaniaCounty
    CKasogs卡索吉亚() 	kasogs.KasogsCounty
    CYegorlyk叶戈尔雷克() 	yegorlyk.YegorlykCounty
}

type 阿兰尼亚AlaniaDuke struct {
	feud.BaseDuke
	阿兰尼亚Alania 	alania.AlaniaCounty
	卡索吉亚Kasogs 	kasogs.KasogsCounty
	叶戈尔雷克Yegorlyk 	yegorlyk.YegorlykCounty
}

func (d *阿兰尼亚AlaniaDuke) CAlania阿兰尼亚() alania.AlaniaCounty {
	return d.阿兰尼亚Alania
}
    
func (d *阿兰尼亚AlaniaDuke) CKasogs卡索吉亚() kasogs.KasogsCounty {
	return d.卡索吉亚Kasogs
}
    
func (d *阿兰尼亚AlaniaDuke) CYegorlyk叶戈尔雷克() yegorlyk.YegorlykCounty {
	return d.叶戈尔雷克Yegorlyk
}
    
var DAlania阿兰尼亚 AlaniaDuke = &阿兰尼亚AlaniaDuke{}

func init() {
	f := DAlania阿兰尼亚.(*阿兰尼亚AlaniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "alania",
		TitleName: "阿兰尼亚",
		TitleCode: "d_alania",
		Counties:  map[string]feud.County{},
	}

	f.阿兰尼亚Alania = alania.CAlania阿兰尼亚
	f.阿兰尼亚Alania.SetParent(f)
	
	f.卡索吉亚Kasogs = kasogs.CKasogs卡索吉亚
	f.卡索吉亚Kasogs.SetParent(f)
	
	f.叶戈尔雷克Yegorlyk = yegorlyk.CYegorlyk叶戈尔雷克
	f.叶戈尔雷克Yegorlyk.SetParent(f)
	
}
