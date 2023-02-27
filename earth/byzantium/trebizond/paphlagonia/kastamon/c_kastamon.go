package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KastamonCounty interface {
    feud.County
    BAboniteichos阿沃努蒂霍斯() 	feud.Barony
    BAbonon阿柏农() 	feud.Barony
    BGermanicopolis日尔曼尼科波利斯() 	feud.Barony
    BIonopolis伊俄诺波利斯() 	feud.Barony
    BKastamon卡斯塔莫努() 	feud.Barony
    BOlgassys奥尔加西斯() 	feud.Barony
    BPompeiopolis庞贝波利斯() 	feud.Barony
}

type 卡斯塔莫努KastamonCounty struct {
	feud.BaseCounty
	阿沃努蒂霍斯Aboniteichos 	feud.Barony
	阿柏农Abonon 	feud.Barony
	日尔曼尼科波利斯Germanicopolis 	feud.Barony
	伊俄诺波利斯Ionopolis 	feud.Barony
	卡斯塔莫努Kastamon 	feud.Barony
	奥尔加西斯Olgassys 	feud.Barony
	庞贝波利斯Pompeiopolis 	feud.Barony
}

func (c *卡斯塔莫努KastamonCounty) BAboniteichos阿沃努蒂霍斯() feud.Barony {
	return c.阿沃努蒂霍斯Aboniteichos
}
    
func (c *卡斯塔莫努KastamonCounty) BAbonon阿柏农() feud.Barony {
	return c.阿柏农Abonon
}
    
func (c *卡斯塔莫努KastamonCounty) BGermanicopolis日尔曼尼科波利斯() feud.Barony {
	return c.日尔曼尼科波利斯Germanicopolis
}
    
func (c *卡斯塔莫努KastamonCounty) BIonopolis伊俄诺波利斯() feud.Barony {
	return c.伊俄诺波利斯Ionopolis
}
    
func (c *卡斯塔莫努KastamonCounty) BKastamon卡斯塔莫努() feud.Barony {
	return c.卡斯塔莫努Kastamon
}
    
func (c *卡斯塔莫努KastamonCounty) BOlgassys奥尔加西斯() feud.Barony {
	return c.奥尔加西斯Olgassys
}
    
func (c *卡斯塔莫努KastamonCounty) BPompeiopolis庞贝波利斯() feud.Barony {
	return c.庞贝波利斯Pompeiopolis
}
    
var CKastamon卡斯塔莫努 KastamonCounty = &卡斯塔莫努KastamonCounty{}

func init() {
	f := CKastamon卡斯塔莫努.(*卡斯塔莫努KastamonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1929",
		Title:     "kastamon",
		TitleName: "卡斯塔莫努",
		TitleCode: "c_kastamon",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿沃努蒂霍斯Aboniteichos = BAboniteichos阿沃努蒂霍斯
	f.阿沃努蒂霍斯Aboniteichos.SetParent(f)
	
	f.阿柏农Abonon = BAbonon阿柏农
	f.阿柏农Abonon.SetParent(f)
	
	f.日尔曼尼科波利斯Germanicopolis = BGermanicopolis日尔曼尼科波利斯
	f.日尔曼尼科波利斯Germanicopolis.SetParent(f)
	
	f.伊俄诺波利斯Ionopolis = BIonopolis伊俄诺波利斯
	f.伊俄诺波利斯Ionopolis.SetParent(f)
	
	f.卡斯塔莫努Kastamon = BKastamon卡斯塔莫努
	f.卡斯塔莫努Kastamon.SetParent(f)
	
	f.奥尔加西斯Olgassys = BOlgassys奥尔加西斯
	f.奥尔加西斯Olgassys.SetParent(f)
	
	f.庞贝波利斯Pompeiopolis = BPompeiopolis庞贝波利斯
	f.庞贝波利斯Pompeiopolis.SetParent(f)
	
}
