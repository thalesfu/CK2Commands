package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Veliky_ustugCounty interface {
    feud.County
    BGleden格列坚() 	feud.Barony
    BKrasavino克拉萨维诺() 	feud.Barony
    BLuza卢扎() 	feud.Barony
    BMaromitsa马罗米察() 	feud.Barony
    BOparino奥帕里诺() 	feud.Barony
    BPinyug皮纽格() 	feud.Barony
    BPodosinovets波多西诺韦茨() 	feud.Barony
    BVelikyustug大乌斯秋格() 	feud.Barony
}

type 大乌斯秋格Veliky_ustugCounty struct {
	feud.BaseCounty
	格列坚Gleden 	feud.Barony
	克拉萨维诺Krasavino 	feud.Barony
	卢扎Luza 	feud.Barony
	马罗米察Maromitsa 	feud.Barony
	奥帕里诺Oparino 	feud.Barony
	皮纽格Pinyug 	feud.Barony
	波多西诺韦茨Podosinovets 	feud.Barony
	大乌斯秋格Velikyustug 	feud.Barony
}

func (c *大乌斯秋格Veliky_ustugCounty) BGleden格列坚() feud.Barony {
	return c.格列坚Gleden
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BKrasavino克拉萨维诺() feud.Barony {
	return c.克拉萨维诺Krasavino
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BLuza卢扎() feud.Barony {
	return c.卢扎Luza
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BMaromitsa马罗米察() feud.Barony {
	return c.马罗米察Maromitsa
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BOparino奥帕里诺() feud.Barony {
	return c.奥帕里诺Oparino
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BPinyug皮纽格() feud.Barony {
	return c.皮纽格Pinyug
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BPodosinovets波多西诺韦茨() feud.Barony {
	return c.波多西诺韦茨Podosinovets
}
    
func (c *大乌斯秋格Veliky_ustugCounty) BVelikyustug大乌斯秋格() feud.Barony {
	return c.大乌斯秋格Velikyustug
}
    
var CVeliky_ustug大乌斯秋格 Veliky_ustugCounty = &大乌斯秋格Veliky_ustugCounty{}

func init() {
	f := CVeliky_ustug大乌斯秋格.(*大乌斯秋格Veliky_ustugCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "402",
		Title:     "veliky_ustug",
		TitleName: "大乌斯秋格",
		TitleCode: "c_veliky_ustug",
		Baronies:  map[string]feud.Barony{},
	}

	f.格列坚Gleden = BGleden格列坚
	f.格列坚Gleden.SetParent(f)
	
	f.克拉萨维诺Krasavino = BKrasavino克拉萨维诺
	f.克拉萨维诺Krasavino.SetParent(f)
	
	f.卢扎Luza = BLuza卢扎
	f.卢扎Luza.SetParent(f)
	
	f.马罗米察Maromitsa = BMaromitsa马罗米察
	f.马罗米察Maromitsa.SetParent(f)
	
	f.奥帕里诺Oparino = BOparino奥帕里诺
	f.奥帕里诺Oparino.SetParent(f)
	
	f.皮纽格Pinyug = BPinyug皮纽格
	f.皮纽格Pinyug.SetParent(f)
	
	f.波多西诺韦茨Podosinovets = BPodosinovets波多西诺韦茨
	f.波多西诺韦茨Podosinovets.SetParent(f)
	
	f.大乌斯秋格Velikyustug = BVelikyustug大乌斯秋格
	f.大乌斯秋格Velikyustug.SetParent(f)
	
}
