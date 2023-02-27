package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZhmudCounty interface {
    feud.County
    BJurbarkas尤尔巴尔卡斯() 	feud.Barony
    BKnituva克尼图瓦() 	feud.Barony
    BKraziai克拉日艾() 	feud.Barony
    BRaseiniai拉塞尼艾() 	feud.Barony
    BSiautuna希奥图纳() 	feud.Barony
    BTverai特韦赖() 	feud.Barony
    BZarenai扎雷奈() 	feud.Barony
}

type 萨莫吉提亚ZhmudCounty struct {
	feud.BaseCounty
	尤尔巴尔卡斯Jurbarkas 	feud.Barony
	克尼图瓦Knituva 	feud.Barony
	克拉日艾Kraziai 	feud.Barony
	拉塞尼艾Raseiniai 	feud.Barony
	希奥图纳Siautuna 	feud.Barony
	特韦赖Tverai 	feud.Barony
	扎雷奈Zarenai 	feud.Barony
}

func (c *萨莫吉提亚ZhmudCounty) BJurbarkas尤尔巴尔卡斯() feud.Barony {
	return c.尤尔巴尔卡斯Jurbarkas
}
    
func (c *萨莫吉提亚ZhmudCounty) BKnituva克尼图瓦() feud.Barony {
	return c.克尼图瓦Knituva
}
    
func (c *萨莫吉提亚ZhmudCounty) BKraziai克拉日艾() feud.Barony {
	return c.克拉日艾Kraziai
}
    
func (c *萨莫吉提亚ZhmudCounty) BRaseiniai拉塞尼艾() feud.Barony {
	return c.拉塞尼艾Raseiniai
}
    
func (c *萨莫吉提亚ZhmudCounty) BSiautuna希奥图纳() feud.Barony {
	return c.希奥图纳Siautuna
}
    
func (c *萨莫吉提亚ZhmudCounty) BTverai特韦赖() feud.Barony {
	return c.特韦赖Tverai
}
    
func (c *萨莫吉提亚ZhmudCounty) BZarenai扎雷奈() feud.Barony {
	return c.扎雷奈Zarenai
}
    
var CZhmud萨莫吉提亚 ZhmudCounty = &萨莫吉提亚ZhmudCounty{}

func init() {
	f := CZhmud萨莫吉提亚.(*萨莫吉提亚ZhmudCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "421",
		Title:     "zhmud",
		TitleName: "萨莫吉提亚",
		TitleCode: "c_zhmud",
		Baronies:  map[string]feud.Barony{},
	}

	f.尤尔巴尔卡斯Jurbarkas = BJurbarkas尤尔巴尔卡斯
	f.尤尔巴尔卡斯Jurbarkas.SetParent(f)
	
	f.克尼图瓦Knituva = BKnituva克尼图瓦
	f.克尼图瓦Knituva.SetParent(f)
	
	f.克拉日艾Kraziai = BKraziai克拉日艾
	f.克拉日艾Kraziai.SetParent(f)
	
	f.拉塞尼艾Raseiniai = BRaseiniai拉塞尼艾
	f.拉塞尼艾Raseiniai.SetParent(f)
	
	f.希奥图纳Siautuna = BSiautuna希奥图纳
	f.希奥图纳Siautuna.SetParent(f)
	
	f.特韦赖Tverai = BTverai特韦赖
	f.特韦赖Tverai.SetParent(f)
	
	f.扎雷奈Zarenai = BZarenai扎雷奈
	f.扎雷奈Zarenai.SetParent(f)
	
}
