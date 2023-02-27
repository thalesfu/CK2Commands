package dege

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/dege/dege"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/dege/lhatok"
	"github.com/thalesfu/CK2Commands/earth/tibet/kham/dege/markam"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DegeDuke interface {
    feud.Duke
    CDege德格() 	dege.DegeCounty
    CLhatok拉多() 	lhatok.LhatokCounty
    CMarkam马儿敢() 	markam.MarkamCounty
}

type 德格DegeDuke struct {
	feud.BaseDuke
	德格Dege 	dege.DegeCounty
	拉多Lhatok 	lhatok.LhatokCounty
	马儿敢Markam 	markam.MarkamCounty
}

func (d *德格DegeDuke) CDege德格() dege.DegeCounty {
	return d.德格Dege
}
    
func (d *德格DegeDuke) CLhatok拉多() lhatok.LhatokCounty {
	return d.拉多Lhatok
}
    
func (d *德格DegeDuke) CMarkam马儿敢() markam.MarkamCounty {
	return d.马儿敢Markam
}
    
var DDege德格 DegeDuke = &德格DegeDuke{}

func init() {
	f := DDege德格.(*德格DegeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dege",
		TitleName: "德格",
		TitleCode: "d_dege",
		Counties:  map[string]feud.County{},
	}

	f.德格Dege = dege.CDege德格
	f.德格Dege.SetParent(f)
	
	f.拉多Lhatok = lhatok.CLhatok拉多
	f.拉多Lhatok.SetParent(f)
	
	f.马儿敢Markam = markam.CMarkam马儿敢
	f.马儿敢Markam.SetParent(f)
	
}
