package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AkroinonCounty interface {
    feud.County
    BAkroinon阿克罗伊农() 	feud.Barony
    BAntiochia_akroinon安提俄基亚() 	feud.Barony
    BApamea_akroinon阿帕美亚() 	feud.Barony
    BCelenaea刻莱奈() 	feud.Barony
    BEumenia尤门尼亚() 	feud.Barony
    BNikopolis_akroinon尼科波利斯() 	feud.Barony
    BSeleucia_akroinon塞琉西亚() 	feud.Barony
}

type 阿克罗伊农AkroinonCounty struct {
	feud.BaseCounty
	阿克罗伊农Akroinon 	feud.Barony
	安提俄基亚Antiochia_akroinon 	feud.Barony
	阿帕美亚Apamea_akroinon 	feud.Barony
	刻莱奈Celenaea 	feud.Barony
	尤门尼亚Eumenia 	feud.Barony
	尼科波利斯Nikopolis_akroinon 	feud.Barony
	塞琉西亚Seleucia_akroinon 	feud.Barony
}

func (c *阿克罗伊农AkroinonCounty) BAkroinon阿克罗伊农() feud.Barony {
	return c.阿克罗伊农Akroinon
}
    
func (c *阿克罗伊农AkroinonCounty) BAntiochia_akroinon安提俄基亚() feud.Barony {
	return c.安提俄基亚Antiochia_akroinon
}
    
func (c *阿克罗伊农AkroinonCounty) BApamea_akroinon阿帕美亚() feud.Barony {
	return c.阿帕美亚Apamea_akroinon
}
    
func (c *阿克罗伊农AkroinonCounty) BCelenaea刻莱奈() feud.Barony {
	return c.刻莱奈Celenaea
}
    
func (c *阿克罗伊农AkroinonCounty) BEumenia尤门尼亚() feud.Barony {
	return c.尤门尼亚Eumenia
}
    
func (c *阿克罗伊农AkroinonCounty) BNikopolis_akroinon尼科波利斯() feud.Barony {
	return c.尼科波利斯Nikopolis_akroinon
}
    
func (c *阿克罗伊农AkroinonCounty) BSeleucia_akroinon塞琉西亚() feud.Barony {
	return c.塞琉西亚Seleucia_akroinon
}
    
var CAkroinon阿克罗伊农 AkroinonCounty = &阿克罗伊农AkroinonCounty{}

func init() {
	f := CAkroinon阿克罗伊农.(*阿克罗伊农AkroinonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1932",
		Title:     "akroinon",
		TitleName: "阿克罗伊农",
		TitleCode: "c_akroinon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克罗伊农Akroinon = BAkroinon阿克罗伊农
	f.阿克罗伊农Akroinon.SetParent(f)
	
	f.安提俄基亚Antiochia_akroinon = BAntiochia_akroinon安提俄基亚
	f.安提俄基亚Antiochia_akroinon.SetParent(f)
	
	f.阿帕美亚Apamea_akroinon = BApamea_akroinon阿帕美亚
	f.阿帕美亚Apamea_akroinon.SetParent(f)
	
	f.刻莱奈Celenaea = BCelenaea刻莱奈
	f.刻莱奈Celenaea.SetParent(f)
	
	f.尤门尼亚Eumenia = BEumenia尤门尼亚
	f.尤门尼亚Eumenia.SetParent(f)
	
	f.尼科波利斯Nikopolis_akroinon = BNikopolis_akroinon尼科波利斯
	f.尼科波利斯Nikopolis_akroinon.SetParent(f)
	
	f.塞琉西亚Seleucia_akroinon = BSeleucia_akroinon塞琉西亚
	f.塞琉西亚Seleucia_akroinon.SetParent(f)
	
}
