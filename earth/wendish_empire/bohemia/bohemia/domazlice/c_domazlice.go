package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DomazliceCounty interface {
    feud.County
    BBudejovice布杰约维采() 	feud.Barony
    BChynov希诺夫() 	feud.Barony
    BDoudleby多乌德莱比() 	feud.Barony
    BGoldenkron戈尔登克龙() 	feud.Barony
    BHohenfurth霍亨富尔特() 	feud.Barony
    BKrumlov克鲁姆洛夫() 	feud.Barony
    BPisek皮塞克() 	feud.Barony
    BPrachen普拉亨尼() 	feud.Barony
    BResenberg雷森贝格() 	feud.Barony
    BRosenberg罗森贝格() 	feud.Barony
}

type 南波希米亚DomazliceCounty struct {
	feud.BaseCounty
	布杰约维采Budejovice 	feud.Barony
	希诺夫Chynov 	feud.Barony
	多乌德莱比Doudleby 	feud.Barony
	戈尔登克龙Goldenkron 	feud.Barony
	霍亨富尔特Hohenfurth 	feud.Barony
	克鲁姆洛夫Krumlov 	feud.Barony
	皮塞克Pisek 	feud.Barony
	普拉亨尼Prachen 	feud.Barony
	雷森贝格Resenberg 	feud.Barony
	罗森贝格Rosenberg 	feud.Barony
}

func (c *南波希米亚DomazliceCounty) BBudejovice布杰约维采() feud.Barony {
	return c.布杰约维采Budejovice
}
    
func (c *南波希米亚DomazliceCounty) BChynov希诺夫() feud.Barony {
	return c.希诺夫Chynov
}
    
func (c *南波希米亚DomazliceCounty) BDoudleby多乌德莱比() feud.Barony {
	return c.多乌德莱比Doudleby
}
    
func (c *南波希米亚DomazliceCounty) BGoldenkron戈尔登克龙() feud.Barony {
	return c.戈尔登克龙Goldenkron
}
    
func (c *南波希米亚DomazliceCounty) BHohenfurth霍亨富尔特() feud.Barony {
	return c.霍亨富尔特Hohenfurth
}
    
func (c *南波希米亚DomazliceCounty) BKrumlov克鲁姆洛夫() feud.Barony {
	return c.克鲁姆洛夫Krumlov
}
    
func (c *南波希米亚DomazliceCounty) BPisek皮塞克() feud.Barony {
	return c.皮塞克Pisek
}
    
func (c *南波希米亚DomazliceCounty) BPrachen普拉亨尼() feud.Barony {
	return c.普拉亨尼Prachen
}
    
func (c *南波希米亚DomazliceCounty) BResenberg雷森贝格() feud.Barony {
	return c.雷森贝格Resenberg
}
    
func (c *南波希米亚DomazliceCounty) BRosenberg罗森贝格() feud.Barony {
	return c.罗森贝格Rosenberg
}
    
var CDomazlice南波希米亚 DomazliceCounty = &南波希米亚DomazliceCounty{}

func init() {
	f := CDomazlice南波希米亚.(*南波希米亚DomazliceCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "362",
		Title:     "domazlice",
		TitleName: "南波希米亚",
		TitleCode: "c_domazlice",
		Baronies:  map[string]feud.Barony{},
	}

	f.布杰约维采Budejovice = BBudejovice布杰约维采
	f.布杰约维采Budejovice.SetParent(f)
	
	f.希诺夫Chynov = BChynov希诺夫
	f.希诺夫Chynov.SetParent(f)
	
	f.多乌德莱比Doudleby = BDoudleby多乌德莱比
	f.多乌德莱比Doudleby.SetParent(f)
	
	f.戈尔登克龙Goldenkron = BGoldenkron戈尔登克龙
	f.戈尔登克龙Goldenkron.SetParent(f)
	
	f.霍亨富尔特Hohenfurth = BHohenfurth霍亨富尔特
	f.霍亨富尔特Hohenfurth.SetParent(f)
	
	f.克鲁姆洛夫Krumlov = BKrumlov克鲁姆洛夫
	f.克鲁姆洛夫Krumlov.SetParent(f)
	
	f.皮塞克Pisek = BPisek皮塞克
	f.皮塞克Pisek.SetParent(f)
	
	f.普拉亨尼Prachen = BPrachen普拉亨尼
	f.普拉亨尼Prachen.SetParent(f)
	
	f.雷森贝格Resenberg = BResenberg雷森贝格
	f.雷森贝格Resenberg.SetParent(f)
	
	f.罗森贝格Rosenberg = BRosenberg罗森贝格
	f.罗森贝格Rosenberg.SetParent(f)
	
}
