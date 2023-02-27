package ruthenia

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/kiev"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/minsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/pereyaslavl"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/polotsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/smolensk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/turov"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RutheniaKingdom interface {
    feud.Kingdom
    DKiev基辅() 	kiev.KievDuke
    DMinsk明斯克() 	minsk.MinskDuke
    DPereyaslavl佩列亚斯拉夫尔() 	pereyaslavl.PereyaslavlDuke
    DPolotsk波洛茨克() 	polotsk.PolotskDuke
    DSmolensk斯摩棱斯克() 	smolensk.SmolenskDuke
    DTurov图罗夫() 	turov.TurovDuke
}

type 鲁塞尼亚RutheniaKingdom struct {
	feud.BaseKingdom
	基辅Kiev 	kiev.KievDuke
	明斯克Minsk 	minsk.MinskDuke
	佩列亚斯拉夫尔Pereyaslavl 	pereyaslavl.PereyaslavlDuke
	波洛茨克Polotsk 	polotsk.PolotskDuke
	斯摩棱斯克Smolensk 	smolensk.SmolenskDuke
	图罗夫Turov 	turov.TurovDuke
}

func (k *鲁塞尼亚RutheniaKingdom) DKiev基辅() kiev.KievDuke {
	return k.基辅Kiev
}
    
func (k *鲁塞尼亚RutheniaKingdom) DMinsk明斯克() minsk.MinskDuke {
	return k.明斯克Minsk
}
    
func (k *鲁塞尼亚RutheniaKingdom) DPereyaslavl佩列亚斯拉夫尔() pereyaslavl.PereyaslavlDuke {
	return k.佩列亚斯拉夫尔Pereyaslavl
}
    
func (k *鲁塞尼亚RutheniaKingdom) DPolotsk波洛茨克() polotsk.PolotskDuke {
	return k.波洛茨克Polotsk
}
    
func (k *鲁塞尼亚RutheniaKingdom) DSmolensk斯摩棱斯克() smolensk.SmolenskDuke {
	return k.斯摩棱斯克Smolensk
}
    
func (k *鲁塞尼亚RutheniaKingdom) DTurov图罗夫() turov.TurovDuke {
	return k.图罗夫Turov
}
    
var KRuthenia鲁塞尼亚 RutheniaKingdom = &鲁塞尼亚RutheniaKingdom{}

func init() {
	f := KRuthenia鲁塞尼亚.(*鲁塞尼亚RutheniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "ruthenia",
		TitleName: "鲁塞尼亚",
		TitleCode: "k_ruthenia",
		Dukes:  map[string]feud.Duke{},
	}

	f.基辅Kiev = kiev.DKiev基辅
	f.基辅Kiev.SetParent(f)
	
	f.明斯克Minsk = minsk.DMinsk明斯克
	f.明斯克Minsk.SetParent(f)
	
	f.佩列亚斯拉夫尔Pereyaslavl = pereyaslavl.DPereyaslavl佩列亚斯拉夫尔
	f.佩列亚斯拉夫尔Pereyaslavl.SetParent(f)
	
	f.波洛茨克Polotsk = polotsk.DPolotsk波洛茨克
	f.波洛茨克Polotsk.SetParent(f)
	
	f.斯摩棱斯克Smolensk = smolensk.DSmolensk斯摩棱斯克
	f.斯摩棱斯克Smolensk.SetParent(f)
	
	f.图罗夫Turov = turov.DTurov图罗夫
	f.图罗夫Turov.SetParent(f)
	
}
