package miass

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MiassCounty interface {
    feud.County
    BKukushi库库希() 	feud.Barony
    BLipikha利皮哈() 	feud.Barony
    BMalyshi马雷希() 	feud.Barony
    BMiass米阿斯() 	feud.Barony
    BRassvet拉斯韦特() 	feud.Barony
    BShatrovo沙特罗沃() 	feud.Barony
    BSingul辛古利() 	feud.Barony
}

type 米阿斯MiassCounty struct {
	feud.BaseCounty
	库库希Kukushi 	feud.Barony
	利皮哈Lipikha 	feud.Barony
	马雷希Malyshi 	feud.Barony
	米阿斯Miass 	feud.Barony
	拉斯韦特Rassvet 	feud.Barony
	沙特罗沃Shatrovo 	feud.Barony
	辛古利Singul 	feud.Barony
}

func (c *米阿斯MiassCounty) BKukushi库库希() feud.Barony {
	return c.库库希Kukushi
}
    
func (c *米阿斯MiassCounty) BLipikha利皮哈() feud.Barony {
	return c.利皮哈Lipikha
}
    
func (c *米阿斯MiassCounty) BMalyshi马雷希() feud.Barony {
	return c.马雷希Malyshi
}
    
func (c *米阿斯MiassCounty) BMiass米阿斯() feud.Barony {
	return c.米阿斯Miass
}
    
func (c *米阿斯MiassCounty) BRassvet拉斯韦特() feud.Barony {
	return c.拉斯韦特Rassvet
}
    
func (c *米阿斯MiassCounty) BShatrovo沙特罗沃() feud.Barony {
	return c.沙特罗沃Shatrovo
}
    
func (c *米阿斯MiassCounty) BSingul辛古利() feud.Barony {
	return c.辛古利Singul
}
    
var CMiass米阿斯 MiassCounty = &米阿斯MiassCounty{}

func init() {
	f := CMiass米阿斯.(*米阿斯MiassCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1857",
		Title:     "miass",
		TitleName: "米阿斯",
		TitleCode: "c_miass",
		Baronies:  map[string]feud.Barony{},
	}

	f.库库希Kukushi = BKukushi库库希
	f.库库希Kukushi.SetParent(f)
	
	f.利皮哈Lipikha = BLipikha利皮哈
	f.利皮哈Lipikha.SetParent(f)
	
	f.马雷希Malyshi = BMalyshi马雷希
	f.马雷希Malyshi.SetParent(f)
	
	f.米阿斯Miass = BMiass米阿斯
	f.米阿斯Miass.SetParent(f)
	
	f.拉斯韦特Rassvet = BRassvet拉斯韦特
	f.拉斯韦特Rassvet.SetParent(f)
	
	f.沙特罗沃Shatrovo = BShatrovo沙特罗沃
	f.沙特罗沃Shatrovo.SetParent(f)
	
	f.辛古利Singul = BSingul辛古利
	f.辛古利Singul.SetParent(f)
	
}
