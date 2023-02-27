package sevede

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SevedeCounty interface {
    feud.County
    BDoderhult多德霍特() 	feud.Barony
    BHogsby赫格斯比() 	feud.Barony
    BHulingsryd胡令斯吕德() 	feud.Barony
    BKronsborg克伦斯堡() 	feud.Barony
    BMonsteras蒙斯特罗斯() 	feud.Barony
    BSimpevarp辛佩瓦普() 	feud.Barony
    BVimmerby维默比() 	feud.Barony
}

type 塞韦德SevedeCounty struct {
	feud.BaseCounty
	多德霍特Doderhult 	feud.Barony
	赫格斯比Hogsby 	feud.Barony
	胡令斯吕德Hulingsryd 	feud.Barony
	克伦斯堡Kronsborg 	feud.Barony
	蒙斯特罗斯Monsteras 	feud.Barony
	辛佩瓦普Simpevarp 	feud.Barony
	维默比Vimmerby 	feud.Barony
}

func (c *塞韦德SevedeCounty) BDoderhult多德霍特() feud.Barony {
	return c.多德霍特Doderhult
}
    
func (c *塞韦德SevedeCounty) BHogsby赫格斯比() feud.Barony {
	return c.赫格斯比Hogsby
}
    
func (c *塞韦德SevedeCounty) BHulingsryd胡令斯吕德() feud.Barony {
	return c.胡令斯吕德Hulingsryd
}
    
func (c *塞韦德SevedeCounty) BKronsborg克伦斯堡() feud.Barony {
	return c.克伦斯堡Kronsborg
}
    
func (c *塞韦德SevedeCounty) BMonsteras蒙斯特罗斯() feud.Barony {
	return c.蒙斯特罗斯Monsteras
}
    
func (c *塞韦德SevedeCounty) BSimpevarp辛佩瓦普() feud.Barony {
	return c.辛佩瓦普Simpevarp
}
    
func (c *塞韦德SevedeCounty) BVimmerby维默比() feud.Barony {
	return c.维默比Vimmerby
}
    
var CSevede塞韦德 SevedeCounty = &塞韦德SevedeCounty{}

func init() {
	f := CSevede塞韦德.(*塞韦德SevedeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1704",
		Title:     "sevede",
		TitleName: "塞韦德",
		TitleCode: "c_sevede",
		Baronies:  map[string]feud.Barony{},
	}

	f.多德霍特Doderhult = BDoderhult多德霍特
	f.多德霍特Doderhult.SetParent(f)
	
	f.赫格斯比Hogsby = BHogsby赫格斯比
	f.赫格斯比Hogsby.SetParent(f)
	
	f.胡令斯吕德Hulingsryd = BHulingsryd胡令斯吕德
	f.胡令斯吕德Hulingsryd.SetParent(f)
	
	f.克伦斯堡Kronsborg = BKronsborg克伦斯堡
	f.克伦斯堡Kronsborg.SetParent(f)
	
	f.蒙斯特罗斯Monsteras = BMonsteras蒙斯特罗斯
	f.蒙斯特罗斯Monsteras.SetParent(f)
	
	f.辛佩瓦普Simpevarp = BSimpevarp辛佩瓦普
	f.辛佩瓦普Simpevarp.SetParent(f)
	
	f.维默比Vimmerby = BVimmerby维默比
	f.维默比Vimmerby.SetParent(f)
	
}
