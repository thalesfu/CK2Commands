package ajmer

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/ajmer/ajayameru"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/ajmer/gwalior"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/ajmer/ranthambore"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/ajmer/vairata"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AjmerDuke interface {
    feud.Duke
    CAjayameru阿阇耶迷卢() 	ajayameru.AjayameruCounty
    CGwalior伽婆梨耶罗() 	gwalior.GwaliorCounty
    CRanthambore罗那萨担婆补罗() 	ranthambore.RanthamboreCounty
    CVairata吠罗多() 	vairata.VairataCounty
}

type 阿阇耶迷楼AjmerDuke struct {
	feud.BaseDuke
	阿阇耶迷卢Ajayameru 	ajayameru.AjayameruCounty
	伽婆梨耶罗Gwalior 	gwalior.GwaliorCounty
	罗那萨担婆补罗Ranthambore 	ranthambore.RanthamboreCounty
	吠罗多Vairata 	vairata.VairataCounty
}

func (d *阿阇耶迷楼AjmerDuke) CAjayameru阿阇耶迷卢() ajayameru.AjayameruCounty {
	return d.阿阇耶迷卢Ajayameru
}
    
func (d *阿阇耶迷楼AjmerDuke) CGwalior伽婆梨耶罗() gwalior.GwaliorCounty {
	return d.伽婆梨耶罗Gwalior
}
    
func (d *阿阇耶迷楼AjmerDuke) CRanthambore罗那萨担婆补罗() ranthambore.RanthamboreCounty {
	return d.罗那萨担婆补罗Ranthambore
}
    
func (d *阿阇耶迷楼AjmerDuke) CVairata吠罗多() vairata.VairataCounty {
	return d.吠罗多Vairata
}
    
var DAjmer阿阇耶迷楼 AjmerDuke = &阿阇耶迷楼AjmerDuke{}

func init() {
	f := DAjmer阿阇耶迷楼.(*阿阇耶迷楼AjmerDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ajmer",
		TitleName: "阿阇耶迷楼",
		TitleCode: "d_ajmer",
		Counties:  map[string]feud.County{},
	}

	f.阿阇耶迷卢Ajayameru = ajayameru.CAjayameru阿阇耶迷卢
	f.阿阇耶迷卢Ajayameru.SetParent(f)
	
	f.伽婆梨耶罗Gwalior = gwalior.CGwalior伽婆梨耶罗
	f.伽婆梨耶罗Gwalior.SetParent(f)
	
	f.罗那萨担婆补罗Ranthambore = ranthambore.CRanthambore罗那萨担婆补罗
	f.罗那萨担婆补罗Ranthambore.SetParent(f)
	
	f.吠罗多Vairata = vairata.CVairata吠罗多
	f.吠罗多Vairata.SetParent(f)
	
}
