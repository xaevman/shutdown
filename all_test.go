//  ---------------------------------------------------------------------------
//
//  all_test.go
//
//  Copyright (c) 2014, Jared Chavez.
//  All rights reserved.
//
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.
//
//  -----------

package shutdown

import (
    "testing"
    "time"
)

func TestShutdown(t *testing.T) {
    s := New()

    if s.IsShutdown() {
        t.Errorf("Shutdown sync option was incorrectly triggered at instantiation")
    }

    go s.Start()

    select {
    case <-s.Signal:
        s.Complete()
    }

    if s.WaitForTimeout() {
        t.Errorf("Shutdown reached timeout despite beging triggered")
    }

    if !s.IsShutdown() {
        t.Errorf("Sync objection has shutdown but is not marked as complete")
    }
}

func TestTimeout(t *testing.T) {
    s := NewTimeout(1)

    ltchan := make(chan interface{}, 1)
    stchan := make(chan bool, 1)

    go func() {
        <-time.After(2 * time.Second)
        ltchan <- nil
    }()

    go func() {
        result := s.WaitForTimeout()
        stchan <- result
    }()

    select {
    case <-ltchan:
        t.Errorf("Time out took longer than expected")
    case val := <-stchan:
        if !val {
            t.Errorf("Sync object completed when timeout was expected")
        }
    }
}
