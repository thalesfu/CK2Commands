package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VikenCounty interface {
    feud.County
    BBagahus巴格赫斯() 	feud.Barony
    BHede赫德() 	feud.Barony
    BKungahalla库耶赫拉() 	feud.Barony
    BKungsviken孔斯维肯() 	feud.Barony
    BMarstrand马斯特兰德() 	feud.Barony
    BOckero厄克勒() 	feud.Barony
    BSvarteborg斯瓦特波雷() 	feud.Barony
    BSvenneby斯维尼比() 	feud.Barony
}

type 博胡斯伦VikenCounty struct {
	feud.BaseCounty
	巴格赫斯Bagahus 	feud.Barony
	赫德Hede 	feud.Barony
	库耶赫拉Kungahalla 	feud.Barony
	孔斯维肯Kungsviken 	feud.Barony
	马斯特兰德Marstrand 	feud.Barony
	厄克勒Ockero 	feud.Barony
	斯瓦特波雷Svarteborg 	feud.Barony
	斯维尼比Svenneby 	feud.Barony
}

func (c *博胡斯伦VikenCounty) BBagahus巴格赫斯() feud.Barony {
	return c.巴格赫斯Bagahus
}
    
func (c *博胡斯伦VikenCounty) BHede赫德() feud.Barony {
	return c.赫德Hede
}
    
func (c *博胡斯伦VikenCounty) BKungahalla库耶赫拉() feud.Barony {
	return c.库耶赫拉Kungahalla
}
    
func (c *博胡斯伦VikenCounty) BKungsviken孔斯维肯() feud.Barony {
	return c.孔斯维肯Kungsviken
}
    
func (c *博胡斯伦VikenCounty) BMarstrand马斯特兰德() feud.Barony {
	return c.马斯特兰德Marstrand
}
    
func (c *博胡斯伦VikenCounty) BOckero厄克勒() feud.Barony {
	return c.厄克勒Ockero
}
    
func (c *博胡斯伦VikenCounty) BSvarteborg斯瓦特波雷() feud.Barony {
	return c.斯瓦特波雷Svarteborg
}
    
func (c *博胡斯伦VikenCounty) BSvenneby斯维尼比() feud.Barony {
	return c.斯维尼比Svenneby
}
    
var CViken博胡斯伦 VikenCounty = &博胡斯伦VikenCounty{}

func init() {
	f := CViken博胡斯伦.(*博胡斯伦VikenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "296",
		Title:     "viken",
		TitleName: "博胡斯伦",
		TitleCode: "c_viken",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴格赫斯Bagahus = BBagahus巴格赫斯
	f.巴格赫斯Bagahus.SetParent(f)
	
	f.赫德Hede = BHede赫德
	f.赫德Hede.SetParent(f)
	
	f.库耶赫拉Kungahalla = BKungahalla库耶赫拉
	f.库耶赫拉Kungahalla.SetParent(f)
	
	f.孔斯维肯Kungsviken = BKungsviken孔斯维肯
	f.孔斯维肯Kungsviken.SetParent(f)
	
	f.马斯特兰德Marstrand = BMarstrand马斯特兰德
	f.马斯特兰德Marstrand.SetParent(f)
	
	f.厄克勒Ockero = BOckero厄克勒
	f.厄克勒Ockero.SetParent(f)
	
	f.斯瓦特波雷Svarteborg = BSvarteborg斯瓦特波雷
	f.斯瓦特波雷Svarteborg.SetParent(f)
	
	f.斯维尼比Svenneby = BSvenneby斯维尼比
	f.斯维尼比Svenneby.SetParent(f)
	
}
