package remote_package_test

import (
	"testing"

	cm "github.com/easuerway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
}
