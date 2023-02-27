package sarkel

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel/chortitza"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel/don_portage"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel/sarkel"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel/sharukan"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel/sugrov"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarkelDuke interface {
    feud.Duke
    CChortitza霍尔季察() 	chortitza.ChortitzaCounty
    CDon_portage顿河水冲() 	don_portage.Don_portageCounty
    CSarkel萨克尔() 	sarkel.SarkelCounty
    CSharukan沙鲁坎() 	sharukan.SharukanCounty
    CSugrov苏格罗夫() 	sugrov.SugrovCounty
}

type 萨克尔SarkelDuke struct {
	feud.BaseDuke
	霍尔季察Chortitza 	chortitza.ChortitzaCounty
	顿河水冲Don_portage 	don_portage.Don_portageCounty
	萨克尔Sarkel 	sarkel.SarkelCounty
	沙鲁坎Sharukan 	sharukan.SharukanCounty
	苏格罗夫Sugrov 	sugrov.SugrovCounty
}

func (d *萨克尔SarkelDuke) CChortitza霍尔季察() chortitza.ChortitzaCounty {
	return d.霍尔季察Chortitza
}
    
func (d *萨克尔SarkelDuke) CDon_portage顿河水冲() don_portage.Don_portageCounty {
	return d.顿河水冲Don_portage
}
    
func (d *萨克尔SarkelDuke) CSarkel萨克尔() sarkel.SarkelCounty {
	return d.萨克尔Sarkel
}
    
func (d *萨克尔SarkelDuke) CSharukan沙鲁坎() sharukan.SharukanCounty {
	return d.沙鲁坎Sharukan
}
    
func (d *萨克尔SarkelDuke) CSugrov苏格罗夫() sugrov.SugrovCounty {
	return d.苏格罗夫Sugrov
}
    
var DSarkel萨克尔 SarkelDuke = &萨克尔SarkelDuke{}

func init() {
	f := DSarkel萨克尔.(*萨克尔SarkelDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sarkel",
		TitleName: "萨克尔",
		TitleCode: "d_sarkel",
		Counties:  map[string]feud.County{},
	}

	f.霍尔季察Chortitza = chortitza.CChortitza霍尔季察
	f.霍尔季察Chortitza.SetParent(f)
	
	f.顿河水冲Don_portage = don_portage.CDon_portage顿河水冲
	f.顿河水冲Don_portage.SetParent(f)
	
	f.萨克尔Sarkel = sarkel.CSarkel萨克尔
	f.萨克尔Sarkel.SetParent(f)
	
	f.沙鲁坎Sharukan = sharukan.CSharukan沙鲁坎
	f.沙鲁坎Sharukan.SetParent(f)
	
	f.苏格罗夫Sugrov = sugrov.CSugrov苏格罗夫
	f.苏格罗夫Sugrov.SetParent(f)
	
}
