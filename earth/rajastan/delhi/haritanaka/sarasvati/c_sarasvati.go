package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarasvatiCounty interface {
    feud.County
    BBhadra跋陀罗() 	feud.Barony
    BBhatnir婆提尼() 	feud.Barony
    BBukar富伽尔() 	feud.Barony
    BChakarpura遮迦罗补罗() 	feud.Barony
    BChandara旃陀罗() 	feud.Barony
    BSarasvati萨罗萨伐底() 	feud.Barony
    BTarain多罗因() 	feud.Barony
}

type 萨罗萨伐底SarasvatiCounty struct {
	feud.BaseCounty
	跋陀罗Bhadra 	feud.Barony
	婆提尼Bhatnir 	feud.Barony
	富伽尔Bukar 	feud.Barony
	遮迦罗补罗Chakarpura 	feud.Barony
	旃陀罗Chandara 	feud.Barony
	萨罗萨伐底Sarasvati 	feud.Barony
	多罗因Tarain 	feud.Barony
}

func (c *萨罗萨伐底SarasvatiCounty) BBhadra跋陀罗() feud.Barony {
	return c.跋陀罗Bhadra
}
    
func (c *萨罗萨伐底SarasvatiCounty) BBhatnir婆提尼() feud.Barony {
	return c.婆提尼Bhatnir
}
    
func (c *萨罗萨伐底SarasvatiCounty) BBukar富伽尔() feud.Barony {
	return c.富伽尔Bukar
}
    
func (c *萨罗萨伐底SarasvatiCounty) BChakarpura遮迦罗补罗() feud.Barony {
	return c.遮迦罗补罗Chakarpura
}
    
func (c *萨罗萨伐底SarasvatiCounty) BChandara旃陀罗() feud.Barony {
	return c.旃陀罗Chandara
}
    
func (c *萨罗萨伐底SarasvatiCounty) BSarasvati萨罗萨伐底() feud.Barony {
	return c.萨罗萨伐底Sarasvati
}
    
func (c *萨罗萨伐底SarasvatiCounty) BTarain多罗因() feud.Barony {
	return c.多罗因Tarain
}
    
var CSarasvati萨罗萨伐底 SarasvatiCounty = &萨罗萨伐底SarasvatiCounty{}

func init() {
	f := CSarasvati萨罗萨伐底.(*萨罗萨伐底SarasvatiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1352",
		Title:     "sarasvati",
		TitleName: "萨罗萨伐底",
		TitleCode: "c_sarasvati",
		Baronies:  map[string]feud.Barony{},
	}

	f.跋陀罗Bhadra = BBhadra跋陀罗
	f.跋陀罗Bhadra.SetParent(f)
	
	f.婆提尼Bhatnir = BBhatnir婆提尼
	f.婆提尼Bhatnir.SetParent(f)
	
	f.富伽尔Bukar = BBukar富伽尔
	f.富伽尔Bukar.SetParent(f)
	
	f.遮迦罗补罗Chakarpura = BChakarpura遮迦罗补罗
	f.遮迦罗补罗Chakarpura.SetParent(f)
	
	f.旃陀罗Chandara = BChandara旃陀罗
	f.旃陀罗Chandara.SetParent(f)
	
	f.萨罗萨伐底Sarasvati = BSarasvati萨罗萨伐底
	f.萨罗萨伐底Sarasvati.SetParent(f)
	
	f.多罗因Tarain = BTarain多罗因
	f.多罗因Tarain.SetParent(f)
	
}
