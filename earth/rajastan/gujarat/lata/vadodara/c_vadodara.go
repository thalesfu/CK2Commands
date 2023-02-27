package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VadodaraCounty interface {
    feud.County
    BAnkotakka安拘吒迦() 	feud.Barony
    BDarbhavati达婆伐底() 	feud.Barony
    BKayavarohan迦耶伐卢汉拿() 	feud.Barony
    BLilam梨蓝() 	feud.Barony
    BNasvadi那萨迪() 	feud.Barony
    BVadodara婆度陀罗() 	feud.Barony
    BVatpatrak跋波罗() 	feud.Barony
}

type 婆度陀罗VadodaraCounty struct {
	feud.BaseCounty
	安拘吒迦Ankotakka 	feud.Barony
	达婆伐底Darbhavati 	feud.Barony
	迦耶伐卢汉拿Kayavarohan 	feud.Barony
	梨蓝Lilam 	feud.Barony
	那萨迪Nasvadi 	feud.Barony
	婆度陀罗Vadodara 	feud.Barony
	跋波罗Vatpatrak 	feud.Barony
}

func (c *婆度陀罗VadodaraCounty) BAnkotakka安拘吒迦() feud.Barony {
	return c.安拘吒迦Ankotakka
}
    
func (c *婆度陀罗VadodaraCounty) BDarbhavati达婆伐底() feud.Barony {
	return c.达婆伐底Darbhavati
}
    
func (c *婆度陀罗VadodaraCounty) BKayavarohan迦耶伐卢汉拿() feud.Barony {
	return c.迦耶伐卢汉拿Kayavarohan
}
    
func (c *婆度陀罗VadodaraCounty) BLilam梨蓝() feud.Barony {
	return c.梨蓝Lilam
}
    
func (c *婆度陀罗VadodaraCounty) BNasvadi那萨迪() feud.Barony {
	return c.那萨迪Nasvadi
}
    
func (c *婆度陀罗VadodaraCounty) BVadodara婆度陀罗() feud.Barony {
	return c.婆度陀罗Vadodara
}
    
func (c *婆度陀罗VadodaraCounty) BVatpatrak跋波罗() feud.Barony {
	return c.跋波罗Vatpatrak
}
    
var CVadodara婆度陀罗 VadodaraCounty = &婆度陀罗VadodaraCounty{}

func init() {
	f := CVadodara婆度陀罗.(*婆度陀罗VadodaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1133",
		Title:     "vadodara",
		TitleName: "婆度陀罗",
		TitleCode: "c_vadodara",
		Baronies:  map[string]feud.Barony{},
	}

	f.安拘吒迦Ankotakka = BAnkotakka安拘吒迦
	f.安拘吒迦Ankotakka.SetParent(f)
	
	f.达婆伐底Darbhavati = BDarbhavati达婆伐底
	f.达婆伐底Darbhavati.SetParent(f)
	
	f.迦耶伐卢汉拿Kayavarohan = BKayavarohan迦耶伐卢汉拿
	f.迦耶伐卢汉拿Kayavarohan.SetParent(f)
	
	f.梨蓝Lilam = BLilam梨蓝
	f.梨蓝Lilam.SetParent(f)
	
	f.那萨迪Nasvadi = BNasvadi那萨迪
	f.那萨迪Nasvadi.SetParent(f)
	
	f.婆度陀罗Vadodara = BVadodara婆度陀罗
	f.婆度陀罗Vadodara.SetParent(f)
	
	f.跋波罗Vatpatrak = BVatpatrak跋波罗
	f.跋波罗Vatpatrak.SetParent(f)
	
}
