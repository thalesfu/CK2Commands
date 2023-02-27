package achaia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/achaia/achaia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/achaia/korinthos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/achaia/methone"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/achaia/monemvasia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AchaiaDuke interface {
    feud.Duke
    CAchaia亚该亚() 	achaia.AchaiaCounty
    CKorinthos科林斯() 	korinthos.KorinthosCounty
    CMethone墨托涅() 	methone.MethoneCounty
    CMonemvasia莫奈姆瓦夏() 	monemvasia.MonemvasiaCounty
}

type 亚该亚AchaiaDuke struct {
	feud.BaseDuke
	亚该亚Achaia 	achaia.AchaiaCounty
	科林斯Korinthos 	korinthos.KorinthosCounty
	墨托涅Methone 	methone.MethoneCounty
	莫奈姆瓦夏Monemvasia 	monemvasia.MonemvasiaCounty
}

func (d *亚该亚AchaiaDuke) CAchaia亚该亚() achaia.AchaiaCounty {
	return d.亚该亚Achaia
}
    
func (d *亚该亚AchaiaDuke) CKorinthos科林斯() korinthos.KorinthosCounty {
	return d.科林斯Korinthos
}
    
func (d *亚该亚AchaiaDuke) CMethone墨托涅() methone.MethoneCounty {
	return d.墨托涅Methone
}
    
func (d *亚该亚AchaiaDuke) CMonemvasia莫奈姆瓦夏() monemvasia.MonemvasiaCounty {
	return d.莫奈姆瓦夏Monemvasia
}
    
var DAchaia亚该亚 AchaiaDuke = &亚该亚AchaiaDuke{}

func init() {
	f := DAchaia亚该亚.(*亚该亚AchaiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "achaia",
		TitleName: "亚该亚",
		TitleCode: "d_achaia",
		Counties:  map[string]feud.County{},
	}

	f.亚该亚Achaia = achaia.CAchaia亚该亚
	f.亚该亚Achaia.SetParent(f)
	
	f.科林斯Korinthos = korinthos.CKorinthos科林斯
	f.科林斯Korinthos.SetParent(f)
	
	f.墨托涅Methone = methone.CMethone墨托涅
	f.墨托涅Methone.SetParent(f)
	
	f.莫奈姆瓦夏Monemvasia = monemvasia.CMonemvasia莫奈姆瓦夏
	f.莫奈姆瓦夏Monemvasia.SetParent(f)
	
}
