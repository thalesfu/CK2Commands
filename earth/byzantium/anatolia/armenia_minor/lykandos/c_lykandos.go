package lykandos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LykandosCounty interface {
    feud.County
    BArabissus阿拉比苏斯() 	feud.Barony
    BCocussus格克孙() 	feud.Barony
    BComanagene科曼格尼() 	feud.Barony
    BGermanikeia日耳曼尼基亚() 	feud.Barony
    BLykandos律坎多斯() 	feud.Barony
    BPapurius拉普乌斯() 	feud.Barony
    BSymposion苏穆波逊() 	feud.Barony
    BTzamandos特泽马多斯() 	feud.Barony
}

type 吕坎多斯LykandosCounty struct {
	feud.BaseCounty
	阿拉比苏斯Arabissus 	feud.Barony
	格克孙Cocussus 	feud.Barony
	科曼格尼Comanagene 	feud.Barony
	日耳曼尼基亚Germanikeia 	feud.Barony
	律坎多斯Lykandos 	feud.Barony
	拉普乌斯Papurius 	feud.Barony
	苏穆波逊Symposion 	feud.Barony
	特泽马多斯Tzamandos 	feud.Barony
}

func (c *吕坎多斯LykandosCounty) BArabissus阿拉比苏斯() feud.Barony {
	return c.阿拉比苏斯Arabissus
}
    
func (c *吕坎多斯LykandosCounty) BCocussus格克孙() feud.Barony {
	return c.格克孙Cocussus
}
    
func (c *吕坎多斯LykandosCounty) BComanagene科曼格尼() feud.Barony {
	return c.科曼格尼Comanagene
}
    
func (c *吕坎多斯LykandosCounty) BGermanikeia日耳曼尼基亚() feud.Barony {
	return c.日耳曼尼基亚Germanikeia
}
    
func (c *吕坎多斯LykandosCounty) BLykandos律坎多斯() feud.Barony {
	return c.律坎多斯Lykandos
}
    
func (c *吕坎多斯LykandosCounty) BPapurius拉普乌斯() feud.Barony {
	return c.拉普乌斯Papurius
}
    
func (c *吕坎多斯LykandosCounty) BSymposion苏穆波逊() feud.Barony {
	return c.苏穆波逊Symposion
}
    
func (c *吕坎多斯LykandosCounty) BTzamandos特泽马多斯() feud.Barony {
	return c.特泽马多斯Tzamandos
}
    
var CLykandos吕坎多斯 LykandosCounty = &吕坎多斯LykandosCounty{}

func init() {
	f := CLykandos吕坎多斯.(*吕坎多斯LykandosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "736",
		Title:     "lykandos",
		TitleName: "吕坎多斯",
		TitleCode: "c_lykandos",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉比苏斯Arabissus = BArabissus阿拉比苏斯
	f.阿拉比苏斯Arabissus.SetParent(f)
	
	f.格克孙Cocussus = BCocussus格克孙
	f.格克孙Cocussus.SetParent(f)
	
	f.科曼格尼Comanagene = BComanagene科曼格尼
	f.科曼格尼Comanagene.SetParent(f)
	
	f.日耳曼尼基亚Germanikeia = BGermanikeia日耳曼尼基亚
	f.日耳曼尼基亚Germanikeia.SetParent(f)
	
	f.律坎多斯Lykandos = BLykandos律坎多斯
	f.律坎多斯Lykandos.SetParent(f)
	
	f.拉普乌斯Papurius = BPapurius拉普乌斯
	f.拉普乌斯Papurius.SetParent(f)
	
	f.苏穆波逊Symposion = BSymposion苏穆波逊
	f.苏穆波逊Symposion.SetParent(f)
	
	f.特泽马多斯Tzamandos = BTzamandos特泽马多斯
	f.特泽马多斯Tzamandos.SetParent(f)
	
}
