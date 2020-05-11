// Copyright (c) 2017 Daniel Oaks
// released under the MIT license

package irc

import (
	"reflect"
	"testing"

	"github.com/oragono/oragono/irc/modes"
)

func TestParseDefaultChannelModes(t *testing.T) {
	nt := "+nt"
	n := "+n"
	empty := ""
	tminusi := "+t -i"

	var parseTests = []struct {
		raw      *string
		expected modes.Modes
	}{
		{&nt, modes.Modes{modes.NoOutside, modes.OpOnlyTopic}},
		{&n, modes.Modes{modes.NoOutside}},
		{&empty, modes.Modes{}},
		{&tminusi, modes.Modes{modes.OpOnlyTopic}},
		{nil, modes.Modes{modes.NoOutside, modes.OpOnlyTopic}},
	}

	for _, testcase := range parseTests {
		result := ParseDefaultChannelModes(testcase.raw)
		if !reflect.DeepEqual(result, testcase.expected) {
			t.Errorf("expected modes %s, got %s", testcase.expected, result)
		}
	}
}

func TestParseDefaultUserModes(t *testing.T) {
	iR := "+iR"
	i := "+i"
	empty := ""
	rminusi := "+R -i"

	var parseTests = []struct {
		raw      *string
		expected modes.ModeChanges
	}{
		{&iR, modes.ModeChanges{{Mode: modes.Invisible, Op: modes.Add}, {Mode: modes.RegisteredOnly, Op: modes.Add}}},
		{&i, modes.ModeChanges{{Mode: modes.Invisible, Op: modes.Add}}},
		{&empty, modes.ModeChanges{}},
		{&rminusi, modes.ModeChanges{{Mode: modes.RegisteredOnly, Op: modes.Add}}},
		{nil, modes.ModeChanges{}},
	}

	for _, testcase := range parseTests {
		result := ParseDefaultUserModes(testcase.raw)
		if !reflect.DeepEqual(result, testcase.expected) {
			t.Errorf("expected modes %v, got %v", testcase.expected, result)
		}
	}
}

func TestUmodeGreaterThan(t *testing.T) {
	if !umodeGreaterThan(modes.Halfop, modes.Voice) {
		t.Errorf("expected Halfop > Voice")
	}

	if !umodeGreaterThan(modes.Voice, modes.Mode(0)) {
		t.Errorf("expected Voice > 0 (the zero value of modes.Mode)")
	}

	if umodeGreaterThan(modes.ChannelAdmin, modes.ChannelAdmin) {
		t.Errorf("modes should not be greater than themselves")
	}
}

func assertEqual(supplied, expected interface{}, t *testing.T) {
	if !reflect.DeepEqual(supplied, expected) {
		t.Errorf("expected %v but got %v", expected, supplied)
	}
}

func TestChannelUserModeHasPrivsOver(t *testing.T) {
	assertEqual(channelUserModeHasPrivsOver(modes.Voice, modes.Halfop), false, t)
	assertEqual(channelUserModeHasPrivsOver(modes.Mode(0), modes.Halfop), false, t)
	assertEqual(channelUserModeHasPrivsOver(modes.Voice, modes.Mode(0)), false, t)
	assertEqual(channelUserModeHasPrivsOver(modes.ChannelAdmin, modes.ChannelAdmin), false, t)
	assertEqual(channelUserModeHasPrivsOver(modes.Halfop, modes.Halfop), false, t)
	assertEqual(channelUserModeHasPrivsOver(modes.Voice, modes.Voice), false, t)

	assertEqual(channelUserModeHasPrivsOver(modes.Halfop, modes.Voice), true, t)
	assertEqual(channelUserModeHasPrivsOver(modes.ChannelFounder, modes.ChannelAdmin), true, t)
	assertEqual(channelUserModeHasPrivsOver(modes.ChannelOperator, modes.ChannelOperator), true, t)
}
