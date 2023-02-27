package nagchu

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/nagchu/nagchu"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/nagchu/nyima"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/nagchu/xainza"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagchuDuke interface {
    feud.Duke
    CNagchu那曲() 	nagchu.NagchuCounty
    CNyima尼玛() 	nyima.NyimaCounty
    CXainza申扎() 	xainza.XainzaCounty
}

type 那曲NagchuDuke struct {
	feud.BaseDuke
	那曲Nagchu 	nagchu.NagchuCounty
	尼玛Nyima 	nyima.NyimaCounty
	申扎Xainza 	xainza.XainzaCounty
}

func (d *那曲NagchuDuke) CNagchu那曲() nagchu.NagchuCounty {
	return d.那曲Nagchu
}
    
func (d *那曲NagchuDuke) CNyima尼玛() nyima.NyimaCounty {
	return d.尼玛Nyima
}
    
func (d *那曲NagchuDuke) CXainza申扎() xainza.XainzaCounty {
	return d.申扎Xainza
}
    
var DNagchu那曲 NagchuDuke = &那曲NagchuDuke{}

func init() {
	f := DNagchu那曲.(*那曲NagchuDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nagchu",
		TitleName: "那曲",
		TitleCode: "d_nagchu",
		Counties:  map[string]feud.County{},
	}

	f.那曲Nagchu = nagchu.CNagchu那曲
	f.那曲Nagchu.SetParent(f)
	
	f.尼玛Nyima = nyima.CNyima尼玛
	f.尼玛Nyima.SetParent(f)
	
	f.申扎Xainza = xainza.CXainza申扎
	f.申扎Xainza.SetParent(f)
	
}
