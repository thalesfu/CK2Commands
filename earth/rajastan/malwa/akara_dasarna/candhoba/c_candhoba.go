package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CandhobaCounty interface {
    feud.County
    BCandhoba旃度婆() 	feud.Barony
    BCharmuria遮罗牟梨耶() 	feud.Barony
    BCharwadih遮罗婆提诃() 	feud.Barony
    BGuna求那() 	feud.Barony
    BKadwaya迦伐耶() 	feud.Barony
    BSipri斯毗() 	feud.Barony
    BTilosingi帝卢僧耆() 	feud.Barony
}

type 旃度婆CandhobaCounty struct {
	feud.BaseCounty
	旃度婆Candhoba 	feud.Barony
	遮罗牟梨耶Charmuria 	feud.Barony
	遮罗婆提诃Charwadih 	feud.Barony
	求那Guna 	feud.Barony
	迦伐耶Kadwaya 	feud.Barony
	斯毗Sipri 	feud.Barony
	帝卢僧耆Tilosingi 	feud.Barony
}

func (c *旃度婆CandhobaCounty) BCandhoba旃度婆() feud.Barony {
	return c.旃度婆Candhoba
}
    
func (c *旃度婆CandhobaCounty) BCharmuria遮罗牟梨耶() feud.Barony {
	return c.遮罗牟梨耶Charmuria
}
    
func (c *旃度婆CandhobaCounty) BCharwadih遮罗婆提诃() feud.Barony {
	return c.遮罗婆提诃Charwadih
}
    
func (c *旃度婆CandhobaCounty) BGuna求那() feud.Barony {
	return c.求那Guna
}
    
func (c *旃度婆CandhobaCounty) BKadwaya迦伐耶() feud.Barony {
	return c.迦伐耶Kadwaya
}
    
func (c *旃度婆CandhobaCounty) BSipri斯毗() feud.Barony {
	return c.斯毗Sipri
}
    
func (c *旃度婆CandhobaCounty) BTilosingi帝卢僧耆() feud.Barony {
	return c.帝卢僧耆Tilosingi
}
    
var CCandhoba旃度婆 CandhobaCounty = &旃度婆CandhobaCounty{}

func init() {
	f := CCandhoba旃度婆.(*旃度婆CandhobaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1299",
		Title:     "candhoba",
		TitleName: "旃度婆",
		TitleCode: "c_candhoba",
		Baronies:  map[string]feud.Barony{},
	}

	f.旃度婆Candhoba = BCandhoba旃度婆
	f.旃度婆Candhoba.SetParent(f)
	
	f.遮罗牟梨耶Charmuria = BCharmuria遮罗牟梨耶
	f.遮罗牟梨耶Charmuria.SetParent(f)
	
	f.遮罗婆提诃Charwadih = BCharwadih遮罗婆提诃
	f.遮罗婆提诃Charwadih.SetParent(f)
	
	f.求那Guna = BGuna求那
	f.求那Guna.SetParent(f)
	
	f.迦伐耶Kadwaya = BKadwaya迦伐耶
	f.迦伐耶Kadwaya.SetParent(f)
	
	f.斯毗Sipri = BSipri斯毗
	f.斯毗Sipri.SetParent(f)
	
	f.帝卢僧耆Tilosingi = BTilosingi帝卢僧耆
	f.帝卢僧耆Tilosingi.SetParent(f)
	
}
