/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"testing"

	cmg "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

const (
	httpPattern            = "/metrics"
	addr                   = "localhost:0"
	processStartTimeMetric = "process_start_time_seconds"
)

func initMgr() (MetricsManager, *http.Server) {
	mgr := NewMetricsManager()
	mux := http.NewServeMux()
	err := mgr.PrepareMetricsPath(mux, httpPattern, nil)
	if err != nil {
		log.Fatalf("failed to start serving [%v]", err)
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on address[%s], error[%v]", addr, err)
	}
	srv := &http.Server{Addr: l.Addr().String(), Handler: mux}
	go func() {
		if err := srv.Serve(l); err != http.ErrServerClosed {
			log.Fatalf("failed to start endpoint at:%s/%s, error: %v", addr, httpPattern, err)
		}
	}()

	return mgr, srv
}

func shutdown(srv *http.Server) {
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}

func TestNew(t *testing.T) {
	mgr, srv := initMgr()
	defer shutdown(srv)
	if mgr == nil {
		t.Errorf("failed testing new")
	}
}

func TestIncrementCount(t *testing.T) {
	mgr, srv := initMgr()
	srvAddr := "http://" + srv.Addr + httpPattern
	defer shutdown(srv)
	mgr.IncrementCount("test_result_1")
	mgr.IncrementCount("test_result_2")
	mgr.IncrementCount("test_result_2")

	expected :=
		`# HELP process_start_time_seconds [ALPHA] Start time of the process since unix epoch in seconds.
        # TYPE process_start_time_seconds gauge
        process_start_time_seconds 0
# HELP volume_data_source_validator_volume_data_source_validator_operation_count [ALPHA] Number of validations operations by result
# TYPE volume_data_source_validator_volume_data_source_validator_operation_count counter
volume_data_source_validator_operation_count{result="test_result_1"} 1
volume_data_source_validator_operation_count{result="test_result_2"} 2
`

	if err := verifyMetric(expected, srvAddr); err != nil {
		t.Errorf("failed testing [%v]", err)
	}
}

