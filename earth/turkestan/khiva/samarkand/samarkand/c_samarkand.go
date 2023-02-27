package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SamarkandCounty interface {
    feud.County
    BAfrasiyab阿夫拉西亚卜() 	feud.Barony
    BIshtikhon瑟底痕() 	feud.Barony
    BKhokand浩罕() 	feud.Barony
    BKoshrabot科什拉巴德() 	feud.Barony
    BLaish莱什() 	feud.Barony
    BSamarkand撒马尔罕() 	feud.Barony
    BUrgut乌尔古特() 	feud.Barony
    BZiadin济亚丁() 	feud.Barony
}

type 撒马尔罕SamarkandCounty struct {
	feud.BaseCounty
	阿夫拉西亚卜Afrasiyab 	feud.Barony
	瑟底痕Ishtikhon 	feud.Barony
	浩罕Khokand 	feud.Barony
	科什拉巴德Koshrabot 	feud.Barony
	莱什Laish 	feud.Barony
	撒马尔罕Samarkand 	feud.Barony
	乌尔古特Urgut 	feud.Barony
	济亚丁Ziadin 	feud.Barony
}

func (c *撒马尔罕SamarkandCounty) BAfrasiyab阿夫拉西亚卜() feud.Barony {
	return c.阿夫拉西亚卜Afrasiyab
}
    
func (c *撒马尔罕SamarkandCounty) BIshtikhon瑟底痕() feud.Barony {
	return c.瑟底痕Ishtikhon
}
    
func (c *撒马尔罕SamarkandCounty) BKhokand浩罕() feud.Barony {
	return c.浩罕Khokand
}
    
func (c *撒马尔罕SamarkandCounty) BKoshrabot科什拉巴德() feud.Barony {
	return c.科什拉巴德Koshrabot
}
    
func (c *撒马尔罕SamarkandCounty) BLaish莱什() feud.Barony {
	return c.莱什Laish
}
    
func (c *撒马尔罕SamarkandCounty) BSamarkand撒马尔罕() feud.Barony {
	return c.撒马尔罕Samarkand
}
    
func (c *撒马尔罕SamarkandCounty) BUrgut乌尔古特() feud.Barony {
	return c.乌尔古特Urgut
}
    
func (c *撒马尔罕SamarkandCounty) BZiadin济亚丁() feud.Barony {
	return c.济亚丁Ziadin
}
    
var CSamarkand撒马尔罕 SamarkandCounty = &撒马尔罕SamarkandCounty{}

func init() {
	f := CSamarkand撒马尔罕.(*撒马尔罕SamarkandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "903",
		Title:     "samarkand",
		TitleName: "撒马尔罕",
		TitleCode: "c_samarkand",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫拉西亚卜Afrasiyab = BAfrasiyab阿夫拉西亚卜
	f.阿夫拉西亚卜Afrasiyab.SetParent(f)
	
	f.瑟底痕Ishtikhon = BIshtikhon瑟底痕
	f.瑟底痕Ishtikhon.SetParent(f)
	
	f.浩罕Khokand = BKhokand浩罕
	f.浩罕Khokand.SetParent(f)
	
	f.科什拉巴德Koshrabot = BKoshrabot科什拉巴德
	f.科什拉巴德Koshrabot.SetParent(f)
	
	f.莱什Laish = BLaish莱什
	f.莱什Laish.SetParent(f)
	
	f.撒马尔罕Samarkand = BSamarkand撒马尔罕
	f.撒马尔罕Samarkand.SetParent(f)
	
	f.乌尔古特Urgut = BUrgut乌尔古特
	f.乌尔古特Urgut.SetParent(f)
	
	f.济亚丁Ziadin = BZiadin济亚丁
	f.济亚丁Ziadin.SetParent(f)
	
}
