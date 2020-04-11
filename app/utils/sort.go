package utils

import (
	"fmt"
	"github.com/rs/xid"
)

func XidSort(ids []string) ([]string, error) {
	idSlice := make([]xid.ID, 0)
	for _, v := range ids {
		id, err := xid.FromString(v)
		if err != nil {
			return nil, err
		}
		idSlice = append(idSlice, id)
	}
	xid.Sort(idSlice)
	result := make([]string, 0)
	for _, v := range idSlice {
		id := fmt.Sprintf("%s", v)
		result = append(result, id)
	}
	return result, nil
}
