package al_alamayn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_alamaynCounty interface {
    feud.County
    BElalamein阿莱曼() 	feud.Barony
    BFuka富凯() 	feud.Barony
    BGhazal加扎勒() 	feud.Barony
    BKatabathmos卡塔布赫摩斯() 	feud.Barony
    BShammas沙姆() 	feud.Barony
    BSidibarrani西迪拜拉尼() 	feud.Barony
    BZygra卒格拉() 	feud.Barony
}

type 卒格拉Al_alamaynCounty struct {
	feud.BaseCounty
	阿莱曼Elalamein 	feud.Barony
	富凯Fuka 	feud.Barony
	加扎勒Ghazal 	feud.Barony
	卡塔布赫摩斯Katabathmos 	feud.Barony
	沙姆Shammas 	feud.Barony
	西迪拜拉尼Sidibarrani 	feud.Barony
	卒格拉Zygra 	feud.Barony
}

func (c *卒格拉Al_alamaynCounty) BElalamein阿莱曼() feud.Barony {
	return c.阿莱曼Elalamein
}
    
func (c *卒格拉Al_alamaynCounty) BFuka富凯() feud.Barony {
	return c.富凯Fuka
}
    
func (c *卒格拉Al_alamaynCounty) BGhazal加扎勒() feud.Barony {
	return c.加扎勒Ghazal
}
    
func (c *卒格拉Al_alamaynCounty) BKatabathmos卡塔布赫摩斯() feud.Barony {
	return c.卡塔布赫摩斯Katabathmos
}
    
func (c *卒格拉Al_alamaynCounty) BShammas沙姆() feud.Barony {
	return c.沙姆Shammas
}
    
func (c *卒格拉Al_alamaynCounty) BSidibarrani西迪拜拉尼() feud.Barony {
	return c.西迪拜拉尼Sidibarrani
}
    
func (c *卒格拉Al_alamaynCounty) BZygra卒格拉() feud.Barony {
	return c.卒格拉Zygra
}
    
var CAl_alamayn卒格拉 Al_alamaynCounty = &卒格拉Al_alamaynCounty{}

func init() {
	f := CAl_alamayn卒格拉.(*卒格拉Al_alamaynCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "804",
		Title:     "al_alamayn",
		TitleName: "卒格拉",
		TitleCode: "c_al_alamayn",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莱曼Elalamein = BElalamein阿莱曼
	f.阿莱曼Elalamein.SetParent(f)
	
	f.富凯Fuka = BFuka富凯
	f.富凯Fuka.SetParent(f)
	
	f.加扎勒Ghazal = BGhazal加扎勒
	f.加扎勒Ghazal.SetParent(f)
	
	f.卡塔布赫摩斯Katabathmos = BKatabathmos卡塔布赫摩斯
	f.卡塔布赫摩斯Katabathmos.SetParent(f)
	
	f.沙姆Shammas = BShammas沙姆
	f.沙姆Shammas.SetParent(f)
	
	f.西迪拜拉尼Sidibarrani = BSidibarrani西迪拜拉尼
	f.西迪拜拉尼Sidibarrani.SetParent(f)
	
	f.卒格拉Zygra = BZygra卒格拉
	f.卒格拉Zygra.SetParent(f)
	
}
