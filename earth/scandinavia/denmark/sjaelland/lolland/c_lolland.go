package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LollandCounty interface {
    feud.County
    BBandholm班霍尔姆() 	feud.Barony
    BLolland洛兰() 	feud.Barony
    BMaribo马里博() 	feud.Barony
    BNacascogh纳克斯考() 	feud.Barony
    BNykobing尼克宾() 	feud.Barony
    BRuthby勒兹比() 	feud.Barony
    BSaxakopingh萨克斯克宾() 	feud.Barony
}

type 洛兰LollandCounty struct {
	feud.BaseCounty
	班霍尔姆Bandholm 	feud.Barony
	洛兰Lolland 	feud.Barony
	马里博Maribo 	feud.Barony
	纳克斯考Nacascogh 	feud.Barony
	尼克宾Nykobing 	feud.Barony
	勒兹比Ruthby 	feud.Barony
	萨克斯克宾Saxakopingh 	feud.Barony
}

func (c *洛兰LollandCounty) BBandholm班霍尔姆() feud.Barony {
	return c.班霍尔姆Bandholm
}
    
func (c *洛兰LollandCounty) BLolland洛兰() feud.Barony {
	return c.洛兰Lolland
}
    
func (c *洛兰LollandCounty) BMaribo马里博() feud.Barony {
	return c.马里博Maribo
}
    
func (c *洛兰LollandCounty) BNacascogh纳克斯考() feud.Barony {
	return c.纳克斯考Nacascogh
}
    
func (c *洛兰LollandCounty) BNykobing尼克宾() feud.Barony {
	return c.尼克宾Nykobing
}
    
func (c *洛兰LollandCounty) BRuthby勒兹比() feud.Barony {
	return c.勒兹比Ruthby
}
    
func (c *洛兰LollandCounty) BSaxakopingh萨克斯克宾() feud.Barony {
	return c.萨克斯克宾Saxakopingh
}
    
var CLolland洛兰 LollandCounty = &洛兰LollandCounty{}

func init() {
	f := CLolland洛兰.(*洛兰LollandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1687",
		Title:     "lolland",
		TitleName: "洛兰",
		TitleCode: "c_lolland",
		Baronies:  map[string]feud.Barony{},
	}

	f.班霍尔姆Bandholm = BBandholm班霍尔姆
	f.班霍尔姆Bandholm.SetParent(f)
	
	f.洛兰Lolland = BLolland洛兰
	f.洛兰Lolland.SetParent(f)
	
	f.马里博Maribo = BMaribo马里博
	f.马里博Maribo.SetParent(f)
	
	f.纳克斯考Nacascogh = BNacascogh纳克斯考
	f.纳克斯考Nacascogh.SetParent(f)
	
	f.尼克宾Nykobing = BNykobing尼克宾
	f.尼克宾Nykobing.SetParent(f)
	
	f.勒兹比Ruthby = BRuthby勒兹比
	f.勒兹比Ruthby.SetParent(f)
	
	f.萨克斯克宾Saxakopingh = BSaxakopingh萨克斯克宾
	f.萨克斯克宾Saxakopingh.SetParent(f)
	
}
