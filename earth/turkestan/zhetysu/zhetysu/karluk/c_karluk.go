package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarlukCounty interface {
    feud.County
    BAlishi阿利施() 	feud.Barony
    BNushibi弩失毕() 	feud.Barony
    BSakla_baga索葛莫贺() 	feud.Barony
    BSaryshaghan萨雷沙甘() 	feud.Barony
    BShagkol沙格阔勒() 	feud.Barony
    BUlug_ok胡禄屋() 	feud.Barony
    BZhaylaukol扎伊劳科利() 	feud.Barony
}

type 葛逻禄KarlukCounty struct {
	feud.BaseCounty
	阿利施Alishi 	feud.Barony
	弩失毕Nushibi 	feud.Barony
	索葛莫贺Sakla_baga 	feud.Barony
	萨雷沙甘Saryshaghan 	feud.Barony
	沙格阔勒Shagkol 	feud.Barony
	胡禄屋Ulug_ok 	feud.Barony
	扎伊劳科利Zhaylaukol 	feud.Barony
}

func (c *葛逻禄KarlukCounty) BAlishi阿利施() feud.Barony {
	return c.阿利施Alishi
}
    
func (c *葛逻禄KarlukCounty) BNushibi弩失毕() feud.Barony {
	return c.弩失毕Nushibi
}
    
func (c *葛逻禄KarlukCounty) BSakla_baga索葛莫贺() feud.Barony {
	return c.索葛莫贺Sakla_baga
}
    
func (c *葛逻禄KarlukCounty) BSaryshaghan萨雷沙甘() feud.Barony {
	return c.萨雷沙甘Saryshaghan
}
    
func (c *葛逻禄KarlukCounty) BShagkol沙格阔勒() feud.Barony {
	return c.沙格阔勒Shagkol
}
    
func (c *葛逻禄KarlukCounty) BUlug_ok胡禄屋() feud.Barony {
	return c.胡禄屋Ulug_ok
}
    
func (c *葛逻禄KarlukCounty) BZhaylaukol扎伊劳科利() feud.Barony {
	return c.扎伊劳科利Zhaylaukol
}
    
var CKarluk葛逻禄 KarlukCounty = &葛逻禄KarlukCounty{}

func init() {
	f := CKarluk葛逻禄.(*葛逻禄KarlukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1432",
		Title:     "karluk",
		TitleName: "葛逻禄",
		TitleCode: "c_karluk",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利施Alishi = BAlishi阿利施
	f.阿利施Alishi.SetParent(f)
	
	f.弩失毕Nushibi = BNushibi弩失毕
	f.弩失毕Nushibi.SetParent(f)
	
	f.索葛莫贺Sakla_baga = BSakla_baga索葛莫贺
	f.索葛莫贺Sakla_baga.SetParent(f)
	
	f.萨雷沙甘Saryshaghan = BSaryshaghan萨雷沙甘
	f.萨雷沙甘Saryshaghan.SetParent(f)
	
	f.沙格阔勒Shagkol = BShagkol沙格阔勒
	f.沙格阔勒Shagkol.SetParent(f)
	
	f.胡禄屋Ulug_ok = BUlug_ok胡禄屋
	f.胡禄屋Ulug_ok.SetParent(f)
	
	f.扎伊劳科利Zhaylaukol = BZhaylaukol扎伊劳科利
	f.扎伊劳科利Zhaylaukol.SetParent(f)
	
}
