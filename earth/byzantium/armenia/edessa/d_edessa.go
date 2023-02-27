package edessa

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/edessa/aintab"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/edessa/edessa"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/edessa/tell_bashir"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EdessaDuke interface {
    feud.Duke
    CAintab艾因塔卜() 	aintab.AintabCounty
    CEdessa埃德萨() 	edessa.EdessaCounty
    CTell_bashir巴希尔丘() 	tell_bashir.Tell_bashirCounty
}

type 埃德萨EdessaDuke struct {
	feud.BaseDuke
	艾因塔卜Aintab 	aintab.AintabCounty
	埃德萨Edessa 	edessa.EdessaCounty
	巴希尔丘Tell_bashir 	tell_bashir.Tell_bashirCounty
}

func (d *埃德萨EdessaDuke) CAintab艾因塔卜() aintab.AintabCounty {
	return d.艾因塔卜Aintab
}
    
func (d *埃德萨EdessaDuke) CEdessa埃德萨() edessa.EdessaCounty {
	return d.埃德萨Edessa
}
    
func (d *埃德萨EdessaDuke) CTell_bashir巴希尔丘() tell_bashir.Tell_bashirCounty {
	return d.巴希尔丘Tell_bashir
}
    
var DEdessa埃德萨 EdessaDuke = &埃德萨EdessaDuke{}

func init() {
	f := DEdessa埃德萨.(*埃德萨EdessaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "edessa",
		TitleName: "埃德萨",
		TitleCode: "d_edessa",
		Counties:  map[string]feud.County{},
	}

	f.艾因塔卜Aintab = aintab.CAintab艾因塔卜
	f.艾因塔卜Aintab.SetParent(f)
	
	f.埃德萨Edessa = edessa.CEdessa埃德萨
	f.埃德萨Edessa.SetParent(f)
	
	f.巴希尔丘Tell_bashir = tell_bashir.CTell_bashir巴希尔丘
	f.巴希尔丘Tell_bashir.SetParent(f)
	
}
