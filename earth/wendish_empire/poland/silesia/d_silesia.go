package silesia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/silesia/cieszyn"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/silesia/lower_silesia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/silesia/opole"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland/silesia/upper_silesia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SilesiaDuke interface {
    feud.Duke
    CCieszyn切申() 	cieszyn.CieszynCounty
    CLower_silesia下西里西亚() 	lower_silesia.Lower_silesiaCounty
    COpole上西里西亚() 	opole.OpoleCounty
    CUpper_silesia弗罗茨瓦夫() 	upper_silesia.Upper_silesiaCounty
}

type 西里西亚SilesiaDuke struct {
	feud.BaseDuke
	切申Cieszyn 	cieszyn.CieszynCounty
	下西里西亚Lower_silesia 	lower_silesia.Lower_silesiaCounty
	上西里西亚Opole 	opole.OpoleCounty
	弗罗茨瓦夫Upper_silesia 	upper_silesia.Upper_silesiaCounty
}

func (d *西里西亚SilesiaDuke) CCieszyn切申() cieszyn.CieszynCounty {
	return d.切申Cieszyn
}
    
func (d *西里西亚SilesiaDuke) CLower_silesia下西里西亚() lower_silesia.Lower_silesiaCounty {
	return d.下西里西亚Lower_silesia
}
    
func (d *西里西亚SilesiaDuke) COpole上西里西亚() opole.OpoleCounty {
	return d.上西里西亚Opole
}
    
func (d *西里西亚SilesiaDuke) CUpper_silesia弗罗茨瓦夫() upper_silesia.Upper_silesiaCounty {
	return d.弗罗茨瓦夫Upper_silesia
}
    
var DSilesia西里西亚 SilesiaDuke = &西里西亚SilesiaDuke{}

func init() {
	f := DSilesia西里西亚.(*西里西亚SilesiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "silesia",
		TitleName: "西里西亚",
		TitleCode: "d_silesia",
		Counties:  map[string]feud.County{},
	}

	f.切申Cieszyn = cieszyn.CCieszyn切申
	f.切申Cieszyn.SetParent(f)
	
	f.下西里西亚Lower_silesia = lower_silesia.CLower_silesia下西里西亚
	f.下西里西亚Lower_silesia.SetParent(f)
	
	f.上西里西亚Opole = opole.COpole上西里西亚
	f.上西里西亚Opole.SetParent(f)
	
	f.弗罗茨瓦夫Upper_silesia = upper_silesia.CUpper_silesia弗罗茨瓦夫
	f.弗罗茨瓦夫Upper_silesia.SetParent(f)
	
}
