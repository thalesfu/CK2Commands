package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Khasagt_khairkhanCounty interface {
    feud.County
    BAltai_khasagt阿尔泰() 	feud.Barony
    BBayan_uul巴彦乌拉() 	feud.Barony
    BGegeen格根() 	feud.Barony
    BGuulin_khasagt古林() 	feud.Barony
    BJargalan扎尔嘎朗() 	feud.Barony
    BKhasagt_khairkhan哈萨格特海尔汗() 	feud.Barony
    BSharga沙尔嘎() 	feud.Barony
}

type 哈萨格特海尔汗Khasagt_khairkhanCounty struct {
	feud.BaseCounty
	阿尔泰Altai_khasagt 	feud.Barony
	巴彦乌拉Bayan_uul 	feud.Barony
	格根Gegeen 	feud.Barony
	古林Guulin_khasagt 	feud.Barony
	扎尔嘎朗Jargalan 	feud.Barony
	哈萨格特海尔汗Khasagt_khairkhan 	feud.Barony
	沙尔嘎Sharga 	feud.Barony
}

func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BAltai_khasagt阿尔泰() feud.Barony {
	return c.阿尔泰Altai_khasagt
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BBayan_uul巴彦乌拉() feud.Barony {
	return c.巴彦乌拉Bayan_uul
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BGegeen格根() feud.Barony {
	return c.格根Gegeen
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BGuulin_khasagt古林() feud.Barony {
	return c.古林Guulin_khasagt
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BJargalan扎尔嘎朗() feud.Barony {
	return c.扎尔嘎朗Jargalan
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BKhasagt_khairkhan哈萨格特海尔汗() feud.Barony {
	return c.哈萨格特海尔汗Khasagt_khairkhan
}
    
func (c *哈萨格特海尔汗Khasagt_khairkhanCounty) BSharga沙尔嘎() feud.Barony {
	return c.沙尔嘎Sharga
}
    
var CKhasagt_khairkhan哈萨格特海尔汗 Khasagt_khairkhanCounty = &哈萨格特海尔汗Khasagt_khairkhanCounty{}

func init() {
	f := CKhasagt_khairkhan哈萨格特海尔汗.(*哈萨格特海尔汗Khasagt_khairkhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1911",
		Title:     "khasagt_khairkhan",
		TitleName: "哈萨格特海尔汗",
		TitleCode: "c_khasagt_khairkhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔泰Altai_khasagt = BAltai_khasagt阿尔泰
	f.阿尔泰Altai_khasagt.SetParent(f)
	
	f.巴彦乌拉Bayan_uul = BBayan_uul巴彦乌拉
	f.巴彦乌拉Bayan_uul.SetParent(f)
	
	f.格根Gegeen = BGegeen格根
	f.格根Gegeen.SetParent(f)
	
	f.古林Guulin_khasagt = BGuulin_khasagt古林
	f.古林Guulin_khasagt.SetParent(f)
	
	f.扎尔嘎朗Jargalan = BJargalan扎尔嘎朗
	f.扎尔嘎朗Jargalan.SetParent(f)
	
	f.哈萨格特海尔汗Khasagt_khairkhan = BKhasagt_khairkhan哈萨格特海尔汗
	f.哈萨格特海尔汗Khasagt_khairkhan.SetParent(f)
	
	f.沙尔嘎Sharga = BSharga沙尔嘎
	f.沙尔嘎Sharga.SetParent(f)
	
}
