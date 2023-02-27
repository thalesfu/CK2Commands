package potapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PotapiCounty interface {
    feud.County
    BChaura招罗() 	feud.Barony
    BKaveripattinam劫吠梨钵胝那摩() 	feud.Barony
    BMando曼度() 	feud.Barony
    BMelpadi摩婆提() 	feud.Barony
    BPotapi菩多毗() 	feud.Barony
    BTirupati帝楼波底() 	feud.Barony
    BVirincipuram毗利西补楞() 	feud.Barony
}

type 菩多毗PotapiCounty struct {
	feud.BaseCounty
	招罗Chaura 	feud.Barony
	劫吠梨钵胝那摩Kaveripattinam 	feud.Barony
	曼度Mando 	feud.Barony
	摩婆提Melpadi 	feud.Barony
	菩多毗Potapi 	feud.Barony
	帝楼波底Tirupati 	feud.Barony
	毗利西补楞Virincipuram 	feud.Barony
}

func (c *菩多毗PotapiCounty) BChaura招罗() feud.Barony {
	return c.招罗Chaura
}
    
func (c *菩多毗PotapiCounty) BKaveripattinam劫吠梨钵胝那摩() feud.Barony {
	return c.劫吠梨钵胝那摩Kaveripattinam
}
    
func (c *菩多毗PotapiCounty) BMando曼度() feud.Barony {
	return c.曼度Mando
}
    
func (c *菩多毗PotapiCounty) BMelpadi摩婆提() feud.Barony {
	return c.摩婆提Melpadi
}
    
func (c *菩多毗PotapiCounty) BPotapi菩多毗() feud.Barony {
	return c.菩多毗Potapi
}
    
func (c *菩多毗PotapiCounty) BTirupati帝楼波底() feud.Barony {
	return c.帝楼波底Tirupati
}
    
func (c *菩多毗PotapiCounty) BVirincipuram毗利西补楞() feud.Barony {
	return c.毗利西补楞Virincipuram
}
    
var CPotapi菩多毗 PotapiCounty = &菩多毗PotapiCounty{}

func init() {
	f := CPotapi菩多毗.(*菩多毗PotapiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1219",
		Title:     "potapi",
		TitleName: "菩多毗",
		TitleCode: "c_potapi",
		Baronies:  map[string]feud.Barony{},
	}

	f.招罗Chaura = BChaura招罗
	f.招罗Chaura.SetParent(f)
	
	f.劫吠梨钵胝那摩Kaveripattinam = BKaveripattinam劫吠梨钵胝那摩
	f.劫吠梨钵胝那摩Kaveripattinam.SetParent(f)
	
	f.曼度Mando = BMando曼度
	f.曼度Mando.SetParent(f)
	
	f.摩婆提Melpadi = BMelpadi摩婆提
	f.摩婆提Melpadi.SetParent(f)
	
	f.菩多毗Potapi = BPotapi菩多毗
	f.菩多毗Potapi.SetParent(f)
	
	f.帝楼波底Tirupati = BTirupati帝楼波底
	f.帝楼波底Tirupati.SetParent(f)
	
	f.毗利西补楞Virincipuram = BVirincipuram毗利西补楞
	f.毗利西补楞Virincipuram.SetParent(f)
	
}
