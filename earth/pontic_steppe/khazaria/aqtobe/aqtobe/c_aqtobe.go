package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AqtobeCounty interface {
    feud.County
    BAqtobe阿克托别() 	feud.Barony
    BAulkutyrtas奥尔_库特尔塔斯() 	feud.Barony
    BBurta布尔塔() 	feud.Barony
    BKunsay昆赛() 	feud.Barony
    BShilikta希利克塔() 	feud.Barony
    BSibiryak失必儿雅克() 	feud.Barony
    BZharlysay扎尔雷_赛() 	feud.Barony
}

type 阿克托别AqtobeCounty struct {
	feud.BaseCounty
	阿克托别Aqtobe 	feud.Barony
	奥尔_库特尔塔斯Aulkutyrtas 	feud.Barony
	布尔塔Burta 	feud.Barony
	昆赛Kunsay 	feud.Barony
	希利克塔Shilikta 	feud.Barony
	失必儿雅克Sibiryak 	feud.Barony
	扎尔雷_赛Zharlysay 	feud.Barony
}

func (c *阿克托别AqtobeCounty) BAqtobe阿克托别() feud.Barony {
	return c.阿克托别Aqtobe
}
    
func (c *阿克托别AqtobeCounty) BAulkutyrtas奥尔_库特尔塔斯() feud.Barony {
	return c.奥尔_库特尔塔斯Aulkutyrtas
}
    
func (c *阿克托别AqtobeCounty) BBurta布尔塔() feud.Barony {
	return c.布尔塔Burta
}
    
func (c *阿克托别AqtobeCounty) BKunsay昆赛() feud.Barony {
	return c.昆赛Kunsay
}
    
func (c *阿克托别AqtobeCounty) BShilikta希利克塔() feud.Barony {
	return c.希利克塔Shilikta
}
    
func (c *阿克托别AqtobeCounty) BSibiryak失必儿雅克() feud.Barony {
	return c.失必儿雅克Sibiryak
}
    
func (c *阿克托别AqtobeCounty) BZharlysay扎尔雷_赛() feud.Barony {
	return c.扎尔雷_赛Zharlysay
}
    
var CAqtobe阿克托别 AqtobeCounty = &阿克托别AqtobeCounty{}

func init() {
	f := CAqtobe阿克托别.(*阿克托别AqtobeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "896",
		Title:     "aqtobe",
		TitleName: "阿克托别",
		TitleCode: "c_aqtobe",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克托别Aqtobe = BAqtobe阿克托别
	f.阿克托别Aqtobe.SetParent(f)
	
	f.奥尔_库特尔塔斯Aulkutyrtas = BAulkutyrtas奥尔_库特尔塔斯
	f.奥尔_库特尔塔斯Aulkutyrtas.SetParent(f)
	
	f.布尔塔Burta = BBurta布尔塔
	f.布尔塔Burta.SetParent(f)
	
	f.昆赛Kunsay = BKunsay昆赛
	f.昆赛Kunsay.SetParent(f)
	
	f.希利克塔Shilikta = BShilikta希利克塔
	f.希利克塔Shilikta.SetParent(f)
	
	f.失必儿雅克Sibiryak = BSibiryak失必儿雅克
	f.失必儿雅克Sibiryak.SetParent(f)
	
	f.扎尔雷_赛Zharlysay = BZharlysay扎尔雷_赛
	f.扎尔雷_赛Zharlysay.SetParent(f)
	
}
