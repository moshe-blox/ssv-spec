package testingutils

import (
	"github.com/moshe-blox/ssv-spec/qbft"
)

func UnknownDutyValueCheck() qbft.ProposedValueCheckF {
	return func(data []byte) error {
		return nil
	}
}
