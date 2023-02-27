package khangai

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/khangai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/otgontenger"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/suvraga_khairkhan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/tarvagatai"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/khangai/terkhiin_tsagaan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhangaiDuke interface {
    feud.Duke
    CKhangai杭海() 	khangai.KhangaiCounty
    COtgontenger鄂特冈腾格里() 	otgontenger.OtgontengerCounty
    CSuvraga_khairkhan苏布日嘎海尔汗() 	suvraga_khairkhan.Suvraga_khairkhanCounty
    CTarvagatai塔尔巴嘎泰() 	tarvagatai.TarvagataiCounty
    CTerkhiin_tsagaan特尔欣查干() 	terkhiin_tsagaan.Terkhiin_tsagaanCounty
}

type 杭海KhangaiDuke struct {
	feud.BaseDuke
	杭海Khangai 	khangai.KhangaiCounty
	鄂特冈腾格里Otgontenger 	otgontenger.OtgontengerCounty
	苏布日嘎海尔汗Suvraga_khairkhan 	suvraga_khairkhan.Suvraga_khairkhanCounty
	塔尔巴嘎泰Tarvagatai 	tarvagatai.TarvagataiCounty
	特尔欣查干Terkhiin_tsagaan 	terkhiin_tsagaan.Terkhiin_tsagaanCounty
}

func (d *杭海KhangaiDuke) CKhangai杭海() khangai.KhangaiCounty {
	return d.杭海Khangai
}
    
func (d *杭海KhangaiDuke) COtgontenger鄂特冈腾格里() otgontenger.OtgontengerCounty {
	return d.鄂特冈腾格里Otgontenger
}
    
func (d *杭海KhangaiDuke) CSuvraga_khairkhan苏布日嘎海尔汗() suvraga_khairkhan.Suvraga_khairkhanCounty {
	return d.苏布日嘎海尔汗Suvraga_khairkhan
}
    
func (d *杭海KhangaiDuke) CTarvagatai塔尔巴嘎泰() tarvagatai.TarvagataiCounty {
	return d.塔尔巴嘎泰Tarvagatai
}
    
func (d *杭海KhangaiDuke) CTerkhiin_tsagaan特尔欣查干() terkhiin_tsagaan.Terkhiin_tsagaanCounty {
	return d.特尔欣查干Terkhiin_tsagaan
}
    
var DKhangai杭海 KhangaiDuke = &杭海KhangaiDuke{}

func init() {
	f := DKhangai杭海.(*杭海KhangaiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khangai",
		TitleName: "杭海",
		TitleCode: "d_khangai",
		Counties:  map[string]feud.County{},
	}

	f.杭海Khangai = khangai.CKhangai杭海
	f.杭海Khangai.SetParent(f)
	
	f.鄂特冈腾格里Otgontenger = otgontenger.COtgontenger鄂特冈腾格里
	f.鄂特冈腾格里Otgontenger.SetParent(f)
	
	f.苏布日嘎海尔汗Suvraga_khairkhan = suvraga_khairkhan.CSuvraga_khairkhan苏布日嘎海尔汗
	f.苏布日嘎海尔汗Suvraga_khairkhan.SetParent(f)
	
	f.塔尔巴嘎泰Tarvagatai = tarvagatai.CTarvagatai塔尔巴嘎泰
	f.塔尔巴嘎泰Tarvagatai.SetParent(f)
	
	f.特尔欣查干Terkhiin_tsagaan = terkhiin_tsagaan.CTerkhiin_tsagaan特尔欣查干
	f.特尔欣查干Terkhiin_tsagaan.SetParent(f)
	
}
