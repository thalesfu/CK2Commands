package alania

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/alania"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/azov"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/caspian_steppe"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/derbent"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlaniaKingdom interface {
    feud.Kingdom
    DAlania阿兰尼亚() 	alania.AlaniaDuke
    DAzov亚速() 	azov.AzovDuke
    DCaspian_steppe里海草原() 	caspian_steppe.Caspian_steppeDuke
    DDerbent杰尔宾特() 	derbent.DerbentDuke
}

type 阿兰尼亚AlaniaKingdom struct {
	feud.BaseKingdom
	阿兰尼亚Alania 	alania.AlaniaDuke
	亚速Azov 	azov.AzovDuke
	里海草原Caspian_steppe 	caspian_steppe.Caspian_steppeDuke
	杰尔宾特Derbent 	derbent.DerbentDuke
}

func (k *阿兰尼亚AlaniaKingdom) DAlania阿兰尼亚() alania.AlaniaDuke {
	return k.阿兰尼亚Alania
}
    
func (k *阿兰尼亚AlaniaKingdom) DAzov亚速() azov.AzovDuke {
	return k.亚速Azov
}
    
func (k *阿兰尼亚AlaniaKingdom) DCaspian_steppe里海草原() caspian_steppe.Caspian_steppeDuke {
	return k.里海草原Caspian_steppe
}
    
func (k *阿兰尼亚AlaniaKingdom) DDerbent杰尔宾特() derbent.DerbentDuke {
	return k.杰尔宾特Derbent
}
    
var KAlania阿兰尼亚 AlaniaKingdom = &阿兰尼亚AlaniaKingdom{}

func init() {
	f := KAlania阿兰尼亚.(*阿兰尼亚AlaniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "alania",
		TitleName: "阿兰尼亚",
		TitleCode: "k_alania",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿兰尼亚Alania = alania.DAlania阿兰尼亚
	f.阿兰尼亚Alania.SetParent(f)
	
	f.亚速Azov = azov.DAzov亚速
	f.亚速Azov.SetParent(f)
	
	f.里海草原Caspian_steppe = caspian_steppe.DCaspian_steppe里海草原
	f.里海草原Caspian_steppe.SetParent(f)
	
	f.杰尔宾特Derbent = derbent.DDerbent杰尔宾特
	f.杰尔宾特Derbent.SetParent(f)
	
}
