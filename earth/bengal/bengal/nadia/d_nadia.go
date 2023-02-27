package nadia

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/nadia/nabadwipa"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/nadia/saptagrama"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/nadia/vijayapura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NadiaDuke interface {
    feud.Duke
    CNabadwipa那婆提鞞波() 	nabadwipa.NabadwipaCounty
    CSaptagrama七村() 	saptagrama.SaptagramaCounty
    CVijayapura毗阇耶城() 	vijayapura.VijayapuraCounty
}

type 那地耶NadiaDuke struct {
	feud.BaseDuke
	那婆提鞞波Nabadwipa 	nabadwipa.NabadwipaCounty
	七村Saptagrama 	saptagrama.SaptagramaCounty
	毗阇耶城Vijayapura 	vijayapura.VijayapuraCounty
}

func (d *那地耶NadiaDuke) CNabadwipa那婆提鞞波() nabadwipa.NabadwipaCounty {
	return d.那婆提鞞波Nabadwipa
}
    
func (d *那地耶NadiaDuke) CSaptagrama七村() saptagrama.SaptagramaCounty {
	return d.七村Saptagrama
}
    
func (d *那地耶NadiaDuke) CVijayapura毗阇耶城() vijayapura.VijayapuraCounty {
	return d.毗阇耶城Vijayapura
}
    
var DNadia那地耶 NadiaDuke = &那地耶NadiaDuke{}

func init() {
	f := DNadia那地耶.(*那地耶NadiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nadia",
		TitleName: "那地耶",
		TitleCode: "d_nadia",
		Counties:  map[string]feud.County{},
	}

	f.那婆提鞞波Nabadwipa = nabadwipa.CNabadwipa那婆提鞞波
	f.那婆提鞞波Nabadwipa.SetParent(f)
	
	f.七村Saptagrama = saptagrama.CSaptagrama七村
	f.七村Saptagrama.SetParent(f)
	
	f.毗阇耶城Vijayapura = vijayapura.CVijayapura毗阇耶城
	f.毗阇耶城Vijayapura.SetParent(f)
	
}
