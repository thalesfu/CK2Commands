package arjin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArjinCounty interface {
    feud.County
    BAqqikkol阿其克库勒() 	feud.Barony
    BArjin阿尔金() 	feud.Barony
    BAyakkum阿牙克库木() 	feud.Barony
    BKumkol库木库里() 	feud.Barony
    BMuztagh木孜塔格() 	feud.Barony
    BYitun依吞() 	feud.Barony
    BZhotsen雪景() 	feud.Barony
}

type 阿尔金ArjinCounty struct {
	feud.BaseCounty
	阿其克库勒Aqqikkol 	feud.Barony
	阿尔金Arjin 	feud.Barony
	阿牙克库木Ayakkum 	feud.Barony
	库木库里Kumkol 	feud.Barony
	木孜塔格Muztagh 	feud.Barony
	依吞Yitun 	feud.Barony
	雪景Zhotsen 	feud.Barony
}

func (c *阿尔金ArjinCounty) BAqqikkol阿其克库勒() feud.Barony {
	return c.阿其克库勒Aqqikkol
}
    
func (c *阿尔金ArjinCounty) BArjin阿尔金() feud.Barony {
	return c.阿尔金Arjin
}
    
func (c *阿尔金ArjinCounty) BAyakkum阿牙克库木() feud.Barony {
	return c.阿牙克库木Ayakkum
}
    
func (c *阿尔金ArjinCounty) BKumkol库木库里() feud.Barony {
	return c.库木库里Kumkol
}
    
func (c *阿尔金ArjinCounty) BMuztagh木孜塔格() feud.Barony {
	return c.木孜塔格Muztagh
}
    
func (c *阿尔金ArjinCounty) BYitun依吞() feud.Barony {
	return c.依吞Yitun
}
    
func (c *阿尔金ArjinCounty) BZhotsen雪景() feud.Barony {
	return c.雪景Zhotsen
}
    
var CArjin阿尔金 ArjinCounty = &阿尔金ArjinCounty{}

func init() {
	f := CArjin阿尔金.(*阿尔金ArjinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1572",
		Title:     "arjin",
		TitleName: "阿尔金",
		TitleCode: "c_arjin",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿其克库勒Aqqikkol = BAqqikkol阿其克库勒
	f.阿其克库勒Aqqikkol.SetParent(f)
	
	f.阿尔金Arjin = BArjin阿尔金
	f.阿尔金Arjin.SetParent(f)
	
	f.阿牙克库木Ayakkum = BAyakkum阿牙克库木
	f.阿牙克库木Ayakkum.SetParent(f)
	
	f.库木库里Kumkol = BKumkol库木库里
	f.库木库里Kumkol.SetParent(f)
	
	f.木孜塔格Muztagh = BMuztagh木孜塔格
	f.木孜塔格Muztagh.SetParent(f)
	
	f.依吞Yitun = BYitun依吞
	f.依吞Yitun.SetParent(f)
	
	f.雪景Zhotsen = BZhotsen雪景
	f.雪景Zhotsen.SetParent(f)
	
}
