package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KanchipuramCounty interface {
    feud.County
    BAlampuravi阿楞补罗毗() 	feud.Barony
    BKanchipuram建志补罗() 	feud.Barony
    BKoodalur句陀楼罗() 	feud.Barony
    BMailapur梅罗补罗() 	feud.Barony
    BRajagiri罗阇耆厘() 	feud.Barony
    BTakkaloma多伽卢摩() 	feud.Barony
    BUttaramerur尤多罗摩罗() 	feud.Barony
}

type 建志补罗KanchipuramCounty struct {
	feud.BaseCounty
	阿楞补罗毗Alampuravi 	feud.Barony
	建志补罗Kanchipuram 	feud.Barony
	句陀楼罗Koodalur 	feud.Barony
	梅罗补罗Mailapur 	feud.Barony
	罗阇耆厘Rajagiri 	feud.Barony
	多伽卢摩Takkaloma 	feud.Barony
	尤多罗摩罗Uttaramerur 	feud.Barony
}

func (c *建志补罗KanchipuramCounty) BAlampuravi阿楞补罗毗() feud.Barony {
	return c.阿楞补罗毗Alampuravi
}
    
func (c *建志补罗KanchipuramCounty) BKanchipuram建志补罗() feud.Barony {
	return c.建志补罗Kanchipuram
}
    
func (c *建志补罗KanchipuramCounty) BKoodalur句陀楼罗() feud.Barony {
	return c.句陀楼罗Koodalur
}
    
func (c *建志补罗KanchipuramCounty) BMailapur梅罗补罗() feud.Barony {
	return c.梅罗补罗Mailapur
}
    
func (c *建志补罗KanchipuramCounty) BRajagiri罗阇耆厘() feud.Barony {
	return c.罗阇耆厘Rajagiri
}
    
func (c *建志补罗KanchipuramCounty) BTakkaloma多伽卢摩() feud.Barony {
	return c.多伽卢摩Takkaloma
}
    
func (c *建志补罗KanchipuramCounty) BUttaramerur尤多罗摩罗() feud.Barony {
	return c.尤多罗摩罗Uttaramerur
}
    
var CKanchipuram建志补罗 KanchipuramCounty = &建志补罗KanchipuramCounty{}

func init() {
	f := CKanchipuram建志补罗.(*建志补罗KanchipuramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1119",
		Title:     "kanchipuram",
		TitleName: "建志补罗",
		TitleCode: "c_kanchipuram",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿楞补罗毗Alampuravi = BAlampuravi阿楞补罗毗
	f.阿楞补罗毗Alampuravi.SetParent(f)
	
	f.建志补罗Kanchipuram = BKanchipuram建志补罗
	f.建志补罗Kanchipuram.SetParent(f)
	
	f.句陀楼罗Koodalur = BKoodalur句陀楼罗
	f.句陀楼罗Koodalur.SetParent(f)
	
	f.梅罗补罗Mailapur = BMailapur梅罗补罗
	f.梅罗补罗Mailapur.SetParent(f)
	
	f.罗阇耆厘Rajagiri = BRajagiri罗阇耆厘
	f.罗阇耆厘Rajagiri.SetParent(f)
	
	f.多伽卢摩Takkaloma = BTakkaloma多伽卢摩
	f.多伽卢摩Takkaloma.SetParent(f)
	
	f.尤多罗摩罗Uttaramerur = BUttaramerur尤多罗摩罗
	f.尤多罗摩罗Uttaramerur.SetParent(f)
	
}
