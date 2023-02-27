package valencia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/valencia/castellon"
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/valencia/denia"
	"github.com/thalesfu/CK2Commands/earth/spain/valencia/valencia/valencia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValenciaDuke interface {
    feud.Duke
    CCastellon卡斯特利翁() 	castellon.CastellonCounty
    CDenia德尼亚() 	denia.DeniaCounty
    CValencia瓦伦西亚() 	valencia.ValenciaCounty
}

type 瓦伦西亚ValenciaDuke struct {
	feud.BaseDuke
	卡斯特利翁Castellon 	castellon.CastellonCounty
	德尼亚Denia 	denia.DeniaCounty
	瓦伦西亚Valencia 	valencia.ValenciaCounty
}

func (d *瓦伦西亚ValenciaDuke) CCastellon卡斯特利翁() castellon.CastellonCounty {
	return d.卡斯特利翁Castellon
}
    
func (d *瓦伦西亚ValenciaDuke) CDenia德尼亚() denia.DeniaCounty {
	return d.德尼亚Denia
}
    
func (d *瓦伦西亚ValenciaDuke) CValencia瓦伦西亚() valencia.ValenciaCounty {
	return d.瓦伦西亚Valencia
}
    
var DValencia瓦伦西亚 ValenciaDuke = &瓦伦西亚ValenciaDuke{}

func init() {
	f := DValencia瓦伦西亚.(*瓦伦西亚ValenciaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "valencia",
		TitleName: "瓦伦西亚",
		TitleCode: "d_valencia",
		Counties:  map[string]feud.County{},
	}

	f.卡斯特利翁Castellon = castellon.CCastellon卡斯特利翁
	f.卡斯特利翁Castellon.SetParent(f)
	
	f.德尼亚Denia = denia.CDenia德尼亚
	f.德尼亚Denia.SetParent(f)
	
	f.瓦伦西亚Valencia = valencia.CValencia瓦伦西亚
	f.瓦伦西亚Valencia.SetParent(f)
	
}
