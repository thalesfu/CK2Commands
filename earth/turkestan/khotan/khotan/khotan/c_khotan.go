package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhotanCounty interface {
    feud.County
    BBojiayi勃伽夷() 	feud.Barony
    BJiandu鞬都() 	feud.Barony
    BJingjue精绝() 	feud.Barony
    BKhotan于阗() 	feud.Barony
    BPishan皮山() 	feud.Barony
    BQira质逻() 	feud.Barony
    BWeiguan苇关() 	feud.Barony
}

type 于阗KhotanCounty struct {
	feud.BaseCounty
	勃伽夷Bojiayi 	feud.Barony
	鞬都Jiandu 	feud.Barony
	精绝Jingjue 	feud.Barony
	于阗Khotan 	feud.Barony
	皮山Pishan 	feud.Barony
	质逻Qira 	feud.Barony
	苇关Weiguan 	feud.Barony
}

func (c *于阗KhotanCounty) BBojiayi勃伽夷() feud.Barony {
	return c.勃伽夷Bojiayi
}
    
func (c *于阗KhotanCounty) BJiandu鞬都() feud.Barony {
	return c.鞬都Jiandu
}
    
func (c *于阗KhotanCounty) BJingjue精绝() feud.Barony {
	return c.精绝Jingjue
}
    
func (c *于阗KhotanCounty) BKhotan于阗() feud.Barony {
	return c.于阗Khotan
}
    
func (c *于阗KhotanCounty) BPishan皮山() feud.Barony {
	return c.皮山Pishan
}
    
func (c *于阗KhotanCounty) BQira质逻() feud.Barony {
	return c.质逻Qira
}
    
func (c *于阗KhotanCounty) BWeiguan苇关() feud.Barony {
	return c.苇关Weiguan
}
    
var CKhotan于阗 KhotanCounty = &于阗KhotanCounty{}

func init() {
	f := CKhotan于阗.(*于阗KhotanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1440",
		Title:     "khotan",
		TitleName: "于阗",
		TitleCode: "c_khotan",
		Baronies:  map[string]feud.Barony{},
	}

	f.勃伽夷Bojiayi = BBojiayi勃伽夷
	f.勃伽夷Bojiayi.SetParent(f)
	
	f.鞬都Jiandu = BJiandu鞬都
	f.鞬都Jiandu.SetParent(f)
	
	f.精绝Jingjue = BJingjue精绝
	f.精绝Jingjue.SetParent(f)
	
	f.于阗Khotan = BKhotan于阗
	f.于阗Khotan.SetParent(f)
	
	f.皮山Pishan = BPishan皮山
	f.皮山Pishan.SetParent(f)
	
	f.质逻Qira = BQira质逻
	f.质逻Qira.SetParent(f)
	
	f.苇关Weiguan = BWeiguan苇关
	f.苇关Weiguan.SetParent(f)
	
}
