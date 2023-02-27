package ossory

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OssoryCounty interface {
    feud.County
    BAghaboe阿哈波() 	feud.Barony
    BCallan卡伦() 	feud.Barony
    BClonmacnoise克隆马克诺伊斯() 	feud.Barony
    BGowran高兰() 	feud.Barony
    BGrannagh格兰纳() 	feud.Barony
    BGrennan格雷南() 	feud.Barony
    BJerpoint杰洛因() 	feud.Barony
    BKilkenny基尔肯尼() 	feud.Barony
}

type 奥索里OssoryCounty struct {
	feud.BaseCounty
	阿哈波Aghaboe 	feud.Barony
	卡伦Callan 	feud.Barony
	克隆马克诺伊斯Clonmacnoise 	feud.Barony
	高兰Gowran 	feud.Barony
	格兰纳Grannagh 	feud.Barony
	格雷南Grennan 	feud.Barony
	杰洛因Jerpoint 	feud.Barony
	基尔肯尼Kilkenny 	feud.Barony
}

func (c *奥索里OssoryCounty) BAghaboe阿哈波() feud.Barony {
	return c.阿哈波Aghaboe
}
    
func (c *奥索里OssoryCounty) BCallan卡伦() feud.Barony {
	return c.卡伦Callan
}
    
func (c *奥索里OssoryCounty) BClonmacnoise克隆马克诺伊斯() feud.Barony {
	return c.克隆马克诺伊斯Clonmacnoise
}
    
func (c *奥索里OssoryCounty) BGowran高兰() feud.Barony {
	return c.高兰Gowran
}
    
func (c *奥索里OssoryCounty) BGrannagh格兰纳() feud.Barony {
	return c.格兰纳Grannagh
}
    
func (c *奥索里OssoryCounty) BGrennan格雷南() feud.Barony {
	return c.格雷南Grennan
}
    
func (c *奥索里OssoryCounty) BJerpoint杰洛因() feud.Barony {
	return c.杰洛因Jerpoint
}
    
func (c *奥索里OssoryCounty) BKilkenny基尔肯尼() feud.Barony {
	return c.基尔肯尼Kilkenny
}
    
var COssory奥索里 OssoryCounty = &奥索里OssoryCounty{}

func init() {
	f := COssory奥索里.(*奥索里OssoryCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "12",
		Title:     "ossory",
		TitleName: "奥索里",
		TitleCode: "c_ossory",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿哈波Aghaboe = BAghaboe阿哈波
	f.阿哈波Aghaboe.SetParent(f)
	
	f.卡伦Callan = BCallan卡伦
	f.卡伦Callan.SetParent(f)
	
	f.克隆马克诺伊斯Clonmacnoise = BClonmacnoise克隆马克诺伊斯
	f.克隆马克诺伊斯Clonmacnoise.SetParent(f)
	
	f.高兰Gowran = BGowran高兰
	f.高兰Gowran.SetParent(f)
	
	f.格兰纳Grannagh = BGrannagh格兰纳
	f.格兰纳Grannagh.SetParent(f)
	
	f.格雷南Grennan = BGrennan格雷南
	f.格雷南Grennan.SetParent(f)
	
	f.杰洛因Jerpoint = BJerpoint杰洛因
	f.杰洛因Jerpoint.SetParent(f)
	
	f.基尔肯尼Kilkenny = BKilkenny基尔肯尼
	f.基尔肯尼Kilkenny.SetParent(f)
	
}
