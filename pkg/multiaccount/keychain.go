package multiaccount

import (
	"os"

	"github.com/majd/ipatool/v2/pkg/keychain"
)

const profileEnvKey = "IPATOOL_PROFILE"

type profileKeychain struct {
	base keychain.Keychain
}

// NewProfileKeychain wraps an existing keychain and prefixes keys by profile.
// When IPATOOL_PROFILE is empty, keys are unchanged for backward compatibility.
func NewProfileKeychain(base keychain.Keychain) keychain.Keychain {
	return &profileKeychain{base: base}
}

func (k *profileKeychain) transform(key string) string {
	profile := os.Getenv(profileEnvKey)
	if profile == "" {
		// keep old behaviour: single global "account" key
		return key
	}

	return profile + ":" + key
}
