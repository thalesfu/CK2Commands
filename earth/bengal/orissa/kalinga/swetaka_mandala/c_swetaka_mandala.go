package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Swetaka_mandalaCounty interface {
    feud.County
    BBadakhemundi婆荼契闷稚() 	feud.Barony
    BKharki伽吉() 	feud.Barony
    BKhed艺德() 	feud.Barony
    BKhijaria弃折梨耶() 	feud.Barony
    BRayagada罗耶迦陀() 	feud.Barony
    BSwetakapura湿吠多迦城() 	feud.Barony
    BUmerkote优迷罗拘吒() 	feud.Barony
}

type 湿吠多迦曼荼罗Swetaka_mandalaCounty struct {
	feud.BaseCounty
	婆荼契闷稚Badakhemundi 	feud.Barony
	伽吉Kharki 	feud.Barony
	艺德Khed 	feud.Barony
	弃折梨耶Khijaria 	feud.Barony
	罗耶迦陀Rayagada 	feud.Barony
	湿吠多迦城Swetakapura 	feud.Barony
	优迷罗拘吒Umerkote 	feud.Barony
}

func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BBadakhemundi婆荼契闷稚() feud.Barony {
	return c.婆荼契闷稚Badakhemundi
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BKharki伽吉() feud.Barony {
	return c.伽吉Kharki
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BKhed艺德() feud.Barony {
	return c.艺德Khed
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BKhijaria弃折梨耶() feud.Barony {
	return c.弃折梨耶Khijaria
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BRayagada罗耶迦陀() feud.Barony {
	return c.罗耶迦陀Rayagada
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BSwetakapura湿吠多迦城() feud.Barony {
	return c.湿吠多迦城Swetakapura
}
    
func (c *湿吠多迦曼荼罗Swetaka_mandalaCounty) BUmerkote优迷罗拘吒() feud.Barony {
	return c.优迷罗拘吒Umerkote
}
    
var CSwetaka_mandala湿吠多迦曼荼罗 Swetaka_mandalaCounty = &湿吠多迦曼荼罗Swetaka_mandalaCounty{}

func init() {
	f := CSwetaka_mandala湿吠多迦曼荼罗.(*湿吠多迦曼荼罗Swetaka_mandalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1228",
		Title:     "swetaka_mandala",
		TitleName: "湿吠多迦曼荼罗",
		TitleCode: "c_swetaka_mandala",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆荼契闷稚Badakhemundi = BBadakhemundi婆荼契闷稚
	f.婆荼契闷稚Badakhemundi.SetParent(f)
	
	f.伽吉Kharki = BKharki伽吉
	f.伽吉Kharki.SetParent(f)
	
	f.艺德Khed = BKhed艺德
	f.艺德Khed.SetParent(f)
	
	f.弃折梨耶Khijaria = BKhijaria弃折梨耶
	f.弃折梨耶Khijaria.SetParent(f)
	
	f.罗耶迦陀Rayagada = BRayagada罗耶迦陀
	f.罗耶迦陀Rayagada.SetParent(f)
	
	f.湿吠多迦城Swetakapura = BSwetakapura湿吠多迦城
	f.湿吠多迦城Swetakapura.SetParent(f)
	
	f.优迷罗拘吒Umerkote = BUmerkote优迷罗拘吒
	f.优迷罗拘吒Umerkote.SetParent(f)
	
}
