package testingutils

import (
	"github.com/mosheblox/ssv-spec/qbft"
)

func UnknownDutyValueCheck() qbft.ProposedValueCheckF {
	return func(data []byte) error {
		return nil
	}
}
