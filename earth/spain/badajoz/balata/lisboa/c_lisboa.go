package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LisboaCounty interface {
    feud.County
    BAlcobaca阿尔科巴萨() 	feud.Barony
    BAlenquer阿伦克尔() 	feud.Barony
    BAtouguia阿托吉亚() 	feud.Barony
    BBatalha巴塔利亚() 	feud.Barony
    BLisboa里斯本() 	feud.Barony
    BSantarem圣塔伦() 	feud.Barony
    BSintra辛特拉() 	feud.Barony
    BTomar托马尔() 	feud.Barony
}

type 里斯本LisboaCounty struct {
	feud.BaseCounty
	阿尔科巴萨Alcobaca 	feud.Barony
	阿伦克尔Alenquer 	feud.Barony
	阿托吉亚Atouguia 	feud.Barony
	巴塔利亚Batalha 	feud.Barony
	里斯本Lisboa 	feud.Barony
	圣塔伦Santarem 	feud.Barony
	辛特拉Sintra 	feud.Barony
	托马尔Tomar 	feud.Barony
}

func (c *里斯本LisboaCounty) BAlcobaca阿尔科巴萨() feud.Barony {
	return c.阿尔科巴萨Alcobaca
}
    
func (c *里斯本LisboaCounty) BAlenquer阿伦克尔() feud.Barony {
	return c.阿伦克尔Alenquer
}
    
func (c *里斯本LisboaCounty) BAtouguia阿托吉亚() feud.Barony {
	return c.阿托吉亚Atouguia
}
    
func (c *里斯本LisboaCounty) BBatalha巴塔利亚() feud.Barony {
	return c.巴塔利亚Batalha
}
    
func (c *里斯本LisboaCounty) BLisboa里斯本() feud.Barony {
	return c.里斯本Lisboa
}
    
func (c *里斯本LisboaCounty) BSantarem圣塔伦() feud.Barony {
	return c.圣塔伦Santarem
}
    
func (c *里斯本LisboaCounty) BSintra辛特拉() feud.Barony {
	return c.辛特拉Sintra
}
    
func (c *里斯本LisboaCounty) BTomar托马尔() feud.Barony {
	return c.托马尔Tomar
}
    
var CLisboa里斯本 LisboaCounty = &里斯本LisboaCounty{}

func init() {
	f := CLisboa里斯本.(*里斯本LisboaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "160",
		Title:     "lisboa",
		TitleName: "里斯本",
		TitleCode: "c_lisboa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔科巴萨Alcobaca = BAlcobaca阿尔科巴萨
	f.阿尔科巴萨Alcobaca.SetParent(f)
	
	f.阿伦克尔Alenquer = BAlenquer阿伦克尔
	f.阿伦克尔Alenquer.SetParent(f)
	
	f.阿托吉亚Atouguia = BAtouguia阿托吉亚
	f.阿托吉亚Atouguia.SetParent(f)
	
	f.巴塔利亚Batalha = BBatalha巴塔利亚
	f.巴塔利亚Batalha.SetParent(f)
	
	f.里斯本Lisboa = BLisboa里斯本
	f.里斯本Lisboa.SetParent(f)
	
	f.圣塔伦Santarem = BSantarem圣塔伦
	f.圣塔伦Santarem.SetParent(f)
	
	f.辛特拉Sintra = BSintra辛特拉
	f.辛特拉Sintra.SetParent(f)
	
	f.托马尔Tomar = BTomar托马尔
	f.托马尔Tomar.SetParent(f)
	
}
