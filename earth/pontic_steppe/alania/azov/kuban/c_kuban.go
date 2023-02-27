package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KubanCounty interface {
    feud.County
    BBeshtau别什陶() 	feud.Barony
    BCoparia科帕里亚() 	feud.Barony
    BKhumar胡马尔() 	feud.Barony
    BKhutor胡托尔() 	feud.Barony
    BKirpili基尔皮利() 	feud.Barony
    BKuban库班() 	feud.Barony
    BPodkumok波德库莫克() 	feud.Barony
    BPsekups普谢库普斯() 	feud.Barony
}

type 库班KubanCounty struct {
	feud.BaseCounty
	别什陶Beshtau 	feud.Barony
	科帕里亚Coparia 	feud.Barony
	胡马尔Khumar 	feud.Barony
	胡托尔Khutor 	feud.Barony
	基尔皮利Kirpili 	feud.Barony
	库班Kuban 	feud.Barony
	波德库莫克Podkumok 	feud.Barony
	普谢库普斯Psekups 	feud.Barony
}

func (c *库班KubanCounty) BBeshtau别什陶() feud.Barony {
	return c.别什陶Beshtau
}
    
func (c *库班KubanCounty) BCoparia科帕里亚() feud.Barony {
	return c.科帕里亚Coparia
}
    
func (c *库班KubanCounty) BKhumar胡马尔() feud.Barony {
	return c.胡马尔Khumar
}
    
func (c *库班KubanCounty) BKhutor胡托尔() feud.Barony {
	return c.胡托尔Khutor
}
    
func (c *库班KubanCounty) BKirpili基尔皮利() feud.Barony {
	return c.基尔皮利Kirpili
}
    
func (c *库班KubanCounty) BKuban库班() feud.Barony {
	return c.库班Kuban
}
    
func (c *库班KubanCounty) BPodkumok波德库莫克() feud.Barony {
	return c.波德库莫克Podkumok
}
    
func (c *库班KubanCounty) BPsekups普谢库普斯() feud.Barony {
	return c.普谢库普斯Psekups
}
    
var CKuban库班 KubanCounty = &库班KubanCounty{}

func init() {
	f := CKuban库班.(*库班KubanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "599",
		Title:     "kuban",
		TitleName: "库班",
		TitleCode: "c_kuban",
		Baronies:  map[string]feud.Barony{},
	}

	f.别什陶Beshtau = BBeshtau别什陶
	f.别什陶Beshtau.SetParent(f)
	
	f.科帕里亚Coparia = BCoparia科帕里亚
	f.科帕里亚Coparia.SetParent(f)
	
	f.胡马尔Khumar = BKhumar胡马尔
	f.胡马尔Khumar.SetParent(f)
	
	f.胡托尔Khutor = BKhutor胡托尔
	f.胡托尔Khutor.SetParent(f)
	
	f.基尔皮利Kirpili = BKirpili基尔皮利
	f.基尔皮利Kirpili.SetParent(f)
	
	f.库班Kuban = BKuban库班
	f.库班Kuban.SetParent(f)
	
	f.波德库莫克Podkumok = BPodkumok波德库莫克
	f.波德库莫克Podkumok.SetParent(f)
	
	f.普谢库普斯Psekups = BPsekups普谢库普斯
	f.普谢库普斯Psekups.SetParent(f)
	
}
