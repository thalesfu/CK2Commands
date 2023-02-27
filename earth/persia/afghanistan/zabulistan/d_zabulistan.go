package zabulistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/zabulistan/ghazna"
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/zabulistan/kalat"
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/zabulistan/zamindawar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZabulistanDuke interface {
    feud.Duke
    CGhazna伽色尼() 	ghazna.GhaznaCounty
    CKalat卡拉特() 	kalat.KalatCounty
    CZamindawar扎敏达瓦尔() 	zamindawar.ZamindawarCounty
}

type 社护罗萨他那ZabulistanDuke struct {
	feud.BaseDuke
	伽色尼Ghazna 	ghazna.GhaznaCounty
	卡拉特Kalat 	kalat.KalatCounty
	扎敏达瓦尔Zamindawar 	zamindawar.ZamindawarCounty
}

func (d *社护罗萨他那ZabulistanDuke) CGhazna伽色尼() ghazna.GhaznaCounty {
	return d.伽色尼Ghazna
}
    
func (d *社护罗萨他那ZabulistanDuke) CKalat卡拉特() kalat.KalatCounty {
	return d.卡拉特Kalat
}
    
func (d *社护罗萨他那ZabulistanDuke) CZamindawar扎敏达瓦尔() zamindawar.ZamindawarCounty {
	return d.扎敏达瓦尔Zamindawar
}
    
var DZabulistan社护罗萨他那 ZabulistanDuke = &社护罗萨他那ZabulistanDuke{}

func init() {
	f := DZabulistan社护罗萨他那.(*社护罗萨他那ZabulistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "zabulistan",
		TitleName: "社护罗萨他那",
		TitleCode: "d_zabulistan",
		Counties:  map[string]feud.County{},
	}

	f.伽色尼Ghazna = ghazna.CGhazna伽色尼
	f.伽色尼Ghazna.SetParent(f)
	
	f.卡拉特Kalat = kalat.CKalat卡拉特
	f.卡拉特Kalat.SetParent(f)
	
	f.扎敏达瓦尔Zamindawar = zamindawar.CZamindawar扎敏达瓦尔
	f.扎敏达瓦尔Zamindawar.SetParent(f)
	
}
