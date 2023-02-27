package gobi_altay

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/gobi_altay/khasagt_khairkhan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/gobi_altay/khokh_serkh"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/gobi_altay/sutai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/gobi_altay/tsagaannuur"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Gobi_altayDuke interface {
    feud.Duke
    CKhasagt_khairkhan哈萨格特海尔汗() 	khasagt_khairkhan.Khasagt_khairkhanCounty
    CKhokh_serkh呼赫色尔赫() 	khokh_serkh.Khokh_serkhCounty
    CSutai苏泰() 	sutai.SutaiCounty
    CTsagaannuur扎不罕() 	tsagaannuur.TsagaannuurCounty
}

type 戈壁阿尔泰Gobi_altayDuke struct {
	feud.BaseDuke
	哈萨格特海尔汗Khasagt_khairkhan 	khasagt_khairkhan.Khasagt_khairkhanCounty
	呼赫色尔赫Khokh_serkh 	khokh_serkh.Khokh_serkhCounty
	苏泰Sutai 	sutai.SutaiCounty
	扎不罕Tsagaannuur 	tsagaannuur.TsagaannuurCounty
}

func (d *戈壁阿尔泰Gobi_altayDuke) CKhasagt_khairkhan哈萨格特海尔汗() khasagt_khairkhan.Khasagt_khairkhanCounty {
	return d.哈萨格特海尔汗Khasagt_khairkhan
}
    
func (d *戈壁阿尔泰Gobi_altayDuke) CKhokh_serkh呼赫色尔赫() khokh_serkh.Khokh_serkhCounty {
	return d.呼赫色尔赫Khokh_serkh
}
    
func (d *戈壁阿尔泰Gobi_altayDuke) CSutai苏泰() sutai.SutaiCounty {
	return d.苏泰Sutai
}
    
func (d *戈壁阿尔泰Gobi_altayDuke) CTsagaannuur扎不罕() tsagaannuur.TsagaannuurCounty {
	return d.扎不罕Tsagaannuur
}
    
var DGobi_altay戈壁阿尔泰 Gobi_altayDuke = &戈壁阿尔泰Gobi_altayDuke{}

func init() {
	f := DGobi_altay戈壁阿尔泰.(*戈壁阿尔泰Gobi_altayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gobi_altay",
		TitleName: "戈壁阿尔泰",
		TitleCode: "d_gobi_altay",
		Counties:  map[string]feud.County{},
	}

	f.哈萨格特海尔汗Khasagt_khairkhan = khasagt_khairkhan.CKhasagt_khairkhan哈萨格特海尔汗
	f.哈萨格特海尔汗Khasagt_khairkhan.SetParent(f)
	
	f.呼赫色尔赫Khokh_serkh = khokh_serkh.CKhokh_serkh呼赫色尔赫
	f.呼赫色尔赫Khokh_serkh.SetParent(f)
	
	f.苏泰Sutai = sutai.CSutai苏泰
	f.苏泰Sutai.SetParent(f)
	
	f.扎不罕Tsagaannuur = tsagaannuur.CTsagaannuur扎不罕
	f.扎不罕Tsagaannuur.SetParent(f)
	
}
