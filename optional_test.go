package golang_optional

import (
	"encoding/json"
	"testing"
)

func TestOptionalString(t *testing.T) {
	t.Run("Unmarshal", func(t *testing.T) {
		t.Run("Unspecified String", func(t *testing.T) {
			type obj struct {
				optionalVal Optional[string]
			}

			var actual obj
			err := json.Unmarshal([]byte(`{}`), &actual)
			if err != nil {
				t.Errorf("Unmarshal failed: %v", err)
			}

			if !actual.optionalVal.IsEmpty() {
				t.Errorf("IsEmpty() should return true")
			}
		})

		t.Run("Null String", func(t *testing.T) {
			type obj struct {
				OptionalVal Optional[string]
			}

			var actual obj
			err := json.Unmarshal([]byte(`{
			"OptionalVal": null
		}`), &actual)
			if err != nil {
				t.Errorf("Unmarshal failed: %v", err)
			}

			if !actual.OptionalVal.IsEmpty() {
				t.Errorf("IsEmpty() should return true")
			}
		})

		t.Run("Empty String", func(t *testing.T) {
			type obj struct {
				OptionalVal Optional[string]
			}

			var actual obj
			err := json.Unmarshal([]byte(`{
			"OptionalVal": ""
		}`), &actual)
			if err != nil {
				t.Errorf("Unmarshal failed: %v", err)
			}

			if actual.OptionalVal.IsEmpty() {
				t.Errorf("IsEmpty() should return false")
			}

			val, ok := actual.OptionalVal.Value()
			if val != "" || !ok {
				t.Errorf("Value() should return '', true")
			}
		})

		t.Run("String with Value", func(t *testing.T) {
			type obj struct {
				OptionalVal Optional[string]
			}

			var actual obj
			err := json.Unmarshal([]byte(`{
			"OptionalVal": "value"
		}`), &actual)
			if err != nil {
				t.Errorf("Unmarshal failed: %v", err)
			}

			if actual.OptionalVal.IsEmpty() {
				t.Errorf("IsEmpty() should return false")
			}

			val, ok := actual.OptionalVal.Value()
			if val != "value" || !ok {
				t.Errorf("Value() should return 'value', true")
			}
		})
	})
}
