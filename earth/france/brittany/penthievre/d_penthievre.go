package penthievre

import (
	"github.com/thalesfu/CK2Commands/earth/france/brittany/penthievre/penthievre"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/penthievre/tregor"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PenthievreDuke interface {
    feud.Duke
    CPenthievre庞蒂耶夫尔() 	penthievre.PenthievreCounty
    CTregor特雷戈尔() 	tregor.TregorCounty
}

type 庞蒂耶夫尔PenthievreDuke struct {
	feud.BaseDuke
	庞蒂耶夫尔Penthievre 	penthievre.PenthievreCounty
	特雷戈尔Tregor 	tregor.TregorCounty
}

func (d *庞蒂耶夫尔PenthievreDuke) CPenthievre庞蒂耶夫尔() penthievre.PenthievreCounty {
	return d.庞蒂耶夫尔Penthievre
}
    
func (d *庞蒂耶夫尔PenthievreDuke) CTregor特雷戈尔() tregor.TregorCounty {
	return d.特雷戈尔Tregor
}
    
var DPenthievre庞蒂耶夫尔 PenthievreDuke = &庞蒂耶夫尔PenthievreDuke{}

func init() {
	f := DPenthievre庞蒂耶夫尔.(*庞蒂耶夫尔PenthievreDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "penthievre",
		TitleName: "庞蒂耶夫尔",
		TitleCode: "d_penthievre",
		Counties:  map[string]feud.County{},
	}

	f.庞蒂耶夫尔Penthievre = penthievre.CPenthievre庞蒂耶夫尔
	f.庞蒂耶夫尔Penthievre.SetParent(f)
	
	f.特雷戈尔Tregor = tregor.CTregor特雷戈尔
	f.特雷戈尔Tregor.SetParent(f)
	
}
