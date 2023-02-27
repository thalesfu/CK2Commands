package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NavasarikaCounty interface {
    feud.County
    BAkruresvara阿拘卢丽湿伐罗() 	feud.Barony
    BBharuch跋禄羯呫婆() 	feud.Barony
    BDharampur达兰补罗() 	feud.Barony
    BNavasarika那婆萨利迦() 	feud.Barony
    BRajpipla罗阇毕钵罗() 	feud.Barony
    BRander兰多尔() 	feud.Barony
    BSurat苏剌吒() 	feud.Barony
    BTadkeshwar呾计湿伐罗() 	feud.Barony
    BVariav伐里瓦() 	feud.Barony
}

type 那婆萨利迦NavasarikaCounty struct {
	feud.BaseCounty
	阿拘卢丽湿伐罗Akruresvara 	feud.Barony
	跋禄羯呫婆Bharuch 	feud.Barony
	达兰补罗Dharampur 	feud.Barony
	那婆萨利迦Navasarika 	feud.Barony
	罗阇毕钵罗Rajpipla 	feud.Barony
	兰多尔Rander 	feud.Barony
	苏剌吒Surat 	feud.Barony
	呾计湿伐罗Tadkeshwar 	feud.Barony
	伐里瓦Variav 	feud.Barony
}

func (c *那婆萨利迦NavasarikaCounty) BAkruresvara阿拘卢丽湿伐罗() feud.Barony {
	return c.阿拘卢丽湿伐罗Akruresvara
}
    
func (c *那婆萨利迦NavasarikaCounty) BBharuch跋禄羯呫婆() feud.Barony {
	return c.跋禄羯呫婆Bharuch
}
    
func (c *那婆萨利迦NavasarikaCounty) BDharampur达兰补罗() feud.Barony {
	return c.达兰补罗Dharampur
}
    
func (c *那婆萨利迦NavasarikaCounty) BNavasarika那婆萨利迦() feud.Barony {
	return c.那婆萨利迦Navasarika
}
    
func (c *那婆萨利迦NavasarikaCounty) BRajpipla罗阇毕钵罗() feud.Barony {
	return c.罗阇毕钵罗Rajpipla
}
    
func (c *那婆萨利迦NavasarikaCounty) BRander兰多尔() feud.Barony {
	return c.兰多尔Rander
}
    
func (c *那婆萨利迦NavasarikaCounty) BSurat苏剌吒() feud.Barony {
	return c.苏剌吒Surat
}
    
func (c *那婆萨利迦NavasarikaCounty) BTadkeshwar呾计湿伐罗() feud.Barony {
	return c.呾计湿伐罗Tadkeshwar
}
    
func (c *那婆萨利迦NavasarikaCounty) BVariav伐里瓦() feud.Barony {
	return c.伐里瓦Variav
}
    
var CNavasarika那婆萨利迦 NavasarikaCounty = &那婆萨利迦NavasarikaCounty{}

func init() {
	f := CNavasarika那婆萨利迦.(*那婆萨利迦NavasarikaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1127",
		Title:     "navasarika",
		TitleName: "那婆萨利迦",
		TitleCode: "c_navasarika",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拘卢丽湿伐罗Akruresvara = BAkruresvara阿拘卢丽湿伐罗
	f.阿拘卢丽湿伐罗Akruresvara.SetParent(f)
	
	f.跋禄羯呫婆Bharuch = BBharuch跋禄羯呫婆
	f.跋禄羯呫婆Bharuch.SetParent(f)
	
	f.达兰补罗Dharampur = BDharampur达兰补罗
	f.达兰补罗Dharampur.SetParent(f)
	
	f.那婆萨利迦Navasarika = BNavasarika那婆萨利迦
	f.那婆萨利迦Navasarika.SetParent(f)
	
	f.罗阇毕钵罗Rajpipla = BRajpipla罗阇毕钵罗
	f.罗阇毕钵罗Rajpipla.SetParent(f)
	
	f.兰多尔Rander = BRander兰多尔
	f.兰多尔Rander.SetParent(f)
	
	f.苏剌吒Surat = BSurat苏剌吒
	f.苏剌吒Surat.SetParent(f)
	
	f.呾计湿伐罗Tadkeshwar = BTadkeshwar呾计湿伐罗
	f.呾计湿伐罗Tadkeshwar.SetParent(f)
	
	f.伐里瓦Variav = BVariav伐里瓦
	f.伐里瓦Variav.SetParent(f)
	
}
