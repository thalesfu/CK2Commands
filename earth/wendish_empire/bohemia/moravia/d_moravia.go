package moravia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/moravia/brno"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/moravia/olomouc"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/moravia/opava"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/moravia/znojmo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoraviaDuke interface {
    feud.Duke
    CBrno布尔诺() 	brno.BrnoCounty
    COlomouc奥洛穆茨() 	olomouc.OlomoucCounty
    COpava奥帕瓦() 	opava.OpavaCounty
    CZnojmo兹诺伊莫() 	znojmo.ZnojmoCounty
}

type 摩拉维亚MoraviaDuke struct {
	feud.BaseDuke
	布尔诺Brno 	brno.BrnoCounty
	奥洛穆茨Olomouc 	olomouc.OlomoucCounty
	奥帕瓦Opava 	opava.OpavaCounty
	兹诺伊莫Znojmo 	znojmo.ZnojmoCounty
}

func (d *摩拉维亚MoraviaDuke) CBrno布尔诺() brno.BrnoCounty {
	return d.布尔诺Brno
}
    
func (d *摩拉维亚MoraviaDuke) COlomouc奥洛穆茨() olomouc.OlomoucCounty {
	return d.奥洛穆茨Olomouc
}
    
func (d *摩拉维亚MoraviaDuke) COpava奥帕瓦() opava.OpavaCounty {
	return d.奥帕瓦Opava
}
    
func (d *摩拉维亚MoraviaDuke) CZnojmo兹诺伊莫() znojmo.ZnojmoCounty {
	return d.兹诺伊莫Znojmo
}
    
var DMoravia摩拉维亚 MoraviaDuke = &摩拉维亚MoraviaDuke{}

func init() {
	f := DMoravia摩拉维亚.(*摩拉维亚MoraviaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "moravia",
		TitleName: "摩拉维亚",
		TitleCode: "d_moravia",
		Counties:  map[string]feud.County{},
	}

	f.布尔诺Brno = brno.CBrno布尔诺
	f.布尔诺Brno.SetParent(f)
	
	f.奥洛穆茨Olomouc = olomouc.COlomouc奥洛穆茨
	f.奥洛穆茨Olomouc.SetParent(f)
	
	f.奥帕瓦Opava = opava.COpava奥帕瓦
	f.奥帕瓦Opava.SetParent(f)
	
	f.兹诺伊莫Znojmo = znojmo.CZnojmo兹诺伊莫
	f.兹诺伊莫Znojmo.SetParent(f)
	
}
