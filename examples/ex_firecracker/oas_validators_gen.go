// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func (s Drive) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.RateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "rate_limiter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s FullVmConfiguration) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.BlockDevices {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "block_devices",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Logger // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "logger",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.MachineConfig // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "machine_config",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.NetDevices {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "net_devices",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.VsockDevice // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "vsock_device",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s InstanceActionInfo) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.ActionType // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "action_type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s InstanceInfo) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.State // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "state",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Logger) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Level // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "level",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s MachineConfiguration) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.CPUTemplate // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "cpu_template",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          1,
			MaxSet:       true,
			Max:          32,
			MinExclusive: false,
			MaxExclusive: false,
		}).Validate(int64(s.VcpuCount)); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "vcpu_count",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s NetworkInterface) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.RxRateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "rx_rate_limiter",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.TxRateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tx_rate_limiter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PartialDrive) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.RateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "rate_limiter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PartialNetworkInterface) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.RxRateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "rx_rate_limiter",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.TxRateLimiter // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "tx_rate_limiter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s RateLimiter) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Bandwidth // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "bandwidth",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Ops // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "ops",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s SnapshotCreateParams) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.SnapshotType // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "snapshot_type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s TokenBucket) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.OneTimeBurst // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "one_time_burst",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: false,
			MaxExclusive: false,
		}).Validate(int64(s.RefillTime)); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "refill_time",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: false,
			MaxExclusive: false,
		}).Validate(int64(s.Size)); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "size",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s VM) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.State // validation expected, but not supported
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "state",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s Vsock) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          3,
			MaxSet:       false,
			Max:          0,
			MinExclusive: false,
			MaxExclusive: false,
		}).Validate(int64(s.GuestCid)); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "guest_cid",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}