package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZyrianeCounty interface {
    feud.County
    BErgach叶尔加奇() 	feud.Barony
    BGari加里() 	feud.Barony
    BKordon科尔东() 	feud.Barony
    BKukushtan库库什坦() 	feud.Barony
    BKungur昆古尔() 	feud.Barony
    BLek列克() 	feud.Barony
    BPosad波萨德() 	feud.Barony
    BSuksun苏克孙() 	feud.Barony
}

type 伊伦斯克ZyrianeCounty struct {
	feud.BaseCounty
	叶尔加奇Ergach 	feud.Barony
	加里Gari 	feud.Barony
	科尔东Kordon 	feud.Barony
	库库什坦Kukushtan 	feud.Barony
	昆古尔Kungur 	feud.Barony
	列克Lek 	feud.Barony
	波萨德Posad 	feud.Barony
	苏克孙Suksun 	feud.Barony
}

func (c *伊伦斯克ZyrianeCounty) BErgach叶尔加奇() feud.Barony {
	return c.叶尔加奇Ergach
}
    
func (c *伊伦斯克ZyrianeCounty) BGari加里() feud.Barony {
	return c.加里Gari
}
    
func (c *伊伦斯克ZyrianeCounty) BKordon科尔东() feud.Barony {
	return c.科尔东Kordon
}
    
func (c *伊伦斯克ZyrianeCounty) BKukushtan库库什坦() feud.Barony {
	return c.库库什坦Kukushtan
}
    
func (c *伊伦斯克ZyrianeCounty) BKungur昆古尔() feud.Barony {
	return c.昆古尔Kungur
}
    
func (c *伊伦斯克ZyrianeCounty) BLek列克() feud.Barony {
	return c.列克Lek
}
    
func (c *伊伦斯克ZyrianeCounty) BPosad波萨德() feud.Barony {
	return c.波萨德Posad
}
    
func (c *伊伦斯克ZyrianeCounty) BSuksun苏克孙() feud.Barony {
	return c.苏克孙Suksun
}
    
var CZyriane伊伦斯克 ZyrianeCounty = &伊伦斯克ZyrianeCounty{}

func init() {
	f := CZyriane伊伦斯克.(*伊伦斯克ZyrianeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "400",
		Title:     "zyriane",
		TitleName: "伊伦斯克",
		TitleCode: "c_zyriane",
		Baronies:  map[string]feud.Barony{},
	}

	f.叶尔加奇Ergach = BErgach叶尔加奇
	f.叶尔加奇Ergach.SetParent(f)
	
	f.加里Gari = BGari加里
	f.加里Gari.SetParent(f)
	
	f.科尔东Kordon = BKordon科尔东
	f.科尔东Kordon.SetParent(f)
	
	f.库库什坦Kukushtan = BKukushtan库库什坦
	f.库库什坦Kukushtan.SetParent(f)
	
	f.昆古尔Kungur = BKungur昆古尔
	f.昆古尔Kungur.SetParent(f)
	
	f.列克Lek = BLek列克
	f.列克Lek.SetParent(f)
	
	f.波萨德Posad = BPosad波萨德
	f.波萨德Posad.SetParent(f)
	
	f.苏克孙Suksun = BSuksun苏克孙
	f.苏克孙Suksun.SetParent(f)
	
}
