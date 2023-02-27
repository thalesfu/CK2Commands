package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GnieznienskieCounty interface {
    feud.County
    BBialaziema比亚瓦杰马() 	feud.Barony
    BKamien卡缅() 	feud.Barony
    BKcynia克齐尼亚() 	feud.Barony
    BLubiewo卢别沃() 	feud.Barony
    BNaklo纳克沃() 	feud.Barony
    BPila皮瓦() 	feud.Barony
    BPodgaje波德加耶() 	feud.Barony
}

type 纳克沃GnieznienskieCounty struct {
	feud.BaseCounty
	比亚瓦杰马Bialaziema 	feud.Barony
	卡缅Kamien 	feud.Barony
	克齐尼亚Kcynia 	feud.Barony
	卢别沃Lubiewo 	feud.Barony
	纳克沃Naklo 	feud.Barony
	皮瓦Pila 	feud.Barony
	波德加耶Podgaje 	feud.Barony
}

func (c *纳克沃GnieznienskieCounty) BBialaziema比亚瓦杰马() feud.Barony {
	return c.比亚瓦杰马Bialaziema
}
    
func (c *纳克沃GnieznienskieCounty) BKamien卡缅() feud.Barony {
	return c.卡缅Kamien
}
    
func (c *纳克沃GnieznienskieCounty) BKcynia克齐尼亚() feud.Barony {
	return c.克齐尼亚Kcynia
}
    
func (c *纳克沃GnieznienskieCounty) BLubiewo卢别沃() feud.Barony {
	return c.卢别沃Lubiewo
}
    
func (c *纳克沃GnieznienskieCounty) BNaklo纳克沃() feud.Barony {
	return c.纳克沃Naklo
}
    
func (c *纳克沃GnieznienskieCounty) BPila皮瓦() feud.Barony {
	return c.皮瓦Pila
}
    
func (c *纳克沃GnieznienskieCounty) BPodgaje波德加耶() feud.Barony {
	return c.波德加耶Podgaje
}
    
var CGnieznienskie纳克沃 GnieznienskieCounty = &纳克沃GnieznienskieCounty{}

func init() {
	f := CGnieznienskie纳克沃.(*纳克沃GnieznienskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "429",
		Title:     "gnieznienskie",
		TitleName: "纳克沃",
		TitleCode: "c_gnieznienskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.比亚瓦杰马Bialaziema = BBialaziema比亚瓦杰马
	f.比亚瓦杰马Bialaziema.SetParent(f)
	
	f.卡缅Kamien = BKamien卡缅
	f.卡缅Kamien.SetParent(f)
	
	f.克齐尼亚Kcynia = BKcynia克齐尼亚
	f.克齐尼亚Kcynia.SetParent(f)
	
	f.卢别沃Lubiewo = BLubiewo卢别沃
	f.卢别沃Lubiewo.SetParent(f)
	
	f.纳克沃Naklo = BNaklo纳克沃
	f.纳克沃Naklo.SetParent(f)
	
	f.皮瓦Pila = BPila皮瓦
	f.皮瓦Pila.SetParent(f)
	
	f.波德加耶Podgaje = BPodgaje波德加耶
	f.波德加耶Podgaje.SetParent(f)
	
}
