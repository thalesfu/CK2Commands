package kastoria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KastoriaCounty interface {
    feud.County
    BArgos_orestikon阿尔戈斯奥雷斯蒂孔() 	feud.Barony
    BChlerinon克勒里农() 	feud.Barony
    BEordaia欧耳代亚() 	feud.Barony
    BKailara凯拉拉() 	feud.Barony
    BKastoria卡斯托里亚() 	feud.Barony
    BMesopotamia_kastoria美索不达米亚() 	feud.Barony
    BOrestis奥雷斯蒂斯() 	feud.Barony
}

type 卡斯托里亚KastoriaCounty struct {
	feud.BaseCounty
	阿尔戈斯奥雷斯蒂孔Argos_orestikon 	feud.Barony
	克勒里农Chlerinon 	feud.Barony
	欧耳代亚Eordaia 	feud.Barony
	凯拉拉Kailara 	feud.Barony
	卡斯托里亚Kastoria 	feud.Barony
	美索不达米亚Mesopotamia_kastoria 	feud.Barony
	奥雷斯蒂斯Orestis 	feud.Barony
}

func (c *卡斯托里亚KastoriaCounty) BArgos_orestikon阿尔戈斯奥雷斯蒂孔() feud.Barony {
	return c.阿尔戈斯奥雷斯蒂孔Argos_orestikon
}
    
func (c *卡斯托里亚KastoriaCounty) BChlerinon克勒里农() feud.Barony {
	return c.克勒里农Chlerinon
}
    
func (c *卡斯托里亚KastoriaCounty) BEordaia欧耳代亚() feud.Barony {
	return c.欧耳代亚Eordaia
}
    
func (c *卡斯托里亚KastoriaCounty) BKailara凯拉拉() feud.Barony {
	return c.凯拉拉Kailara
}
    
func (c *卡斯托里亚KastoriaCounty) BKastoria卡斯托里亚() feud.Barony {
	return c.卡斯托里亚Kastoria
}
    
func (c *卡斯托里亚KastoriaCounty) BMesopotamia_kastoria美索不达米亚() feud.Barony {
	return c.美索不达米亚Mesopotamia_kastoria
}
    
func (c *卡斯托里亚KastoriaCounty) BOrestis奥雷斯蒂斯() feud.Barony {
	return c.奥雷斯蒂斯Orestis
}
    
var CKastoria卡斯托里亚 KastoriaCounty = &卡斯托里亚KastoriaCounty{}

func init() {
	f := CKastoria卡斯托里亚.(*卡斯托里亚KastoriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1883",
		Title:     "kastoria",
		TitleName: "卡斯托里亚",
		TitleCode: "c_kastoria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔戈斯奥雷斯蒂孔Argos_orestikon = BArgos_orestikon阿尔戈斯奥雷斯蒂孔
	f.阿尔戈斯奥雷斯蒂孔Argos_orestikon.SetParent(f)
	
	f.克勒里农Chlerinon = BChlerinon克勒里农
	f.克勒里农Chlerinon.SetParent(f)
	
	f.欧耳代亚Eordaia = BEordaia欧耳代亚
	f.欧耳代亚Eordaia.SetParent(f)
	
	f.凯拉拉Kailara = BKailara凯拉拉
	f.凯拉拉Kailara.SetParent(f)
	
	f.卡斯托里亚Kastoria = BKastoria卡斯托里亚
	f.卡斯托里亚Kastoria.SetParent(f)
	
	f.美索不达米亚Mesopotamia_kastoria = BMesopotamia_kastoria美索不达米亚
	f.美索不达米亚Mesopotamia_kastoria.SetParent(f)
	
	f.奥雷斯蒂斯Orestis = BOrestis奥雷斯蒂斯
	f.奥雷斯蒂斯Orestis.SetParent(f)
	
}
