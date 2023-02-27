package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SocotraCounty interface {
    feud.County
    BAsma阿斯马() 	feud.Barony
    BQadub加迪卜() 	feud.Barony
    BQashiu加希乌() 	feud.Barony
    BQualnsiyah古兰西耶() 	feud.Barony
    BSiqirah锡吉拉() 	feud.Barony
    BSteroh斯泰鲁赫() 	feud.Barony
    BTamrida塔姆利达() 	feud.Barony
}

type 索科特拉SocotraCounty struct {
	feud.BaseCounty
	阿斯马Asma 	feud.Barony
	加迪卜Qadub 	feud.Barony
	加希乌Qashiu 	feud.Barony
	古兰西耶Qualnsiyah 	feud.Barony
	锡吉拉Siqirah 	feud.Barony
	斯泰鲁赫Steroh 	feud.Barony
	塔姆利达Tamrida 	feud.Barony
}

func (c *索科特拉SocotraCounty) BAsma阿斯马() feud.Barony {
	return c.阿斯马Asma
}
    
func (c *索科特拉SocotraCounty) BQadub加迪卜() feud.Barony {
	return c.加迪卜Qadub
}
    
func (c *索科特拉SocotraCounty) BQashiu加希乌() feud.Barony {
	return c.加希乌Qashiu
}
    
func (c *索科特拉SocotraCounty) BQualnsiyah古兰西耶() feud.Barony {
	return c.古兰西耶Qualnsiyah
}
    
func (c *索科特拉SocotraCounty) BSiqirah锡吉拉() feud.Barony {
	return c.锡吉拉Siqirah
}
    
func (c *索科特拉SocotraCounty) BSteroh斯泰鲁赫() feud.Barony {
	return c.斯泰鲁赫Steroh
}
    
func (c *索科特拉SocotraCounty) BTamrida塔姆利达() feud.Barony {
	return c.塔姆利达Tamrida
}
    
var CSocotra索科特拉 SocotraCounty = &索科特拉SocotraCounty{}

func init() {
	f := CSocotra索科特拉.(*索科特拉SocotraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1369",
		Title:     "socotra",
		TitleName: "索科特拉",
		TitleCode: "c_socotra",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯马Asma = BAsma阿斯马
	f.阿斯马Asma.SetParent(f)
	
	f.加迪卜Qadub = BQadub加迪卜
	f.加迪卜Qadub.SetParent(f)
	
	f.加希乌Qashiu = BQashiu加希乌
	f.加希乌Qashiu.SetParent(f)
	
	f.古兰西耶Qualnsiyah = BQualnsiyah古兰西耶
	f.古兰西耶Qualnsiyah.SetParent(f)
	
	f.锡吉拉Siqirah = BSiqirah锡吉拉
	f.锡吉拉Siqirah.SetParent(f)
	
	f.斯泰鲁赫Steroh = BSteroh斯泰鲁赫
	f.斯泰鲁赫Steroh.SetParent(f)
	
	f.塔姆利达Tamrida = BTamrida塔姆利达
	f.塔姆利达Tamrida.SetParent(f)
	
}
