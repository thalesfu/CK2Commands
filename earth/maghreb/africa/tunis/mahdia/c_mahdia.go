package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MahdiaCounty interface {
    feud.County
    BAgareb阿加里卜() 	feud.Barony
    BLacroix拉克鲁瓦() 	feud.Barony
    BLalaouza劳欧扎() 	feud.Barony
    BMahdia马赫迪耶() 	feud.Barony
    BMonastir莫纳斯提尔() 	feud.Barony
    BMsaken姆萨肯() 	feud.Barony
    BNeffatia勒法蒂亚() 	feud.Barony
    BSfax斯法克斯() 	feud.Barony
    BSidimansour西迪曼苏尔() 	feud.Barony
}

type 马赫迪耶MahdiaCounty struct {
	feud.BaseCounty
	阿加里卜Agareb 	feud.Barony
	拉克鲁瓦Lacroix 	feud.Barony
	劳欧扎Lalaouza 	feud.Barony
	马赫迪耶Mahdia 	feud.Barony
	莫纳斯提尔Monastir 	feud.Barony
	姆萨肯Msaken 	feud.Barony
	勒法蒂亚Neffatia 	feud.Barony
	斯法克斯Sfax 	feud.Barony
	西迪曼苏尔Sidimansour 	feud.Barony
}

func (c *马赫迪耶MahdiaCounty) BAgareb阿加里卜() feud.Barony {
	return c.阿加里卜Agareb
}
    
func (c *马赫迪耶MahdiaCounty) BLacroix拉克鲁瓦() feud.Barony {
	return c.拉克鲁瓦Lacroix
}
    
func (c *马赫迪耶MahdiaCounty) BLalaouza劳欧扎() feud.Barony {
	return c.劳欧扎Lalaouza
}
    
func (c *马赫迪耶MahdiaCounty) BMahdia马赫迪耶() feud.Barony {
	return c.马赫迪耶Mahdia
}
    
func (c *马赫迪耶MahdiaCounty) BMonastir莫纳斯提尔() feud.Barony {
	return c.莫纳斯提尔Monastir
}
    
func (c *马赫迪耶MahdiaCounty) BMsaken姆萨肯() feud.Barony {
	return c.姆萨肯Msaken
}
    
func (c *马赫迪耶MahdiaCounty) BNeffatia勒法蒂亚() feud.Barony {
	return c.勒法蒂亚Neffatia
}
    
func (c *马赫迪耶MahdiaCounty) BSfax斯法克斯() feud.Barony {
	return c.斯法克斯Sfax
}
    
func (c *马赫迪耶MahdiaCounty) BSidimansour西迪曼苏尔() feud.Barony {
	return c.西迪曼苏尔Sidimansour
}
    
var CMahdia马赫迪耶 MahdiaCounty = &马赫迪耶MahdiaCounty{}

func init() {
	f := CMahdia马赫迪耶.(*马赫迪耶MahdiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "816",
		Title:     "mahdia",
		TitleName: "马赫迪耶",
		TitleCode: "c_mahdia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加里卜Agareb = BAgareb阿加里卜
	f.阿加里卜Agareb.SetParent(f)
	
	f.拉克鲁瓦Lacroix = BLacroix拉克鲁瓦
	f.拉克鲁瓦Lacroix.SetParent(f)
	
	f.劳欧扎Lalaouza = BLalaouza劳欧扎
	f.劳欧扎Lalaouza.SetParent(f)
	
	f.马赫迪耶Mahdia = BMahdia马赫迪耶
	f.马赫迪耶Mahdia.SetParent(f)
	
	f.莫纳斯提尔Monastir = BMonastir莫纳斯提尔
	f.莫纳斯提尔Monastir.SetParent(f)
	
	f.姆萨肯Msaken = BMsaken姆萨肯
	f.姆萨肯Msaken.SetParent(f)
	
	f.勒法蒂亚Neffatia = BNeffatia勒法蒂亚
	f.勒法蒂亚Neffatia.SetParent(f)
	
	f.斯法克斯Sfax = BSfax斯法克斯
	f.斯法克斯Sfax.SetParent(f)
	
	f.西迪曼苏尔Sidimansour = BSidimansour西迪曼苏尔
	f.西迪曼苏尔Sidimansour.SetParent(f)
	
}