func verifyMetric(expected, srvAddr string) error {
	rsp, err := http.Get(srvAddr)
	if err != nil {
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get response from serve: %s", http.StatusText(rsp.StatusCode))
	}
	r, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	format := expfmt.ResponseFormat(rsp.Header)
	expectedReader := strings.NewReader(expected)
	expectedDecoder := expfmt.NewDecoder(expectedReader, format)
	expectedMfs := []*cmg.MetricFamily{}
	for {
		mf := &cmg.MetricFamily{}
		if err := expectedDecoder.Decode(mf); err != nil {
			// return correctly if EOF
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		expectedMfs = append(expectedMfs, mf)
	}

	gotReader := strings.NewReader(string(r))
	gotDecoder := expfmt.NewDecoder(gotReader, format)
	gotMfs := []*cmg.MetricFamily{}
	for {
		mf := &cmg.MetricFamily{}
		if err := gotDecoder.Decode(mf); err != nil {
			// return correctly if  EOF
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		gotMfs = append(gotMfs, mf)
	}

	if !containsMetrics(expectedMfs, gotMfs) {
		return fmt.Errorf("failed testing, expected\n%s\n, got\n%s\n", expected, string(r))
	}

	return nil
}

// sortMfs, sorts metric families in alphabetical order by type.
// currently only supports counter and histogram
func sortMfs(mfs []*cmg.MetricFamily) []*cmg.MetricFamily {
	var sortedMfs []*cmg.MetricFamily

	// Sort first by type
	sort.Slice(mfs, func(i, j int) bool {
		return *mfs[i].Type < *mfs[j].Type
	})

	// Next, sort by length of name
	sort.Slice(mfs, func(i, j int) bool {
		return len(*mfs[i].Name) < len(*mfs[j].Name)
	})

	return sortedMfs
}

func containsMetrics(expectedMfs, gotMfs []*cmg.MetricFamily) bool {
	if len(gotMfs) != len(expectedMfs) {
		fmt.Printf("Not same length: expected and got metrics families: %v vs. %v\n", len(expectedMfs), len(gotMfs))
		return false
	}

	// sort metric families for deterministic comparison.
	sortedExpectedMfs := sortMfs(expectedMfs)
	sortedGotMfs := sortMfs(gotMfs)

	// compare expected vs. sorted actual metrics
	for k, got := range sortedGotMfs {
		matchCount := 0
		expected := sortedExpectedMfs[k]

		if (got.Name == nil || *(got.Name) != *(expected.Name)) ||
			(got.Type == nil || *(got.Type) != *(expected.Type)) ||
			(got.Help == nil || *(got.Help) != *(expected.Help)) {
			fmt.Printf("invalid header info: got: %v, expected: %v\n", *got.Name, *expected.Name)
			fmt.Printf("invalid header info: got: %v, expected: %v\n", *got.Type, *expected.Type)
			fmt.Printf("invalid header info: got: %v, expected: %v\n", *got.Help, *expected.Help)
			return false
		}

		numRecords := len(expected.Metric)
		if len(got.Metric) < numRecords {
			fmt.Printf("Not the same number of records: got.Metric: %v, numRecords: %v\n", len(got.Metric), numRecords)
			return false
		}
		for i := 0; i < len(got.Metric); i++ {
			for j := 0; j < numRecords; j++ {
				if got.Metric[i].Histogram == nil && expected.Metric[j].Histogram != nil ||
					got.Metric[i].Histogram != nil && expected.Metric[j].Histogram == nil {
					fmt.Printf("got metric and expected metric histogram type mismatch")
					return false
				}

				// labels should be the same
				if !reflect.DeepEqual(got.Metric[i].Label, expected.Metric[j].Label) {
					continue
				}

				// metric type specific checks
				switch {
				case got.Metric[i].Histogram != nil && expected.Metric[j].Histogram != nil:
					gh := got.Metric[i].Histogram
					eh := expected.Metric[j].Histogram
					if gh == nil || eh == nil {
						continue
					}
					if !reflect.DeepEqual(gh.Bucket, eh.Bucket) {
						fmt.Println("got and expected histogram bucket not equal")
						continue
					}

					// this is a sum record, expecting a latency which is more than the
					// expected one. If the sum is smaller than expected, it will be considered
					// as NOT a match
					if gh.SampleSum == nil || eh.SampleSum == nil || *(gh.SampleSum) < *(eh.SampleSum) {
						fmt.Println("difference in sample sum")
						continue
					}
					if gh.SampleCount == nil || eh.SampleCount == nil || *(gh.SampleCount) != *(eh.SampleCount) {
						fmt.Println("difference in sample count")
						continue
					}

				case got.Metric[i].Counter != nil && expected.Metric[j].Counter != nil:
					gc := got.Metric[i].Counter
					ec := expected.Metric[j].Counter
					if gc.Value == nil || *(gc.Value) != *(ec.Value) {
						fmt.Println("difference in counter values")
						continue
					}
				}

				// this is a match
				matchCount = matchCount + 1
				break
			}
		}

		if matchCount != numRecords {
			fmt.Printf("matchCount %v, numRecords %v\n", matchCount, numRecords)
			return false
		}
	}

	return true
}

func TestProcessStartTimeMetricExist(t *testing.T) {
	mgr, srv := initMgr()
	defer shutdown(srv)
	metricsFamilies, err := mgr.GetRegistry().Gather()
	if err != nil {
		t.Fatalf("Error fetching metrics: %v", err)
	}

	for _, metricsFamily := range metricsFamilies {
		if metricsFamily.GetName() == processStartTimeMetric {
			return
		}
		m := metricsFamily.GetMetric()
		if m[0].GetGauge().GetValue() <= 0 {
			t.Fatalf("Expected non zero timestamp for process start time")
		}
	}

	t.Fatalf("Metrics does not contain %v. Scraped content: %v", processStartTimeMetric, metricsFamilies)
}
