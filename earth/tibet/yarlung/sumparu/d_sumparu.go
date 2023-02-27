package sumparu

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/sumparu/amdo"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/sumparu/arjin"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/sumparu/qangtang"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/sumparu/samtho"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SumparuDuke interface {
    feud.Duke
    CAmdo安多() 	amdo.AmdoCounty
    CArjin阿尔金() 	arjin.ArjinCounty
    CQangtang羌塘() 	qangtang.QangtangCounty
    CSamtho萨托() 	samtho.SamthoCounty
}

type 孙波如SumparuDuke struct {
	feud.BaseDuke
	安多Amdo 	amdo.AmdoCounty
	阿尔金Arjin 	arjin.ArjinCounty
	羌塘Qangtang 	qangtang.QangtangCounty
	萨托Samtho 	samtho.SamthoCounty
}

func (d *孙波如SumparuDuke) CAmdo安多() amdo.AmdoCounty {
	return d.安多Amdo
}
    
func (d *孙波如SumparuDuke) CArjin阿尔金() arjin.ArjinCounty {
	return d.阿尔金Arjin
}
    
func (d *孙波如SumparuDuke) CQangtang羌塘() qangtang.QangtangCounty {
	return d.羌塘Qangtang
}
    
func (d *孙波如SumparuDuke) CSamtho萨托() samtho.SamthoCounty {
	return d.萨托Samtho
}
    
var DSumparu孙波如 SumparuDuke = &孙波如SumparuDuke{}

func init() {
	f := DSumparu孙波如.(*孙波如SumparuDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sumparu",
		TitleName: "孙波如",
		TitleCode: "d_sumparu",
		Counties:  map[string]feud.County{},
	}

	f.安多Amdo = amdo.CAmdo安多
	f.安多Amdo.SetParent(f)
	
	f.阿尔金Arjin = arjin.CArjin阿尔金
	f.阿尔金Arjin.SetParent(f)
	
	f.羌塘Qangtang = qangtang.CQangtang羌塘
	f.羌塘Qangtang.SetParent(f)
	
	f.萨托Samtho = samtho.CSamtho萨托
	f.萨托Samtho.SetParent(f)
	
}
