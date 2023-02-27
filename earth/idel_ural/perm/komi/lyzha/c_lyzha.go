package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LyzhaCounty interface {
    feud.County
    BAkis阿基斯() 	feud.Barony
    BLyzha雷扎() 	feud.Barony
    BMaterik马捷里克() 	feud.Barony
    BMutnyy穆特内() 	feud.Barony
    BNovikobzh诺维科布日() 	feud.Barony
    BUst_usa乌斯季_乌萨() 	feud.Barony
    BZakharvan扎哈尔万() 	feud.Barony
}

type 雷扎LyzhaCounty struct {
	feud.BaseCounty
	阿基斯Akis 	feud.Barony
	雷扎Lyzha 	feud.Barony
	马捷里克Materik 	feud.Barony
	穆特内Mutnyy 	feud.Barony
	诺维科布日Novikobzh 	feud.Barony
	乌斯季_乌萨Ust_usa 	feud.Barony
	扎哈尔万Zakharvan 	feud.Barony
}

func (c *雷扎LyzhaCounty) BAkis阿基斯() feud.Barony {
	return c.阿基斯Akis
}
    
func (c *雷扎LyzhaCounty) BLyzha雷扎() feud.Barony {
	return c.雷扎Lyzha
}
    
func (c *雷扎LyzhaCounty) BMaterik马捷里克() feud.Barony {
	return c.马捷里克Materik
}
    
func (c *雷扎LyzhaCounty) BMutnyy穆特内() feud.Barony {
	return c.穆特内Mutnyy
}
    
func (c *雷扎LyzhaCounty) BNovikobzh诺维科布日() feud.Barony {
	return c.诺维科布日Novikobzh
}
    
func (c *雷扎LyzhaCounty) BUst_usa乌斯季_乌萨() feud.Barony {
	return c.乌斯季_乌萨Ust_usa
}
    
func (c *雷扎LyzhaCounty) BZakharvan扎哈尔万() feud.Barony {
	return c.扎哈尔万Zakharvan
}
    
var CLyzha雷扎 LyzhaCounty = &雷扎LyzhaCounty{}

func init() {
	f := CLyzha雷扎.(*雷扎LyzhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1827",
		Title:     "lyzha",
		TitleName: "雷扎",
		TitleCode: "c_lyzha",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿基斯Akis = BAkis阿基斯
	f.阿基斯Akis.SetParent(f)
	
	f.雷扎Lyzha = BLyzha雷扎
	f.雷扎Lyzha.SetParent(f)
	
	f.马捷里克Materik = BMaterik马捷里克
	f.马捷里克Materik.SetParent(f)
	
	f.穆特内Mutnyy = BMutnyy穆特内
	f.穆特内Mutnyy.SetParent(f)
	
	f.诺维科布日Novikobzh = BNovikobzh诺维科布日
	f.诺维科布日Novikobzh.SetParent(f)
	
	f.乌斯季_乌萨Ust_usa = BUst_usa乌斯季_乌萨
	f.乌斯季_乌萨Ust_usa.SetParent(f)
	
	f.扎哈尔万Zakharvan = BZakharvan扎哈尔万
	f.扎哈尔万Zakharvan.SetParent(f)
	
}
