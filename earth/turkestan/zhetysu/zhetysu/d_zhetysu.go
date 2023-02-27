package zhetysu

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/zhetysu/almaty"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/zhetysu/karluk"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/zhetysu/zhetysu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZhetysuDuke interface {
    feud.Duke
    CAlmaty阿拉木图() 	almaty.AlmatyCounty
    CKarluk葛逻禄() 	karluk.KarlukCounty
    CZhetysu卡拉塔尔() 	zhetysu.ZhetysuCounty
}

type 七河ZhetysuDuke struct {
	feud.BaseDuke
	阿拉木图Almaty 	almaty.AlmatyCounty
	葛逻禄Karluk 	karluk.KarlukCounty
	卡拉塔尔Zhetysu 	zhetysu.ZhetysuCounty
}

func (d *七河ZhetysuDuke) CAlmaty阿拉木图() almaty.AlmatyCounty {
	return d.阿拉木图Almaty
}
    
func (d *七河ZhetysuDuke) CKarluk葛逻禄() karluk.KarlukCounty {
	return d.葛逻禄Karluk
}
    
func (d *七河ZhetysuDuke) CZhetysu卡拉塔尔() zhetysu.ZhetysuCounty {
	return d.卡拉塔尔Zhetysu
}
    
var DZhetysu七河 ZhetysuDuke = &七河ZhetysuDuke{}

func init() {
	f := DZhetysu七河.(*七河ZhetysuDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "zhetysu",
		TitleName: "七河",
		TitleCode: "d_zhetysu",
		Counties:  map[string]feud.County{},
	}

	f.阿拉木图Almaty = almaty.CAlmaty阿拉木图
	f.阿拉木图Almaty.SetParent(f)
	
	f.葛逻禄Karluk = karluk.CKarluk葛逻禄
	f.葛逻禄Karluk.SetParent(f)
	
	f.卡拉塔尔Zhetysu = zhetysu.CZhetysu卡拉塔尔
	f.卡拉塔尔Zhetysu.SetParent(f)
	
}
