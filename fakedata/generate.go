package fakedata

import (
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

const HyphenTruncatedUUIDTagName = "hyphenTruncatedUUID"
const RFC3339TagName = "rfc3339"

func init() {
	_ = faker.AddProvider(HyphenTruncatedUUIDTagName, func(v reflect.Value) (any, error) {
		return strings.ReplaceAll(uuid.NewString(), "-", ""), nil
	})
	_ = faker.AddProvider(RFC3339TagName, func (v reflect.Value) (any, error)  {
		return time.Unix(faker.RandomUnixTime(), 0).Format(time.RFC3339), nil
	})
	// structのsliceにlenを指定するとUnknown Typeエラーになるので最大、最小を固定する
	faker.SetRandomMapAndSliceMaxSize(3)
	faker.SetRandomMapAndSliceMinSize(1)
}

func GenerateFakeData(model any) (any, error) {
	err := faker.FakeData(&model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
