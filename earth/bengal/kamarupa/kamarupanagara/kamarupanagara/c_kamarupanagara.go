package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KamarupanagaraCounty interface {
    feud.County
    BDariyaon陀利耶乌那() 	feud.Barony
    BDurjaya突阇耶() 	feud.Barony
    BHajo诃阇乌() 	feud.Barony
    BKamarupanagara迦摩缕波城() 	feud.Barony
    BMadan_kamdev摩檀那迦摩提婆() 	feud.Barony
    BManikuta摩尼矩吒() 	feud.Barony
    BPragyotisapura东光城() 	feud.Barony
}

type 迦摩缕波城KamarupanagaraCounty struct {
	feud.BaseCounty
	陀利耶乌那Dariyaon 	feud.Barony
	突阇耶Durjaya 	feud.Barony
	诃阇乌Hajo 	feud.Barony
	迦摩缕波城Kamarupanagara 	feud.Barony
	摩檀那迦摩提婆Madan_kamdev 	feud.Barony
	摩尼矩吒Manikuta 	feud.Barony
	东光城Pragyotisapura 	feud.Barony
}

func (c *迦摩缕波城KamarupanagaraCounty) BDariyaon陀利耶乌那() feud.Barony {
	return c.陀利耶乌那Dariyaon
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BDurjaya突阇耶() feud.Barony {
	return c.突阇耶Durjaya
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BHajo诃阇乌() feud.Barony {
	return c.诃阇乌Hajo
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BKamarupanagara迦摩缕波城() feud.Barony {
	return c.迦摩缕波城Kamarupanagara
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BMadan_kamdev摩檀那迦摩提婆() feud.Barony {
	return c.摩檀那迦摩提婆Madan_kamdev
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BManikuta摩尼矩吒() feud.Barony {
	return c.摩尼矩吒Manikuta
}
    
func (c *迦摩缕波城KamarupanagaraCounty) BPragyotisapura东光城() feud.Barony {
	return c.东光城Pragyotisapura
}
    
var CKamarupanagara迦摩缕波城 KamarupanagaraCounty = &迦摩缕波城KamarupanagaraCounty{}

func init() {
	f := CKamarupanagara迦摩缕波城.(*迦摩缕波城KamarupanagaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1321",
		Title:     "kamarupanagara",
		TitleName: "迦摩缕波城",
		TitleCode: "c_kamarupanagara",
		Baronies:  map[string]feud.Barony{},
	}

	f.陀利耶乌那Dariyaon = BDariyaon陀利耶乌那
	f.陀利耶乌那Dariyaon.SetParent(f)
	
	f.突阇耶Durjaya = BDurjaya突阇耶
	f.突阇耶Durjaya.SetParent(f)
	
	f.诃阇乌Hajo = BHajo诃阇乌
	f.诃阇乌Hajo.SetParent(f)
	
	f.迦摩缕波城Kamarupanagara = BKamarupanagara迦摩缕波城
	f.迦摩缕波城Kamarupanagara.SetParent(f)
	
	f.摩檀那迦摩提婆Madan_kamdev = BMadan_kamdev摩檀那迦摩提婆
	f.摩檀那迦摩提婆Madan_kamdev.SetParent(f)
	
	f.摩尼矩吒Manikuta = BManikuta摩尼矩吒
	f.摩尼矩吒Manikuta.SetParent(f)
	
	f.东光城Pragyotisapura = BPragyotisapura东光城
	f.东光城Pragyotisapura.SetParent(f)
	
}
