package york

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/york/hull"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/york/leeds"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/york/yoredale"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/york/york"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YorkDuke interface {
    feud.Duke
    CHull赫尔() 	hull.HullCounty
    CLeeds利兹() 	leeds.LeedsCounty
    CYoredale约代尔() 	yoredale.YoredaleCounty
    CYork约克() 	york.YorkCounty
}

type 约克YorkDuke struct {
	feud.BaseDuke
	赫尔Hull 	hull.HullCounty
	利兹Leeds 	leeds.LeedsCounty
	约代尔Yoredale 	yoredale.YoredaleCounty
	约克York 	york.YorkCounty
}

func (d *约克YorkDuke) CHull赫尔() hull.HullCounty {
	return d.赫尔Hull
}
    
func (d *约克YorkDuke) CLeeds利兹() leeds.LeedsCounty {
	return d.利兹Leeds
}
    
func (d *约克YorkDuke) CYoredale约代尔() yoredale.YoredaleCounty {
	return d.约代尔Yoredale
}
    
func (d *约克YorkDuke) CYork约克() york.YorkCounty {
	return d.约克York
}
    
var DYork约克 YorkDuke = &约克YorkDuke{}

func init() {
	f := DYork约克.(*约克YorkDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "york",
		TitleName: "约克",
		TitleCode: "d_york",
		Counties:  map[string]feud.County{},
	}

	f.赫尔Hull = hull.CHull赫尔
	f.赫尔Hull.SetParent(f)
	
	f.利兹Leeds = leeds.CLeeds利兹
	f.利兹Leeds.SetParent(f)
	
	f.约代尔Yoredale = yoredale.CYoredale约代尔
	f.约代尔Yoredale.SetParent(f)
	
	f.约克York = york.CYork约克
	f.约克York.SetParent(f)
	
}
