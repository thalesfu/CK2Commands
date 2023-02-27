package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KajaaniCounty interface {
    feud.County
    BHaslam哈斯拉姆() 	feud.Barony
    BKainuu凯努() 	feud.Barony
    BKajaani卡亚尼() 	feud.Barony
    BKiehiman基耶希梅() 	feud.Barony
    BNyflax尼夫拉克斯() 	feud.Barony
    BPaltamo帕尔塔莫() 	feud.Barony
    BSarismaki萨里斯马基() 	feud.Barony
}

type 凯努KajaaniCounty struct {
	feud.BaseCounty
	哈斯拉姆Haslam 	feud.Barony
	凯努Kainuu 	feud.Barony
	卡亚尼Kajaani 	feud.Barony
	基耶希梅Kiehiman 	feud.Barony
	尼夫拉克斯Nyflax 	feud.Barony
	帕尔塔莫Paltamo 	feud.Barony
	萨里斯马基Sarismaki 	feud.Barony
}

func (c *凯努KajaaniCounty) BHaslam哈斯拉姆() feud.Barony {
	return c.哈斯拉姆Haslam
}
    
func (c *凯努KajaaniCounty) BKainuu凯努() feud.Barony {
	return c.凯努Kainuu
}
    
func (c *凯努KajaaniCounty) BKajaani卡亚尼() feud.Barony {
	return c.卡亚尼Kajaani
}
    
func (c *凯努KajaaniCounty) BKiehiman基耶希梅() feud.Barony {
	return c.基耶希梅Kiehiman
}
    
func (c *凯努KajaaniCounty) BNyflax尼夫拉克斯() feud.Barony {
	return c.尼夫拉克斯Nyflax
}
    
func (c *凯努KajaaniCounty) BPaltamo帕尔塔莫() feud.Barony {
	return c.帕尔塔莫Paltamo
}
    
func (c *凯努KajaaniCounty) BSarismaki萨里斯马基() feud.Barony {
	return c.萨里斯马基Sarismaki
}
    
var CKajaani凯努 KajaaniCounty = &凯努KajaaniCounty{}

func init() {
	f := CKajaani凯努.(*凯努KajaaniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1601",
		Title:     "kajaani",
		TitleName: "凯努",
		TitleCode: "c_kajaani",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈斯拉姆Haslam = BHaslam哈斯拉姆
	f.哈斯拉姆Haslam.SetParent(f)
	
	f.凯努Kainuu = BKainuu凯努
	f.凯努Kainuu.SetParent(f)
	
	f.卡亚尼Kajaani = BKajaani卡亚尼
	f.卡亚尼Kajaani.SetParent(f)
	
	f.基耶希梅Kiehiman = BKiehiman基耶希梅
	f.基耶希梅Kiehiman.SetParent(f)
	
	f.尼夫拉克斯Nyflax = BNyflax尼夫拉克斯
	f.尼夫拉克斯Nyflax.SetParent(f)
	
	f.帕尔塔莫Paltamo = BPaltamo帕尔塔莫
	f.帕尔塔莫Paltamo.SetParent(f)
	
	f.萨里斯马基Sarismaki = BSarismaki萨里斯马基
	f.萨里斯马基Sarismaki.SetParent(f)
	
}
