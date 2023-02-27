package barcelona

import (
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/barcelona"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/empuries"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/lleida"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/rosello"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/tarragona"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona/urgell"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarcelonaDuke interface {
    feud.Duke
    CBarcelona巴塞罗那() 	barcelona.BarcelonaCounty
    CEmpuries恩普里耶斯() 	empuries.EmpuriesCounty
    CLleida列伊达() 	lleida.LleidaCounty
    CRosello罗塞约() 	rosello.RoselloCounty
    CTarragona塔拉戈纳() 	tarragona.TarragonaCounty
    CUrgell乌尔赫尔() 	urgell.UrgellCounty
}

type 巴塞罗那BarcelonaDuke struct {
	feud.BaseDuke
	巴塞罗那Barcelona 	barcelona.BarcelonaCounty
	恩普里耶斯Empuries 	empuries.EmpuriesCounty
	列伊达Lleida 	lleida.LleidaCounty
	罗塞约Rosello 	rosello.RoselloCounty
	塔拉戈纳Tarragona 	tarragona.TarragonaCounty
	乌尔赫尔Urgell 	urgell.UrgellCounty
}

func (d *巴塞罗那BarcelonaDuke) CBarcelona巴塞罗那() barcelona.BarcelonaCounty {
	return d.巴塞罗那Barcelona
}
    
func (d *巴塞罗那BarcelonaDuke) CEmpuries恩普里耶斯() empuries.EmpuriesCounty {
	return d.恩普里耶斯Empuries
}
    
func (d *巴塞罗那BarcelonaDuke) CLleida列伊达() lleida.LleidaCounty {
	return d.列伊达Lleida
}
    
func (d *巴塞罗那BarcelonaDuke) CRosello罗塞约() rosello.RoselloCounty {
	return d.罗塞约Rosello
}
    
func (d *巴塞罗那BarcelonaDuke) CTarragona塔拉戈纳() tarragona.TarragonaCounty {
	return d.塔拉戈纳Tarragona
}
    
func (d *巴塞罗那BarcelonaDuke) CUrgell乌尔赫尔() urgell.UrgellCounty {
	return d.乌尔赫尔Urgell
}
    
var DBarcelona巴塞罗那 BarcelonaDuke = &巴塞罗那BarcelonaDuke{}

func init() {
	f := DBarcelona巴塞罗那.(*巴塞罗那BarcelonaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "barcelona",
		TitleName: "巴塞罗那",
		TitleCode: "d_barcelona",
		Counties:  map[string]feud.County{},
	}

	f.巴塞罗那Barcelona = barcelona.CBarcelona巴塞罗那
	f.巴塞罗那Barcelona.SetParent(f)
	
	f.恩普里耶斯Empuries = empuries.CEmpuries恩普里耶斯
	f.恩普里耶斯Empuries.SetParent(f)
	
	f.列伊达Lleida = lleida.CLleida列伊达
	f.列伊达Lleida.SetParent(f)
	
	f.罗塞约Rosello = rosello.CRosello罗塞约
	f.罗塞约Rosello.SetParent(f)
	
	f.塔拉戈纳Tarragona = tarragona.CTarragona塔拉戈纳
	f.塔拉戈纳Tarragona.SetParent(f)
	
	f.乌尔赫尔Urgell = urgell.CUrgell乌尔赫尔
	f.乌尔赫尔Urgell.SetParent(f)
	
}
