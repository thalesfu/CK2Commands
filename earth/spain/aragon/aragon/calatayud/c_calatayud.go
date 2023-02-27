package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CalatayudCounty interface {
    feud.County
    BAlhamadearagon阿拉马德亚拉贡() 	feud.Barony
    BCalatayud卡拉塔尤() 	feud.Barony
    BCalmarza卡尔马尔萨() 	feud.Barony
    BCimballa辛巴利亚() 	feud.Barony
    BDaroca达罗卡() 	feud.Barony
    BMunebrega穆内夫雷加() 	feud.Barony
    BNuevalos努埃瓦洛斯() 	feud.Barony
    BPiedra彼德拉() 	feud.Barony
}

type 卡拉塔尤CalatayudCounty struct {
	feud.BaseCounty
	阿拉马德亚拉贡Alhamadearagon 	feud.Barony
	卡拉塔尤Calatayud 	feud.Barony
	卡尔马尔萨Calmarza 	feud.Barony
	辛巴利亚Cimballa 	feud.Barony
	达罗卡Daroca 	feud.Barony
	穆内夫雷加Munebrega 	feud.Barony
	努埃瓦洛斯Nuevalos 	feud.Barony
	彼德拉Piedra 	feud.Barony
}

func (c *卡拉塔尤CalatayudCounty) BAlhamadearagon阿拉马德亚拉贡() feud.Barony {
	return c.阿拉马德亚拉贡Alhamadearagon
}
    
func (c *卡拉塔尤CalatayudCounty) BCalatayud卡拉塔尤() feud.Barony {
	return c.卡拉塔尤Calatayud
}
    
func (c *卡拉塔尤CalatayudCounty) BCalmarza卡尔马尔萨() feud.Barony {
	return c.卡尔马尔萨Calmarza
}
    
func (c *卡拉塔尤CalatayudCounty) BCimballa辛巴利亚() feud.Barony {
	return c.辛巴利亚Cimballa
}
    
func (c *卡拉塔尤CalatayudCounty) BDaroca达罗卡() feud.Barony {
	return c.达罗卡Daroca
}
    
func (c *卡拉塔尤CalatayudCounty) BMunebrega穆内夫雷加() feud.Barony {
	return c.穆内夫雷加Munebrega
}
    
func (c *卡拉塔尤CalatayudCounty) BNuevalos努埃瓦洛斯() feud.Barony {
	return c.努埃瓦洛斯Nuevalos
}
    
func (c *卡拉塔尤CalatayudCounty) BPiedra彼德拉() feud.Barony {
	return c.彼德拉Piedra
}
    
var CCalatayud卡拉塔尤 CalatayudCounty = &卡拉塔尤CalatayudCounty{}

func init() {
	f := CCalatayud卡拉塔尤.(*卡拉塔尤CalatayudCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "175",
		Title:     "calatayud",
		TitleName: "卡拉塔尤",
		TitleCode: "c_calatayud",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉马德亚拉贡Alhamadearagon = BAlhamadearagon阿拉马德亚拉贡
	f.阿拉马德亚拉贡Alhamadearagon.SetParent(f)
	
	f.卡拉塔尤Calatayud = BCalatayud卡拉塔尤
	f.卡拉塔尤Calatayud.SetParent(f)
	
	f.卡尔马尔萨Calmarza = BCalmarza卡尔马尔萨
	f.卡尔马尔萨Calmarza.SetParent(f)
	
	f.辛巴利亚Cimballa = BCimballa辛巴利亚
	f.辛巴利亚Cimballa.SetParent(f)
	
	f.达罗卡Daroca = BDaroca达罗卡
	f.达罗卡Daroca.SetParent(f)
	
	f.穆内夫雷加Munebrega = BMunebrega穆内夫雷加
	f.穆内夫雷加Munebrega.SetParent(f)
	
	f.努埃瓦洛斯Nuevalos = BNuevalos努埃瓦洛斯
	f.努埃瓦洛斯Nuevalos.SetParent(f)
	
	f.彼德拉Piedra = BPiedra彼德拉
	f.彼德拉Piedra.SetParent(f)
	
}
