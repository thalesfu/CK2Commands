package courland

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/courland/kurs"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/courland/vanema"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CourlandDuke interface {
    feud.Duke
    CKurs库尔泽梅() 	kurs.KursCounty
    CVanema瓦内马() 	vanema.VanemaCounty
}

type 库尔兰CourlandDuke struct {
	feud.BaseDuke
	库尔泽梅Kurs 	kurs.KursCounty
	瓦内马Vanema 	vanema.VanemaCounty
}

func (d *库尔兰CourlandDuke) CKurs库尔泽梅() kurs.KursCounty {
	return d.库尔泽梅Kurs
}
    
func (d *库尔兰CourlandDuke) CVanema瓦内马() vanema.VanemaCounty {
	return d.瓦内马Vanema
}
    
var DCourland库尔兰 CourlandDuke = &库尔兰CourlandDuke{}

func init() {
	f := DCourland库尔兰.(*库尔兰CourlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "courland",
		TitleName: "库尔兰",
		TitleCode: "d_courland",
		Counties:  map[string]feud.County{},
	}

	f.库尔泽梅Kurs = kurs.CKurs库尔泽梅
	f.库尔泽梅Kurs.SetParent(f)
	
	f.瓦内马Vanema = vanema.CVanema瓦内马
	f.瓦内马Vanema.SetParent(f)
	
}
