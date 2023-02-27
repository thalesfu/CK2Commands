package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Syr_daryaCounty interface {
    feud.County
    BAkmechet阿克_梅切特() 	feud.Barony
    BDzhanadzhol贾纳_焦尔() 	feud.Barony
    BKazaly哈扎雷() 	feud.Barony
    BKyzylorda克孜勒奥尔达() 	feud.Barony
    BSarytogay萨雷托盖() 	feud.Barony
    BSyganak昔格纳黑() 	feud.Barony
    BTerenuzyak捷连_乌兹亚克() 	feud.Barony
    BZhalagash扎拉加什() 	feud.Barony
}

type 昔格纳黑Syr_daryaCounty struct {
	feud.BaseCounty
	阿克_梅切特Akmechet 	feud.Barony
	贾纳_焦尔Dzhanadzhol 	feud.Barony
	哈扎雷Kazaly 	feud.Barony
	克孜勒奥尔达Kyzylorda 	feud.Barony
	萨雷托盖Sarytogay 	feud.Barony
	昔格纳黑Syganak 	feud.Barony
	捷连_乌兹亚克Terenuzyak 	feud.Barony
	扎拉加什Zhalagash 	feud.Barony
}

func (c *昔格纳黑Syr_daryaCounty) BAkmechet阿克_梅切特() feud.Barony {
	return c.阿克_梅切特Akmechet
}
    
func (c *昔格纳黑Syr_daryaCounty) BDzhanadzhol贾纳_焦尔() feud.Barony {
	return c.贾纳_焦尔Dzhanadzhol
}
    
func (c *昔格纳黑Syr_daryaCounty) BKazaly哈扎雷() feud.Barony {
	return c.哈扎雷Kazaly
}
    
func (c *昔格纳黑Syr_daryaCounty) BKyzylorda克孜勒奥尔达() feud.Barony {
	return c.克孜勒奥尔达Kyzylorda
}
    
func (c *昔格纳黑Syr_daryaCounty) BSarytogay萨雷托盖() feud.Barony {
	return c.萨雷托盖Sarytogay
}
    
func (c *昔格纳黑Syr_daryaCounty) BSyganak昔格纳黑() feud.Barony {
	return c.昔格纳黑Syganak
}
    
func (c *昔格纳黑Syr_daryaCounty) BTerenuzyak捷连_乌兹亚克() feud.Barony {
	return c.捷连_乌兹亚克Terenuzyak
}
    
func (c *昔格纳黑Syr_daryaCounty) BZhalagash扎拉加什() feud.Barony {
	return c.扎拉加什Zhalagash
}
    
var CSyr_darya昔格纳黑 Syr_daryaCounty = &昔格纳黑Syr_daryaCounty{}

func init() {
	f := CSyr_darya昔格纳黑.(*昔格纳黑Syr_daryaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "900",
		Title:     "syr_darya",
		TitleName: "昔格纳黑",
		TitleCode: "c_syr_darya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克_梅切特Akmechet = BAkmechet阿克_梅切特
	f.阿克_梅切特Akmechet.SetParent(f)
	
	f.贾纳_焦尔Dzhanadzhol = BDzhanadzhol贾纳_焦尔
	f.贾纳_焦尔Dzhanadzhol.SetParent(f)
	
	f.哈扎雷Kazaly = BKazaly哈扎雷
	f.哈扎雷Kazaly.SetParent(f)
	
	f.克孜勒奥尔达Kyzylorda = BKyzylorda克孜勒奥尔达
	f.克孜勒奥尔达Kyzylorda.SetParent(f)
	
	f.萨雷托盖Sarytogay = BSarytogay萨雷托盖
	f.萨雷托盖Sarytogay.SetParent(f)
	
	f.昔格纳黑Syganak = BSyganak昔格纳黑
	f.昔格纳黑Syganak.SetParent(f)
	
	f.捷连_乌兹亚克Terenuzyak = BTerenuzyak捷连_乌兹亚克
	f.捷连_乌兹亚克Terenuzyak.SetParent(f)
	
	f.扎拉加什Zhalagash = BZhalagash扎拉加什
	f.扎拉加什Zhalagash.SetParent(f)
	
}
