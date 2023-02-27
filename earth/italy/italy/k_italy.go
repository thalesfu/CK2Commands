package italy

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/genoa"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/modena"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/pisa"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/susa"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/toscana"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/verona"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ItalyKingdom interface {
    feud.Kingdom
    DGenoa热那亚() 	genoa.GenoaDuke
    DLombardia米兰() 	lombardia.LombardiaDuke
    DModena摩德纳() 	modena.ModenaDuke
    DPisa比萨() 	pisa.PisaDuke
    DSusa苏萨() 	susa.SusaDuke
    DToscana托斯卡纳() 	toscana.ToscanaDuke
    DVerona维罗纳() 	verona.VeronaDuke
}

type 意大利ItalyKingdom struct {
	feud.BaseKingdom
	热那亚Genoa 	genoa.GenoaDuke
	米兰Lombardia 	lombardia.LombardiaDuke
	摩德纳Modena 	modena.ModenaDuke
	比萨Pisa 	pisa.PisaDuke
	苏萨Susa 	susa.SusaDuke
	托斯卡纳Toscana 	toscana.ToscanaDuke
	维罗纳Verona 	verona.VeronaDuke
}

func (k *意大利ItalyKingdom) DGenoa热那亚() genoa.GenoaDuke {
	return k.热那亚Genoa
}
    
func (k *意大利ItalyKingdom) DLombardia米兰() lombardia.LombardiaDuke {
	return k.米兰Lombardia
}
    
func (k *意大利ItalyKingdom) DModena摩德纳() modena.ModenaDuke {
	return k.摩德纳Modena
}
    
func (k *意大利ItalyKingdom) DPisa比萨() pisa.PisaDuke {
	return k.比萨Pisa
}
    
func (k *意大利ItalyKingdom) DSusa苏萨() susa.SusaDuke {
	return k.苏萨Susa
}
    
func (k *意大利ItalyKingdom) DToscana托斯卡纳() toscana.ToscanaDuke {
	return k.托斯卡纳Toscana
}
    
func (k *意大利ItalyKingdom) DVerona维罗纳() verona.VeronaDuke {
	return k.维罗纳Verona
}
    
var KItaly意大利 ItalyKingdom = &意大利ItalyKingdom{}

func init() {
	f := KItaly意大利.(*意大利ItalyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "italy",
		TitleName: "意大利",
		TitleCode: "k_italy",
		Dukes:  map[string]feud.Duke{},
	}

	f.热那亚Genoa = genoa.DGenoa热那亚
	f.热那亚Genoa.SetParent(f)
	
	f.米兰Lombardia = lombardia.DLombardia米兰
	f.米兰Lombardia.SetParent(f)
	
	f.摩德纳Modena = modena.DModena摩德纳
	f.摩德纳Modena.SetParent(f)
	
	f.比萨Pisa = pisa.DPisa比萨
	f.比萨Pisa.SetParent(f)
	
	f.苏萨Susa = susa.DSusa苏萨
	f.苏萨Susa.SetParent(f)
	
	f.托斯卡纳Toscana = toscana.DToscana托斯卡纳
	f.托斯卡纳Toscana.SetParent(f)
	
	f.维罗纳Verona = verona.DVerona维罗纳
	f.维罗纳Verona.SetParent(f)
	
}
