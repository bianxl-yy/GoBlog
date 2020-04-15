package utils

import (
	"fmt"
	"github.com/rs/xid"
	"testing"
)

func TestXidSort(t *testing.T) {
	/*
		xids := make([]string, 0)
		for i := 0; i < 10; i++ {
			uid := xid.New()
			xids = append(xids, uid.String())
		}
		fmt.Printf("%#v", xids)
	*/
	// []string{"bq7f7jmeivh52q2hcsvg", "bq7f7jmeivh52q2hct00", "bq7f7jmeivh52q2hct0g", "bq7f7jmeivh52q2hct10", "bq7f7jmeivh52q2hct1g", "bq7f7jmeivh52q2hct20", "bq7f7jmeivh52q2hct2g", "bq7f7jmeivh52q2hct30", "bq7f7jmeivh52q2hct3g", "bq7f7jmeivh52q2hct40"}--- PASS: TestXidSort (0.00s)
	ids := []string{"bq7f7jmeivh52q2hcsvg", "bq7f7jmeivh52q2hct10", "bq7f7jmeivh52q2hct00", "bq7f7jmeivh52q2hct30", "bq7f7jmeivh52q2hct0g", "bq7f7jmeivh52q2hct1g", "bq7f7jmeivh52q2hct20", "bq7f7jmeivh52q2hct40", "bq7f7jmeivh52q2hct2g", "bq7f7jmeivh52q2hct3g"}
	idsSort, err := XidSort(ids)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%#v", idsSort)

	xids := xid.New()
	fmt.Println(xids.String(), "-", xids.String())

}
