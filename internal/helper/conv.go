package helper

import "strconv"

func StrSliceToUint64SliceMap(ids []string) (map[uint64]struct{}, error) {
	set := make(map[uint64]struct{})
	for _, str := range ids {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		set[i] = struct{}{}
	}
	return set, nil
}

func StrSliceToUint64Slice(ids []string) ([]uint64, error) {
	set := make([]uint64, 0, len(ids))
	for _, str := range ids {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		set = append(set, i)
	}
	return set, nil
}
