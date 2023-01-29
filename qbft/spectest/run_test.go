package spectest

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	json "github.com/bytedance/sonic"

	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/timeout"

	"github.com/bloxapp/ssv-spec/qbft"
	tests2 "github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller/futuremsg"
	"github.com/bloxapp/ssv-spec/types/testingutils"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	for _, test := range AllTests {
		t.Run(test.TestName(), func(t *testing.T) {
			test.Run(t)
		})
	}
}

func TestJson(t *testing.T) {
	basedir, _ := os.Getwd()
	path := filepath.Join(basedir, "generate")
	fileName := "tests.json"
	untypedTests := map[string]interface{}{}
	byteValue, err := os.ReadFile(path + "/" + fileName)
	if err != nil {
		panic(err.Error())
	}

	if err := json.ConfigFastest.Unmarshal(byteValue, &untypedTests); err != nil {
		panic(err.Error())
	}

	tests := make(map[string]SpecTest)
	for name, test := range untypedTests {
		testName := test.(map[string]interface{})["Name"].(string)
		t.Run(testName, func(t *testing.T) {
			testType := strings.Split(name, "_")[0]
			switch testType {
			case reflect.TypeOf(&tests2.MsgProcessingSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &tests2.MsgProcessingSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				// a little trick we do to instantiate all the internal instance params
				preByts, _ := typedTest.Pre.Encode()
				pre := qbft.NewInstance(
					testingutils.TestingConfig(testingutils.KeySetForShare(typedTest.Pre.State.Share)),
					typedTest.Pre.State.Share,
					typedTest.Pre.State.ID,
					typedTest.Pre.State.Height,
				)
				err = pre.Decode(preByts)
				require.NoError(t, err)
				typedTest.Pre = pre

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&tests2.MsgSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &tests2.MsgSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&tests2.ControllerSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &tests2.ControllerSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&tests2.CreateMsgSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &tests2.CreateMsgSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&tests2.RoundRobinSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &tests2.RoundRobinSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&futuremsg.ControllerSyncSpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &futuremsg.ControllerSyncSpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				tests[testName] = typedTest
				typedTest.Run(t)
			case reflect.TypeOf(&timeout.SpecTest{}).String():
				byts, err := json.ConfigFastest.Marshal(test)
				require.NoError(t, err)
				typedTest := &timeout.SpecTest{}
				require.NoError(t, json.ConfigFastest.Unmarshal(byts, &typedTest))

				// a little trick we do to instantiate all the internal instance params
				preByts, _ := typedTest.Pre.Encode()
				pre := qbft.NewInstance(
					testingutils.TestingConfig(testingutils.KeySetForShare(typedTest.Pre.State.Share)),
					typedTest.Pre.State.Share,
					typedTest.Pre.State.ID,
					typedTest.Pre.State.Height,
				)
				err = pre.Decode(preByts)
				require.NoError(t, err)
				typedTest.Pre = pre

				tests[testName] = typedTest
				typedTest.Run(t)
			default:
				panic("unsupported test type " + testType)
			}
		})
	}
}
