package config

import "testing"

func TestCreateWithError(t *testing.T) {
	factory := EnvConfigFactory{}
	_, err := factory.Create()
	if err != nil {
		if err.Error() != ErrEnvEmptyToken.Error() {
			t.Errorf("unexpected error [%v]", err.Error())
		}
	} else {
		t.Errorf("misconfigured test environment")
	}

}
